package controller

import (
	"demo/helper"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"column:name"`
}

func (User) TableName() string {
	return "tb_user"
}

func MysqlTest(ctx *gin.Context) {
	var users []User
	helper.DB.Model(&User{}).Find(&users)
	ctx.JSON(200, gin.H{
		"code":   0,
		"msg":    "succ",
		"result": users,
	})
}
