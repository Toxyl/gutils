package gutils

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

func StringToSha256(s string) string {
	data := []byte(s)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash[:])
}

func StringToSha1(s string) string {
	data := []byte(s)
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x", hash[:])
}

func StringSliceToSha256(s []string) string {
	sort.Strings(s)
	return StringToSha256(strings.Join(s, ","))
}
