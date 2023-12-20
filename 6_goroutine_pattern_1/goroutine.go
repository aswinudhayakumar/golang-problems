package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgCh  chan Message
	quitCh chan struct{}
}

func NewServer() *Server {
	return &Server{
		msgCh:  make(chan Message),
		quitCh: make(chan struct{}),
	}
}

func (s *Server) StartAndListen() {
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Printf("received msg from: %s and payload is: %s\n", msg.From, msg.Payload)
		case <-s.quitCh:
			fmt.Println("GraceFully shuting down....")
			// some code to gracefully disconnect connections like DB
			fmt.Println("Server shutdown")
			return
		default:
		}
	}
}

func (s *Server) SendMessageToServer(msgCh chan Message, payload string) {
	msg := Message{
		From:    "client",
		Payload: payload,
	}

	msgCh <- msg
}

func GraceFullShutDown(quitCh chan struct{}) {
	close(quitCh)
}

func main() {
	s := NewServer()

	go s.StartAndListen()

	go func() {
		time.Sleep(1 * time.Second)
		s.SendMessageToServer(s.msgCh, "Hello world!")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		s.SendMessageToServer(s.msgCh, "Hello coder!")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		GraceFullShutDown(s.quitCh)
	}()

	// using this select will block the main method till goroutine runs
	// just for testing purpose
	select {}
}
