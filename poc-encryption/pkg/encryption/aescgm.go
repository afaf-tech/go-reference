package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

type AESGCMEncryptor struct {
	keyProvider KeyProvider
	keyCache    KeyCache

	localKey     *CachedKey
	localKeyLock sync.RWMutex
}

func NewEncryptor(provider KeyProvider, cache KeyCache) Encryptor {
	instance := &AESGCMEncryptor{keyProvider: provider, keyCache: cache}

	if err := instance.warmKey(); err != nil {
		fmt.Println("Failed to warm up key:", err)
	}
	return instance
}

func (a *AESGCMEncryptor) getKey() ([]byte, error) {
	// First, check in-memory cache
	a.localKeyLock.RLock()
	if a.localKey != nil && a.localKey.ExpiresAt > time.Now().Unix() {
		defer a.localKeyLock.RUnlock()
		return a.localKey.Value, nil
	}
	a.localKeyLock.RUnlock()

	// Fallback to Redis cache
	// k, ok, err := a.keyCache.GetKey()
	// if err != nil {
	// 	fmt.Println("err on get Key", err.Error())
	// }
	// if ok && k.ExpiresAt > time.Now().Unix() {
	// 	a.localKeyLock.Lock()
	// 	a.localKey = k
	// 	a.localKeyLock.Unlock()
	// 	return k.Value, nil
	// }

	// Fallback to provider
	freshKey, err := a.keyProvider.FetchKey()
	if err != nil {
		return nil, err
	}

	a.keyCache.SetKey(freshKey)

	a.localKeyLock.Lock()
	a.localKey = &CachedKey{
		Value:     freshKey.Value,
		ExpiresAt: freshKey.ExpiresAt,
	}
	a.localKeyLock.Unlock()

	return freshKey.Value, nil

}

func (a *AESGCMEncryptor) warmKey() error {
	_, err := a.getKey()
	return err
}

func (a *AESGCMEncryptor) Encrypt(plain []byte) ([]byte, error) {
	key, err := a.getKey()
	if err != nil {
		return nil, err
	}
	// fmt.Println("encrypt with key =", base64.StdEncoding.EncodeToString(key))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("unexpected key length: got %d, want 32", len(key))
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := aesgcm.Seal(nil, nonce, plain, nil)
	return append(nonce, ciphertext...), nil
}

func (a *AESGCMEncryptor) Decrypt(ciphertext []byte) ([]byte, error) {
	key, err := a.getKey()
	if err != nil {
		return nil, err
	}
	// fmt.Println("decrypt with key =", base64.StdEncoding.EncodeToString(key))

	if len(key) != 32 {
		return nil, fmt.Errorf("unexpected key length: got %d, want 32", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
