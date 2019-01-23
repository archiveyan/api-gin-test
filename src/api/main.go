package main

import (
	_ "fmt"
	_ "log"
	_ "os"

	"github.com/weet/router"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
