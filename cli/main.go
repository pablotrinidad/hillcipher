package main

import (
	"flag"
	"fmt"
	"os"

	hcipher "github.com/pablotrinidad/hillcipher/cipher"
)

// mode controls the CLI operation mode, either encryption or decryption.
type mode uint8

const (
	modeUnknown mode = 0
	modeEncrypt mode = 1
	modeDecrypt mode = 2
)

var (
	text, key, alphabet string
	excMode             mode
	validModes          = map[string]mode{
		"e": modeEncrypt, "encrypt": modeEncrypt,
		"d": modeDecrypt, "decrypt": modeDecrypt,
	}
)

func init() {
	flag.StringVar(&text, "t", "", "the text that will be used in the cipher")
	flag.StringVar(&key, "k", "", "the key that will be used in the cipher")
	flag.StringVar(&alphabet, "a", "", "the alphabet that will be used in the cipher")
	flagMode := flag.String("m", "", "the cipher mode, either 'encrypt'/'e' or 'decrypt'/'d'")

	flag.Parse()

	flagsSet := true
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			flagsSet = false
			fmt.Fprintf(os.Stderr, "missing required -%s argument (%s)\n", f.Name, f.Usage)
		}
	})
	if !flagsSet {
		os.Exit(2)
	}

	if _, found := validModes[*flagMode]; !found {
		fmt.Fprintf(os.Stderr, "got invalid cipher mode %s", flagMode)
	}
	excMode = validModes[*flagMode]
}

func main() {
	cipher, err := hcipher.NewCipher(hcipher.NewAlphabet(alphabet))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var op func(string, string) (string, error)
	switch excMode {
	case modeEncrypt:
		op = cipher.Encrypt
	case modeDecrypt:
		op = cipher.Decrypt
	default:
		// This is impossible since flags are parsed at the begining
		fmt.Fprintf(os.Stderr, "got invalid execution mode %v\n", excMode)
		os.Exit(1)
	}

	result, err := op(text, key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred during cipher execution\n%v", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, result)
}
