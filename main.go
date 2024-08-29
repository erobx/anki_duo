package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/erobx/anki_duo/cmd"
	"github.com/joho/godotenv"
)

type Cmdr struct {
	Out  io.Writer
	Args []string
	cmd  *exec.Cmd
}

func NewCmdr(args []string) Cmdr {
	return Cmdr{
		Args: args,
	}
}

func main() {
	args := []string{"python3", "scripts/__main__.py"}
	c := NewCmdr(args)
	c.runDuo()

}

func runAnki() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	path := os.Getenv("IMPORT_PATH")
	anki := cmd.NewAnki()
	_ = path

	anki.ImportFile(path)
}

func (c *Cmdr) runDuo() {
	cmd := exec.Command("python3", "scripts/__main__.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}
