package main

import (
	"fmt"
	
)


type Contato struct {
	nome  string
	email string
}


func (c *Contato) alterarEmail(novoEmail string) {
	c.email = novoEmail
}


func adicionarContato(contato Contato, listaContatos *[5]Contato) {
	for i := range listaContatos {
		if listaContatos[i].nome == "" && listaContatos[i].email == "" {
			listaContatos[i] = contato
			fmt.Println("Contato adicionado com sucesso!")
			return
		}
	}
	fmt.Println("Não há espaço disponível para adicionar o contato.")
}


func excluirContato(listaContatos *[5]Contato) {
	for i := len(*listaContatos) - 1; i >= 0; i-- {
		if listaContatos[i].nome != "" || listaContatos[i].email != "" {
			listaContatos[i] = Contato{} 
			return
		}
	}
	fmt.Println("Não há contatos para excluir.")
}

func main() {
	var listaContatos [5]Contato

	for {
		fmt.Println("\nSelecione a operação:")
		fmt.Println("1. Adicionar Contato")
		fmt.Println("2. Excluir Último Contato")
		fmt.Println("3. Sair")

		var escolha int
		fmt.Print("Escolha: ")
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			var nome, email string
			fmt.Print("Nome do Contato: ")
			fmt.Scanln(&nome)
			fmt.Print("E-mail do Contato: ")
			fmt.Scanln(&email)

			novoContato := Contato{nome, email}
			adicionarContato(novoContato, &listaContatos)

		case 2:
			excluirContato(&listaContatos)
			fmt.Println("Último contato excluído com sucesso!")
		case 3:
			fmt.Println("Saindo.")
			break
		default:
			fmt.Println("Escolha inválida. Tente novamente.")
		}

		fmt.Println("\nLista de contatos atualizada:")
		for _, contato := range listaContatos {
			if contato.nome != "" || contato.email != "" {
				fmt.Println("Nome:", contato.nome, "- Email:", contato.email)
			}
		}
	}
}
