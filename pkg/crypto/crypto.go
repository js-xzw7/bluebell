package crypto

import (
	"bluebell/settings"
	"crypto/md5"
	"encoding/hex"
)

func Encrypt(data string) string {
	h := md5.New()
	h.Write([]byte(settings.Conf.Secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
