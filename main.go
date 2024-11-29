package main

import (
	"ewallet-framework/cmd"
	"ewallet-framework/helpers"
)

func main() {
	// Load configuration
	helpers.SetupConfig()
	helpers.SetupLogger()
	// helpers.SetupMySQL()

	// Run server
	go cmd.ServeGRPC() // run in different thread
	cmd.ServeHTTP()    // run in the main thread
}
