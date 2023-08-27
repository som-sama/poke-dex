package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Pokemon struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
}

func getPokemonDetails(name string) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func startingRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("pokedex> ")

		scanner.Scan()
		text := scanner.Text()

		if text == "help" {
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("Usage: ")
			fmt.Println()
			fmt.Println("help: Displays a help message")
			fmt.Println("type pokemon name to get details about it")
			fmt.Println("exit or quit: Exit the Pokedex")
		} else if text == "exit" || text == "quit" {
			os.Exit(0)
		} else {
			pokemonDetails, err := getPokemonDetails(text)
			if err != nil {
				fmt.Println("Error fetching Pokémon details:", err)
				return
			}

			fmt.Printf("Name: %s, ID: %d\n", pokemonDetails.Name, pokemonDetails.ID)

			// Displaying types
			fmt.Print("Type: ")
			for _, t := range pokemonDetails.Types {
				fmt.Printf("%s ", t.Type.Name)
			}
			fmt.Println()

			// Displaying moves (let's limit to first 5 moves for brevity)
			fmt.Print("Moves: ")
			for i, m := range pokemonDetails.Moves {
				if i >= 5 {
					break
				}
				fmt.Printf("%s, ", m.Move.Name)
			}
			fmt.Println()
		}
	}
}

func cleaned(str string) []string {
	lower := strings.ToLower(str)

	words := strings.Fields(lower)

	return words
}
