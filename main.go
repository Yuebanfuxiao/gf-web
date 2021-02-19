package main

import (
	_ "gf-web/boot"
	_ "gf-web/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
