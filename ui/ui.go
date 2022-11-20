package ui

import (
	"LAB_02/client"
	"LAB_02/utils"
	"errors"
	"github.com/rivo/tview"
)

var app *tview.Application
var tcpClient *client.TcpChatClient

func Init() {
	app = tview.NewApplication()
	tcpClient = client.NewClient()

	flex, err := getConnectFlex()
	utils.HandleErrFatal(err)
	setFlexAsWindow(flex)
}

func setFlexAsWindow(flex *tview.Flex) {
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		utils.HandleErrFatal(err)
	}
}

func getConnectFlex() (*tview.Flex, error) {
	thisFlex := tview.NewFlex()

	inputIP := tview.NewInputField().SetLabel("IP:PORT").SetText("127.0.0.1:3333")
	inputName := tview.NewInputField().SetLabel("Name")

	form := tview.NewForm().
		AddFormItem(inputIP).
		AddFormItem(inputName).
		AddButton("Connect", func() {
			ip := inputIP.GetText()
			name := inputName.GetText()
			err := tcpClient.DialAndSetName(ip, name)
			if err != nil {
				showMessage(thisFlex, "OK", err.Error(), messageTypeError)
			} else {
				showMessage(thisFlex, "OK", "Connected", messageTypeInfo)
			}
			//handleErrorMessageUI(err, thisFlex)
		})

	thisFlex.AddItem(form, 0, 1, true)

	return thisFlex, nil
}

func getMessageFlex() (*tview.Flex, error) {
	return nil, errors.New("implement me")
}

func getRecipientTree() (*tview.TreeView, error) {
	return nil, errors.New("implement me")
}

func getMessageTextView() (*tview.TextView, error) {
	return nil, errors.New("implement me")
}

func getTextInput() (*tview.InputField, error) {
	return nil, errors.New("implement me")
}
