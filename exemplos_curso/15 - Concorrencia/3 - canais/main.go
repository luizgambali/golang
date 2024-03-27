package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //cria um canal de string

	//goroutine: vai entrar na rotina, mas não vai esperar o fim de sua execução
	go escrever("Olá mundo", canal)

	//vai imprimir essa mensagem
	fmt.Println("Mensagem logo após passar pela rotina go escrever")

	//OPÇÃO 1: fazer um loop infinito e testar as mensagens

	//loop infinito para testes; aqui, ele vai ficar esperando o canal receber
	//as mensagens enviadas lá dentro de escrever... a medida que as mensagens
	//vão chegando, ele vai mostrando na tela. Lá dentro de escrever, tem um
	//close(canal), que vai parar de enviar mensagens para o canal. Isso é pego
	//aqui no for pela segunda variavel, 'aberto', que verifica se o canal está
	//ainda aberto para enviar/receber mensagens. Se estiver aberto, executa o
	//print logo abaixo, senão sai do for com o break

	// for {
	// 	//recupera o valor do canal. Prestar atenção na sintaxe
	// 	mensagem, aberto := <-canal

	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	//OPÇÃO 2: for range. MUITO MAIS ELEGANTE! Enquanto o canal estiver aberto,
	//ele fica no for, se não estiver mais aberto, já sai fora

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

//recebe o canal de comunicação de mensagens (aparentemente por referencia... mas sem ser ponteiro? estranho hein)
func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		//insere um dado no canal. Usa-se a sintaxe abaixo
		// canal recebe valor: canal <- valor
		// recupara valor de um canal: <- canal
		canal <- texto
		time.Sleep(time.Second)
	}

	//fecha o canal de comunicação
	close(canal)
}
