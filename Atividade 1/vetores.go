package main

import "fmt"

// Q001
func somaElementos() int {
	array := []int{1, 2, 3, 4}
	soma := 0
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}
	return soma
}

// Q002
func contaPares() int {
	array := []int{1, 2, 3, 4, 5, 6}
	cont := 0

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			cont++
		}
	}

	return cont
}

// Q003
func maiorNum() int {
	array := []int{3, 7, 2, 9, 5}
	var maior int

	for i := 0; i < len(array); i++ {
		if array[i] > maior {
			maior = array[i]
		}
	}

	return maior
}

// Q004
func contMaiorQueX() int {
	array := []int{1, 5, 8, 2, 10}
	x := 5
	cont := 0

	for i := 0; i < len(array); i++ {
		if array[i] > x {
			cont++
		}
	}

	return cont
}

// Q005
func somaPares() int {
	array := []int{1, 2, 3, 4, 5, 6}
	soma := 0

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			soma += array[i]
		}
	}

	return soma
}

// Q006
func verificaSeX() string {
	array := []int{4, 7, 1, 9}
	x := 7

	for i := 0; i < len(array); i++ {
		if array[i] == x {
			return "verdadeiro"
		}
	}
	return "falso"
}

// Q007
func inverteLista() []int {
	array := []int{1, 2, 3, 4}
	var invert []int
	cont := len(array) - 1

	for i := 0; i < len(array); i++ {
		invert = append(invert, array[cont])
		cont--
	}

	return invert
}

// Q008
func contaOcX() int {
	array := []int{1, 2, 2, 3, 2, 4}
	x := 2
	cont := 0

	for i := 0; i < len(array); i++ {
		if array[i] == x {
			cont++
		}
	}

	return cont
}

// Q009
func somaPositivo() int {
	array := []int{-1, 2, -3, 4, 5}
	soma := 0

	for i := 0; i < len(array); i++ {
		if array[i] > 0 {
			soma += array[i]
		}
	}
	return soma
}

// Q10
func produto() int {
	array := []int{1, 2, 3, 4}
	mult := 1

	for i := 0; i < len(array); i++ {
		mult *= array[i]
	}

	return mult
}

func main() {
	fmt.Println(somaElementos())
	fmt.Println(contaPares())
	fmt.Println(maiorNum())
	fmt.Println(contMaiorQueX())
	fmt.Println(somaPares())
	fmt.Println(verificaSeX())
	fmt.Println(inverteLista())
	fmt.Println(contaOcX())
	fmt.Println(somaPositivo())
	fmt.Println(produto())
}
