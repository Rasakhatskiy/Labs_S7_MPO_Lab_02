package protocol

import (
	"LAB_02/utils"
	"errors"
	"fmt"
	"io"
)

var (
	UnknownCommand = errors.New("Unknown command")
)

// BroadcastCommand is used for sending new message from client
type BroadcastCommand struct {
	ReceiverName string
	SenderName   string
	Time         string
	Message      string
} // NameCommand is used for setting client display name
type NameCommand struct {
	Name string
} // MessageCommand is used for notifying new messages
type MessageCommand struct {
	ReceiverName string
	SenderName   string
	Time         string
	Message      string
}

type CommandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: writer,
	}
}
func (w *CommandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	return err
}
func (w *CommandWriter) Write(command interface{}) error {
	// naive implementation ...
	var err error
	switch v := command.(type) {
	case BroadcastCommand:
		err = w.writeString(fmt.Sprintf("%s %v %s %s", CmdBroadcast, v.Message, v.Time, DelimNLS))
	case MessageCommand:
		err = w.writeString(fmt.Sprintf("%s %v %v %v %s", CmdMessage, v.ReceiverName, v.Message, v.Time, DelimNLS))
	case NameCommand:
		err = w.writeString(fmt.Sprintf("%s %v%s", CmdName, v.Name, DelimNLS))
	default:
		err = UnknownCommand
	}
	return err
}

type CommandReader struct {
	reader *utils.MyReader
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: utils.NewReader(reader),
	}
}
func (r *CommandReader) Read() (interface{}, error) {
	// Read the command
	commandName, err := r.reader.ReadString(DelimSP)
	if err != nil {
		return nil, err
	}

	switch commandName {
	case CmdMessage:
		user, err := r.reader.ReadString(DelimSP)
		if err != nil {
			return nil, err
		}
		time, err := r.reader.ReadString(DelimSP)
		if err != nil {
			return nil, err
		}
		message, err := r.reader.ReadString(DelimNL)
		if err != nil {
			return nil, err
		}
		return MessageCommand{
			ReceiverName: user,
			Time:         time,
			Message:      message,
		}, nil

	case CmdBroadcast:
		message, err := r.reader.ReadString(DelimNL)
		if err != nil {
			return nil, err
		}
		time, err := r.reader.ReadString(DelimSP)
		if err != nil {
			return nil, err
		}
		return BroadcastCommand{
			Time:    time,
			Message: message,
		}, nil

	case CmdName:
		name, err := r.reader.ReadString(DelimNL)
		if err != nil {
			return nil, err
		}
		return NameCommand{
			Name: name,
		}, nil
	}

	return nil, UnknownCommand
}
