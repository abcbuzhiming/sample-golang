package main

import (
    "fmt"
    "github.com/gogf/gf"
    //_ "sample-goframe/boot"
	_ "sample-goframe/router"
)

func main() {
    fmt.Println("hello GF", gf.VERSION)
    g.Server().Run()
}