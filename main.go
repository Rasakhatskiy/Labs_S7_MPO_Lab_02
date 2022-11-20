package main

import (
	"LAB_02/server"
	"LAB_02/ui"
	"LAB_02/utils"
	"os"
)

func main() {
	//box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	//if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	//	panic(err)
	//}

	args := os.Args
	if args[1] == "server" {
		var s server.ChatServer
		s = server.NewServer()
		err := s.Listen(":3333") // start the server
		utils.HandleErrFatal(err)

		s.Start()
	} else {
		ui.Init()
		//reader := bufio.NewReader(os.Stdin)

		//address := "127.0.0.1:3333"
		//var err error

		//chatClient := client.NewClient()

		//err = chatClient.Dial(address)

		//for err != nil {
		//	utils.HandleErr(err)
		//	time.Sleep(time.Second * 4)
		//	err = chatClient.Dial(address)
		//}
		//
		//defer chatClient.Close()
		//
		//// start the client to listen for incoming message
		//go chatClient.Start()
		//
		//fmt.Println("Your name: ")
		//myName, _ := reader.ReadString('\n')
		//myName = strings.Replace(myName, " ", "_", -1)
		//err = chatClient.SetName(myName)
		//utils.HandleErr(err)
		//
		//for true {
		//	fmt.Println("To: ")
		//	recipientName, _ := reader.ReadString('\n')
		//	recipientName = strings.Replace(recipientName, "\r", "", -1)
		//	recipientName = strings.Replace(recipientName, "\n", "", -1)
		//
		//	fmt.Println("Enter message: ")
		//	messageText, _ := reader.ReadString('\n')
		//	messageText = strings.Replace(messageText, "\r", "", -1)
		//	messageText = strings.Replace(messageText, "\n", "", -1)
		//
		//	if recipientName == "@everyone" {
		//		err = chatClient.Broadcast(messageText)
		//	} else {
		//		err = chatClient.SendMessage(recipientName, messageText)
		//	}
		//	utils.HandleErr(err)
		//	time.Sleep(time.Second * 4)
		//}

	}

}
