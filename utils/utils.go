package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// Min returns the minimum value of x and y
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// Max returns the maximum value of x and y
func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// Md5Hash returns a Md5 hash
func Md5Hash(value []byte) string {
	checksumData := md5.Sum(value)
	return hex.EncodeToString(checksumData[:])
}

// Random returns a random number in between min and max
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// StringInSlice will return true if string is in slice
func StringInSlice(s string, slice []string) bool {
	for _, elem := range slice {
		if s == elem {
			return true
		}
	}

	return false
}
