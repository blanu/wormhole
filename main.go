package main

import (
	"flag"
	"github.com/blanu/radiowave"
	"github.com/blanu/wormhole/internal/shuttle"
	"log"
	"os"
)

func main() {
	logFile, openError := os.OpenFile("wormhole.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openError != nil {
		os.Exit(4)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println("<~=wormhole=~>")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalln("not enough arguments")
		os.Exit(5)
	} else {
		log.Printf("args: %v", args)
	}

	client := radiowave.NewFileFromFd[shuttle.Message, shuttle.Message](shuttle.NewMessageFactory(), 3)

	coreFactory := core.NewFactory(args)
	core1, allocationError := coreFactory.Allocate()
	if allocationError != nil {
		log.Fatalln("Error allocating core")
		os.Exit(6)
	} else {
		log.Println("Core launched.")
	}

	text1, dialError := radiowave.Dial[shuttle.Message, shuttle.Message](shuttle.NewMessageFactory(), "localhost:7000")
	if dialError != nil {
		log.Fatalln("Error connecting to service")
		os.Exit(7)
	} else {
		log.Println("Connected.")
	}

	shuttle1 := shuttle.New(client, *core1, *text1)
	shuttle1.Service()
}
