package main

import (
	"gonovel/model"
	"gonovel/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
