package stringutils

import (
	"unicode/utf8"
	"crypto/sha1"
	"encoding/hex"
	"crypto/sha256"
	"crypto/md5"
	"crypto/sha512"
	"hash/crc32"
	"hash/crc64"
	"strconv"
)

var crc32Table = crc32.MakeTable(0xD5828281)
var crc64Table = crc64.MakeTable(0xC96C5795D7870F42)

func Reverse_string (s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func MD5_hash (bv []byte) string {
	md5HashInBytes := md5.Sum(bv)
	return hex.EncodeToString(md5HashInBytes[:])
}

func Sha1_hash (bv []byte) string {
	hasher := sha1.New()
	hasher.Write(bv)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha224_hash (bv []byte) string {
	res := sha256.Sum224(bv)
	return hex.EncodeToString(res[:])
}

func Sha256_hash (bv []byte) string {
	res := sha256.Sum256(bv)
	return hex.EncodeToString(res[:])
}

func Sha512_hash (bv []byte) string {
	res := sha512.Sum512(bv)
	return hex.EncodeToString(res[:])
}

func Crc32 (bv []byte) string {
	res := crc32.Checksum(bv, crc32Table)
	return strconv.FormatUint(uint64(res), 16)
}

func Crc64 (bv []byte) string {
	res := crc64.Checksum(bv, crc64Table)
	return strconv.FormatUint(res, 16)
}
