package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	Server *gin.Engine
}

func New() *App {
	a := App{Server: gin.New()}

	a.Server.Use(gin.Recovery())
	a.Server.Use(gin.Logger())

	a.Server.GET("/hello", HelloFunc)

	return &a
}

func (s *App) Run(port string) {
	err := s.Server.Run(port)
	fmt.Print(err)
}

type Hello struct {
	Msg string `json:"msg"`
}

func HelloFunc(c *gin.Context) {
	res := Hello{Msg: "Hi"}
	fmt.Print(res)
	c.JSON(200, res)
}
