package main

import (
	"log"
	"sync"
)

type EventBroker struct {
	mu          sync.Mutex
	Subscribers map[int]Subscriber
	Messages    chan Event
}

// For events that subscribe to changes in gamestate, use the id of the card being played.
type Subscriber struct {
	id              int
	messageChan     chan string
	conditionMet    string
	conditionFailed string
	unsubChan       chan Event
}

type Event struct {
	// for unsubscribing to changes. not really needed for publishing them atm.
	senderId int
	message  string
	// for things like the user info and turn ordering. to prevent issues with locks.
	data interface{}
}

func (eb *EventBroker) StartPublishing() {
	for {
		m := <-eb.Messages
		if m.message == "unsub" {
			eb.Subscribers[m.senderId].messageChan <- "unsub ack"
			eb.Unsubscribe(m.senderId)
		} else {
			for _, s := range eb.Subscribers {
				Logger("MESSAGE OUT " + m.message)
				s.messageChan <- m.message
				Logger("MESSAGE SENT" + m.message)
			}
		}
	}
}

func (eb *EventBroker) Unsubscribe(id int) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	delete(eb.Subscribers, id)
}

func (eb *EventBroker) Subscribe(sub Subscriber) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.Subscribers[sub.id] = sub
}

// change to have a result chan instead of returning a value.
func (s *Subscriber) Receive(resChan chan bool) {
	exit := false
	for {
		m, ok := <-s.messageChan
		if !ok {
			log.Println("err receiving message,", s.id)
		}
		switch m {
		case s.conditionMet:
			go func() {
				resChan <- true
			}()
		case s.conditionFailed:
			// send unsub request in go routine to avoid blocking.
			go func() {
				s.unsubChan <- Event{senderId: s.id, message: "unsub"}
			}()
			// wait for ack before quitting
		case "unsub ack":
			exit = true
		default:
		}

		if exit {
			break
		}
	}

	resChan <- false
}

// *****************************************************
// Here's some predefined common events to use

var EndTurnEvent = Event{senderId: -1, message: "end turn"}
var LocationAddedEvent = Event{senderId: -1, message: "location added"}
var LocationRemovedEvent = Event{senderId: -1, message: "location removed"}
var PlayerDiscarded = Event{senderId: -1, message: "player discarded"}
var DoloresUmbridgeTrigger = Event{senderId: -1, message: "umbridge condition"}
