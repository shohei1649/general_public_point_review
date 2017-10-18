package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

func ToMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ToJSONIndent(jsonBytes []byte) string {

	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "  ")
	return out.String()
}

func EncodeBase64(rowStr string) string {

	return base64.StdEncoding.EncodeToString([]byte(rowStr))
}

func DecodeBase64(base64Str string) string {

	rowStr, _ := base64.StdEncoding.DecodeString(base64Str)
	return string(rowStr)
}
