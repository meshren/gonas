package main

import (
	_ "gonas/config"
	"gonas/routes"
)

func main() {
	routes.Run()
}
