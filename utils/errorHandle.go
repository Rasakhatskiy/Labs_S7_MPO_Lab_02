package utils

import (
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func HandleErrFatal(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
