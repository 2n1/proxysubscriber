package main

import (
	"log"
	"net/http"

	"github.com/2n1/proxysubscriber/app"
	"github.com/2n1/proxysubscriber/app/cfg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	if err := cfg.Init(); err != nil {
		log.Fatal("Init config failed:", err)
	}
}
func main() {
	gin.SetMode(cfg.Cfg.Mode)
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.StaticFS("/static", http.Dir("./static"))
	engine.LoadHTMLGlob("./templates/*")

	store := cookie.NewStore([]byte(cfg.Cfg.Session.SecretKey))
	engine.Use(sessions.Sessions(cfg.Cfg.Session.Name, store))

	app.InitRouter(engine)
	log.Fatalln(engine.Run(cfg.Cfg.Addr))
}
