package main

import (
	"fmt"
	"strconv"
)

// Q011
func contaVogais() int {
	s := "Programacao"
	char := []rune(s)
	vogais := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	cont := 0

	for i := 0; i < len(char); i++ {
		for j := 0; j < len(vogais); j++ {
			if char[i] == vogais[j] {
				cont++
				break
			}
		}
	}

	return cont
}

// Q012
func contaChars() int {
	s := "teste"
	char := []rune(s)
	cont := 0

	for i := 0; i < len(char); i++ {
		cont++
	}

	return cont
}

// Q013
func verificaPalindromo() []string {
	s := []string{"arara", "casa"}

	char1 := []rune(s[0])
	char2 := []rune(s[1])

	var inv []string

	cont1 := len(char1) - 1
	var sInv1 string

	for i := 0; i < len(char1); i++ {
		sInv1 += string(char1[cont1])
		cont1--
	}
	inv = append(inv, sInv1)

	cont2 := len(char2) - 1
	var sInv2 string

	for i := 0; i < len(char2); i++ {
		sInv2 += string(char2[cont2])
		cont2--
	}
	inv = append(inv, sInv2)

	var resp []string

	if inv[0] == s[0] {
		resp = append(resp, "verdadeiro")
	} else {
		resp = append(resp, "falso")
	}
	if inv[1] == s[1] {
		resp = append(resp, "verdadeiro")
		return resp
	}
	resp = append(resp, "falso")
	return resp
}

// Q014
func contaAparicaoCaractere() int {
	s := "banana"
	c := 'a'
	charsS := []rune(s)
	cont := 0

	for i := 0; i < len(charsS); i++ {
		if charsS[i] == c {
			cont++
		}
	}

	return cont
}

// Q015
func trocaCaractere() string {
	s := "banana"
	c := 'a'
	troca := 'o'
	charsS := []rune(s)
	var novaStr string

	for i := 0; i < len(charsS); i++ {
		if charsS[i] == c {
			novaStr += string(troca)
			continue
		}
		novaStr += string(charsS[i])
	}
	return novaStr
}

// Q016
func contaLetrasMaisculasEMinusculas() string {
	s := "AbCde"
	chars := []rune(s)
	alfabetoMin := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	maiu := 0
	minu := 0
	eMinu := false

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(alfabetoMin); j++ {
			if chars[i] == alfabetoMin[j] {
				eMinu = true
				break
			}
		}
		if eMinu {
			minu++
			eMinu = false
			continue
		}
		maiu++
	}
	return "Maiusculas: " + strconv.Itoa(maiu) + " | Minusculas: " + strconv.Itoa(minu)
}

// Q017
func removerEspacos() string {
	s := "ola mundo"
	chars := []rune(s)
	var novaStr string

	for i := 0; i < len(chars); i++ {
		if chars[i] == ' ' {
			continue
		}
		novaStr += string(chars[i])
	}
	return novaStr
}

// Q018
func primeiroChar(s string) string {
	if s == "" {
		return ""
	}

	chars := []rune(s)

	return string(chars[0])
}

func main() {
	fmt.Println(contaVogais())
	fmt.Println(contaChars())
	fmt.Println(verificaPalindromo())
	fmt.Println(contaAparicaoCaractere())
	fmt.Println(trocaCaractere())
	fmt.Println(contaLetrasMaisculasEMinusculas())
	fmt.Println(removerEspacos())
	fmt.Println(primeiroChar("python"))
}
