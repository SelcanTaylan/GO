package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sayi := rand.Intn(5) + 1
	var tahmin int
	for {

		fmt.Print("Bir sayı gir: ")
		fmt.Scan(&tahmin)

		if tahmin == sayi {
			fmt.Println("Doğru tahmin!")
			break
		}

		fmt.Println("Yanlış, tekrar dene.")
	}

}
