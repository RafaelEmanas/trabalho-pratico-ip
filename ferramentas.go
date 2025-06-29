package main

import (
	"math/rand"
)

func CriaMatriz(numLinhas int, numColunas int) [][]int {
	var matriz [][]int
	var i int

	matriz = make([][]int, numLinhas)
	for i=0;i<numLinhas;i++ {
		matriz[i] = make([]int, numColunas)
	}
	return matriz
}

func IniciaMatrizAleatoria(mat [][]int) [][]int {

	var i, j int
	var tamanhoMat, quadradoTamanhoMat int

	tamanhoMat = len(mat)
	quadradoTamanhoMat = len(mat) * len(mat)

	for i = 0; i < tamanhoMat; i++ {
		for j = 0; j < tamanhoMat; j++ {
			mat[i][j] = rand.Intn(quadradoTamanhoMat/2 + 1)
		}
	}

	return mat
}

func CopiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int) {
	var contAi, contAj, contBi, contBj, temp, numL, numC int
	numL = len(menor)
	numC = len(menor[0])

	//Ai e Aj são os índices da matriz maior
	//Bi e Bj são os índices da matriz menor

	contAi = 0
	for contBi = 0; contBi < numL; contBi++ {

		//se for a linha que não poode, pula
		if contAi == isqn {
			contAi++
		}
		contAj = 0
		for contBj = 0; contBj < numC; contBj++ {

			//se for a coluna que não pode, pula
			if contAj == jsqn {
				contAj++
			}

			//elemento da matriz menor recebe da maior
			temp = maior[contAi][contAj]
			menor[contBi][contBj] = temp
			contAj++
		}
		contAi++
	}
}

func CalculaSinal(indiceL int, indiceC int) int {
	var sinal int

	sinal = -1
	if ((indiceL + indiceC) % 2) == 0 {
		sinal = 1
	}

	return sinal
}

func VerificaQuadradaOrdem(mat [][]int) (bool, int) {
	var numLinhas int
	var numColunas int
	var ehQuadrada bool

	numLinhas = len(mat)
	numColunas = len(mat[0])

	ehQuadrada = false
	if numLinhas == numColunas {
		ehQuadrada = true
	}

	return ehQuadrada, numLinhas
}

func DetOrdem1(mat [][]int) int {
	return mat[0][0]
}

func DetOrdem2(mat [][]int) int {
	var diagonalP int
	var diagonalI int

	diagonalP = mat[0][0] * mat[1][1]
	diagonalI = mat[1][0] * mat[0][1]
	return (diagonalP - diagonalI)
}
