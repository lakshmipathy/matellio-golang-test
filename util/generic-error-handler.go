package util

import "log"

func HandleGenericError(err error, s string) {
	if err != nil {
		log.Fatalln(s, err)
	}
}
