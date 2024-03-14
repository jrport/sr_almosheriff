package main

import (
	"fmt"

	"jrport.almosheriff/api"
)


func main(){
	fmt.Println("Starting stuff...")
	api := api.NewApiServer(":8080")
	api.Run()
}
