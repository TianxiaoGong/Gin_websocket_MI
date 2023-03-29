package main

import (
	app "Gin_WebSocket_IM/router"
	"Gin_WebSocket_IM/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := app.Router()
	r.Run(":8081")
}
