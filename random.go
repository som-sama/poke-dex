package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var cities = []string{"Delhi", "Mumbai", "Kolkata", "Guwahati", "Patna",
	"Bangalore", "Jaipur", "Howrah", "Chennai", "Jammu"}

var allPokemon = []string{}

func random_pokemons() {
	rand.Seed(time.Now().UnixNano())

	// Load Pokémon names from the text file
	loadPokemonFromFile("pokemon.txt")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			switch {
			case input == "explore":
				for _, city := range cities {
					fmt.Println(">", city)
				}
			case strings.HasPrefix(input, "explore "):
				city := strings.TrimPrefix(input, "explore ")
				if cityInList(city, cities) {
					displayRandomPokemon()
				} else {
					fmt.Println("Unknown city")
				}
			case strings.HasPrefix(input, "catch "):
				pokemon := strings.TrimPrefix(input, "catch ")
				if pokemonInList(pokemon, allPokemon) {
					fmt.Printf("You have caught %s!\n", pokemon)
				} else {
					fmt.Println("Unknown Pokemon")
				}
			case input == "help":
				fmt.Println("Welcome to the Pokedex!")
				fmt.Println("Usage: ")
				fmt.Println()
				fmt.Println("help: Displays a help message")
				fmt.Println("type <start> to enter pokedex")
				fmt.Println("type <explore> to explore areas")
				fmt.Println("type <explore area-name> to see the pokemons available")
				fmt.Println("type <pokemon name> to get details about it")
				fmt.Println("type catch to catch the pokemon")
				fmt.Println("exit: Exit the Pokedex")
			case input == "exit":
				os.Exit(0)
			case input == "start":
				startingRepl()
			default:
				fmt.Print("Please type a valid command")
			}
		}
	}
}

func loadPokemonFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allPokemon = append(allPokemon, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
	}
}

func cityInList(city string, cities []string) bool {
	for _, c := range cities {
		if c == city {
			return true
		}
	}
	return false
}

func pokemonInList(pokemon string, pokemons []string) bool {
	for _, p := range pokemons {
		if p == pokemon {
			return true
		}
	}
	return false
}

func displayRandomPokemon() {
	shufflePokemon := rand.Perm(len(allPokemon))
	for i := 0; i < 5 && i < len(allPokemon); i++ {
		fmt.Println(allPokemon[shufflePokemon[i]])
	}
}
