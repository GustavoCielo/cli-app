package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
	greetFormalCmd := flag.NewFlagSet("formal", flag.ExitOnError)
	greetFormalCmd.Usage = func() {
		fmt.Printf("Usage: %s greet formal --name=<name>\n", os.Args[0])
		greetFormalCmd.PrintDefaults()
	}
	greetNamePtr := greetFormalCmd.String("name", "Go CLI", "name to greet.")

	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(2)
	}

	switch os.Args[1] {
	case "greet":
		greetCmd.Parse(os.Args[2:])
		fmt.Println("Hi there!")
		if greetCmd.Arg(0) == "formal" {
			greetFormalCmd.Parse(os.Args[3:])
			fmt.Printf("How do you do, %s?", *greetNamePtr)
		}
	default:
		fmt.Println("Unknown command.")
		os.Exit(2)
	}
}
