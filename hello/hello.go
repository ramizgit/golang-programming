package hello

import "rsc.io/quote"

func Quote() string {
	return quote.Hello()
}

func Hello() string {
	return GetMessage()
}

func GetMessage() string {
	return "Hello, world."
}
