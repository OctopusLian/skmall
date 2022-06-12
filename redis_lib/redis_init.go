/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 17:48:11
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 17:48:22
 */
package redis_lib

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var Conn redis.Conn
var Err error

func init() {
	Conn, Err = redis.Dial("tcp", "127.0.0.1:6379")
	if Err != nil {
		fmt.Println(Err)
		panic(Err)
	}
}
