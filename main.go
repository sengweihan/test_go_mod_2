package main

import (
	"log"

	p "github.com/sengweihan/test_go_mod"
)

func main() {
	age := 50
	r := p.T2{Age: &age}

	log.Printf("Age is: %v", *r.Age)
}
