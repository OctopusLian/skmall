/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-11 23:27:08
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-11 23:27:09
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
