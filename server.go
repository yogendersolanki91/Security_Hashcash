package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
	"flag"
	"crypto/sha256"
	"strconv"
)
var useproof = flag.Bool("proof", false, "use proof of work to defend agaist Spam")
type POWMessage struct {
	Message string
	Nonce   int
	Zeroes int
}
func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}
func verify(workedmessage POWMessage) (bool) {
	digest := sha256.Sum256([]byte(workedmessage.Message + strconv.Itoa(workedmessage.Nonce)))
	// fmt.Println(digest)
	for i := 0; i < workedmessage.Zeroes; i++ {
		if digest[i] != 0 {
			return false
		}
	}

	return true
}
func handleServerConnection(c net.Conn) {
	// receive the message
	var msg POWMessage

	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	}
		if *useproof {
			fmt.Println("CHCK")
			msg.Zeroes=2
			if verify(msg){
				fmt.Println("OK")
				time.Sleep(time.Millisecond * 300)
			}else{
				fmt.Println("Rejected")
			}
		} else {
			time.Sleep(time.Millisecond * 300)
			fmt.Println("Received :", msg.Message, " at", time.Now())
		}
		c.Close()

}
func main(){

	flag.Parse()
	fmt.Println(*useproof)
	server()
	time.Sleep(time.Hour)
}
