package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type POWMessage struct {
	message int
	nonce   int
	zeroes  int
}

func work(message int, zeroes int) POWMessage {
	nonce := 0
	digest := sha256.Sum256([]byte(strconv.Itoa(message) + strconv.Itoa(nonce)))
	fmt.Println(string(digest[:32]))
	for {
		nonce = nonce + 1
		digest = sha256.Sum256([]byte(strconv.Itoa(message) + strconv.Itoa(nonce)))
		fmt.Println(string(digest[:32]))
		done := true
		count :=0
		for i := 0; i < zeroes; i++ {
			if digest[i] != 0 {
			count++
			}
		}
		if done {
			break
		}
	}
	return POWMessage{message, nonce, zeroes}
}

func verify(workedmessage POWMessage) bool {
	digest := sha256.Sum256([]byte(strconv.Itoa(workedmessage.message) + strconv.Itoa(workedmessage.nonce)))
	fmt.Println(digest)
	for i := 0; i < workedmessage.zeroes; i++ {
		if digest[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	message := 1000010
	workedmessage := work(message, 2)
	fmt.Println(workedmessage)
	fmt.Println(verify(workedmessage))
}
