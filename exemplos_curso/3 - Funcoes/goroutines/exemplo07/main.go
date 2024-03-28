package main

func main() {

	//canal é uma variável do tipo channel que vai receber um int
	//ela vai permitir a troca de mensagens entre as goroutines
	var canal = make(chan int)

	//funcao anonima
	go func() {
		canal <- 42
	}()

	go func() {
		canal <- 43
	}()

	x := <-canal

	println(x)
}
