/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:19:13
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:19:17
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5pwd(pwd string) string {
	h := md5.New()
	md5code := "12ffjkjru98hsjq%*"
	h.Write([]byte(pwd + md5code))
	return hex.EncodeToString(h.Sum(nil))
}
