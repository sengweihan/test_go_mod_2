package tester

import (
	"log"

	p "github.com/sengweihan/test_go_mod"
)

func Alibaba() {
	age := 110
	r := p.T2{Age: &age}

	log.Printf("Age is: %v", *r.Age)
}
