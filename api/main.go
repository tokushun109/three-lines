package main

import (
	"api/app/controllers"
	_ "api/app/db/connector"
)

func main() {
	controllers.StartMainServer()
}
