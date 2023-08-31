package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("> Welcome to pokedex, type <help> to know more")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">  ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		switch text {
		case "explore":
			random_pokemons()
		case "start":
			startingRepl()
		case "help":
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
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}
