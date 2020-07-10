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

	fmt.Println(`
__________           .___               __________                      .__    ___________                                      
\______   \ ____   __| _/___________    \______   \_____  ______   ____ |  |   \__    ___/___   __________  __ ______________   
 |     ___// __ \ / __ |\_  __ \__  \    |     ___/\__  \ \____ \_/ __ \|  |     |    |_/ __ \ /  ___/  _ \|  |  \_  __ \__  \  
 |    |   \  ___// /_/ | |  | \// __ \_  |    |     / __ \|  |_> >  ___/|  |__   |    |\  ___/ \___ (  <_> )  |  /|  | \// __ \_
 |____|    \___  >____ | |__|  (____  /  |____|    (____  /   __/ \___  >____/   |____| \___  >____  >____/|____/ |__|  (____  /
               \/     \/            \/                  \/|__|        \/                    \/     \/                        \/ 
 `)

	fmt.Println("Introduz o teu nome")

	reader := bufio.NewReader(os.Stdin)
	nome, _ := reader.ReadString('\n')
	nome = strings.Replace(nome, "\n", "", -1)

	fmt.Printf("%s %s\n", "\nBem-vindo, escolhe uma opção", nome)
	fmt.Printf("\n")
	for {

		fmt.Println(`Clica "c" para começar`)
		fmt.Println(`Clica "s" para sair`)
		fmt.Println(`Clica "v" para ver a pontuação`)

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case 'c':
			fmt.Println("\nEscolhe uma jogada:")
			startGame()
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

func printScore(nome string) {
	fmt.Printf("\n%s  %d : %d  %s\n\n", nome, pontosJogador, pontosComputador, "Computador")
}

func startGame() {
	fmt.Println(`Clica "r"  para Pedra`)
	fmt.Println(`Clica "p"  para Papel`)
	fmt.Println(`Clica "t"  para Tesoura`)

	reader := bufio.NewReader(os.Stdin)
	escolhaUtilizador, _, error := reader.ReadRune()

	if error != nil {
		fmt.Println(error)
	}

	escolhaComputador := getescolhaComputador()

	escolhas := map[string]string{"r": "Pedra", "p": "Papel", "t": "Tesoura"}

	fmt.Println("O computador escolheu " + escolhas[string(escolhaComputador)])

	resultados(escolhaUtilizador, escolhaComputador)
}

func getescolhaComputador() rune {
	escolhas := [3]rune{'r', 'p', 't'}
	return escolhas[rand.Intn(len(escolhas))]
}

func resultados(utilizador rune, computador rune) {

	resultado := string(utilizador) + string(computador)
	switch resultado {
	case "rt", "pr", "tp":
		fmt.Println("Ganhaste :)")
		pontosJogador++
	case "rp", "pt", "tr":
		fmt.Println("Perdeste :(")
		pontosComputador++
	case "rr", "pp", "tt":
		fmt.Println("Empate :|")
	}
}