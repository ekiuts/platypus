package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Provide and highlight an informational message
func inform(message string) {
	Yellow.Printf("%s", "** ")
	fmt.Print(message)
	Yellow.Println(" **")
}

// Print a colourized error message
func Alert(message string) {
	Red.Printf("\n%s", "Error: ")
	fmt.Printf("%s", message)
	BGRed.Println(halt)
	inform("Use -h to display help information")
	os.Exit(0)
}

// Remove files or directories
func Cleanup(cut string) {
	Inspect(os.Remove(cut))
}

// Run a terminal command, then capture and return the output as a byte
func Capture(task string, args ...string) ([]byte, error) {
	return exec.Command(task, args...).CombinedOutput()
}

// Check for errors, halt the program if found, and log the result
func Inspect(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// A generic GET request which captures the response
func Get(target string) []byte {
	resp, err := http.Get(target)
	Inspect(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Inspect(err)

	return body
}

// Run the Linux mail command and email the result to the configured recipent(s)
func Mailing(list string) {
	cmd := exec.Command("mail", "-s", "WordPress updates for "+Ecosystem.Address, "-r", "Delivery Cactuar <"+Ecosystem.Sender+">", Ecosystem.Recipient)
	stdin, err := cmd.StdinPipe()
	Inspect(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, "Below is the current list of plugins requiring updates for "+Ecosystem.Address+". Have a magical day!\n\n"+list)
		Inspect(err)
	}()

	out, _ := cmd.CombinedOutput() // Necessary although valueless

	log.Println(string(out))
}

// Pipe together commands using the exec.Command function
func Concat(method, flag, task, pipe string) []byte {
	cmd := exec.Command(method, flag, task)
	stdin, err := cmd.StdinPipe()
	Inspect(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, pipe)
		Inspect(err)
	}()

	out, _ := cmd.CombinedOutput()
	return out
}
