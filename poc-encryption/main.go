package main

import (
	"context"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"poc-encryption/pkg/config"
	"poc-encryption/pkg/connector"
	"poc-encryption/pkg/encryption"
	"time"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	redisClient, err := connector.ConnectToRedis(context.Background(), connector.RedisConfig{
		Addr:     cfg.Redis.Addr,
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect to Redis: %w", err))
	}
	defer redisClient.Close()

	keyCache := encryption.NewRedisKeyCache(redisClient, cfg.KeyCachePrefix)
	keyProvider := encryption.NewMockKeyProvider()
	encryptor := encryption.NewEncryptor(keyProvider, keyCache)

	file, err := os.Open(cfg.CSVPath)
	if err != nil {
		panic(fmt.Errorf("failed to open CSV: %w", err))
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var data []string
	rowIndex := 0

	for rowIndex < cfg.MaxRowsToRead {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error reading row %d: %v\n", rowIndex, err)
			continue
		}
		if len(record) < 8 {
			fmt.Printf("skipping row %d: not enough columns\n", rowIndex)
			continue
		}
		data = append(data, record[7])
		rowIndex++
	}
	if len(data) == 0 {
		panic("no valid data found in column 8 (H)")
	}
	fmt.Printf("Loaded %d rows from CSV\n", len(data))

	// Parallel encryption
	startEnc := time.Now()
	encrypted := make([][]byte, len(data))
	// var wg sync.WaitGroup
	// var mu sync.Mutex

	for i, plaintext := range data {
		// wg.Add(1)
		// go func(i int, plaintext string) {
		// defer wg.Done()
		ciphertext, err := encryptor.Encrypt([]byte(plaintext))
		if err != nil {
			fmt.Printf("error encrypting row %d: %v\n", i, err)
			return
		}

		// mu.Lock()
		encrypted[i] = ciphertext
		// mu.Unlock()
		// }(i, plaintext)
	}

	// wg.Wait()
	fmt.Printf("Encryption completed in %v\n", time.Since(startEnc))

	newFile, err := os.Create(cfg.OutputCSVPath)
	if err != nil {
		panic(fmt.Errorf("failed to create output CSV: %w", err))
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	defer writer.Flush()

	for i, enc := range encrypted {
		if enc == nil {
			continue
		}
		row := []string{
			data[i],
			base64.StdEncoding.EncodeToString(enc),
		}
		if err := writer.Write(row); err != nil {
			fmt.Printf("failed to write row %d: %v\n", i, err)
		}
	}
	fmt.Println("Data and encrypted columns saved to:", cfg.OutputCSVPath)

	startDec := time.Now()
	for i, ciphertext := range encrypted {
		if ciphertext == nil {
			continue
		}
		decrypted, err := encryptor.Decrypt(ciphertext)
		if err != nil {
			fmt.Printf("decryption failed at row %d: %v\n", i, err)
			panic(err)
		}
		_ = decrypted
		// fmt.Println("decrypt data", string(decrypted))
	}
	fmt.Printf("Decryption completed in %v\n", time.Since(startDec))
}
