package main

import "github.com/tarm/serial"

const (
	BufferSize = 512

	Stx = 0xea
	Etx = 0xee
)

func recv(validatorQ chan []byte, s *serial.Port) {
	var tmp = make([]byte, BufferSize)
	var buf = make([]byte, BufferSize)
	for {
		tmpSize, _ := s.Read(tmp)
		// fmt.Println(convert.BytesToHexStrings(buf[:n]))

		if len(buf) == 0 {
			if tmp[0] == Stx {

			}
			//
			// if buf[0] == Stx && buf[len(buf)-1] == Etx {
			//
			// }
		}

		validatorQ <- tmp[:tmpSize]

		buf = []byte{}
	}
}
