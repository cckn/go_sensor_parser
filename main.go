package main

import (
	"fmt"
	"log"

	"github.com/enitt-dev/go-utils/convert"
	"github.com/tarm/serial"
)

func main() {
	config := &serial.Config{
		Name: "/dev/ttyUSB2",
		Baud: 115200,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(stream)
	}

	err = stream.Flush()
	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	var (
		validatorQ = make(chan []byte)
	)

	go recv(validatorQ, stream)

	for {
		select {
		case recvBytes := <-validatorQ:
			fmt.Println(convert.BytesToHexStrings(recvBytes))
			// case msg2 := <-validatorQ:
			//     fmt.Println("received", msg2)
		}
	}
}
