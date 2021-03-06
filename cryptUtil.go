package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// TripleDesEncrypt 3DS???
func TripleDesEncrypt(origData []byte) ([]byte, error) {
	key := []byte("sr$*)(ruan$@lx100$#365#$")
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte("01234567"))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//PKCS5Padding ??
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
