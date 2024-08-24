package main

import (
	"log"

	p "github.com/sengweihan/test_go_mod"
)

func main() {
	r := p.Match{Text: "hello world"}
	log.Printf("%+v", r)
}
