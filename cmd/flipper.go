package main

import "regexp"

// Regex patterns for the flipper
var (
	// Matches 'First, Last' or 'First Middle, Last'
	namePattern = regexp.MustCompile(`^(\w+), (\w+)$`)

	//Matches an email address with 'First_Last@...'
	emailPattern = regexp.MustCompile(`^(\w+)_(\w+)@`)
)

func FlipName(name string) (string, error) {
	return "", nil
}

func EmailToName(email string) (string, error) {
	return "", nil
}
