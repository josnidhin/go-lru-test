/**
 * @author Jose Nidhin
 */
package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const CACHE_SIZE = 10_000

var secret []byte

func init() {
	secret = make([]byte, 256)
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}
}

func main() {
	for i := 0; i < CACHE_SIZE; i++ {
		key := uuid.New().String()
		rawVal := uuid.New()

		mac := hmac.New(sha256.New, secret)
		rawValByte, _ := rawVal.MarshalText()
		mac.Write(rawValByte)

		val := mac.Sum(nil)

		fmt.Printf("Key: %s\tValue: %x\n", key, val)
	}

	fmt.Println("Sleep for 10s")
	time.Sleep(10 * time.Second)
}
