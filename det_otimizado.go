package main

import "fmt"

type FileiraMaisZeros struct{
	indice int
	quantidadeZeros int
}

func retornaFileiraMaisZeros(matriz [][]int) (string, int) {

	var i, j int
	var linhaMaisZeros, colunaMaisZeros FileiraMaisZeros
	var acumuladorZerosLinha, acumuladorZerosColuna int
	var numL, numC int
	var tipoFileira string
	var indiceFileiraMaisZeros int

	numL = len(matriz)
	numC = len(matriz[0])

	linhaMaisZeros = FileiraMaisZeros{
		indice: 0,
		quantidadeZeros: 0,
	}
	colunaMaisZeros = FileiraMaisZeros{
		indice: 0,
		quantidadeZeros: 0,
	}

	acumuladorZerosLinha = 0
	acumuladorZerosColuna = 0

	for i = 0; i < numL; i++ {
		for j = 0; j < numC; j++ {
			if matriz[i][j] == 0 {
				acumuladorZerosLinha++
			}
			if matriz[j][i] == 0 {
				acumuladorZerosColuna++
			}
		}

		if acumuladorZerosLinha > linhaMaisZeros.quantidadeZeros {
			linhaMaisZeros.indice = i
			linhaMaisZeros.quantidadeZeros = acumuladorZerosLinha
		}
		if acumuladorZerosColuna > colunaMaisZeros.quantidadeZeros {
			colunaMaisZeros.indice = i
			colunaMaisZeros.quantidadeZeros = acumuladorZerosColuna
		}

		acumuladorZerosLinha = 0
		acumuladorZerosColuna = 0
	}

	if linhaMaisZeros.quantidadeZeros >= colunaMaisZeros.quantidadeZeros {
		tipoFileira = "linha"
		indiceFileiraMaisZeros = linhaMaisZeros.indice
	} else {
		tipoFileira = "coluna"
		indiceFileiraMaisZeros = colunaMaisZeros.indice
	}

	return tipoFileira, indiceFileiraMaisZeros

}

func detOrdemNOtimizado(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC int
	var matMenor [][]int
	var tipoFileiraMaisZeros string
	var indiceFileiraMaisZeros int

	numL = len(mat)
	numC = len(mat[0])
	resposta = 0

	tipoFileiraMaisZeros, indiceFileiraMaisZeros = retornaFileiraMaisZeros(mat)

	if tipoFileiraMaisZeros == "linha" {
		contL = indiceFileiraMaisZeros

		for contC = 0; contC < numC; contC++ {
			cofator = mat[contL][contC]
			if cofator != 0 {
				sinal = CalculaSinal(contL, contC)

				matMenor = CriaMatriz(numL-1, numC-1)

				CopiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
				detTemp = CalculaDeterminanteOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)
			}
		}
	} else {
		contC = indiceFileiraMaisZeros

		for contL = 0; contL < numL; contL++ {
			cofator = mat[contL][contC]
			if cofator != 0 {
				sinal = CalculaSinal(contL, contC)

				matMenor = CriaMatriz(numL-1, numC-1)

				CopiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
				detTemp = CalculaDeterminanteOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)
			}
		}
	}

	return resposta
}

func CalculaDeterminanteOtimizado(mat [][]int) int {
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
			det = detOrdemNOtimizado(mat)
		}

	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}
