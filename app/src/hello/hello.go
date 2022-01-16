package hello

import "rsc.io/quote"

func Hello() string {
	return "Hello, World!"
}

func HelloImport() string {
	return quote.Hello()
}
