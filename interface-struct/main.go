package main

import (
	"fmt"

	"github.com/hugosjoberg/interface-struct/cache"
	"github.com/hugosjoberg/interface-struct/session"
)

func main() {
	c := cache.NewCache()
	s := session.NewSessionService(c)

	s.SetSession("john", "john-session")
	fmt.Println(s.GetSession("john"))
}
