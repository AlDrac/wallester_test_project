package main

import (
	"github.com/AlDrac/wallister_test_project/app/web/configs"
	"github.com/AlDrac/wallister_test_project/app/web/kernels"
)

func main() {
	config := configs.New()

	kernel := kernels.Initialise(config)
	kernel.Run()
}
