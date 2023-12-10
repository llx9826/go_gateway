package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/llx9826/gin_gateway/public"
	"time"
)

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"管理员用户名" example:"admin" validate:"valid_username"` //管理员用户名
	PassWord string `json:"passWord" form:"password" comment:"密码" example:"123456" validate:"required"`
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"123456" validate:""`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}
