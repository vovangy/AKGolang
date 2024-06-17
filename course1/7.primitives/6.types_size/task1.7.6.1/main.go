package main

import (
	"fmt"
	"unsafe"
)

func sizeOfBool(b bool) int {
	return int(unsafe.Sizeof(b))
}

func sizeOfInt(n int) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt8(n int8) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt16(n int16) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt32(n int32) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt64(n int64) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint(n uint) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint8(n uint8) int {
	return int(unsafe.Sizeof(n))
}

func main() {
	fmt.Println("bool size: ", sizeOfBool(true))
	fmt.Println("int size: ", sizeOfInt(1))
	fmt.Println("int8 size: ", sizeOfInt8(1))
	fmt.Println("int16 size: ", sizeOfInt16(1))
	fmt.Println("int32 size: ", sizeOfInt32(1))
	fmt.Println("int64 size: ", sizeOfInt64(1))
	fmt.Println("uint size: ", sizeOfUint(1))
	fmt.Println("uint8 size: ", sizeOfUint8(1))
	return
}
