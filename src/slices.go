package main

import "fmt"

func main()  {
	primes := [6]int{1, 3, 5, 4, 8, 9}
	var s []int = primes[0:4]
	fmt.Println(s)
}
