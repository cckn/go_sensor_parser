package main

import (
	"github.com/tarm/serial"
)

const (
	BufferSize = 512

	Stx = 0xea
	Etx = 0xee
)

func recv(validatorQ chan []byte, s *serial.Port) {
	var temp = make([]byte, BufferSize)
	var buff = make([]byte, BufferSize)

	var bufferInit = func() { buff = []byte{} }

	for {
		n, _ := s.Read(temp)

		temp = temp[:n]

		if len(buff) == 0 {
			if temp[0] == Stx {
				buff = temp
			}
		} else {
			if buff[0] != Stx {
				bufferInit()
			}
			buff = append(buff, temp...)
		}

		if len(buff) != 0 {
			if buff[0] == Stx && buff[len(buff)-1] == Etx {
				length := len(buff) - 7
				expectedLength := int(buff[5])

				if expectedLength == len(buff)-7 {
					validatorQ <- buff[6 : len(buff)-1]
					bufferInit()

				} else if length > expectedLength {
					bufferInit()
				}
				bufferInit()
			}
		} else if len(buff) > 100 {
			buff = []byte{}

		}
	}
}
