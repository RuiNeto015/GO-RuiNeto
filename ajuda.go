package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var pontosJogador int
var pontosComputador int

func main() {

//http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20

	fmt.Println(` 
__________           .___               __________                      .__    ___________                                      
\______   \ ____   __| _/___________    \______   \_____  ______   ____ |  |   \__    ___/___   __________  __ ______________   
 |     ___// __ \ / __ |\_  __ \__  \    |     ___/\__  \ \____ \_/ __ \|  |     |    |_/ __ \ /  ___/  _ \|  |  \_  __ \__  \  
 |    |   \  ___// /_/ | |  | \// __ \_  |    |     / __ \|  |_> >  ___/|  |__   |    |\  ___/ \___ (  <_> )  |  /|  | \// __ \_
 |____|    \___  >____ | |__|  (____  /  |____|    (____  /   __/ \___  >____/   |____| \___  >____  >____/|____/ |__|  (____  /
               \/     \/            \/                  \/|__|        \/                    \/     \/                        \/ 
 `)

	fmt.Println("Introduz o teu nome")

// := serve para declarar e atribuir logo um valor a uma variável
// O pacote bufio é um implementador. Ele envolve um objeto io.Reader ou io.Writer, criando outro objeto (Reader ou Writer)
	reader := bufio.NewReader(os.Stdin)

// Lê a string implementada no Reader 
	nome, _ := reader.ReadString('\n')

// Transforma aquela variável nome, no nome que nós introduzimos
	nome = strings.Replace(nome, "\n", "", -1)

	fmt.Printf("%s %s\n", "\nBem-vindo, escolhe uma opção", nome)
	fmt.Printf("\n")
	for {

		fmt.Println(`Clica "c" para começar`)
		fmt.Println(`Clica "s" para sair`)
		fmt.Println(`Clica "v" para ver a pontuação`)

		reader := bufio.NewReader(os.Stdin)

// O ReadRune lê um único caracter codificado em UTF-8 e o seu tamanho em bytes.
		char, _, err := reader.ReadRune()

// Nil basicamente é um limite, uma limitação, por exemplo encontrar erros. Para usos simples, um scanner pode ser mais conveniente.
		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case 'c':
			fmt.Println("\nEscolhe uma jogada:")
			startGame()
// Mostra os resultados (func printScore) com a variável que tem o nome escolhido
			printScore(nome)
		case 's':
			fmt.Println("A sair...")
			os.Exit(1)
		case 'v':
			printScore(nome)
		default:
			fmt.Println("Erro! Escolhe um valor válido")
		}

	}
}

// Função que mostra as pontuações do jogo
func printScore(nome string) {
	fmt.Printf("\n%s  %d : %d  %s\n\n", nome, pontosJogador, pontosComputador, "Computador")
}

// Função para começar o jogo
func startGame() {
	fmt.Println(`Clica "r"  para Pedra`)
	fmt.Println(`Clica "p"  para Papel`)
	fmt.Println(`Clica "t"  para Tesoura`)

// mais uma vez vai ler e se nao houverem erros, procede o programa
	reader := bufio.NewReader(os.Stdin)
	escolhaUtilizador, _, error := reader.ReadRune()

	if error != nil {
		fmt.Println(error)
	}

	escolhaComputador := getescolhaComputador()

	escolhas := map[string]string{"r": "Pedra", "p": "Papel", "t": "Tesoura"}

	fmt.Println("O computador escolheu " + escolhas[string(escolhaComputador)])

// armazena a escolha do utilizador e a escolha random do computador
	resultados(escolhaUtilizador, escolhaComputador)
}

// Função para esolher uma letra random de entre aquelas 3 [3] para o computador
func getescolhaComputador() rune {
	escolhas := [3]rune{'r', 'p', 't'}
	return escolhas[rand.Intn(len(escolhas))]
}

// Função que mostra o resultado final, se se ganhou, perdeu ou empatou
func resultados(utilizador rune, computador rune) {

// guarda as letras, primeiro a jogada do utilizador, depois a do computador
	resultado := string(utilizador) + string(computador)
	switch resultado {
	case "rt", "pr", "tp":
		fmt.Println("Ganhaste :)")
		pontosJogador++
	case "rp", "pt", "tr":
		fmt.Println("Perdeste :(")
		pontosComputador++
	case "rr", "pp", "st":
		fmt.Println("Empate :|")
	}
}