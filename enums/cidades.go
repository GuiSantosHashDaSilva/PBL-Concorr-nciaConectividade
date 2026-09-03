package enums

type Cidade int

const (
	Feira Cidade = iota
	Salvador
	Alagoinhas
	Cajazeiras
	Berimbau
	Cruz
	Cachoeira
	PauloAfonso
	Barreiras
	Lauro
	Filadelfia
	Extrema
	XiqueXique
	Mucuri
	Biritinga
	Itamaraju
)

var Cidades = map[Cidade]string{
	Feira:       "Feira",
	Salvador:    "Salvador",
	Alagoinhas:  "Alagoinhas",
	Cajazeiras:  "Cajazeiras",
	Berimbau:    "Berimbau",
	Cruz:        "Cruz",
	Cachoeira:   "Cachoeira",
	PauloAfonso: "Paulo Afonso",
	Barreiras:   "Barreiras",
	Lauro:       "Lauro",
	Filadelfia:  "Filadelfia",
	Extrema:     "Extrema",
	XiqueXique:  "Xique-Xique",
	Mucuri:      "Mucuri",
	Biritinga:   "Biritinga",
	Itamaraju:   "Itamaraju",
}
