package main

import (
	"fmt"

	_"almosheriff.api/src/database"
	"almosheriff.api/src/server"
)

func main(){
	fmt.Println("Starting up...")
	api := server.NewApi(":8000")
	api.Run()
	// println("Connecting...")
	// dbCon := database.NewDbConn(5000)
	// println("Creating tables...")
	// err := dbCon.SetupTable()
	// if err != nil {
		// fmt.Println(err.Error())
	// }
	// println("Finished setting tables...")
}
