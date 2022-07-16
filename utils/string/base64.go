package string

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
)

func DecodeBase64(s string) ([]byte, error) {
	return ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(s)))
}
