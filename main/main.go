package main

import "Introducao-Backend/conn"

func main() {
	_, err := conn.ConnectDatabase()
	if err != nil {
		panic(err)
	}
}