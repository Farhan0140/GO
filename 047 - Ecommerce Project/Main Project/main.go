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

	// jwt, err := util.Create_JWT("Yoo-Mother-Fkr", util.Payload{
	// 	ID: 34,
	// 	FirstName: "Farhan",
	// 	LastName: "Nadim",
	// 	Email: "farhan@gmail.com",
	// 	Password: "34-sdf@#",
	// 	IsAdmin: true,
	// })
	// if err != nil {
	// 	fmt.Println("Fuck you")
	// }

	// fmt.Println(jwt)
}
