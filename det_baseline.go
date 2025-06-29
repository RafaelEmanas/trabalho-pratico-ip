package main

import "fmt"

func detOrdemN(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC int
	var matMenor [][]int
	numL = len(mat)
	numC = len(mat[0])

	resposta = 0
	contL = 0
	for contC = 0; contC < numC; contC++ {
		cofator = mat[contL][contC]
		sinal = CalculaSinal(contL, contC)

		matMenor = CriaMatriz(numL-1, numC-1)

		CopiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
		detTemp = CalculaDeterminanteBaseline(matMenor)
		resposta = resposta + (cofator * sinal * detTemp)
	}

	return resposta
}

func CalculaDeterminanteBaseline(mat [][]int) int {
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = VerificaQuadradaOrdem(mat)
	det = 0
	if ehQuadrada {
		switch ordem {
			case 1:
				det = DetOrdem1(mat)
			case 2:
				det = DetOrdem2(mat)
			default:
				det = detOrdemN(mat)
		}

	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}
