package app

import (
	"github.com/2n1/proxysubscriber/app/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	man := e.Group("/man")
	man.Use(handler.H(handler.AuthMiddleWare))
	{
		man.GET("", handler.H(handler.Index))
		man.GET("/group/add", handler.H(handler.GroupInput))
		man.POST("/group/add", handler.H(handler.GroupInput))
		man.GET("/group/edit/:id", handler.H(handler.GroupInput))
		man.GET("/group/url/:id", handler.H(handler.GroupUrl))
		man.POST("/group/refresh-url/:id", handler.H(handler.RefreshGroupUrl))
		man.POST("/group/edit/:id", handler.H(handler.GroupInput))
		man.GET("/node/add", handler.H(handler.NodeInput))
		man.POST("/node/add", handler.H(handler.NodeInput))
		man.GET("/node/edit/:id", handler.H(handler.NodeInput))
		man.GET("/node/del/:id", handler.H(handler.NodeDel))
		man.POST("/node/edit/:id", handler.H(handler.NodeInput))
		man.GET("/group", handler.H(handler.Group))
		man.GET("/group/del/:id", handler.H(handler.GroupDel))
		man.GET("/node", handler.H(handler.Node))
		man.GET("/cfip", handler.H(handler.Cfips))
		man.POST("/cfip", handler.H(handler.Cfips))
		man.GET("/auth", handler.H(handler.UpdateAuth))
		man.POST("/auth", handler.H(handler.UpdateAuth))
		man.GET("/logout", handler.H(handler.Logout))
	}
	frontend := e.Group("/")
	{
		frontend.GET("", handler.H(handler.FrontendIndex))
		frontend.GET("s/:id", handler.H(handler.UrlHandler))
		frontend.GET("login", handler.H(handler.LoginHandler))
		frontend.POST("login", handler.H(handler.LoginHandler))
		frontend.POST("install", handler.H(handler.Install))
		frontend.GET("install", handler.H(handler.Install))
	}
}
