package main

import (
	_ "pinyougou/routers"
	"github.com/astaxie/beego"
	_ "pinyougou/models"
)

func main() {
	beego.Run()
}

