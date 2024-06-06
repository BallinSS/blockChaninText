package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 将int64转换为字节数组
func IntToHex(num int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
