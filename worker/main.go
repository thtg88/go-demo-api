package main

import (
	"context"
	"flag"
	"goDemoApi/queue"
	"goDemoApi/worker/tasks"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load .env file")
	}

	flag.Parse()

	c := context.Background()

	queue.Factory.StartConsumers(c)

	sig := tasks.WaitSignal()
	log.Println(sig.String())

	err = queue.Factory.Close()
	if err != nil {
		log.Fatal(err)
	}
}
