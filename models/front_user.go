/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 15:43:20
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 15:43:21
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
