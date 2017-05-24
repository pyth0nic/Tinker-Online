package controllers

import (
	"golang.org/x/net/websocket"
	"github.com/revel/revel"
	"fmt"
	"Tinker_Online/app/controllers/tinker"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) ExpressionSocket(ws *websocket.Conn) revel.Result {
	var a tinker.Tree
	a= tinker.Tree{nil,"a",nil}
	fmt.Println(a.Value)
	
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()
	
	var t tinker.Export

	for {
		msg, ok := <-newMessages
		if !ok {
			return nil
		}
		
		if msg != "old" { 
			postfix := t.Postfix(msg)
			fmt.Println(msg)
			fmt.Println(postfix)
			if websocket.JSON.Send(ws, postfix) != nil {
					return nil
			}
			msg="old"
		}
	}
	
	//how to keep websocket open?
	//closes socket?
	return nil
}

func process_input() {

}

/*
func (c WebSocket) Room(user string) revel.Result {
	return c.Render(user)
}

func (c WebSocket) RoomSocket(user string, ws *websocket.Conn) revel.Result {
	// Join the room.
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user)
	defer chatroom.Leave(user)

	// Send down the archive.
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroom.
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			chatroom.Say(user, msg)
		}
	}
	return nil
}
*/