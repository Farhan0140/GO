package main

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func main() {
	cnf := config.GetConfig()
	// fmt.Println(cnf.HttpPort)
	// fmt.Println(cnf.ServiceName)
	// fmt.Println(cnf.Version)

	rest.Start(cnf)
}
