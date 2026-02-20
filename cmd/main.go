package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("AgentDeploy starting...")
	if err := initialize(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("AgentDeploy ready")
}

func initialize() error {
	return nil
}
