package main

import (
	"os/exec"
	"os"

)

func main() {
	pscmd := exec.Command("ps", "-ef")
	grepcmd := exec.Command("grep", "jearry")

	stdinpip, _:= grepcmd.StdinPipe()
	pscmd.Stdout = stdinpip

	grepcmd.Stdout = os.Stdout
	pscmd.Start()


	grepcmd.Start()

	pscmd.Wait()


}
