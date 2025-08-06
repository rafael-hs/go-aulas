package main

import "fmt"

func main() {
	err, q, r := dividir(5, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(q, r)
}

func dividir(dividendo, divisor int) (err error, quociente, resto int) {
	if divisor == 0 {
		err = fmt.Errorf("Divisão por zero é indefinida!")
		return err, 0, 0
	}
	quociente = dividendo / divisor
	resto = dividendo % divisor

	return err, quociente, resto

}
