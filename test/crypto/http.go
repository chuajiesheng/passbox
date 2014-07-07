package crypto

import (
	"fmt"
	"net/http"
	h "helper"
)

func init() {
	http.HandleFunc("/roadtest/crypto", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", "test1")
	test1(w, r)
	fmt.Fprintf(w, "\n%s\n", "test2")
	test2(w, r)
	fmt.Fprintf(w, "\n%s\n", "test3")
	test3(w, r)
	fmt.Fprintf(w, "\n%s\n", "test4")
	test4(w, r)
}

func test1(w http.ResponseWriter, r *http.Request) {
	key := []byte("a very very very very secret key") // 32 bytes
	plaintext := []byte("some really really really long plaintext")
	fmt.Fprintf(w, "%s\n", plaintext)
	ciphertext := h.Encrypt(key, plaintext)
	fmt.Fprintf(w, "%x\n", ciphertext)
	result, err := h.Decrypt(key, ciphertext)
	if err != nil {
		fmt.Fprintf(w, "%s\n", "decipher error")
	} else {
		fmt.Fprintf(w, "%s\n", result)
	}
}

func test2(w http.ResponseWriter, r *http.Request) {
	key := []byte("a very very very very secret key") // 32 bytes
	key2 := []byte("1 very very very very secret key") // 32 bytes
	plaintext := []byte("some really really really long plaintext")
	fmt.Fprintf(w, "%s\n", plaintext)
	ciphertext := h.Encrypt(key, plaintext)
	fmt.Fprintf(w, "%x\n", ciphertext)
	result, err := h.Decrypt(key2, ciphertext)
	if err != nil {
		fmt.Fprintf(w, "%s\n", "decipher error")
	} else {
		fmt.Fprintf(w, "%s\n", result)
	}
}

func test3(w http.ResponseWriter, r *http.Request) {
	key := []byte("a very very very secret key") // 32 bytes
	plaintext := []byte("some really really really long plaintext")
	fmt.Fprintf(w, "%s\n", plaintext)
	ciphertext := h.Encrypt(h.Pad(key), plaintext)
	fmt.Fprintf(w, "%x\n", ciphertext)
	result, err := h.Decrypt(h.Pad(key), ciphertext)
	if err != nil {
		fmt.Fprintf(w, "%s\n", "decipher error")
	} else {
		fmt.Fprintf(w, "%s\n", result)
	}
}

func test4(w http.ResponseWriter, r *http.Request) {
	c := 32
	key := h.GetRand(c)
	plaintext := []byte("some really really really long plaintext")
	fmt.Fprintf(w, "key: %s\n", key)
	fmt.Fprintf(w, "%s\n", plaintext)
	ciphertext := h.Encrypt(h.Pad(key), plaintext)
	fmt.Fprintf(w, "%x\n", ciphertext)
	result, err := h.Decrypt(h.Pad(key), ciphertext)
	if err != nil {
		fmt.Fprintf(w, "%s\n", "decipher error")
	} else {
		fmt.Fprintf(w, "%s\n", result)
	}
}
