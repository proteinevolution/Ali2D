package main

import "log"

const newline = "\n"

// If error occurred, the Program wil exit and error will be logged
func check(err error) {

	if err != nil {

		log.Fatal(err)
	}
}
