/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:09:31
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:09:33
 */
package utils

import "regexp"

// 校验邮箱
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
