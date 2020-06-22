package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// request key
	var requestKey string = "1234privy5678"
	// request plaint text
	var requestPlaintext string = "Selamat Datang"

	// mengubah paksa panjangan key menjadi 32 menggunakan sha256
	// (32 byte key for AES-256)
	// dari hasil sha kita ubah ke byte
	sum := sha256.Sum256([]byte(requestKey))
	bytes := []byte(string(sum[:]))

	// cek apakah panjang key sudah sesuai pada data hasil sha256 dalam bentuk byte
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	// hasil key yg sudah di hash sha256 dalam byte kita ubah ke string untuk melihat hasilnya
	// dan digunakan untuk proses selanjutnya
	key := hex.EncodeToString(bytes)
	fmt.Printf("kunci encrypt/decrypt '" + requestKey + "' untuk encrypt/decrypt kalimat '" + requestPlaintext + "'\n")

	// process enkripsi aes256-gcm
	encrypted := encrypt(requestPlaintext, key)
	fmt.Printf("encrypted : %s\n", encrypted)

	// process dekripsi aes256-gcm
	decrypted := decrypt(encrypted, key)
	fmt.Printf("decrypted : %s\n", decrypted)
}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	// sebelum digunakan untuk process kita ubah string ke byte
	// referensi tutorial
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	// membuat chiper block dari key
	// referensi tutorial
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// buat gcm dari block
	// referensi tutorial
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// buat nonce. nonce harus dari GCM
	// referensi tutorial
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypt data memakai aesGCM.Seal
	// Karena tidak ingin menyimpan nonce di tempat lain dalam kasus ini,
	// menambahkannya sebagai awalan ke data yang dienkripsi. Argumen nonce pertama di Seal adalah awalan.
	// referensi tutorial
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	// kembalikan hasil enkrip
	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(encryptedString string, keyString string) (decryptedString string) {

	// sebelum digunakan untuk process kita ubah string ke byte
	// referensi tutorial
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	// membuat chiper block dari key
	// referensi tutorial
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// buat gcm dari block
	// referensi tutorial
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// ambil ukuran size nonce
	// referensi tutorial
	nonceSize := aesGCM.NonceSize()

	// Ekstrak nonce dari data terenkripsi
	// referensi tutorial
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// dekrip menjadi plaintext memakai gcm
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	// kembalikan hasil dekrip
	return fmt.Sprintf("%s", plaintext)
}
