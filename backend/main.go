package main

import (
	_ "gonas/config"
	"gonas/routes"
)

func main() {
	// todo 检测必要端口是否可用 redis mysql ...
	routes.Run()
}
