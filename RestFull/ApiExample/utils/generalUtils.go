package utils

import "log"

//CheckError metotu
func CheckError(err error) {
	if err != nil {
		log.Print(err)
	}
}
