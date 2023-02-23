package main

import (
	"facade/compression"
	"facade/crypt"
	"fmt"
	"os"
)

/*
Provided a simple interface to read and write file,
and hiding the complex parts (encrypting/decrypting
and compressing/uncompressing) from the user.
*/
func main() {
	bits := []byte("0000000000000000111111110000000011111111")
	fmt.Println("wrote:", string(bits))

	writeFile("misc.dat", &bits)
	bits2 := readFile("misc.dat")

	fmt.Println("read: ", string(*bits2))
}

/*************************************************
 * A Facade used for writing and reading a file.
 *************************************************/

/*
The content is first encrypted and compressed
before writing to a file.
*/
func writeFile(filename string, data *[]byte) {
	// performs encryption
	var key byte = 55
	cryptor := crypt.Xor{}
	cipher := cryptor.Encrypt(key, data)

	// performs compression
	rle := compression.Rle{}
	zip := rle.Compress(cipher)

	os.WriteFile(filename, *zip, 0644)
}

/*
The content, read in from the file, is first
decrypted and decompressed before passing back
to the user
*/
func readFile(filename string) *[]byte {
	bytes, _ := os.ReadFile(filename)

	// performs decompression
	rle := compression.Rle{}
	unzip := rle.Uncompress(&bytes)

	// performs encryption
	var key byte = 55
	cryptor := crypt.Xor{}
	plain := cryptor.Decrypt(key, unzip)

	return plain
}
