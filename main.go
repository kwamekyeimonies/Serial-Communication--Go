package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {

	//nozzleID := []byte{0x01}
	//_, _, _ := protocol.CmdChangeNozzlePrize(nozzleID, 0x00000011)
	// xor_output_cheker := protocol.XORChecker(nozzleID)

	// fmt.Println(func_test)
	// fmt.Println(dex)
	// fmt.Println(kwame)
	// fmt.Println(xor_output_cheker)

	for {
		c := &serial.Config{
			Name:        "ttyUSB0",
			Baud:        9600,
			StopBits:    1,
			Parity:      'E',
			ReadTimeout: time.Millisecond * 200,
		}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}

		n, err := s.Write([]byte{0xF5, 0x01, 0xA2, 0xD5, 0x03})
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])
	}
}
