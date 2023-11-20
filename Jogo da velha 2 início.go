package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	JogadorX  = "X"
	JogadorO  = "O"
	Empate    = "Empate"
	Ninguem   = "Ninguem"
	VencedorX = "VencedorX"
	VencedorO = "VencedorO"
)

var tabuleiro [3][3]string
var jogadorAtual string

func main() {
	inicializarTabuleiro()
	jogadorAtual = JogadorX

	var resultadoJogo string

	for {
		limparTela()
		imprimirTabuleiro()

		fmt.Printf("É a vez do jogador %s\n", jogadorAtual)
		fmt.Print("Digite a linha (0, 1 ou 2): ")
		linha := obterEntradaUsuario()
		fmt.Print("Digite a coluna (0, 1 ou 2): ")
		coluna := obterEntradaUsuario()

		if movimentoValido(linha, coluna) {
			realizarMovimento(linha, coluna)
			resultadoJogo = verificarResultado()
			if resultadoJogo == VencedorX || resultadoJogo == VencedorO {
				limparTela()
				imprimirTabuleiro()
				fmt.Printf("Parabéns! O jogador %s venceu!\n", resultadoJogo)
				break
			} else if resultadoJogo == Empate {
				limparTela()
				imprimirTabuleiro()
				fmt.Println("O jogo terminou em empate!")
				break
			}

			alternarJogador()
		} else {
			fmt.Println("Movimento inválido. Tente novamente.")
		}
	}
}

func inicializarTabuleiro() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tabuleiro[i][j] = " "
		}
	}
}

func imprimirTabuleiro() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tabuleiro[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -----")
		}
	}
}

func obterEntradaUsuario() int {
	var entrada int
	fmt.Scan(&entrada)
	return entrada
}

func movimentoValido(linha, coluna int) bool {
	if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 {
		return false
	}
	return tabuleiro[linha][coluna] == " "
}

func realizarMovimento(linha, coluna int) {
	tabuleiro[linha][coluna] = jogadorAtual
}

func alternarJogador() {
	if jogadorAtual == "X" {
		jogadorAtual = "O"
	} else {
		jogadorAtual = "X"
	}
}

func verificarEmpate() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tabuleiro[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func verificarResultado() string {
	for i := 0; i < 3; i++ {
		if tabuleiro[i][0] == jogadorAtual && tabuleiro[i][1] == jogadorAtual && tabuleiro[i][2] == jogadorAtual {
			return jogadorAtual
		}
		if tabuleiro[0][i] == jogadorAtual && tabuleiro[1][i] == jogadorAtual && tabuleiro[2][i] == jogadorAtual {
			return jogadorAtual
		}
	}
	if tabuleiro[0][0] == jogadorAtual && tabuleiro[1][1] == jogadorAtual && tabuleiro[2][2] == jogadorAtual {
		return jogadorAtual
	}
	if tabuleiro[0][2] == jogadorAtual && tabuleiro[1][1] == jogadorAtual && tabuleiro[2][0] == jogadorAtual {
		return jogadorAtual
	}
	if verificarEmpate() {
		return Empate
	}
	return Ninguem
}

func limparTela() {
	cmd := exec.Command("clear") // Linux/OSX
	cmd.Stdout = os.Stdout
	cmd.Run()
}
