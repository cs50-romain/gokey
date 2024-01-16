package clipboard

import (
	"log"
	"os/exec"
)

func CopytoClipboard(output string) {
	var copyCmd *exec.Cmd

	copyCmd = exec.Command("xclip", "-selection", "C")

	in, err := copyCmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := copyCmd.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := in.Write([]byte(output)); err != nil {
		log.Fatal(err)
	}

	if err := in.Close(); err != nil {
		log.Fatal(err)
	}

	copyCmd.Wait()
}
