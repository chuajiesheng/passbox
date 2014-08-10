package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

const HashSize = sha256.Size

func GetRand(len int) []byte {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return b
}

func Pad(key []byte) []byte {
	var b []byte
	b = make([]byte, 32, 32)
	for i := 0; i < 32; i++ {
		b[i] = key[i % len(key)]
	}
	return b
}

func Encrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext
}

func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		err = fmt.Errorf("decrypt: %s", "ciphertext too short")
		return nil, err
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	res, err := decodeBase64(string(text))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Sum256(data []byte) [sha256.Size]byte {
	return sha256.Sum256(data)
}

func GenerateMAC(key []byte, in []byte) []byte {
	hash := sha256.New
	mac := hmac.New(hash, key)
	n, err := mac.Write(in)
	if n != len(in) || err != nil {
		panic(err)
	}
	return mac.Sum(nil)
}

// CheckMAC returns true if messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func GenEncryptMAC(length int, key []byte) (plaintext []byte, ciphertext []byte, mac []byte) {
	plaintext = GetRand(length)
	ciphertext = Encrypt(Pad(key), plaintext)
	mac = GenerateMAC(key, plaintext)
	return
}

func DecryptVerify(key []byte, ciphertext []byte, mac []byte) (plaintext []byte, err error, verified bool) {
	plaintext, err = Decrypt(Pad(key), ciphertext)
	verified = CheckMAC(plaintext, mac, key)
	return
}
