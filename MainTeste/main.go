package Main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ==========================================
// MODELOS DE DADOS
// ==========================================

type User struct {
	ID       int
	Username string
	Password string
}

type Itinerary struct {
	ID          int
	Origin      string
	Destination string
	Date        string
}

type Reservation struct {
	ID          int
	UserID      int
	ItineraryID int
	Status      string // "Confirmada" ou "Cancelada"
}

// ==========================================
// BANCO DE DADOS MOCK (Em memória)
// ==========================================

var users = map[string]User{
	"passageiro1": {ID: 1, Username: "passageiro1", Password: "123"},
}

var itineraries = []Itinerary{
	{ID: 1, Origin: "SAO PAULO", Destination: "RIO DE JANEIRO", Date: "10/10/2026"},
	{ID: 2, Origin: "SAO PAULO", Destination: "RIO DE JANEIRO", Date: "10/10/2026"},
	{ID: 3, Origin: "FEIRA DE SANTANA", Destination: "SALVADOR", Date: "15/10/2026"},
}

var reservations = make(map[int]*Reservation)
var nextReservationID = 1

// ==========================================
// VARIÁVEIS GLOBAIS DE ESTADO
// ==========================================

var currentUser *User
var scanner = bufio.NewScanner(os.Stdin)

// ==========================================
// FUNÇÃO PRINCIPAL E MENUS
// ==========================================

func main() {
	fmt.Println("Bem-vindo ao Sistema de Reserva de Passagens!")

	for {
		if currentUser == nil {
			mainMenu()
		} else {
			passengerMenu()
		}
	}
}

func mainMenu() {
	fmt.Println("\n--- MENU PRINCIPAL ---")
	fmt.Println("1. Autenticar (Login)")
	fmt.Println("2. Sair")
	fmt.Print("Escolha uma opção: ")

	option := readInput()

	switch option {
	case "1":
		authenticate()
	case "2":
		fmt.Println("Encerrando o sistema. Até logo!")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida.")
	}
}

func passengerMenu() {
	fmt.Printf("\n--- PAINEL DO PASSAGEIRO (%s) ---\n", currentUser.Username)
	fmt.Println("1. Buscar e Reservar Itinerários")
	fmt.Println("2. Consultar Minhas Reservas")
	fmt.Println("3. Cancelar Reserva")
	fmt.Println("4. Sair da Conta (Logout)")
	fmt.Print("Escolha uma opção: ")

	option := readInput()

	switch option {
	case "1":
		searchAndBookItinerary()
	case "2":
		viewReservations()
	case "3":
		cancelReservation()
	case "4":
		currentUser = nil
		fmt.Println("Logout realizado com sucesso.")
	default:
		fmt.Println("Opção inválida.")
	}
}

// ==========================================
// LÓGICA DE NEGÓCIO
// ==========================================

func authenticate() {
	fmt.Print("Usuário: ")
	username := readInput()
	fmt.Print("Senha: ")
	password := readInput()

	user, exists := users[username]
	if exists && user.Password == password {
		currentUser = &user
		fmt.Println("Login realizado com sucesso!")
	} else {
		fmt.Println("Erro: Usuário ou senha incorretos.")
	}
}

func searchAndBookItinerary() {
	fmt.Println("\n--- BUSCAR ITINERÁRIOS ---")
	fmt.Print("Origem: ")
	origin := strings.ToUpper(readInput())
	fmt.Print("Destino: ")
	dest := strings.ToUpper(readInput())
	fmt.Print("Data (DD/MM/AAAA): ")
	date := readInput()

	var results []Itinerary
	for _, it := range itineraries {
		if it.Origin == origin && it.Destination == dest && it.Date == date {
			results = append(results, it)
		}
	}

	if len(results) == 0 {
		fmt.Println("Nenhum itinerário encontrado para esses dados.")
		return
	}

	fmt.Println("\nItinerários Encontrados:")
	for _, it := range results {
		fmt.Printf("ID: %d | %s -> %s | Data: %s\n", it.ID, it.Origin, it.Destination, it.Date)
	}

	fmt.Print("\nDeseja fazer a reserva de algum itinerário? (S/N): ")
	choice := strings.ToUpper(readInput())
	if choice == "S" {
		fmt.Print("Digite o ID do itinerário que deseja reservar: ")
		idStr := readInput()
		itID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("ID inválido.")
			return
		}

		// Validar se o ID existe nos resultados
		valid := false
		for _, it := range results {
			if it.ID == itID {
				valid = true
				break
			}
		}

		if !valid {
			fmt.Println("Erro: ID de itinerário não corresponde à sua busca.")
			return
		}

		// Confirmar reserva
		confirmReservation(itID)
	}
}

func confirmReservation(itineraryID int) {
	res := &Reservation{
		ID:          nextReservationID,
		UserID:      currentUser.ID,
		ItineraryID: itineraryID,
		Status:      "Confirmada",
	}
	reservations[nextReservationID] = res
	fmt.Printf("Sucesso! Reserva confirmada. Seu código de reserva é: %d\n", nextReservationID)
	nextReservationID++
}

func viewReservations() {
	fmt.Println("\n--- MINHAS RESERVAS ---")
	found := false

	for _, res := range reservations {
		if res.UserID == currentUser.ID {
			found = true
			// Buscar os detalhes do itinerário
			var itDetails string
			for _, it := range itineraries {
				if it.ID == res.ItineraryID {
					itDetails = fmt.Sprintf("%s -> %s (%s)", it.Origin, it.Destination, it.Date)
					break
				}
			}
			fmt.Printf("Reserva ID: %d | Status: %s | Itinerário: %s\n", res.ID, res.Status, itDetails)
		}
	}

	if !found {
		fmt.Println("Você ainda não possui reservas.")
	}
}

func cancelReservation() {
	fmt.Println("\n--- CANCELAR RESERVA ---")
	fmt.Print("Digite o ID da reserva que deseja cancelar (ou '0' para voltar): ")
	idStr := readInput()

	if idStr == "0" {
		return
	}

	resID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	res, exists := reservations[resID]
	if !exists || res.UserID != currentUser.ID {
		fmt.Println("Erro: Reserva não encontrada ou não pertence a você.")
		return
	}

	if res.Status == "Cancelada" {
		fmt.Println("Aviso: Esta reserva já está cancelada.")
		return
	}

	res.Status = "Cancelada"
	fmt.Printf("Sucesso! A reserva %d foi cancelada.\n", resID)
}

// ==========================================
// FUNÇÕES UTILITÁRIAS
// ==========================================

func readInput() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
