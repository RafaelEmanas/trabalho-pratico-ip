package main

import (
	"fmt"
	"time"
)

func imprimeDados(matriz [][]int, ordem int, iteracao int, tempoBaseline int64, tempoOtimizado int64){
	fmt.Println("Ordem:", ordem)
	fmt.Println("Iteração:", iteracao)
	ImprimeMatriz(matriz)
	fmt.Println("Tempo Baseline:",tempoBaseline)
	fmt.Println("Tempo Otimizado:", tempoOtimizado)
	fmt.Println()
}

func retornaMedias() ([5]int64, [5]int64) {

	var matriz [][]int
	var inicioBaseline, inicioOtimizado time.Time
	var ordemMatriz int
	var i, j int
	var duracaoBaseline, acumTempoBaseline, duracaoOtimizado, acumTempoOtimizado int64
	var mediasTemposBaseline, mediasTemposOtimizado [5]int64

	acumTempoBaseline = 0
	acumTempoOtimizado = 0

	i = 0
	for ordemMatriz = 3; ordemMatriz <= 11; ordemMatriz += 2 {
		for j = 0; j < 3; j++ {

			matriz = CriaMatriz(ordemMatriz, ordemMatriz)
			matriz = IniciaMatrizAleatoria(matriz)

			inicioBaseline = time.Now()
			CalculaDeterminanteBaseline(matriz)
			duracaoBaseline = time.Since(inicioBaseline).Nanoseconds()
			acumTempoBaseline = acumTempoBaseline + duracaoBaseline

			inicioOtimizado = time.Now()
			CalculaDeterminanteOtimizado(matriz)
			duracaoOtimizado = time.Since(inicioOtimizado).Nanoseconds()
			acumTempoOtimizado = acumTempoOtimizado + duracaoOtimizado

			imprimeDados(matriz, ordemMatriz,j,duracaoBaseline,duracaoOtimizado)
		}

		mediasTemposBaseline[i] = acumTempoBaseline / 3
		mediasTemposOtimizado[i] = acumTempoOtimizado / 3

		acumTempoBaseline = 0
		acumTempoOtimizado = 0
		i++
	}

	return mediasTemposBaseline, mediasTemposOtimizado
}

func main() {

	var mediasTemposBaseline, mediasTemposOtimizado [5]int64

	mediasTemposBaseline, mediasTemposOtimizado = retornaMedias()

	fmt.Println("Média Tempos Det Baseline:", mediasTemposBaseline)
	fmt.Println("Média Tempos Det Otimizado:", mediasTemposOtimizado)
}
