package main

import (
	"flag"
	"github.com/blanu/wormhole"
	"log"
	"os"
	"strings"
)

func main() {
	logFile, openError := os.OpenFile("wormhole.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openError != nil {
		os.Exit(4)
	}
	defer func() {
		_ = logFile.Close()
	}()

	logger := log.New(logFile, "wormhole:", log.LstdFlags|log.Lshortfile)

	logger.SetOutput(logFile)

	logger.Println("<<~=wormhole=~>>")

	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		logger.Fatalln("not enough arguments")
	} else {
		logger.Printf("args: %v", args)
	}
	args = args[1:] // Discard the path to this executable, which is in the first argument.

	client := os.NewFile(3, "incoming")
	if client == nil {
		logger.Fatalln("invalid file descriptor for client")
	}

	registerFactory := wormhole.NewRegisterFactory(args, logger)
	text1, allocationError := registerFactory.Allocate()
	if allocationError != nil {
		log.Fatalln("Error allocating core")
	} else {
		log.Printf("Launched: %s\n", strings.Join(args, " "))
	}

	shuttle1 := wormhole.NewShuttle(client, *text1, logger)
	shuttle1.Service()
}
