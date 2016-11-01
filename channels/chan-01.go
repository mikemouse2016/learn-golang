package main

import (
	"log"
	"fmt"
	"time"
)

type Disp struct {
	clients    map[chan string]bool
	newClient  chan chan string
	deadClient chan chan string
	message    chan string
}

func (d *Disp) Start() {
	go func() {
		for {
			select {
			case c := <-d.newClient:
				d.clients[c] = true
				log.Println("New client added")

			case c := <-d.deadClient:
				delete(d.clients, c)
				close(c)
				log.Println("Client deleted")

			case message := <-d.message:
				for c := range d.clients {
					c <- message
				}
				log.Printf("Sent message to %d clients", len(d.clients))
			}
		}
	}()
}

func (d *Disp) dispAdd(i int) {
	for n := 1; n <= i; n++ {
		messageChan := make(chan string)
		d.newClient <- messageChan
		go func(id int) {
			//j := i
			msg := <-messageChan
			log.Printf(" %d msg: %v", id, msg)
		}(n)
	}
}

func main() {
	d := &Disp{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}
	d.Start()
	d.dispAdd(3)
	fmt.Println("Enter text: ")
	text := ""
	fmt.Scanln(&text)
	fmt.Println(text)
	d.message <- text
	time.Sleep(5 * time.Second)
}
