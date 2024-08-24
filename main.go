package main

import (
	"log"

	p "github.com/sengweihan/test_go_mod"
)

func main() {
	r := p.TTT{
		Email: "test@gmail.com",
	}
	log.Println("rtests")
	log.Printf("%+v", r)
}
