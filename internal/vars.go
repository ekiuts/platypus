package internal

import "fmt"

type Core struct {
	Offers []struct {
		Current string `json:"current"`
	} `json:"offers"`
}

// Agile builds a collection of urls to target changelogs
type Agile struct {
	ACFPro    string `json:"acfpro"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	Spotlight string `json:"spotlight"`
	WordPress string `json:"wordpress"`
	WPExport  string `json:"wpexport"`
}

// Values specific tp the WordPress installation
type Environment struct {
	Address   string `json:"address"`
	Install   string `json:"install"`
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
	Server    string `json:"server"`
	User      string `json:"user"`
}

type Color string

const (
	Reset          = "\033[0m"
	Red     Color  = "\033[31m"
	Green   Color  = "\033[32m"
	Yellow  Color  = "\033[33m"
	Magenta Color  = "\033[35m"
	BGRed   Color  = "\033[41m"
	BV      string = "1.0.0"
	halt    string = "program halted"
	meta    string = "/data/automation/jsons/"
	temp    string = "/data/automation/temp/"
)

var (
	Version   Core
	Changelog Agile
	Ecosystem Environment
	Jsons     = []string{meta + "changelogs.json", meta + "test.json"}
	Remains   = []string{temp + "grepped.txt", temp + "temp.json", temp + "webscrape.txt"}
)

// Println function for colourized text
func (c Color) Println(text string) {
	fmt.Println(string(c) + text + Reset)
}

// Printf function for colourized text
func (c Color) Printf(format string, a ...any) {
	fmt.Printf(string(c)+format+Reset, a...)
}
