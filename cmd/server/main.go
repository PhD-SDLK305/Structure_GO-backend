package main

import (
	"gobackend/iternal/initialize"
)

func main() {
	// r := routes.NewRouter()
	// if err := r.Run(); err != nil {
	// 	log.Fatalf("failed to run server: %v", err)
	// }
	initialize.Run()
}
