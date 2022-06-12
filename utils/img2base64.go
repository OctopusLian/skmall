/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 23:18:04
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 23:18:08
 */
package utils

import (
	"encoding/base64"
	"os"
)

func Img2Base64(imgPath string) string {
	file, _ := os.Open(imgPath)

	defer file.Close()

	bufByte := make([]byte, 100000)
	n, _ := file.Read(bufByte)

	imgBase64Str := base64.StdEncoding.EncodeToString(bufByte[:n])

	return imgBase64Str

}
