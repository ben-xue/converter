package main

import (
	"fmt"
	"github.com/gohouse/converter"
	"os"
)

func main() {
	if len(os.Args) != 3{
		fmt.Println("Usage: convert dbName tableName")
		return
	}

	var b byte = 1
	var bb uint8 = 2

	b = bb
	fmt.Println(b)

	dbName := os.Args[1]
	tableName := os.Args[2]

	err := converter.NewTable2Struct().
		SavePath("./model.go").
		Dsn("root:unixc@tcp(192.168.100.13:3306)/"+dbName+"?charset=utf8").
		TagKey("xorm").
		EnableJsonTag(false).
		Table(tableName).
		Run()

	fmt.Println(err)
}
