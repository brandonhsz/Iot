package services

import (
	"os/exec"
)

func Reboot() {
	err := exec.Command("bash", "-c", "reboot ").Run()

	if err != nil {
		panic(err)
	}
}
