package internal

import "fmt"

type WPLatest struct {
	Offers []struct {
		Current string `json:"current"`
	} `json:"offers"`
}

// Changelogs builds a collection of urls to target changelogs
type Changelogs struct {
	ACFPro    string `json:"acfpro"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	Spotlight string `json:"spotlight"`
	WordPress string `json:"wordpress"`
	WPExport  string `json:"wpexport"`
}

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
	Lists   string = "/data/automation/lists/"
	meta    string = "/data/automation/jsons/"
	temp    string = "/data/automation/temp/"
)

var (
	Version   WPLatest
	News      Changelogs
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
