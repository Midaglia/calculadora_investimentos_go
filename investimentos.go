//Projeto de Investimentos em Go, baseado em renda fixa e tempo de investimento. Início 20.11.23 - V 0.01

package main

import "fmt"

func main() {

	var invest_inicial, valor_aport, invest_selic_anual, invest_final float32
	var invest_meses int
	// const ano = 12 // 12 meses dentro de um ano

	fmt.Println("Digite o valor inicial do investimento: ")
	fmt.Scan(&invest_inicial)

	fmt.Println("Digite os aportes mensais que você pretende fazer: ")
	fmt.Scan(&valor_aport)

	fmt.Println("Digite por quantos meses você pretende investir: ")
	fmt.Scan(&invest_meses)

	fmt.Println("Informe o valor da SELIC atual: ")
	fmt.Scan(&invest_selic_anual)

	valor_inicial := (invest_inicial * invest_selic_anual / 100) / 12
	valor_inicial = invest_inicial + valor_inicial
	lucro_mensal := (invest_inicial * invest_selic_anual / 100) / 12
	rendimentos_totais := lucro_mensal
	fmt.Println("1 º Mês | Valor acumulado: R$", valor_inicial, "| Lucro mensal: R$", lucro_mensal, "| Rendimento total: R$", rendimentos_totais)

	invest_final = valor_inicial + valor_aport

	contador := 2

	for contador <= invest_meses {

		lucro_mensal = ((invest_final * invest_selic_anual / 100) / 12)
		rendimentos_totais = rendimentos_totais + lucro_mensal

		invest_final = ((invest_final * invest_selic_anual / 100) / 12) + invest_final

		fmt.Println(contador, "º Mês | Valor acumulado: R$", invest_final, "| Lucro mensal: R$", lucro_mensal, "| Rendimentos total: R$", rendimentos_totais)

		invest_final = invest_final + valor_aport

		contador = contador + 1
	}

	var pause string
	fmt.Scan(&pause)

}
