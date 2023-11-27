package main

import (
	"fmt"
	"strings"
	"time"
)

var tabuleirop [3][3]string
var resposta string
var jogadorAtual string
var tab1 [3][3]string
var tab2 [3][3]string
var tab3 [3][3]string
var tab4 [3][3]string
var tab5 [3][3]string
var tab6 [3][3]string
var tab7 [3][3]string
var tab8 [3][3]string
var tab9 [3][3]string
var linha, coluna int
var principal string
var miniempate string
var vencedor string
var escolhajogador string
var novo string

func inicio() bool {
	for {
		fmt.Print("\nVocê conhece o famoso jogo da velha 2.0?? Caso não conheca e desaja conhece-lo estarei deixando um link abaixo de como ele funciona.")
		fmt.Print("\n\nhttps://www.tiktok.com/@clebitoyt/video/7280661786972491013\n\n")
		fmt.Print("Agora que já conhece o jogo, deseja joga-lo? (Vazio e espeços serão contados como Não querer jogar) \nR:")
		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta) // Converta a resposta para minúsculas
		if resposta == "sim" || resposta == "y" || resposta == "yes" || resposta == "s" {
			fmt.Print("\nE vamos de\n")
			return true
		}
		if resposta == "não" || resposta == "n" || resposta == "no" || resposta == "nao" {
			fmt.Print("\nAté a próxima então!\n")
			return false
		}
		fmt.Print("\nLarga de palhaçada e digida SIM ou NÃO pf.\n")
	}
}

func inicializarTabuleirop() {
	tabuleirop = [3][3]string{{"#1", "#2", "#3"}, {"#4", "#5", "#6"}, {"#7", "#8", "#9"}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab1[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab2[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab3[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab4[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab5[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab6[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab7[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab8[i][j] = " "
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tab9[i][j] = " "
		}
	}
}

func imprimirTabuleirop() {
	for i := 0; i < 3; i++ {
		fmt.Printf("          ")
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tabuleirop[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("          --------")
		}
	}
}

func imprimirTab1() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO SUPERIOR ESQUERDO-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab1[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}

func imprimirTab2() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO SUPERIOR CENTRAL-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab2[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}

func imprimirTab3() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO SUPERIOR DIREITO-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab3[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}

func imprimirTab4() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO MEIO ESQUERDO-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab4[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}

func imprimirTab5() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO MEIO CENTRAL-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab5[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}

func imprimirTab6() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO MEIO DIREITO-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab6[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}

func imprimirTab7() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO INFERIOR ESQUERDO-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab7[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}

}
func imprimirTab8() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO INFERIOR CENTRAL-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab8[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}
func imprimirTab9() {
	fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
	fmt.Print("\n -----JOGO INFERIOR DIREITO-----\n")
	fmt.Println("            0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("          %d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", tab9[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("            ------")
		}
	}
}
func vereficarvencedorreal() bool {
	for i := 0; i < 3; i++ {
		if tabuleirop[i][0] == tabuleirop[i][1] && tabuleirop[i][1] == tabuleirop[i][2] {
			return true
		}
		if tabuleirop[i][0] == tabuleirop[i][1] && tabuleirop[i][2] == "\u001B[35mØ\u001B[0m" {
			return true
		}
		if tabuleirop[i][2] == tabuleirop[i][1] && tabuleirop[i][0] == "\u001B[35mØ\u001B[0m" {
			return true
		}
		if tabuleirop[i][2] == tabuleirop[i][0] && tabuleirop[i][1] == "\u001B[35mØ\u001B[0m" {
			return true
		}
		if tabuleirop[0][i] == tabuleirop[1][i] && tabuleirop[1][i] == tabuleirop[2][i] {
			return true
		}
		if tabuleirop[0][i] == tabuleirop[1][i] && tabuleirop[2][i] == "\u001B[35mØ\u001B[0m" {
			return true
		}
		if tabuleirop[2][i] == tabuleirop[1][i] && tabuleirop[0][i] == "\u001B[35mØ\u001B[0m" {
			return true
		}
		if tabuleirop[0][i] == tabuleirop[2][i] && tabuleirop[1][i] == "\u001B[35mØ\u001B[0m" {
			return true
		}
	}
	if tabuleirop[0][0] == tabuleirop[1][1] && tabuleirop[1][1] == tabuleirop[2][2] {
		return true
	}
	if tabuleirop[0][0] == tabuleirop[1][1] && tabuleirop[2][2] == "\u001B[35mØ\u001B[0m" {
		return true
	}
	if tabuleirop[0][0] == tabuleirop[2][2] && tabuleirop[1][1] == "\u001B[35mØ\u001B[0m" {
		return true
	}
	if tabuleirop[2][2] == tabuleirop[1][1] && tabuleirop[0][0] == "\u001B[35mØ\u001B[0m" {
		return true
	}
	if tabuleirop[0][2] == tabuleirop[1][1] && tabuleirop[1][1] == tabuleirop[2][0] {
		return true
	}
	if tabuleirop[0][2] == tabuleirop[1][1] && tabuleirop[2][0] == "\u001B[35mØ\u001B[0m" {
		return true
	}
	if tabuleirop[2][0] == tabuleirop[1][1] && tabuleirop[0][2] == "\u001B[35mØ\u001B[0m" {
		return true
	}
	if tabuleirop[2][0] == tabuleirop[0][2] && tabuleirop[1][1] == "\u001B[35mØ\u001B[0m" {
		return true
	}
	return false
}

func vereficarvencedor1() bool {
	for i := 0; i < 3; i++ {
		if tab1[i][0] != " " && tab1[i][0] == tab1[i][1] && tab1[i][1] == tab1[i][2] {
			return true
		}
		if tab1[0][i] != " " && tab1[0][i] == tab1[1][i] && tab1[1][i] == tab1[2][i] {
			return true
		}
	}
	if tab1[0][0] != " " && tab1[0][0] == tab1[1][1] && tab1[1][1] == tab1[2][2] {
		return true
	}
	if tab1[0][2] != " " && tab1[0][2] == tab1[1][1] && tab1[1][1] == tab1[2][0] {
		return true
	}
	return false
}
func vereficarvencedor2() bool {
	for i := 0; i < 3; i++ {
		if tab2[i][0] != " " && tab2[i][0] == tab2[i][1] && tab2[i][1] == tab2[i][2] {
			return true
		}
		if tab2[0][i] != " " && tab2[0][i] == tab2[1][i] && tab2[1][i] == tab2[2][i] {
			return true
		}
	}
	if tab2[0][0] != " " && tab2[0][0] == tab2[1][1] && tab2[1][1] == tab2[2][2] {
		return true
	}
	if tab2[0][2] != " " && tab2[0][2] == tab2[1][1] && tab2[1][1] == tab2[2][0] {
		return true
	}
	return false
}
func vereficarvencedor3() bool {
	for i := 0; i < 3; i++ {
		if tab3[i][0] != " " && tab3[i][0] == tab3[i][1] && tab3[i][1] == tab3[i][2] {
			return true
		}
		if tab3[0][i] != " " && tab3[0][i] == tab3[1][i] && tab3[1][i] == tab3[2][i] {
			return true
		}
	}
	if tab3[0][0] != " " && tab3[0][0] == tab3[1][1] && tab3[1][1] == tab3[2][2] {
		return true
	}
	if tab3[0][2] != " " && tab3[0][2] == tab3[1][1] && tab3[1][1] == tab3[2][0] {
		return true
	}
	return false
}
func vereficarvencedor4() bool {
	for i := 0; i < 3; i++ {
		if tab4[i][0] != " " && tab4[i][0] == tab4[i][1] && tab4[i][1] == tab4[i][2] {
			return true
		}
		if tab4[0][i] != " " && tab4[0][i] == tab4[1][i] && tab4[1][i] == tab4[2][i] {
			return true
		}
	}
	if tab4[0][0] != " " && tab4[0][0] == tab4[1][1] && tab4[1][1] == tab4[2][2] {
		return true
	}
	if tab4[0][2] != " " && tab4[0][2] == tab4[1][1] && tab4[1][1] == tab4[2][0] {
		return true
	}
	return false
}
func vereficarvencedor5() bool {
	for i := 0; i < 3; i++ {
		if tab5[i][0] != " " && tab5[i][0] == tab5[i][1] && tab5[i][1] == tab5[i][2] {
			return true
		}
		if tab5[0][i] != " " && tab5[0][i] == tab5[1][i] && tab5[1][i] == tab5[2][i] {
			return true
		}
	}
	if tab5[0][0] != " " && tab5[0][0] == tab5[1][1] && tab5[1][1] == tab5[2][2] {
		return true
	}
	if tab5[0][2] != " " && tab5[0][2] == tab5[1][1] && tab5[1][1] == tab5[2][0] {
		return true
	}
	return false
}
func vereficarvencedor6() bool {
	for i := 0; i < 3; i++ {
		if tab6[i][0] != " " && tab6[i][0] == tab6[i][1] && tab6[i][1] == tab6[i][2] {
			return true
		}
		if tab6[0][i] != " " && tab6[0][i] == tab6[1][i] && tab6[1][i] == tab6[2][i] {
			return true
		}
	}
	if tab6[0][0] != " " && tab6[0][0] == tab6[1][1] && tab6[1][1] == tab6[2][2] {
		return true
	}
	if tab6[0][2] != " " && tab6[0][2] == tab6[1][1] && tab6[1][1] == tab6[2][0] {
		return true
	}
	return false
}
func vereficarvencedor7() bool {
	for i := 0; i < 3; i++ {
		if tab7[i][0] != " " && tab7[i][0] == tab7[i][1] && tab7[i][1] == tab7[i][2] {
			return true
		}
		if tab7[0][i] != " " && tab7[0][i] == tab7[1][i] && tab7[1][i] == tab7[2][i] {
			return true
		}
	}
	if tab7[0][0] != " " && tab7[0][0] == tab7[1][1] && tab7[1][1] == tab7[2][2] {
		return true
	}
	if tab7[0][2] != " " && tab7[0][2] == tab7[1][1] && tab7[1][1] == tab7[2][0] {
		return true
	}
	return false
}
func vereficarvencedor8() bool {
	for i := 0; i < 3; i++ {
		if tab8[i][0] != " " && tab8[i][0] == tab8[i][1] && tab8[i][1] == tab8[i][2] {
			return true
		}
		if tab8[0][i] != " " && tab8[0][i] == tab8[1][i] && tab8[1][i] == tab8[2][i] {
			return true
		}
	}
	if tab8[0][0] != " " && tab8[0][0] == tab8[1][1] && tab8[1][1] == tab8[2][2] {
		return true
	}
	if tab8[0][2] != " " && tab8[0][2] == tab8[1][1] && tab8[1][1] == tab8[2][0] {
		return true
	}
	return false
}
func vereficarvencedor9() bool {
	for i := 0; i < 3; i++ {
		if tab9[i][0] != " " && tab9[i][0] == tab9[i][1] && tab9[i][1] == tab9[i][2] {
			return true
		}
		if tab9[0][i] != " " && tab9[0][i] == tab9[1][i] && tab9[1][i] == tab9[2][i] {
			return true
		}
	}
	if tab9[0][0] != " " && tab9[0][0] == tab9[1][1] && tab9[1][1] == tab9[2][2] {
		return true
	}
	if tab9[0][2] != " " && tab9[0][2] == tab9[1][1] && tab9[1][1] == tab9[2][0] {
		return true
	}
	return false
}
func EMPATE1() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab1[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE2() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab2[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE3() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab3[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE4() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab4[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE5() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab5[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE6() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab6[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE7() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab7[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE8() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab8[i][j] == " " {
				return false
			}
		}
	}
	return true
}
func EMPATE9() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab9[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func movimentopodetp(principal string) bool {
	for {
		if principal == "#1" || principal == "#2" || principal == "#3" || principal == "#4" || principal == "#5" || principal == "#6" || principal == "#7" || principal == "#8" || principal == "#9" {
			return true
		}
		if principal == "1" || principal == "2" || principal == "3" || principal == "4" || principal == "5" || principal == "6" || principal == "7" || principal == "8" || principal == "9" {
			return true
		}
		return false
	}
}

func movimentopode(linha, coluna int) bool {
	if linha < 0 || linha > 2 || coluna < 0 || coluna > 2 {
		return false
	}
	return true
}

func realizarMovimentop(principal string) {
	switch principal {
	case "#1", "1":
		imprimirTab1()
	case "#2", "2":
		imprimirTab2()
	case "#3", "3":
		imprimirTab3()
	case "#4", "4":
		imprimirTab4()
	case "#5", "5":
		imprimirTab5()
	case "#6", "6":
		imprimirTab6()
	case "#7", "7":
		imprimirTab7()
	case "#8", "8":
		imprimirTab8()
	case "#9", "9":
		imprimirTab9()
	}
}

func realizamovimentopode(principal string) {
	switch principal {
	case "#1", "1":
		tab1[linha][coluna] = jogadorAtual
		if vereficarvencedor1() {
			tabuleirop[0][0] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE1() {
			tabuleirop[0][0] = miniempate
			imprimirTabuleirop()
		}
	case "#2", "2":
		tab2[linha][coluna] = jogadorAtual
		if vereficarvencedor2() {
			tabuleirop[0][1] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE2() {
			tabuleirop[0][1] = miniempate
			imprimirTabuleirop()
		}
	case "#3", "3":
		tab3[linha][coluna] = jogadorAtual
		if vereficarvencedor3() {
			tabuleirop[0][2] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE3() {
			tabuleirop[0][2] = miniempate
			imprimirTabuleirop()
		}
	case "#4", "4":
		tab4[linha][coluna] = jogadorAtual
		if vereficarvencedor4() {
			tabuleirop[1][0] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE4() {
			tabuleirop[1][0] = miniempate
			imprimirTabuleirop()
		}
	case "#5", "5":
		tab5[linha][coluna] = jogadorAtual
		if vereficarvencedor5() {
			tabuleirop[1][1] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE5() {
			tabuleirop[1][1] = miniempate
			imprimirTabuleirop()
		}
	case "#6", "6":
		tab6[linha][coluna] = jogadorAtual
		if vereficarvencedor6() {
			tabuleirop[1][2] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE6() {
			tabuleirop[1][2] = miniempate
			imprimirTabuleirop()
		}
	case "#7", "7":
		tab7[linha][coluna] = jogadorAtual
		if vereficarvencedor7() {
			tabuleirop[2][0] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE7() {
			tabuleirop[2][0] = miniempate
			imprimirTabuleirop()
		}
	case "#8", "8":
		tab8[linha][coluna] = jogadorAtual
		if vereficarvencedor8() {
			tabuleirop[2][1] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE8() {
			tabuleirop[2][1] = miniempate
			imprimirTabuleirop()
		}
	case "#9", "9":
		tab9[linha][coluna] = jogadorAtual
		if vereficarvencedor9() {
			tabuleirop[2][2] = jogadorAtual
			imprimirTabuleirop()
		}
		if EMPATE9() {
			tabuleirop[2][2] = miniempate
			imprimirTabuleirop()
		}
	}
}

func verificarposicao() {
	if linha == 0 && coluna == 0 {
		principal = "#1"
	}
	if linha == 0 && coluna == 1 {
		principal = "#2"
	}
	if linha == 0 && coluna == 2 {
		principal = "#3"
	}
	if linha == 1 && coluna == 0 {
		principal = "#4"
	}
	if linha == 1 && coluna == 1 {
		principal = "#5"
	}
	if linha == 1 && coluna == 2 {
		principal = "#6"
	}
	if linha == 2 && coluna == 0 {
		principal = "#7"
	}
	if linha == 2 && coluna == 1 {
		principal = "#8"
	}
	if linha == 2 && coluna == 2 {
		principal = "#9"
	}
}

func trocadejogador() {
	if jogadorAtual == "\u001B[31mX\u001B[0m" {
		jogadorAtual = "\u001B[34mO\u001B[0m"
	} else {
		jogadorAtual = "\u001B[31mX\u001B[0m"
	}
}

func main() {
	if inicio() {
		inicializarTabuleirop()

		t := time.Now()
		if t.Second()%2 == 0 {
			jogadorAtual = "\u001B[34mO\u001B[0m"
		} else {
			jogadorAtual = "\u001B[31mX\u001B[0m"
		}

		miniempate = "\u001B[35mØ\u001B[0m"

		for {
			fmt.Print("\n -----JOGO DA VELHA 2.0-----\n")
			imprimirTabuleirop()
			fmt.Printf("Jogador %s, escolha por onde deseja começar: ", jogadorAtual)
			fmt.Scan(&principal)
			if principal == "1" || principal == "2" || principal == "3" || principal == "4" || principal == "5" ||
				principal == "6" || principal == "7" || principal == "8" || principal == "9" || principal == "#1" ||
				principal == "#2" || principal == "#3" || principal == "#4" || principal == "#5" || principal == "#6" ||
				principal == "#7" || principal == "#8" || principal == "#9" {
				break
			}
			fmt.Print("\nMermão digita direito isso ae bora\n")
			continue
		}
		if movimentopodetp(principal) {
			realizarMovimentop(principal)
			for {
				if tabuleirop[linha][coluna] == "\u001B[31mX\u001B[0m" || tabuleirop[linha][coluna] == "\u001B[34mO\u001B[0m" || tabuleirop[linha][coluna] == "\u001B[35mØ\u001B[0m" {
					imprimirTabuleirop()

					fmt.Printf("Jogador %s, escolha por onde deseja jogar: ", jogadorAtual)
					fmt.Scan(&principal)
					movimentopodetp(principal)
					realizarMovimentop(principal)
				}

				for {
					fmt.Printf("Jogador %s, digite a linha que deseja escolher para sua jogada: ", jogadorAtual)
					fmt.Scan(&linha)
					if linha < 0 || linha >= 3 {
						fmt.Println("Linha inválida. Tente novamente.")
						continue
					}
					break
				}
				for {
					fmt.Printf("Jogador %s, digite a coluna que deseja escolher para sua jogada: ", jogadorAtual)
					fmt.Scan(&coluna)
					if coluna < 0 || coluna >= 3 {
						fmt.Println("Coluna inválida. Tente novamente.")
						continue
					}
					break
				}

				switch principal {
				case "#1", "1":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab1[linha][coluna] == "\u001B[31mX\u001B[0m" || tab1[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#2", "2":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab2[linha][coluna] == "\u001B[31mX\u001B[0m" || tab2[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#3", "3":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab3[linha][coluna] == "\u001B[31mX\u001B[0m" || tab3[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#4", "4":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab4[linha][coluna] == "\u001B[31mX\u001B[0m" || tab4[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#5", "5":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab5[linha][coluna] == "\u001B[31mX\u001B[0m" || tab5[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#6", "6":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab6[linha][coluna] == "\u001B[31mX\u001B[0m" || tab6[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#7", "7":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab7[linha][coluna] == "\u001B[31mX\u001B[0m" || tab7[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#8", "8":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab8[linha][coluna] == "\u001B[31mX\u001B[0m" || tab8[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				case "#9", "9":
					if linha < 0 || linha >= 3 || coluna < 0 || coluna >= 3 || tab9[linha][coluna] == "\u001B[31mX\u001B[0m" || tab9[linha][coluna] == "\u001B[34mO\u001B[0m" {
						fmt.Println("Posição inválida. Tente novamente.")
						continue
					}
				}

				if movimentopode(linha, coluna) {
					realizamovimentopode(principal)
					realizarMovimentop(principal)
					switch principal {
					case "#1", "1":
						verificarposicao()
					case "#2", "2":
						verificarposicao()
					case "#3", "3":
						verificarposicao()
					case "#4", "4":
						verificarposicao()
					case "#5", "5":
						verificarposicao()
					case "#6", "6":
						verificarposicao()
					case "#7", "7":
						verificarposicao()
					case "#8", "8":
						verificarposicao()
					case "#9", "9":
						verificarposicao()
					}
					realizarMovimentop(principal)

					if vereficarvencedorreal() {
						vencedor = jogadorAtual
						break
					}
					trocadejogador()
				} else {
					fmt.Print("Movimento inválido, tente novamente")
				}

			}
			if vencedor != "" {
				for {
					fmt.Printf("Parabéns jogador %s, você acaba de ganhar o JOGO DA VELHA 2.0!", jogadorAtual)
					fmt.Print("Deseja jogar essa DESGRAMA novamente? ")
					fmt.Scan(&novo)
					resposta = strings.ToLower(resposta) // Converta a resposta para minúsculas
					if novo == "sim" || novo == "y" || novo == "yes" || novo == "s" {
						inicio()
					} else {
						fmt.Print("ADEUS SEU MERDA")
					}
				}
			} else {
				for {
					fmt.Print("Lamento jogadores, mas aparentemente deu velha. ")
					fmt.Print("Deseja jogar essa DESGRAMA novamente? ")
					fmt.Scan(&novo)
					resposta = strings.ToLower(resposta) // Converta a resposta para minúsculas
					if novo == "sim" || novo == "y" || novo == "yes" || novo == "s" {
						inicio()
						break
					} else {
						fmt.Print("ADEUS SEU MERDA")
						break
					}

				}
			}
		}
	}
}
