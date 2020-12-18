package vbox

import (
	"log"
	"os/exec"
)

func Manage(args ...string) string {
	out, err := exec.Command("VBoxManage", args...).Output()
	if err != nil {
		log.Fatalf("Error during getting list vms: %v", err)
	}
	return string(out)
}
