/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:17:58
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:18:00
 */
package models

import "time"

type AdminUser struct {
	Id         int
	UserName   string
	Password   string
	Desc       string
	Status     int
	CreateTime time.Time
}

func (AdminUser) TableName() string {
	return "admin_users"
}
