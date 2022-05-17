/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:18:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:19:01
 */
package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenRandNum(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
