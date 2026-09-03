package Motorista

import (
	"fmt"
	"gomgs/enums"
)

func definirTrajeto() {
	// Variável iniciada como true para o loop rodar
	definicaoaberta := true

	for definicaoaberta {
		// Substitua as aspas duplas (") por crases (`)
		fmt.Println(`
		================================================
		MENU DE CIDADES
		================================================
		[ 1 ] Feira Cidade      [ 9 ] Barreiras
		[ 2 ] Salvador          [ 10 ] Lauro
		[ 3 ] Alagoinhas        [ 11 ] Filadelfia
		[ 4 ] Cajazeiras        [ 12 ] Extrema
		[ 5 ] Berimbau          [ 13 ] XiqueXique
		[ 6 ] Cruz              [ 14 ] Mucuri
		[ 7 ] Cachoeira         [ 15 ] Biritinga
		[ 8 ] PauloAfonso       [ 16 ] Itamaraju
		================================================
		[ 0 ] Sair
		================================================
		`)

		var cidadetrajeto int
		fmt.Println("Escolha uma opção: ")
		fmt.Scan(&cidadetrajeto)

		if cidadetrajeto == 0 {
			break
		} else {
			fmt.Printf(enums.Cidades[enums.Cidade(cidadetrajeto-1)])
			definicaoaberta = false
		}

	}
}
