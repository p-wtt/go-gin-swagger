package main

import (
	routes "github.com/p-wtt/go-gin-swagger/routesV1"
)

func main() {
	r := routes.NewRoute()
	routes.Start(r)
}
