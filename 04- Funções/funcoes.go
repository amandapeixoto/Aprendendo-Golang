package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//mais de um retorno por função por connta desse (int8, int8)
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(25, 95)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	var r = func(txt string) {
		fmt.Println(txt)
	}

	r("Texto da função r")

	var t = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := t("Texto da função t")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(25, 95)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//quando quiser usar apenas uma das funções utilizar o underline
	resultadoSoma1, _ := calculosMatematicos(25, 95)
	fmt.Println(resultadoSoma1)

}
