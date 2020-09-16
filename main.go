package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/enitt-dev/go-utils/convert"
	"github.com/tarm/serial"
)

var sensorTypeMap = map[int]string{
	0x10: "CO2",
	0x11: "TVOC",
	0x20: "Humidity",
	0x30: "Temperature",
	0x31: "Temperature_object",
	0x40: "GyroX",
	0x41: "GyroY",
	0x42: "GyroZ",
}

type sensorData struct {
	DeviceId uint `json:"deviceId"`
	// data     map[string]float32
	SensorType string  `json:"sensor"`
	Value      float32 `json:"value"`
}

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

			log.Println()

			for i := 0; i < len(recvBytes); i = i + 10 {
				bytes := recvBytes[i : i+10]

				sd := sensorData{
					DeviceId:   uint(bytes[1])*10000 + uint(bytes[3]),
					SensorType: sensorTypeMap[int(bytes[4])],
					Value:      convert.BytesToFloat32(bytes[6:]),
				}
				// log.Println(sd)

				b, _ := json.Marshal(sd)
				j := string(append(b, '\n'))
				log.Print(j)
				writeToFile(j)
				// sensorDataArr = append(sensorDataArr, sd)
			}
			// log.Println(sensorDataArr)

		}
	}
}

func writeToFile(msg string) {
	const fileName = "/home/pi/sensorlog/sensorData6.log"
	//
	// // file, error := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0777)
	// // if error != nil {
	// //     panic(error)
	// // }
	// //
	// // defer file.Close()
	// err := ioutil.WriteFile(fileName, []byte(msg), os.FileMode(777))
	// // _, err := file.WriteString(msg)
	// if err != nil {
	//     log.Println(err)
	// }

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(msg); err != nil {
		panic(err)
	}

}
