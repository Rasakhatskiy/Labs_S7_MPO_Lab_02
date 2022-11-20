package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MessageType int

const (
	messageTypeInfo  MessageType = 0
	messageTypeError MessageType = 1
)

func showMessage(returnFlex *tview.Flex, btnText, message string, messageType MessageType) {
	flex := tview.NewFlex()
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{btnText}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == btnText {
				if returnFlex != nil {
					setFlexAsWindow(returnFlex)
				} else {
					app.Stop()
				}
			}
		})

	switch messageType {
	case messageTypeInfo:
		modal.SetBackgroundColor(tcell.ColorBlue)
	case messageTypeError:
		modal.SetBackgroundColor(tcell.ColorRed)
	}

	flex.AddItem(modal, 0, 1, true)

	setFlexAsWindow(flex)
}

func handleErrorMessageUI(err error, returnFlex *tview.Flex) {
	if err != nil {
		showMessage(returnFlex, "OK", err.Error(), messageTypeError)
	}
}
