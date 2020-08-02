package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func main() {
	fmt.Println("Hello.. Hit the Server!")

	request := gorequest.New()
	//resp, _, _ := request.Get("https://cheoeum.herokuapp.com").End()
	//fmt.Print("Call heroku status code: ")
	//fmt.Println(resp.StatusCode)

	resp, _, _ := request.Get("https://kingofweather-api.herokuapp.com/check").End()
	fmt.Print("Call KingOfWeather-Api status code: ")
	fmt.Println(resp.StatusCode)


}
