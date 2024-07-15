package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func binaryStringToFloat(binary string) float32 {
	bits, err := strconv.ParseUint(binary, 2, 32)
	if err != nil {
		return 0
	}

	var number uint32 = uint32(bits)

	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000"))
	return
}
