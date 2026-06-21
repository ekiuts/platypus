package main

import (
	"flag"
	"fmt"

	"github.com/farghul/platypus/internal"
	"github.com/farghul/platypus/workers"
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
			// workers.Serialize()
			workers.Plugin()
		case *versionshort, *versionlong:
			internal.Magenta.Println(internal.BV)
		}
	} else {
		internal.Alert("No arguments found - ")
	}
}

// Print the help information
func help() {
	internal.Yellow.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	internal.Yellow.Println("\n Flags:")
	internal.Green.Printf("%s", "  -h, -help")
	fmt.Println("	   Help Information")
	internal.Green.Printf("%s", "  -s, -search")
	fmt.Println("	   Run main program")
	internal.Green.Printf("%s", "  -v, -version")
	fmt.Println("	   Display Program Version")
	internal.Yellow.Println("\nRun:")
	fmt.Println("  From the folder containing your compiled executable, run:")
	internal.Green.Printf("%s", "    platypus -s")
	internal.Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	internal.Green.Println("    https://github.com/farghul/platypus.git")
}

func logo() {
	internal.Magenta.Println("▗▄▄▖ ▗▖    ▗▄▖▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▖ ▗▖ ▗▄▄▖")
	internal.Magenta.Println("▐▌ ▐▌▐▌   ▐▌ ▐▌ █   ▝▚▞▘ ▐▌ ▐▌▐▌ ▐▌▐▌   ")
	internal.Magenta.Println("▐▛▀▘ ▐▌   ▐▛▀▜▌ █    ▐▌  ▐▛▀▘ ▐▌ ▐▌ ▝▀▚▖")
	internal.Magenta.Println("▐▌   ▐▙▄▄▖▐▌ ▐▌ █    ▐▌  ▐▌   ▝▚▄▞▘▗▄▄▞▘")
	internal.Magenta.Println(internal.BV)
}

func credits() {
	fmt.Println("\nAn update search tool for WordPress")
	fmt.Println("Created by Byron Stuike")
}
