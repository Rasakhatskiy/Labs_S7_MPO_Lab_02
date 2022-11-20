package client

import (
	"LAB_02/protocol"
	"LAB_02/utils"
	"io"
	"log"
	"net"
)

type messageHandler func(string)

type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
	Send(command interface{}) error
	SetName(name string) error
	SendMessage(message string) error
	Incoming() chan protocol.MessageCommand
}

type TcpChatClient struct {
	conn      net.Conn
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandWriter
	name      string
	//incoming  chan protocol.MessageCommand
	inbox *Inbox
}

func NewClient() *TcpChatClient {
	return &TcpChatClient{
		//incoming: make(chan protocol.MessageCommand),
		inbox: NewInbox(),
	}
}

func (c *TcpChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	}

	c.cmdReader = protocol.NewCommandReader(conn)
	c.cmdWriter = protocol.NewCommandWriter(conn)

	return err
}

func (c *TcpChatClient) DialAndSetName(address, name string) error {
	err := c.Dial(address)
	if err != nil {
		return err
	}

	err = c.SetName(name)
	if err != nil {
		return err
	}

	return nil
}

func (c *TcpChatClient) Start() {
	for {
		cmd, err := c.cmdReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Read error %v", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.MessageCommand:
				log.Printf("From: %s:\n%s", v.ReceiverName, v.Message)
				m := Message{
					Time:   v.Time,
					Author: v.SenderName,
					Text:   v.Message,
				}
				err = c.inbox.putMessageFrom(m)
				utils.HandleErr(err)

			case protocol.BroadcastCommand:
				m := Message{
					Time:   v.Time,
					Author: v.SenderName,
					Text:   v.Message,
				}
				c.inbox.putMessageEveryone(m)

			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}

func (c *TcpChatClient) Close() {
	c.conn.Close()
}

//func (c *TcpChatClient) Incoming() chan protocol.MessageCommand {
//	return c.incoming
//}

func (c *TcpChatClient) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}

func (c *TcpChatClient) SetName(name string) error {
	return c.Send(protocol.NameCommand{
		Name: name,
	})
}

func (c *TcpChatClient) SendMessage(receiverName, message string) error {
	return c.Send(protocol.MessageCommand{
		ReceiverName: receiverName,
		Message:      message,
	})
}

func (c *TcpChatClient) Broadcast(message string) error {
	return c.Send(protocol.BroadcastCommand{
		Message: message,
	})
}
