package main

import (
	"atm/global"
	"atm/initializer"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := initializer.Init()
	engine.(*gin.Engine).Run(fmt.Sprintf(":%v", global.Conf.ServerConf.Port))
}
