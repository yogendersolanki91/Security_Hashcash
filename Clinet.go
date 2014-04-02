package main

import (
	"net"
	"fmt"
	"encoding/gob"
	"time"
	"flag"
	"strconv"
	"os"
	"crypto/sha256"
)
var id = flag.Int("id", 0, "id of the client")
var attack = flag.Bool("attack", false, "Whether to attack or not")
var proof=flag.Bool("proof", false, "Whether to use proof of work or not")
func work(message POWMessage, zeroes int) (POWMessage){
	nonce := 0
	digest := sha256.Sum256([]byte(message.Message + strconv.Itoa(nonce)))
	for true {

		nonce = nonce + 1
		digest = sha256.Sum256([]byte(message.Message + strconv.Itoa(nonce)))

		done := true
		for i := 0; i < zeroes; i++ {
			if digest[i] != 0 {
				done = false
			}
		}
		if done {
			break
		}
	}
	fmt.Println(nonce)
	message.Nonce=nonce;
	message.Zeroes=zeroes;
	return message
}
type POWMessage struct {
	Message string
	Nonce   int
	Zeroes 	int
}

func main() {
	flag.Parse()
	if *id == 0 {
		fmt.Println("tell me my id")
		os.Exit(2)
	}
	// connect to the server



	i:=0
	for {
		i++
		var msg POWMessage
		// send the message
		msg.Message ="from:"+ strconv.Itoa(*id)+"-"+"#"+strconv.Itoa(i)+" "+time.Now().String()
		fmt.Println(msg.Message)
		c, errr := net.Dial("tcp", "127.0.0.1:9999")
		if errr != nil {
			fmt.Println(errr)
			return
		}
		if *proof{
			msg=work(msg,2)
			fmt.Println("doi")
		}
		err := gob.NewEncoder(c).Encode(&msg)
		//fmt.Println("sending")


		if !*attack{
			time.Sleep(500*time.Millisecond);
		}

		if err != nil {
			fmt.Println(err)
		}
		c.Close()
	}


}


