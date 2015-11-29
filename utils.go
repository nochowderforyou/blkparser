package blkparser

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

// Get the Tx count, decode the variable length integer
// https://en.bitcoin.it/wiki/Protocol_specification#Variable_length_integer
func DecodeVariableLengthInteger(raw []byte) (cnt int, cnt_size int) {
	if raw[0] < 0xfd {
		return int(raw[0]), 1
	}
	cnt_size = 1 + (2 << (2 - (0xff - raw[0])))
	if len(raw) < 1+cnt_size {
		return
	}

	res := uint64(0)
	for i := 1; i < cnt_size; i++ {
		res |= (uint64(raw[i]) << uint64(8*(i-1)))
	}

	cnt = int(res)
	return
}

// GetShaString returns the hex-encoded sha256d hash of a byte slice.
func GetShaString(data []byte) (res string) {
	sha := sha256.New()
	sha.Write(data[:])
	tmp := sha.Sum(nil)
	sha.Reset()
	sha.Write(tmp)
	hash := sha.Sum(nil)
	res = HashString(hash)
	return
}

// GetScryptString returns the hex-encoded scrypt hash of a byte slice.
func GetScryptString(data []byte) (res string, err error) {
	hash, err := scrypt.Key(data, data, 1024, 1, 1, 32)
	if err != nil {
		return "", err
	}
	res = HashString(hash)
	return
}

// HashString returns the little-endian hex-encoded format of a 32 byte hash.
func HashString(data []byte) (res string) {
	for i := 0; i < 32; i++ {
		res += fmt.Sprintf("%02x", data[31-i])
	}
	return
}
