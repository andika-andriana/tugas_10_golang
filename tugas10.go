package main

import "fmt"
import "time"

func main() {

	var data = map[string]string{
		"nama":          "Andika Andriana",
		"hobi":          "Belajar Pemograman",
		"tanggal_lahir": "05-05-1995",
	}

	fmt.Println("------------------------>>")

	for key, val := range data {
		fmt.Println(key, ":\t", val)
		time.Sleep(time.Second * 5)
	}

	fmt.Println("<<------------------------")
}
