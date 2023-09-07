package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	info, err := tickets.ReadFIleCSV("tickets.csv")
	if err != nil {
		panic(err)

	}

	total, err := tickets.GetTotalTickets(info, "Brazil")
	if err != nil {
		panic("error")
	}

	fmt.Println("Brazil", total)

	_, err2 := tickets.GetMornings(info, "hoas")
	if err2 != nil {
		panic("error")
	}

	val, err3 := tickets.GetCountByPeriod(info, tickets.Madr)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(val)
	val2, err4 := tickets.GetCountByPeriod(info, tickets.Mania)
	if err4 != nil {
		panic(err3)
	}
	fmt.Println(val2)
	val3, err5 := tickets.GetCountByPeriod(info, tickets.Tarde)
	if err5 != nil {
		panic(err3)
	}
	fmt.Println(val3)
	val4, err6 := tickets.GetCountByPeriod(info, tickets.Noche)
	if err6 != nil {
		panic(err3)
	}
	fmt.Println(val4)
	fmt.Println(val + val2 + val3 + val4)

	//
	val6, err8 := tickets.AverageDestination(info, "China", 100)
	if err != nil {
		panic(err8)
	}
	fmt.Println(val6)
}
