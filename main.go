package main

import (
	"fmt"
	"prediksi-kelulusan-mahasiswa-naive-bayes/config"
	"prediksi-kelulusan-mahasiswa-naive-bayes/routes"
)

func main() {
	fmt.Println("Final Project Alterra Academy Backend Golang Batch 2")
	config.InitDB()
	fmt.Println("connection db success")

	e := routes.New()

	e.Start(":1234")
}
