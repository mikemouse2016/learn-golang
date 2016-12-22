package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

// Dispatch type holds a map of chan as clients, a newClint chan to process new clients, a deadClient chan
// to process dead clients and a message chan to send messages to active clients
type Dispatch struct {
	//map key is the message chan, map value not used
	clients    map[chan string]bool
	newClient  chan chan string
	deadClient chan chan string
	message    chan string
}

// Start launches a go routine to process client events (new, dead) and messages
func (d *Dispatch) Start() {
	go func() {
		for {
			select {
			case c := <-d.newClient:
				// add new client to dispatch client map
				d.clients[c] = true
				log.Println("New client added")

			case c := <-d.deadClient:
				// delete dead client from dispatch client map
				delete(d.clients, c)
				close(c)
				log.Println("Client deleted")

			case message := <-d.message:
				// loop trough client map and send the message
				for c := range d.clients {
					c <- message
				}
				log.Printf("Sent message to %d clients", len(d.clients))
			}
		}
	}()
}

// Register i workers
func (d *Dispatch) registerWorker(i int) {
	for n := 1; n <= i; n++ {
		// make new chan and send it to dispatcher, acting as an worker id.
		// Channels are first class functions
		messageChan := make(chan string)
		d.newClient <- messageChan
		// launch a go routine for each worker
		go func(nth int) {
			// self destruct each 10th worker, random timer
			if math.Mod(float64(nth), 10) == 0 {
				go destroyWorker(d.deadClient, messageChan)
			}

			// process messages received trough messageChan in a loop
			for {
				msg, ok := <-messageChan
				if !ok {
					// If our messageChan was closed, this means that the worker has
					// disconnected.
					break
				}
				//log.Printf(" %d msg: %v", id, msg)
				if ok && (msg == "0") {
					break
				}
				//TODO
				// do some work
			}
		}(n)
		//time.Sleep(500 * time.Millisecond)

	}
}

// destroyWorker will send a dead client notification to dispatcher after a random time of 5-100 sec
func destroyWorker(deadClient chan chan string, messageChan chan string) {
	rnd := time.Duration(rand.Intn(95)+5) * time.Second
	log.Println("chan:", messageChan, "destruct in:", rnd)
	//TODO
	// check if worker is destroyed from other reasons
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in destroy worker", r)
		}
	}()
	func () {
		select {
		case <-time.After(rnd):
			deadClient <- messageChan
		}
	}()
}

func main() {

	// make dispatcher struct
	d := &Dispatch{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}

	// start the dispatcher go routine
	d.Start()

	// register 100 workers
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

	go func() {
		time.Sleep(10 * time.Second)
		d.message <- "0"
		d.registerWorker(1)
	}()
	// send a message to workers each 5 seconds
	for i := 0; ; i++ {
		d.message <- fmt.Sprintf("test %v", i)
		time.Sleep(5 * time.Second)
	}
}
