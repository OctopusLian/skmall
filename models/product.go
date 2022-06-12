/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 15:43:48
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 18:34:00
 */
package models

import "time"

type Products struct {
	Id         int
	Name       string
	Pname      string
	Price      float32 `gorm:"type:decimal(11,2)"`
	Num        int
	Unit       string
	Pic        string
	Desc       string
	CreateTime time.Time

	SecKills []SecKills `gorm:"ForeignKey:PId;AssiciationForeignKey:Id"`
}

func (Products) TableName() string {
	return "products"
}
