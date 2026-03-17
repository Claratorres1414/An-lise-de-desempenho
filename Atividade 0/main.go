package main

import (
	"fmt"
	"strconv"
)

func primo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primosAteN(n int) []int {
	var primos []int
	if n == 2 {
		return []int{2}
	}
	for i := 2; i < n; i++ {
		if primo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}

func main() {
	var n string

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	num, err := strconv.Atoi(n)

	if err == nil && num > 0 {
		fmt.Println("ok")
		if num == 1 {
			fmt.Println("Não existem números primos neste intervalo")
			return
		}
	} else if err != nil {
		fmt.Println("Não é um número")
		return
	} else {
		fmt.Println("Número menor que 1")
		fmt.Println("Não existem números primos neste intervalo")
		return
	}

	primos := primosAteN(num)

	for i, primo := range primos {
		if i == len(primos)-1 {
			fmt.Print(primo)
		} else {
			fmt.Print(primo, " - ")
		}
	}
}
