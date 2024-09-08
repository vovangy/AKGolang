package main

import (
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HashFunc func(data string) uint64

type HashMap struct {
	data     map[uint64]interface{}
	hashFunc HashFunc
}

func NewHashMap(options ...func(*HashMap)) *HashMap {
	hm := &HashMap{
		data: make(map[uint64]interface{}),
		hashFunc: func(data string) uint64 {
			h := fnv.New64()
			h.Write([]byte(data))
			return h.Sum64()
		},
	}

	for _, option := range options {
		option(hm)
	}

	return hm
}

func (h *HashMap) Set(key string, value interface{}) {
	hash := h.hashFunc(key)
	h.data[hash] = value
}

func (h *HashMap) Get(key string) (interface{}, bool) {
	hash := h.hashFunc(key)
	value, ok := h.data[hash]
	return value, ok
}

func WithHashCRC64() func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = func(data string) uint64 {
			return crc64.Checksum([]byte(data), crc64.MakeTable(crc64.ECMA))
		}
	}
}

func WithHashCRC32() func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = func(data string) uint64 {
			return uint64(crc32.ChecksumIEEE([]byte(data)))
		}
	}
}

func WithHashCRC16() func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = func(data string) uint64 {
			var crc16 uint16 = 0xFFFF
			for i := 0; i < len(data); i++ {
				crc16 = crc16 ^ uint16(data[i])
				for j := 0; j < 8; j++ {
					if (crc16 & 0x0001) != 0 {
						crc16 = (crc16 >> 1) ^ 0xA001
					} else {
						crc16 >>= 1
					}
				}
			}
			return uint64(crc16)
		}
	}
}

func WithHashCRC8() func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = func(data string) uint64 {
			var crc8 uint8 = 0
			for i := 0; i < len(data); i++ {
				crc8 ^= uint8(data[i])
				for j := 0; j < 8; j++ {
					if crc8&0x80 != 0 {
						crc8 = (crc8 << 1) ^ 0x07
					} else {
						crc8 <<= 1
					}
				}
			}
			return uint64(crc8)
		}
	}
}
