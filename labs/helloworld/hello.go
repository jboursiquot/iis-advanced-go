package helloworld

import "rsc.io/quote"

// Hello returns a greeting.
func Hello() string {
	return "Hello, world!"
}

// HelloFromRSC returns greeting from rsc.io/quote.
func HelloFromRSC() string {
	return quote.Hello()
}
