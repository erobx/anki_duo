package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/erobx/anki_duo/cmd"
	"github.com/joho/godotenv"
)

type Cmdr struct {
	cmd *exec.Cmd
}

func NewCmdr(args []string) Cmdr {
	return Cmdr{
		cmd: exec.Command("python3", args...),
	}
}

func main() {
	args := []string{"scripts/__main__.py"}
	c := NewCmdr(args)
	//c.runDuo()
	fmt.Println("Importing to Anki")
	c.runAnki()
}

func (c *Cmdr) runDuo() {
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	log.Println(c.cmd.Run())
}

func (c *Cmdr) runAnki() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	path := os.Getenv("IMPORT_PATH")
	anki := cmd.NewAnki()
	anki.ImportFile(path)
}
