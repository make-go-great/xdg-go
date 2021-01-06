package main

import (
	"fmt"

	"github.com/haunt98/xdg"
)

func main() {
	fmt.Println("Home", xdg.GetHome())
	fmt.Println("DataHome", xdg.GetDataHome())
	fmt.Println("ConfigHome", xdg.GetConfigHome())
}
