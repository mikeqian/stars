package main

import (
	"log"
	"net/http"
)

func Read() {
	_, e := http.Get("http://www.cnblogs.com/bnbqian/p/4923597.html")
	if e != nil {
		log.Println(e)
	}
}
