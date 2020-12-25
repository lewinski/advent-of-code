package main

import (
	"fmt"
)

func main() {
	subject := 7

	cardPublicKey := 16915772
	doorPublicKey := 18447943

	cardLoopSize := determineLoopSize(subject, cardPublicKey)
	doorLoopSize := determineLoopSize(subject, doorPublicKey)

	encryptionKey := encrypt(subject, cardLoopSize, doorLoopSize)
	fmt.Println("part1:", encryptionKey)
}

func determineLoopSize(subject int, publicKey int) int {
	test := 1
	for loopSize := 1; ; loopSize++ {
		test *= subject
		test %= 20201227
		if test == publicKey {
			return loopSize
		}
	}
}

func transform(subject int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subject
		value %= 20201227
	}
	return value
}

func encrypt(subject int, cardLoopSize, doorLoopSize int) int {
	return transform(transform(subject, cardLoopSize), doorLoopSize)
}
