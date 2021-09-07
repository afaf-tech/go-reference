package belajar_golang_mbed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}
