package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/llx9826/gin_gateway/dto"
	"github.com/llx9826/gin_gateway/middleware"
)

type AdminLoginController struct{}

func RegisterAdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.POST("/logout", adminLogin.AdminLoginOut)

}

func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	err := params.BindValidParam(c)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	middleware.ResponseSuccess(c, nil)
}

// AdminLoginOut /godoc
//
func (adminLogin *AdminLoginController) AdminLoginOut(c *gin.Context) {
	middleware.ResponseSuccess(c, nil)
}
