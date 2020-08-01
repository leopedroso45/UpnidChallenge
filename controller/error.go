package controller

import (
	"log"
)

/*CheckError: Simple method to handle errors. For simplicity, it fires a Panic, but a customization can be made to identify the type of error.*/
func CheckError(err error) {
	if err != nil {
		log.Panicln(err.Error())
	}
}
