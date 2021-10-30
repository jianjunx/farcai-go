package main

import (
	_ "farcai-go/boot"
	_ "farcai-go/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
