/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 16:19:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 16:19:56
 */
package utils

import (
	"fmt"
	"time"
)

func AddHour(h int) string {
	// h = 24
	now_time := time.Now()
	after_hour := fmt.Sprintf("+%dh", h)
	duration_hour, _ := time.ParseDuration(after_hour)
	ret_fmt_time := now_time.Add(duration_hour).Format("2006-01-02 15:04:05")
	return ret_fmt_time
}
