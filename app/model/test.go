package model

import (
	"gorm.io/gorm"
	"ldap-server/app/common"
	"ldap-server/app/pagination"
)

type Test struct {
	Id   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}

type TestQ struct {
	Test
	pagination.PaginationQ
	FromTime string `form:"from_time"` //搜索开始时间
	ToTime   string `form:"to_time"`   //搜索结束时候
}

func (test TestQ) Search() (list []Test, total int64, err error) {
	list = make([]Test, 0)
	tx := common.DB.Table("test").Find(&list)
	total, err = crudAll(&test.PaginationQ, tx, list)
	return
}

func crudAll(p *pagination.PaginationQ, queryTx *gorm.DB, list interface{}) (int64, error) {
	//1.默认参数
	if p.Size < 1 {
		p.Size = 5
	}
	if p.Page < 1 {
		p.Page = 1
	}

	//2.部搜索数量
	var total int64
	err := queryTx.Count(&total).Error
	if err != nil {
		return 0, err
	}
	offset := p.Size * (p.Page - 1)

	//3.偏移量的数据
	err = queryTx.Limit(p.Size).Offset(offset).Find(list).Error
	if err != nil {
		return 0, err
	}

	return total, err
}
