// Package main contain usage examples of github.com/pablotrinidad/hillcipher/cipher
package main

import (
	"fmt"

	"github.com/gookit/color"
	hcipher "github.com/pablotrinidad/hillcipher/cipher"
)

type keyTextPair [2]string

func main() {
	examples := []struct{
		name, alphabet string
		samples []keyTextPair
	}{
		{
			name: "Spanish alphabet (uppercase) without diacritics",
			alphabet: "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			samples: []keyTextPair{
				keyTextPair{"FORTALEZA", "CONSUL"},
				keyTextPair{"FORTALEZA", "UUNAMFCIENCIASS"},
				keyTextPair{"IKEY", "CRIPTOGRAFIA"},
				keyTextPair{"IAMAVERYLOONGKEY", "CRIPTOGRAFIA"},
				keyTextPair{"IAMAVERYLOONGKEYINFACTLONGERTHANPAST", "CRIPTOGRAFIA"},
			},
		},
	}

	for i, example := range examples {
		color.Bold.Println(example.name)

		// Create alphabet
		alphabet := hcipher .NewAlphabet(example.alphabet)
		fmt.Printf("Alphabet: %s\n", alphabet)

		// Create cipher
		cipher, err := hcipher.NewCipher(alphabet)
		if err != nil {
			color.Comment.Printf("An error ocurred creating the cipher instance: ")
			fmt.Printf("%s\n", err)
			continue
		}

		// Encrypt and decrypt messages
		for j, s := range example.samples {
			if j != 0 {
				fmt.Println()
			}

			key, msg := s[0], s[1]

			// Encrypt
			cipherText, err := cipher.Encrypt(msg, key)
			if err != nil {
				color.Comment.Printf("\t%d) Failed to encrypt %q using %q: ", j+1, msg, key)
				fmt.Printf("%s\n", err)
				continue
			}

			// Decrypt
			plainText, err := cipher.Decrypt(cipherText, key)
			if err != nil {
				color.Comment.Printf("\t%d) Failed to decrypt %q using %q: ", j+1, cipherText, key)
				fmt.Printf("%s\n", err)
				continue
			}

			color.Success.Println("\tSUCCESS")
			fmt.Printf("\tE(msg:%q, key:%q) = %s\n", msg, key, cipherText)
			fmt.Printf("\tD(msg:%q, key:%q) = %s\n", cipherText, key, plainText)
		}

		// Add one blank line between example
		if i != 0 {
			fmt.Println()
		}
	}
}