package main

import "os/exec"

func main()  {
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")
	if err := cmd0.Start(); err != nil {
		
	}
}
