package main

import "fmt"

func vetor(arr []int, sep string) {
	for i, valor := range arr {
		if i != 0 {
			fmt.Println(sep)
		}
		fmt.Printf("%v", valor)
	}
	fmt.Print("\n")
}

func main() {
	var qtd int
	fmt.Scan(&qtd)
	var arr []int = make([]int, qtd)
	for i := range arr {
		fmt.Scan(&arr[i])

	}
	vetor(arr, "")
}
