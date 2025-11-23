package tests

import (
	"os"
	"testing"
)

func testWrite(t *testing.T) {
	files, err := os.ReadDir("entries")
	if err != nil {
		return nil, err
	}
}

func testRead(t *testing.T) {
	files, err := os.ReadDir("entries")
	if err != nil {
		return nil, err
	}
}

