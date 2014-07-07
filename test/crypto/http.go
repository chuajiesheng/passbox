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
	plaintext := []byte("some really really long text wor")
	fmt.Fprintf(w, "key:\t\t %s [%d]\n", key, len(key))
	fmt.Fprintf(w, "plaintext:\t %s [%d]\n", plaintext, len(plaintext))
	ciphertext := h.Encrypt(h.Pad(key), plaintext)
	fmt.Fprintf(w, "ciphertext:\t %x [%d]\n", ciphertext, len(ciphertext))
	result, err := h.Decrypt(h.Pad(key), ciphertext)
	if err != nil {
		fmt.Fprintf(w, "decipher:\t %s\n", "decipher error")
	} else {
		fmt.Fprintf(w, "decipher:\t %s [%d]\n", result, len(result))
	}

	mac := h.GenerateMAC(key, plaintext)
	fmt.Fprintf(w, "mac:\t\t %x [%d]\n", mac, len(mac))
	fmt.Fprintf(w, "check:\t\t %s\n", h.CheckMAC(plaintext, mac, key))
}
