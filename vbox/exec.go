package vbox

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Cmd struct {
	Cmd *exec.Cmd
	In  io.WriteCloser
	Out io.ReadCloser
}

func NewCmd(name string, args ...string) (cmd Cmd) {
	var err error
	cmd.Cmd = exec.Command(name, args...)
	// cmd.Out, err = cmd.Cmd.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	cmd.Cmd.Stdout = os.Stdout
	cmd.In, err = cmd.Cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return cmd
}

func (cmd Cmd) Write(in string) {
	io.WriteString(cmd.In, in)
}

func (cmd Cmd) Read() string {
	buf := new(strings.Builder)
	// first value is number of readed bytes
	_, err := io.Copy(buf, cmd.Out)
	if err != nil {
		log.Fatal(err)
	}
	str := buf.String()
	return str
}

func (cmd Cmd) Close() {
	cmd.In.Close()
	if err := cmd.Cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
