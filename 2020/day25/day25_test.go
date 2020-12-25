package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	subject := 7

	cardPublicKey := 5764801
	cardLoopSize := determineLoopSize(subject, cardPublicKey)
	if cardLoopSize != 8 {
		t.Errorf("wanted cardLoopSize = 8, got %d", cardLoopSize)
	}

	doorPublicKey := 17807724
	doorLoopSize := determineLoopSize(subject, doorPublicKey)
	if doorLoopSize != 11 {
		t.Errorf("wanted doorLoopSize = 11, got %d", doorLoopSize)
	}

	encryptionKey := encrypt(subject, cardLoopSize, doorLoopSize)

	if encryptionKey != 14897079 {
		t.Errorf("wanted encryptionKey = 14897079, got %d", encryptionKey)
	}
}
