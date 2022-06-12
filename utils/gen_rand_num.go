/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 15:45:03
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 15:45:04
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
