package main

import (
	"fmt"

	"github.com/ppthana/cinema/movie"
	"github.com/ppthana/cinema/ticket"
)

func main() {
	movieName := movie.FindName("tt4145796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
