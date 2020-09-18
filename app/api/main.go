package main

import (
	"log"

	"github.com/AlDrac/wallister_test_project/app/api/configs"
	"github.com/AlDrac/wallister_test_project/app/api/kernels"
)

func main() {
	config := configs.New(false)

	kernel, err := kernels.Initialise(config)
	if err != nil {
		log.Fatal(err)
	}

	kernel.Run()
}
