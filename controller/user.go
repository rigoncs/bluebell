package controller

import (
	"bluebell/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1. 获取参数和参数校验
	p := new(models.ParamSignup)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Signup with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//2. 业务处理
	if len(p.Username) == 0 || len(p.Password) == 0 || p.Password != p.RePassword {
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	fmt.Println(p)
	//3. 返回响应
}
