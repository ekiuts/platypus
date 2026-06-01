package main

import (
	"flag"
	"fmt"
)

// Launch the program and execute as directed by the supplied flag
func main() {
	helpshort := flag.Bool("h", false, "Display help information")
	helplong := flag.Bool("help", false, "Display help information")
	searchshort := flag.Bool("s", false, "Display help information")
	searchlong := flag.Bool("search", false, "Display help information")
	versionshort := flag.Bool("v", false, "Display Program Version")
	versionlong := flag.Bool("version", false, "Display Program Version")
	flag.Parse()

	if flag.NFlag() > 0 {
		switch {
		case *helpshort, *helplong:
			help()
		case *searchshort, *searchlong:
			logo()
			credits()
			serialize()
			plugin()
		case *versionshort, *versionlong:
			Magenta.Println(bv)
		}
	} else {
		alert("No arguments found - ")
	}
}

// Print the help information
func help() {
	Yellow.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	Yellow.Println("\n Flags:")
	Green.Printf("%s", "  -h, -help")
	fmt.Println("	   Help Information")
	Green.Printf("%s", "  -s, -search")
	fmt.Println("	   Run main program")
	Green.Printf("%s", "  -v, -version")
	fmt.Println("	   Display Program Version")
	Yellow.Println("\nRun:")
	fmt.Println("  From the folder containing your compiled executable, run:")
	Green.Printf("%s", "    platypus -s")
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
