# Hill Cipher

Hill Cipher algorithm implementation in Golang. Hill Cipher
is a polygraphic substitution cipher that encrypts messages through
matrix transformations.

* Matrix operations and modular arithmetic
definitions are implemented and exported as part of the package.
* Hill Cipher is broken by cipher-text only attacks.
* DON'T use this package for any purpose other than experimentation and educational purposes.

Auto-generated docs [**here**](https://godoc.org/github.com/pablotrinidad/hillcipher/cipher).

## About this repo

This repo contains 3 Go packages:
* `cipher` which contains the cipher implementation
* `cli` which is a command line interface that uses de `cipher` package
* `examples` which just prints the result of encrypting and decrypting some pre-defined messages using the `cipher` package.

Coverage of package `cipher` is **100%**.

## Using the `cipher` package

Please refer to the [**docs**](https://godoc.org/github.com/pablotrinidad/hillcipher/cipher), and **DON'T** use this in any production code. Here's an example nonetheless:

```go
// Initialize cipher
alp := cipher.NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
cip, err := cipher.NewCipher(alph)
if err != nil {
  return err
}

// Encrypt and decrypt messages
key := "IKEY"
msg := "CRIPTOGRAFIA"
cipherText, err := cip.Encrypt(msg, key)
...
plainText, err := cip.Decrypt(cipherText, key)
```

Please note that key must be invertible modulo size of alphabet. See `examples` and unit tests for more information.

## Using the CLI

Run: `$ go run main.go -m MODE -a ALPHABET -t TEXT -k KEY` where mode is either `e` or `d` for encryption and decryption respectively.

## Running examples

Run `$ go run main.go`

## LICENSE

Read [**LICENSE**](LICENSE.md)
