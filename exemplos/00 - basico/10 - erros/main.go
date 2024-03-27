package main

import "errors"

func main() {

	resultado, err := divisao(10, 2)

	if err != nil {
		panic(err)
	}

	println(resultado)

}

func divisao(x, y int) (float32, error) {

	if y == 0 {
		return 0, errors.New("Nao e possivel dividir por zero")
	}

	return float32(x / y), nil
}
