package opsmantest

import (
	"bytes"
	"crypto/md5"
)

func CompareMd5(buffer *bytes.Buffer, b_array *[]byte) bool {
	slice1 := md5.Sum(*b_array)
	slice2 := md5.Sum(buffer.Bytes())
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
