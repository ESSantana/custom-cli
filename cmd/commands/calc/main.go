package calc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var operation string

var calcCommand = &cobra.Command{
	Use:     "calcular",
	Aliases: []string{"calc"},
	Short:   "Este comando executa uma operacao matematica",
	Long:    "Este comando executa uma operacao matematica usando os dois valores fornecidos",
	Example: `custom-cli calcular --operacao="adicao" 4 5`,
	Run:     Calculate,
}

func init() {
	calcCommand.Flags().StringVarP(&operation, "operacao", "o", "", "Define a operacao do calculo")
	calcCommand.MarkFlagRequired("operacao")
}

func Command() *cobra.Command {
	return calcCommand
}

func Calculate(_ *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Printf("Comando precisa de dois argumentos, mas recebeu %v\n", len(args))
		return
	}

	x, y := args[0], args[1]
	floatX, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Println("Erro ao converter o argumento fornecido")
		return
	}

	floatY, err := strconv.ParseFloat(y, 64)
	if err != nil {
		fmt.Println("Erro ao converter o argumento fornecido")
		return
	}

	op := strings.ToLower(operation)

	var result float64
	switch op {
	case "adicao":
		result = floatX + floatY
	case "subtracao":
		result = floatX - floatY
	case "divisao":
		result = floatX / floatY
	case "multiplicacao":
		result = floatX * floatY
	default:
		fmt.Printf("Operacao %s nao suportada\n", op)
		return
	}

	fmt.Printf("O resultado do calculo da operacao de %s entre os numeros %v e %v foi igual a %v\n", op, floatX, floatY, result)
}
