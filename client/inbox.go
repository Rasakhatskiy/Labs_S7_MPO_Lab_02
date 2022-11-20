package client

import (
	"LAB_02/protocol"
	"LAB_02/utils"
)

type Message struct {
	Time   string
	Author string
	Text   string
}

type Inbox struct {
	Recipients    []string
	MessagesTable [][]Message
}

func NewInbox() *Inbox {
	return &Inbox{
		Recipients:    []string{protocol.ServerName, protocol.EveryoneName},
		MessagesTable: [][]Message{{}, {}},
	}
}

func (in *Inbox) getMessagesFrom(name string) []Message {
	i, err := utils.Find(in.Recipients, name)
	if err != nil {
		return []Message{}
	}
	return in.MessagesTable[i]
}

func (in *Inbox) putMessageFrom(message Message) error {
	i, err := utils.Find(in.Recipients, message.Author)
	if err != nil {
		_, ok := err.(utils.ItemNotFoundError)
		if ok {
			in.Recipients = append(in.Recipients, message.Author)
			messages := []Message{message}
			in.MessagesTable = append(in.MessagesTable, messages)
		} else {
			return err
		}
	}
	in.MessagesTable[i] = append(in.MessagesTable[i], message)
	return nil
}

func (in *Inbox) putMessageEveryone(message Message) {
	in.MessagesTable[protocol.EveryoneInboxID] = append(in.MessagesTable[protocol.EveryoneInboxID], message)
}
