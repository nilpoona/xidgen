package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/rs/xid"
)

func main() {
	if runtime.GOOS != "darwin" {
		fmt.Println("Operating systems other than darwin are not supported.")
		return
	}
	id := xid.New().String()
	echoCmd := exec.Command("echo", "-n", id)
	pbcopyCmd := exec.Command("pbcopy")

	reader, writer := io.Pipe()
	var buf bytes.Buffer

	echoCmd.Stdout = writer

	pbcopyCmd.Stdin = reader

	pbcopyCmd.Stdout = &buf

	echoCmd.Start()

	pbcopyCmd.Start()

	echoCmd.Wait()
	writer.Close()

	pbcopyCmd.Wait()
	reader.Close()

	io.Copy(os.Stdout, &buf)

	fmt.Println("copied the following id to the clipboard")
	fmt.Println(fmt.Sprintf("xid: %s", id))
}
