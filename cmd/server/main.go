package main

import (
	"fmt"

	"github.com/ThiagoKufa/ref_api_go/configs"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(configs.DBName)
}
