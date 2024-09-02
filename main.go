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

// go run . JWT file_name
func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatal("Please provide the correct arguments")
	}
	jwt := os.Args[1]
	file_name := os.Args[2]
	args := []string{"scripts/__main__.py", jwt, file_name}
	c := NewCmdr(args)
	c.runDuo()
	c.runAnki(file_name)
}

func (c *Cmdr) runDuo() {
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	err := c.cmd.Run()
	if err != nil {
		log.Fatal("Could not get vocab from Duolingo")
	}
}

func (c *Cmdr) runAnki(file_name string) {
	fmt.Println("Importing to Anki")
	anki := cmd.NewAnki()
	anki.ImportFile(file_name)
}
