package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ldap-server/app/model"
	"ldap-server/app/pagination"
)

func GetList(c *gin.Context) {
	query := &model.TestQ{}
	err := c.ShouldBindQuery(query) //开始绑定url-query 参数到结构体
	if err != nil {
		fmt.Println(err)
		return
	}
	list, total, err := query.Search() //开始mysql 业务搜索查询
	if err != nil {
		fmt.Println(err)
		return
	}
	//返回数据开始拼装分页json
	jsonPagination(c, list, total, &query.PaginationQ)
}

// 4.json 分页数据
func jsonPagination(c *gin.Context, list interface{}, total int64, query *pagination.PaginationQ) {
	c.AbortWithStatusJSON(200, gin.H{
		"ok":    true,
		"data":  list,
		"total": total,
		"page":  query.Page,
		"size":  query.Size,
	})
}
