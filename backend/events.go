package game

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
	senderId int
	message  string
}

func (eb *EventBroker) StartPublishing() {
	for {
		m := <-eb.Messages
		if m.message == "unsub" {
			log.Println("sending unsub ack")
			eb.Subscribers[m.senderId].messageChan <- "unsub ack"
			eb.Unsubscribe(m.senderId)
		} else {
			for _, s := range eb.Subscribers {
				log.Println("sending", s.id, "message: ", m.message)
				s.messageChan <- m.message
			}
		}
	}
}

func (eb *EventBroker) Unsubscribe(id int) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	delete(eb.Subscribers, id)
	log.Println("unsubscribing ", id)
}

func (eb *EventBroker) Subscribe(sub Subscriber) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.Subscribers[sub.id] = sub
}

func (s *Subscriber) Receive() bool {
	res := false
	exit := false
	for {
		m, ok := <-s.messageChan
		if !ok {
			log.Println("err receiving message,", s.id)
		}
		switch m {
		case s.conditionMet:
			res = true
			// send unsub request in go routine to avoid blocking.
			go func() {
				s.unsubChan <- Event{senderId: s.id, message: "unsub"}
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

	return res
}
