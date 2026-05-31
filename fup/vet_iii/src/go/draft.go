package main

import (
	"fmt"
)

// função que formata o vetor
func formatarVetor(v []int) string {
	if len(v) == 0 {
		return "[]"
	}

	resultado := "["

	for i := 0; i < len(v); i++ {
		resultado += fmt.Sprint(v[i])

		// adiciona ", " apenas entre elementos
		if i < len(v)-1 {
			resultado += ", "
		}
	}

	resultado += "]"
	return resultado
}

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Println(formatarVetor(vetor))
}