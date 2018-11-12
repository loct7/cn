package cn

import "log"

func CheckError(err error) {
	// Error
	if err != nil {
		log.Fatal(err)
	}
}
