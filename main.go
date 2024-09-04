package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/erobx/anki_duo/cmd"
)

type Cmdr struct {
	cmd *exec.Cmd
}

func NewCmdr(args []string) Cmdr {
	return Cmdr{
		cmd: exec.Command("python3", args...),
	}
}

// go run . JWT
func main() {
	var jwt string
	fmt.Print("Enter your JWT from Duolingo: ")
	fmt.Scan(&jwt)

	args := []string{"scripts/__main__.py", jwt}
	c := NewCmdr(args)
	c.runDuo()
	c.runAnki()
}

func (c *Cmdr) runDuo() {
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	err := c.cmd.Run()
	if err != nil {
		log.Fatal("Could not get vocab from Duolingo")
	}
}

func (c *Cmdr) runAnki() {
	fmt.Println("Importing to Anki")
	anki := cmd.NewAnki()
	anki.ImportFile()
}
