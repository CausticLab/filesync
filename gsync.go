package main

import (
	"filesync/config"
	vars "filesync/vars"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("CPUs: ", runtime.NumCPU())
	done := make(chan bool)
	start(done)
	<-done
}

func start(done chan bool) {
	vars.Init()
	vars := vars.GetConfig()
	log.Printf("Fileshare Config:\n%+v\n", vars)

	if vars.Mode == "server" {
		config.StartServer()
	} else if vars.Mode == "client" {
		config.StartClient(done)
	}
}
