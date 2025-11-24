package tests

import (
	"testing"

	"github.com/dlambda/zup/pkg/zup"
)

const (
	TEST_KEY = "284a04ad4b11573905b1e4028d6568434e999437a74810fa52f7c9df1b2dc5ea"
)

/*
	func TestOpenZup(t *testing.T) {
		name := "testfile"
		key, err := zup.InitZup(name)
		if err != nil {

		}
	}
*/
func TestEncryption(t *testing.T) {
	testStrings := [...]string{
		"Do you think love can bloom? even on a battlefield?",
		"It was the best of times, it was the worst of times.",
		"A Monad is just a monoid in the category of endofunctors",
		"Were you rushing or were you dragging? ANSWER!",
		"Where'd everybody go? Bingo?",
	}

	for _, testString := range testStrings {
		key, err := zup.GenerateKey(zup.KEY_SIZE)
		if err != nil {
			t.Fatalf("Failed to decode string: %v", err)
		}

		encrypted, err := zup.Encrypt(testString, &key)
		if err != nil {
			t.Fatalf("Failed to encrypt content: %v", err)
		}

		decrypted, err := zup.Decrypt(encrypted, &key)
		if err != nil {
			t.Fatalf("failed to decrypt content: %v", err)
		}

		if decrypted != testString {
			t.Fatalf("Decrypted content does not match original. %s != %s", testString, decrypted)
		}
	}
	
}
