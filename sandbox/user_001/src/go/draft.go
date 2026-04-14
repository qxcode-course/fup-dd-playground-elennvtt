package main

import "fmt"

func main() {
	//var qtd int
	//fmt.Scan(&qtd)
	//var idades []int = make([]int, qtd)
	//fmt.Println(idades) 
    arr := []int{9,8,4,5,6,8,3}
    fmt.Print("[")
    for _, valor := range arr {
        fmt.Printf("%v", valor)

    }
    fmt.Print("]\n")
}
