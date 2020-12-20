package vbox

import (
	"log"
	"os/exec"
)

func VBoxCommand(args ...string) string {
	out, err := exec.Command("VBoxManage", args...).Output()
	if err != nil {
		log.Fatalf("Error during getting list vms: %v", err)
	}
	return string(out)
}

type VBoxManager struct {
	Command func(args ...string) string
}

func NewVBoxManager() (vbmng VBoxManager) {
	return VBoxManager{
		Command: VBoxCommand,
	}
}

func (vbmng VBoxManager) VMSpec(vm VM) string {
	return vbmng.Command("showvminfo", vm.UUID)
}
