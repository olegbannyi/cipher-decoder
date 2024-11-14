package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	willYouMarryMe          = "Will you marry me?"
	yes                     = "Yeah, okay!"
	letsBeFriends           = "Let's be friends."
	enthusiasticAffirmation = "Great!"
	disappointmentMessage   = "What a pity!"
)

const (
	numberOfLetters = 26
)

func main() {
	var g, p, A int

	fmt.Scanf("g is %d and p is %d", &g, &p)
	fmt.Println("OK")
	fmt.Scanf("A is %d", &A)

	b := 7

	B := calc(g, p, b)
	var S int = calc(A, p, b)

	fmt.Printf("B is %d\n", B)

	fmt.Println(encryptCaesar(willYouMarryMe, S))

	answer := getAnswer(S)

	switch answer {
	case yes:
		fmt.Println(encryptCaesar(enthusiasticAffirmation, S))
	case letsBeFriends:
		fmt.Println(encryptCaesar(disappointmentMessage, S))
	}
}

func getAnswer(shift int) string {
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	return decryptCaesar(strings.TrimSpace(answer), shift)
}

func calc(g, p, n int) int {
	c := 1

	for i := 0; i < n; i++ {
		c = (c * g) % p
	}

	return c
}

func encryptCaesar(originalText string, shift int) string {
	cipherText := ""
	shift = shift % numberOfLetters

	for _, char := range originalText {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+rune(shift))%26
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+rune(shift))%26
		}

		cipherText += string(char)
	}

	return cipherText
}

func decryptCaesar(cipherText string, shift int) string {
	plainText := ""
	shift = shift % numberOfLetters

	for _, char := range cipherText {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'-rune(shift)+26)%26
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'-rune(shift)+26)%26
		}

		plainText += string(char)
	}

	return plainText
}
