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
				keyTextPair{"FORTALEZA", "CONSUL"}, // N: 3
				keyTextPair{"FORTALEZA", "UUNAMFCIENCIASS"}, // N: 3
				keyTextPair{"IKEY", "CRIPTOGRAFIA"}, // N: 2
				keyTextPair{"IAMAVERYLOONGKEY", "CRIPTOGRAFIA"}, // N: 4
				keyTextPair{"NHWTTNRPOHVZOYRNDBMAXIJXLWVSMJOGKUSJ", "CRIPTOGRAFIA"}, // N: 6
				keyTextPair{"NHWTTNRPOHVZOYRNDBMAXIJXLWVSMJOGKUSJ", "CRIPTOGRAFIAYSEGURIDADESUNCURSOIMPARTIDOENLAFACULTADDECIENCIASUNAM"}, // N: 6
			},
		},
		{
			name: "Binary alphabet",
			alphabet: "01",
			samples: []keyTextPair{
				keyTextPair{"1011", "0110101100101101"},
				keyTextPair{"0000101111100011100111010", "11111"},
			},
		},
		{
			name: "ASCII alphabet with digits and punctuation",
			alphabet: `0123456789 :abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ()'`,
			samples: []keyTextPair{
				keyTextPair{`jb98v4 e0GgIJ5kKxVLtyRzdc`, "I'm a message that want's to be a key :)"}, // N: 5
				keyTextPair{`FORTAleza`, "CONSUL"},
				keyTextPair{
					`O3f0URuwcmMwyJ iO2Cc2zksDH1789(oAI1IelV5uLe'twQjELyh3FS1TyE'c5UCpAGTL2L3vl01wn5TBt66rpHBXGUeQ4yZ2C66`, // N: 10
					`Said you wouldnt be home late tonight I gave up waiting at seventeen past midnight Now my only company's a half full glass of wine`,
				},
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