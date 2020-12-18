package vbox

import (
	"log"
)

type VM struct {
	Name string
	UUID string
}

func ParseVMsList(vmList string) (vms []VM) {
	return vms
}

func VMs() {
	listVMs := Manage("list", "vms")
	log.Println(listVMs)
}
