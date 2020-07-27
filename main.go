package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func main() {
	fmt.Println("Hello World!")

	request := gorequest.New()
	resp, _, _ := request.Get("https://cheoeum.herokuapp.com").End()
	fmt.Print("Call heroku status code: ")
	fmt.Println(resp.StatusCode)
}
