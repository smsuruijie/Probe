package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xos/probe/model"
	"github.com/xos/probe/pkg/mygin"
	"github.com/xos/probe/service/singleton"
)

type guestPage struct {
	r *gin.Engine
}

func (gp *guestPage) serve() {
	gr := gp.r.Group("")
	gr.Use(mygin.Authorize(mygin.AuthorizeOption{
		Guest:    true,
		IsPage:   true,
		Msg:      "您已登录",
		Btn:      "返回首页",
		Redirect: "/",
	}))

	gr.GET("/login", gp.login)

	oauth := &oauth2controller{
		r: gr,
	}
	oauth.serve()
}

func (gp *guestPage) login(c *gin.Context) {
	LoginType := "GitHub"
	RegistrationLink := "https://github.com/join"
	if singleton.Conf.Oauth2.Type == model.ConfigTypeGitee {
		LoginType = "Gitee"
		RegistrationLink = "https://gitee.com/signup"
	}
	c.HTML(http.StatusOK, "dashboard/login", mygin.CommonEnvironment(c, gin.H{
		"Title":            "登录",
		"LoginType":        LoginType,
		"RegistrationLink": RegistrationLink,
	}))
}
