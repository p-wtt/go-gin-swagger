package main

import (
	routes "github.com/p-wtt/go-gin-swagger/src/routes"
)


func main() {
	r := routes.Initialize()
	routes.Start(r)
}
