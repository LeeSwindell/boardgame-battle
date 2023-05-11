package game

import (
	"log"
	"sync"
	"testing"
)

func TestEvents(t *testing.T) {
	eb := EventBroker{
		mu:          sync.Mutex{},
		Subscribers: make(map[int]Subscriber),
		Messages:    make(chan Event),
	}

	sub1 := Subscriber{
		id:              1,
		messageChan:     make(chan string),
		conditionMet:    "villain killed",
		conditionFailed: "end turn",
		unsubChan:       eb.Messages,
	}

	sub2 := Subscriber{
		id:              2,
		messageChan:     make(chan string),
		conditionMet:    "4 spells cast",
		conditionFailed: "end turn",
		unsubChan:       eb.Messages,
	}

	go eb.StartPublishing()

	go sub1.Receive()
	go sub2.Receive()

	eb.Subscribe(sub1)
	eb.Subscribe(sub2)

	// log.Println(eb.Subscribers)

	event1 := Event{senderId: 0, message: "villain killed"}
	event2 := Event{senderId: 0, message: "4 spells cast"}
	endevent := Event{senderId: 0, message: "end turn"}

	log.Println(eb.Messages)

	eb.Messages <- event1
	eb.Messages <- event1
	eb.Messages <- event2

	eb.Messages <- endevent
	eb.Messages <- event1
	eb.Messages <- endevent
}
