package main

import (
	"fmt"
	// "encoding/json"
	// "net/http"

)

type Item struct {
	shortDescription string
	price string
}

type Receipt struct {
	retailer string
	purchaseDate string
	purchaseTime string
	items []Item
	total string
}

func main() {
	fmt.Println("Hello")
}
