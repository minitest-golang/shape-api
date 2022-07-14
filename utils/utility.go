package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

var (
	AppMode         = ""
	EnableCORS      = false
	TokenExpireTime = int64(6 * 60)
	JwtKey          = "JwtKey"
)

func PrintJson(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err == nil {
		InfoLog("%s\n", b)
	} else {
		ErrorLog("Bad Json Format!\n")
	}
}

func Md5(s string) string {
	x := md5.Sum([]byte(s))
	return hex.EncodeToString(x[:])
}
