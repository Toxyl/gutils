package gutils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"strings"
)

func EncodeBase64String(src string) string {
	res := base64.StdEncoding.EncodeToString([]byte(strings.TrimSpace(src)))
	return res
}

func DecodeBase64String(src string) (string, error) {
	srcdec, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	res := strings.TrimSpace(string(srcdec))
	return res, nil
}

func EncodeGzBase64String(src string) string {
	src = strings.TrimSpace(src)
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(src)); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	return EncodeBase64String(b.String())
}

func DecodeGzBase64String(src string) (string, error) {
	src = strings.TrimSpace(src)
	data, err := DecodeBase64String(src)
	if err != nil {
		return "", err
	}
	rdata := bytes.NewReader([]byte(data))
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
