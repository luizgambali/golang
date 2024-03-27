package main

import "fmt"

func main() {

	//int8, int16, int32, int64 -> inteiros
	//uint8, uint16, uint32, uint64 -> inteiros sem sinal

	//float32, float64 -> ponto flutuante
	//byte, bool, string -> tipos básicos

	//int -> depende da arquitetura do sistema

	//rune -> representa um mapeamento de caracteres da tabela Unicode (int32)

	//complex64, complex128 -> números complexos

	//arrays -> são homogêneos e estáticos

	//slices -> são homogêneos e dinâmicos

	//maps -> são heterogêneos e dinâmicos

	//structs -> são heterogêneos e estáticos

	//ponteiros -> são referências de memória

	inteiro := 10
	var flutuante float64 = 10.45
	var texto string = "Texto"
	var booleano bool = true
	var inteiro8 int8 = 127
	var inteiro16 int16 = 32767
	var inteiro32 int32 = 2147483647
	var inteiro64 int64 = 9223372036854775807
	var inteiroSemSinal8 uint8 = 255
	var inteiroSemSinal16 uint16 = 65535
	var inteiroSemSinal32 uint32 = 4294967295
	var inteiroSemSinal64 uint64 = 18446744073709551615
	var pontoFlutuante32 float32 = 3.40282346638528859811704183484516925440e+38
	var pontoFlutuante64 float64 = 1.797693134862315708145274237317043567981e+308
	var caractere byte = 'A'         //byte é um alias para uint8, vai imprimir o código ASCII correspondente ao caractere!
	var caractereEspecial rune = 'Ç' //rune é um alias para int32, vai imprimir o caractere correspondente ao código Unicode!
	var numeroComplexo64 complex64 = 1 + 2i
	var numeroComplexo128 complex128 = 1 + 2i
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var slice []int = []int{1, 2, 3, 4, 5}
	var mapa map[string]string = map[string]string{
		"chave": "valor",
	}

	fmt.Println("Inteiro:", inteiro)
	fmt.Println("Ponto flutuante:", flutuante)
	fmt.Println("String:", texto)
	fmt.Println("Booleano:", booleano)
	fmt.Println("Inteiro8:", inteiro8)
	fmt.Println("Inteiro16:", inteiro16)
	fmt.Println("Inteiro32:", inteiro32)
	fmt.Println("Inteiro64:", inteiro64)
	fmt.Println("inteiroSemSinal8:", inteiroSemSinal8)
	fmt.Println("inteiroSemSinal16:", inteiroSemSinal16)
	fmt.Println("inteiroSemSinal32:", inteiroSemSinal32)
	fmt.Println("inteiroSemSinal64:", inteiroSemSinal64)
	fmt.Println("pontoFlutuante32:", pontoFlutuante32)
	fmt.Println("pontoFlutuante64:", pontoFlutuante64)
	fmt.Println("caractere:", caractere)                 //byte é um alias para uint8, vai imprimir o código ASCII correspondente ao caractere!
	fmt.Println("caractereEspecial:", caractereEspecial) //rune é um alias para int32, vai imprimir o caractere correspondente ao código Unicode!
	fmt.Println("numeroComplexo64:", numeroComplexo64)
	fmt.Println("numeroComplexo128:", numeroComplexo128)
	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
	fmt.Println("mapa:", mapa)

}
