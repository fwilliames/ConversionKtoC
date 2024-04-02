package main

import "fmt"

// Valor da constante de conversao de ºK para ºC
const conversionConst float64 = 273.0

func main() {

	//Valor do ponto de ebulição da agua em ºK
	var ebulicaoK float64 = 373.15

	//Variavel que recebe o valor da conversao de ºK para ºC
	var ebulicaoC float64 = ebulicaoK - conversionConst

	//Imprimindo no Terminal o resultado da conversao
	fmt.Printf("O valor do ponto de ebulição em ºK é: %.2f e em ºC é: %.2f\n", ebulicaoK, ebulicaoC)
}
