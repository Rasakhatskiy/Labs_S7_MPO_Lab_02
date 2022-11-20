package server

import (
	"LAB_02/protocol"
	"LAB_02/utils"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"syscall"
	"time"
)

type ChatServer interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Start()
	Close()
}

type TcpChatServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex
}
type client struct {
	conn   net.Conn
	name   string
	writer *protocol.CommandWriter
}

func (s *TcpChatServer) Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err == nil {
		s.listener = l
	}
	log.Printf("Listening on %v", address)
	return err
}
func (s *TcpChatServer) Close() {
	s.listener.Close()
}
func (s *TcpChatServer) Start() {
	for {
		// XXX: need a way to break the loop
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print(err)
		} else {
			// handle connection
			client := s.accept(conn)
			go s.serve(client)
		}
	}
}

func (s *TcpChatServer) accept(conn net.Conn) *client {
	log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.clients)+1)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	client := &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	}
	s.clients = append(s.clients, client)
	return client
}
func (s *TcpChatServer) remove(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock() // remove the connections from clients array
	for i, check := range s.clients {
		if check == client {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}
	log.Printf("Closing connection from %v", client.conn.RemoteAddr().String())
	client.conn.Close()
}

func (s *TcpChatServer) serve(client *client) {
	cmdReader := protocol.NewCommandReader(client.conn)
	defer s.remove(client)
	for {
		cmd, err := cmdReader.Read()

		// check disconnect
		opError, ok := err.(*net.OpError)
		if ok {
			syscallError, ok := opError.Err.(*os.SyscallError)
			if ok {
				if syscallError.Err == syscall.ECONNRESET ||
					syscallError.Err == syscall.WSAECONNRESET {
					s.remove(client)
					log.Printf("Client %s disconnected\n", client.name)
					return
				}
			}
		}

		if err != nil && err != io.EOF {
			log.Printf("Read error: %v", err)
		}

		if cmd != nil {
			timeString := time.Now().Format("15:04:05")
			switch v := cmd.(type) {
			case protocol.MessageCommand:
				fmt.Printf("%s From: %s\nTo: %s\nReceived: %s\n", timeString, client.name, v.ReceiverName, v.Message)

				receiver, err := s.getClientByName(v.ReceiverName)
				if err != nil {
					utils.HandleErr(err)
					writer := protocol.NewCommandWriter(client.conn)
					err = writer.Write(protocol.MessageCommand{
						SenderName:   "@Server",
						ReceiverName: client.name,
						Time:         timeString,
						Message:      fmt.Sprintf("Client %s is not available", v.ReceiverName),
					})
					utils.HandleErr(err)
				} else {
					writer := protocol.NewCommandWriter(receiver.conn)
					v.SenderName = client.name
					err = writer.Write(v)
					utils.HandleErr(err)
				}

			case protocol.BroadcastCommand:
				fmt.Printf("%s From: %s\nTo: %s\nReceived: %s\n", timeString, client.name, protocol.EveryoneName, v.Message)
				broadcastMessage := protocol.BroadcastCommand{
					ReceiverName: protocol.EveryoneName,
					SenderName:   client.name,
					Time:         timeString,
					Message:      v.Message,
				}

				for i := range s.clients {
					writer := protocol.NewCommandWriter(s.clients[i].conn)
					err = writer.Write(broadcastMessage)
					utils.HandleErr(err)
				}

			case protocol.NameCommand:
				client.name = v.Name
			}
		}
		if err == io.EOF {
			break
		}
	}
}

func (s *TcpChatServer) Broadcast(command interface{}) error {
	for _, client := range s.clients {
		// TODO: handle error here?
		client.writer.Write(command)
	}
	return nil
}

func NewServer() *TcpChatServer {
	return &TcpChatServer{
		mutex: &sync.Mutex{},
	}
}

func (s *TcpChatServer) getClientByName(name string) (*client, error) {
	for i, c := range s.clients {
		if c.name == name {
			return s.clients[i], nil
		}
	}
	return nil, errors.New("no such client")
}
