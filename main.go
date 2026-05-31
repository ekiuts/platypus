package main

import (
	"fmt"
)

// Launch the program and execute as directed by the supplied flag
func main() {
	flag := flags()
	logo()

	switch flag {
	case "-h", "--help":
		help()
	case "-s", "--search":
		credits()
		serialize()
		plugin()
	case "-v", "--version":
	default:
		alert("Unknown argument(s) - ")
	}
}

// Print the help information
func help() {
	Yellow.Println("\nUsage:")
	fmt.Println("  [program] [optional flag]")
	Yellow.Println("\n Flags:")
	Green.Printf("%s", "  -h, --help")
	fmt.Println("	   Help Information")
	Green.Printf("%s", "  -v, --version")
	fmt.Println("	   Display Program Version")
	Yellow.Println("\nRun:")
	fmt.Println("  From the folder containing your compiled executable, run:")
	Green.Printf("%s", "    platypus")
	Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	Green.Println("    https://github.com/farghul/platypus.git")
}

func logo() {
	Magenta.Println("▗▄▄▖ ▗▖    ▗▄▖▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▖ ▗▖ ▗▄▄▖")
	Magenta.Println("▐▌ ▐▌▐▌   ▐▌ ▐▌ █   ▝▚▞▘ ▐▌ ▐▌▐▌ ▐▌▐▌   ")
	Magenta.Println("▐▛▀▘ ▐▌   ▐▛▀▜▌ █    ▐▌  ▐▛▀▘ ▐▌ ▐▌ ▝▀▚▖")
	Magenta.Println("▐▌   ▐▙▄▄▖▐▌ ▐▌ █    ▐▌  ▐▌   ▝▚▄▞▘▗▄▄▞▘")
	Magenta.Println(bv)
}

func credits() {
	fmt.Println("\nAn update search tool for WordPress")
	fmt.Println("Created by Byron Stuike")
}
