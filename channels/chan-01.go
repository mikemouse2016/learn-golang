package main

import (
	"log"
	"time"
	"fmt"
	"math/rand"
	"math"
)

type Dispatch struct {
	clients    map[chan string]bool
	newClient  chan chan string
	deadClient chan chan string
	message    chan string
}

func (d *Dispatch) Start() {
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

func (d *Dispatch) registerWorker(i int) {
	for n := 1; n <= i; n++ {
		messageChan := make(chan string)
		d.newClient <- messageChan
		// worker
		go func(id int) {
			// self destruct worker, random timer
			if math.Mod(float64(id), 10) == 0 {
				go destroyWorker(d.deadClient, messageChan)
			}
			for {
				msg, ok := <-messageChan
				if !ok {
					// If our messageChan was closed, this means that the worker has
					// disconnected.
					break
				}
				//log.Printf(" %d msg: %v", id, msg)
				if msg == "0" {
					break
				}
			}
		}(n)
		//time.Sleep(500 * time.Millisecond)

	}
}

func destroyWorker(cha chan chan string, chb chan string) {
	rnd := time.Duration(rand.Intn(95) + 5) * time.Second
	log.Println("chan:", chb, "destruct in:", rnd)
	select {
	case <-time.After(rnd):
		cha <- chb
	}
}

func main() {
	d := &Dispatch{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}
	d.Start()
	go func() {
		//for {
		d.registerWorker(100)
		//time.Sleep(20 + time.Second)
		//d.message <- "0"
		//}
	}()
	//fmt.Println("Enter text: ")
	//text := ""
	//fmt.Scanln(&text)
	//fmt.Println(text)
	//d.message <- text

	for i := 0; ; i++ {
		d.message <- fmt.Sprintf("test %v", i)
		time.Sleep(5 * time.Second)
	}
}
