package main

import (
	"fmt"
	"math/rand"
)

func imprimeMatriz(mat [][]int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			fmt.Print(mat[contI][contJ], " ")
		}
		fmt.Println()
	}
}

func criaMatriz(numLinhas, numColunas int) [][]int {
	var matriz [][]int
	matriz = make([][]int, numLinhas)
	for i := range matriz {
		matriz[i] = make([]int, numColunas)
	}
	return matriz
}

func iniciaMatrizAleatoria(mat [][]int) [][]int {

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

func copiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int) {
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

func calculaSinal(indiceL int, indiceC int) int {
	var sinal int

	sinal = -1
	if ((indiceL + indiceC) % 2) == 0 {
		sinal = 1
	}

	return sinal
}

func verificaQuadradaOrdem(mat [][]int) (bool, int) {
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

func detOrdem1(mat [][]int) int {
	return mat[0][0]
}

func detOrdem2(mat [][]int) int {
	var diagonalP int
	var diagonalI int

	diagonalP = mat[0][0] * mat[1][1]
	diagonalI = mat[1][0] * mat[0][1]
	return (diagonalP - diagonalI)
}

func detOrdemN(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC int
	var matMenor [][]int
	numL = len(mat)
	numC = len(mat[0])

	resposta = 0
	contL = 0
	for contC = 0; contC < numC; contC++ {
		cofator = mat[contL][contC]
		sinal = calculaSinal(contL, contC)

		matMenor = criaMatriz(numL-1,numC-1)

		copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
		detTemp = calculaDeterminanteBaseline(matMenor)
		resposta = resposta + (cofator * sinal * detTemp)
	}

	return resposta
}

func calculaDeterminanteBaseline(mat [][]int) int {
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = verificaQuadradaOrdem(mat)
	det = 0
	if ehQuadrada {
		switch ordem {
			case 1:
				det = detOrdem1(mat)
			case 2:
				det = detOrdem2(mat)
			default:
				det = detOrdemN(mat)
		}

	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}

func retornaFileiraMaisZeros(matriz [][]int)(string,int){

	var i,j int
	var linhaMaisZeros []int
	var colunaMaisZeros []int
	var acumuladorZerosLinha, acumuladorZerosColuna int
	var numL,numC int
	var tipoFileira string
	var indiceFileiraMaisZeros int

	numL = len(matriz)
	numC = len(matriz[0])

	linhaMaisZeros = make([]int, 2)
	colunaMaisZeros = make([]int, 2)

	acumuladorZerosLinha = 0
	acumuladorZerosColuna = 0

	for i=0;i<numL;i++{
		for j=0;j<numC;j++{
			if(matriz[i][j]==0){
				acumuladorZerosLinha++
			}
			if(matriz[j][i]==0){
				acumuladorZerosColuna++
			}
		}
		
		if(acumuladorZerosLinha>linhaMaisZeros[1]){
			linhaMaisZeros[0] = i
			linhaMaisZeros[1] = acumuladorZerosLinha
		}
		if(acumuladorZerosColuna>colunaMaisZeros[1]){
			colunaMaisZeros[0] = i
			colunaMaisZeros[1] = acumuladorZerosColuna
		}

		acumuladorZerosLinha = 0
		acumuladorZerosColuna = 0
	}

	if(linhaMaisZeros[1]>=colunaMaisZeros[1]){
		tipoFileira = "linha"
		indiceFileiraMaisZeros = linhaMaisZeros[0]
	} else{
		tipoFileira = "coluna"
		indiceFileiraMaisZeros = colunaMaisZeros[0]
	}

	return tipoFileira,indiceFileiraMaisZeros

}

func detOrdemNOtimizado(mat [][]int) int{
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC int
	var matMenor [][]int
	var tipoFileiraMaisZeros string
	var indiceFileiraMaisZeros int
	
	numL = len(mat)
	numC = len(mat[0])
	resposta = 0

	tipoFileiraMaisZeros,indiceFileiraMaisZeros = retornaFileiraMaisZeros(mat)

	if(tipoFileiraMaisZeros=="linha"){
		contL = indiceFileiraMaisZeros

		for contC = 0; contC < numC; contC++ {
			cofator = mat[contL][contC]
			if(cofator!=0){
				sinal = calculaSinal(contL, contC)

				matMenor = criaMatriz(numL-1,numC-1)

				copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
				detTemp = calculaDeterminanteOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)
			}
		}
	} else{
		contC = indiceFileiraMaisZeros

		for contL=0;contL<numL;contL++{
			cofator = mat[contL][contC]
			if(cofator!=0){
				sinal = calculaSinal(contL,contC)

				matMenor = criaMatriz(numL-1,numC-1)

				copiaMatrizMaiorParaMenor(mat,matMenor,contL,contC)
				detTemp = calculaDeterminanteOtimizado(matMenor)
				resposta = resposta + (cofator*sinal*detTemp)
			}
		}
	}

	return resposta
}

func calculaDeterminanteOtimizado(mat [][]int) int{
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = verificaQuadradaOrdem(mat)
	det = 0
	if ehQuadrada {
		switch ordem {
			case 1:
				det = detOrdem1(mat)
			case 2:
				det = detOrdem2(mat)
			default:
				det = detOrdemNOtimizado(mat)
		}

	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}

func main() {

	var numLinhas, numColunas int
	var matriz [][]int
	var determinanteBaseline, determinanteOtimizado int

	numColunas = 7
	numLinhas = 7

	matriz = criaMatriz(numLinhas,numColunas)
	matriz = iniciaMatrizAleatoria(matriz)

	determinanteBaseline = calculaDeterminanteBaseline(matriz)
	determinanteOtimizado = calculaDeterminanteOtimizado(matriz)

	imprimeMatriz(matriz)

	fmt.Println("Determinante Baseline:", determinanteBaseline)
	fmt.Println("Determinante Otimizado:",determinanteOtimizado)

}
