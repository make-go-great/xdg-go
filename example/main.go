package main

import (
	"fmt"

	"github.com/make-go-great/xdg-go"
)

func main() {
	fmt.Println("Home", xdg.GetHome())
	fmt.Println("DataHome", xdg.GetDataHome())
	fmt.Println("ConfigHome", xdg.GetConfigHome())
}
