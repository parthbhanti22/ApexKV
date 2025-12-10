package main

import (
	"encoding/binary"
	"time"
)

//HeaderSize = 4bytes(key)+4bytes(value size)+8bytes(timestamp)

const HeaderSize = 16

type Entry struct {
	Key string
	Value string
	Timestamp time.Time
}
//take a value/key and return byte slice for disk
func Encode(key,value string) []byte {
	timestamp := uint(time.Now().UnixNano())
	keyLen := len(key)
	valueLen := len(value)


	totalSize := HeaderSize+keyLen+valueLen
	buf := make([]byte,totalSize)

	//Header : Using littleEndian standard
	binary.LittleEndian.PutUint32(buf, uint32(keyLen))
	binary.LittleEndian.PutUint32(buf[4:], uint32(valueLen))
	binary.LittleEndian.PutUint64(buf[8:], uint64(timestamp))

	//now we will write data into the slice
	copy(buf[16:], key)
	copy(buf[16+keyLen:], value)

	return buf
}

//decoding function: takes a raw byte and returns the Entry
func Decode(buf []byte) Entry {
	keyLen := binary.LittleEndian.Uint32(buf[0:4])
	valueLen := binary.LittleEndian.Uint32(buf[4:8])
	timestamp := binary.LittleEndian.Uint64(buf[8:16])

    // 1. Slice the buffer to get the raw bytes, then convert to string
	key := string(buf[16 : 16+keyLen])
	value := string(buf[16+keyLen : 16+keyLen+valueLen])

	return Entry{
		Key:       key,
		Value:     value,
        // 2. Convert the nanoseconds integers back to a time.Time object
		Timestamp: time.Unix(0, int64(timestamp)),
	}
}
