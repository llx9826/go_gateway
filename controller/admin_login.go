package controller

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/llx9826/gin_gateway/dao"
	"github.com/llx9826/gin_gateway/dto"
	"github.com/llx9826/gin_gateway/middleware"
	"github.com/llx9826/gin_gateway/public"
	"time"
)

type AdminLoginController struct{}

func RegisterAdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.POST("/logout", adminLogin.AdminLoginOut)

}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 用户
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body dto.AdminLoginInput  true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	err := params.BindValidParam(c)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	//1.验证用户人账号密码
	admin := dao.Admin{}
	adminInfo, err := admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//2.设置session
	sessionInfo := &dto.AdminSessionInfo{
		ID:        adminInfo.Id,
		UserName:  adminInfo.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessionInfo)
	if err != nil {
		return
	}
	session := sessions.Default(c)
	session.Set(public.AdminSessionInfoKey, string(sessBts))
	err = session.Save()
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// AdminLoginOut /godoc
//
func (adminLogin *AdminLoginController) AdminLoginOut(c *gin.Context) {
	middleware.ResponseSuccess(c, nil)
}
