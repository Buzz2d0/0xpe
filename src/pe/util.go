package pe

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ToC8bytes(s string) [8]byte {
	r := [8]byte{}
	off := len(s)
	if off >= 7 {
		off = 7
	}
	for i := 0; i < 8; i++ {
		// 填充
		if i < off {
			r[i] = s[i]
		} else {
			// 补0
			r[i] = 0x00
		}
	}
	return r
}

// Align 对齐排列
func Align(idx uint, aligment uint) uint {
	return (idx + aligment) & (^(aligment - 1))
}

// FillZeroByte 填充 0x00
func FillZeroByte(len int) []byte {
	return make([]byte, len, len)
}

func StrConv2Bytes(s string) []byte {
	r := []byte(s)
	return append(r, 0x00)
}

func GetBinaryBytes(order binary.ByteOrder, data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, order, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MustGetBinaryBytes(order binary.ByteOrder, data interface{}) []byte {
	if raw, err := GetBinaryBytes(order, data); err != nil {
		log.Fatalln(err)
	} else {
		return raw
	}
	return nil
}
