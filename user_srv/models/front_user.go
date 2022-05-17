/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:17:40
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:17:42
 */
package models

import "time"

type FrontUser struct {
	Id         int
	Email      string
	Password   string
	Desc       string
	Status     int
	CreateTime time.Time
}

func (FrontUser) TableName() string {
	return "front_user"
}
