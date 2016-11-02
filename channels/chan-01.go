package main

import (
	"log"
	"time"
	"fmt"
	"math/rand"
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

func (d *Disp) registerChan(i int) {
	for n := 1; n <= i; n++ {
		messageChan := make(chan string)
		d.newClient <- messageChan
		go func(id int) {
			//self destruct
			go func() {
				rnd := time.Duration(rand.Intn(95) + 5) * time.Second
				log.Println("chan:", messageChan, "destruct in:", rnd)
				select {
				case <-time.After(rnd):
					d.deadClient <- messageChan
				}
			}()

			for {
				msg, ok := <-messageChan
				if !ok {
					// If our messageChan was closed, this means that the client has
					// disconnected.
					break
				}
				//log.Printf(" %d msg: %v", id, msg)
				if msg == "0" {
					break
				}
			}
		}(n)
		time.Sleep(1 * time.Second)

	}
}

func (d *Disp) chanDel() {

}

func main() {
	d := &Disp{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}
	d.Start()
	go func() {
		for {
			d.registerChan(100)
		}
	}()
	//fmt.Println("Enter text: ")
	//text := ""
	//fmt.Scanln(&text)
	//fmt.Println(text)
	//d.message <- text
	//d.message <- "test 1"
	//time.Sleep(5 * time.Second)
	//d.message <- "test 2"
	//time.Sleep(5 * time.Second)
	for i := 0; ; i++ {
		d.message <- fmt.Sprintf("test %v", i)
		time.Sleep(5 * time.Second)
	}
}
