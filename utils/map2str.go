/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 23:18:34
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 23:18:37
 */
package utils

import "encoding/json"

func MapToStr(m map[string]interface{}) string {
	byte_data, _ := json.Marshal(m)

	str := string(byte_data)
	return str
}

func StrToMap(str string) map[string]interface{} {
	var map_data map[string]interface{}
	err := json.Unmarshal([]byte(str), &map_data)

	if err != nil {
		panic(err)
	}
	return map_data
}
