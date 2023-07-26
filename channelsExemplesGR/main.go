package main

import "fmt"

type Message struct {
	from    string
	payload string
}

type Server struct {
	msgchan chan Message
}

func (s *Server) startandListen() {
	fmt.Println("server is running ⚡")

	for {
		message := <-s.msgchan
		fmt.Printf("message from : %s\ncontent : %s \n", message.from, message.payload)
	}
}

func (s *Server) sendMessage(msg string) {
	message := Message{
		from:    "Iskander abbassi ⚡",
		payload: msg,
	}

	s.msgchan <- message
}

func main() {
	// s := &Server{
	// 	msgchan: make(chan Message),
	// }

	// go s.startandListen()

	// s.sendMessage("Hello chabeb 🚀")

	// fmt.Scanf("\n ")

	var data int

	go func() { 
		data++ }()
	if data == 0 {
	 fmt.Println("the value is 0.")
	} else {
	 fmt.Printf("the value is %v.\n", data)
	}
}
