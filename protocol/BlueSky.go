package protocol

import "github.com/albenik/bcd"

const BlueSkyPrecursor byte = 0xF5
const WriteUnitPriceCmd byte = 0xB2
const ControlCodeA5 byte = 0xA2

func XORChecker(d []byte) uint {
	const two int = 2
	const highestByte uint = 0xFF
	const lowestByte uint = 0x01

	if len(d) < two {
		return highestByte + lowestByte
	}
	x := d[0]
	for i := 1; i < len(d); i++ {
		x ^= d[i]
	}
	r := (x & 0x7F)
	return uint(r)
}

func CmdChangeNozzlePrize(nozID []byte, price uint32) (data []byte, size int, msg string) {
	const outputSize = 6

	z := []byte{BlueSkyPrecursor}
	z = append(z, nozID...)
	z = append(z, ControlCodeA5)
	p := bcd.FromUint32(price)[1:]
	z = append(z, p...)
	z = append(z, WriteUnitPriceCmd)
	check_code := XORChecker(z)
	z = append(z, byte(check_code))

	return z, outputSize, "Change Nozzle Price"
}
