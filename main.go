package main

import (
	"fmt"
)

type Event struct {
	Name string
}

func main() {
	events := make(chan Event)

	// listener (подписчик)
	go func() {
		for e := range events {
			fmt.Println("received:", e.Name)
		}
	}()

	// emitter (отправитель)
	events <- Event{Name: "user_created"}
}
