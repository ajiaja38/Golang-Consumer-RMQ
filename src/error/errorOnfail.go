package error

import "log"

func FailOnError(err error, mes string) {
	if err != nil {
		log.Panicf("%s: %s", err, mes)
	}
}
