package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Encode(msg string,salt string) string{
	new_msg:=msg+salt
	encoder:=md5.New()
	encoder.Write([]byte(new_msg))
	tmp:=encoder.Sum(nil)
	return hex.EncodeToString(tmp)
}

func Verify(msg string,salt string,crypto_msg string) bool{
	return Encode(msg,salt)==crypto_msg
}




