package main

import (
	"fmt"

	"github.com/hugosjoberg/interface-struct-bad-example/cache"
	"github.com/hugosjoberg/interface-struct-bad-example/session"
)

func main() {
	c := cache.NewCache()
	s := session.NewSessionService(c)

	s.SetSession("john", "john-session")
	fmt.Println(s.GetSession("john"))
}
