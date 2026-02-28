package main


type User interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type 