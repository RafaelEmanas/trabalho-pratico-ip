package main

import (
	"fmt"
	"os"
	"time"
	"text/tabwriter"
)

type Resultado struct{
	Ordem int
	TempoBaseline int64
	TempoOtimizado int64
}

func imprimeMatrizEmArquivo(file *os.File, mat [][]int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			fmt.Fprint(file,mat[contI][contJ], " ")
		}
		fmt.Fprintln(file)
	}
}

func imprimeExperimentoEmArquivo(file *os.File,matriz [][]int ,ordemMatriz int, iteracao int, tempoBaseline int64, tempoOtimizado int64){

	fmt.Fprintln(file,"Ordem:", ordemMatriz)
	fmt.Fprintln(file,"Iteração:",iteracao)
	imprimeMatrizEmArquivo(file, matriz)
	fmt.Fprintln(file,"Tempo Baseline:", tempoBaseline)
	fmt.Fprintln(file,"Tempo Otimizado:", tempoOtimizado)
	fmt.Fprintln(file)
}

func inicializaTabWriter(formatacao string) *tabwriter.Writer{
	var writer *tabwriter.Writer

	writer = tabwriter.NewWriter(os.Stdout,0,0,2,' ',0)
	fmt.Fprintln(writer, formatacao)
	
	return writer
}

func inicializaArquivoExperimento(nomeArquivo string) (*os.File){
	var file *os.File
	var err error

	file, err = os.OpenFile(nomeArquivo, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }

	return file
}

func retornaTempoDeterminanteBaseline(matriz [][]int) int64{
	var inicioBaseline time.Time
	var duracaoBaseline int64

	inicioBaseline = time.Now()
	CalculaDeterminanteBaseline(matriz)
	duracaoBaseline = time.Since(inicioBaseline).Nanoseconds()

	return duracaoBaseline
}

func retornaTempoDeterminanteOtimizado(matriz [][]int) int64{
	var inicioOtimizado time.Time
	var duracaoOtimizado int64

	inicioOtimizado = time.Now()
	CalculaDeterminanteOtimizado(matriz)
	duracaoOtimizado = time.Since(inicioOtimizado).Nanoseconds()

	return duracaoOtimizado
}

func imprimeResultadoTerminal(resultados []Resultado, writer *tabwriter.Writer) {

	var diferencaTempos int64
	var porcentagemDiferencaTempos float64
	var elementoResultado Resultado
	var i int


	for i=0;i<len(resultados);i++{
		elementoResultado = resultados[i]

		diferencaTempos = elementoResultado.TempoBaseline - elementoResultado.TempoOtimizado
		porcentagemDiferencaTempos = (float64(diferencaTempos) / float64(elementoResultado.TempoBaseline))*100

		fmt.Fprintf(writer,"%d\t%d\t%d\t%d\t%.2f%%\n",
			elementoResultado.Ordem,
			elementoResultado.TempoBaseline,
			elementoResultado.TempoOtimizado,
			diferencaTempos,
			porcentagemDiferencaTempos)
	}
	
	writer.Flush()
}

func realizaExperimento(){

	var matriz [][]int
	var ordemMatriz int
	var i, j int
	var duracaoBaseline, acumTempoBaseline, duracaoOtimizado, acumTempoOtimizado int64
	var resultados [5]Resultado
	var mediaTempoBaseline, mediaTempoOtimizado int64

	var arquivoExperimento *os.File
	var writer *tabwriter.Writer

	acumTempoBaseline = 0
	acumTempoOtimizado = 0

	arquivoExperimento = inicializaArquivoExperimento("experimento.txt")
	writer = inicializaTabWriter("Ordem\tTempo Baseline (ns)\tTempo Otimizado (ns)\tDiferença (ns)\t% Diferença")

	i = 0
	for ordemMatriz = 3; ordemMatriz <= 11; ordemMatriz += 2 {
		for j = 0; j < 3; j++ {

			matriz = CriaMatriz(ordemMatriz, ordemMatriz)
			matriz = IniciaMatrizAleatoria(matriz)

			duracaoBaseline = retornaTempoDeterminanteBaseline(matriz)
			acumTempoBaseline = acumTempoBaseline + duracaoBaseline

			duracaoOtimizado = retornaTempoDeterminanteOtimizado(matriz)
			acumTempoOtimizado = acumTempoOtimizado + duracaoOtimizado

			imprimeExperimentoEmArquivo(arquivoExperimento,matriz,ordemMatriz,j,duracaoBaseline,duracaoOtimizado)
		}

		mediaTempoBaseline = acumTempoBaseline/3
		mediaTempoOtimizado = acumTempoOtimizado/3

		resultados[i] = Resultado{
			Ordem: ordemMatriz,
			TempoBaseline: mediaTempoBaseline,
			TempoOtimizado: mediaTempoOtimizado,
		}

		acumTempoBaseline = 0
		acumTempoOtimizado = 0
		i++
	}

	//[:] cria uma slice para que a função possa aceitar a array resultados
	imprimeResultadoTerminal(resultados[:], writer)
	defer arquivoExperimento.Close()
}

func main() {

	realizaExperimento()

}
