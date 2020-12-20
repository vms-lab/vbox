package vbox

import (
	"bufio"
	"log"
	"regexp"
	"strings"
)

type VM struct {
	Name string
	UUID string
}

type VMSpec struct {
	Name     string `x:"name"`
	OSType   string `x:"ostype"`
	UUID     string `x:"UUID"`
	CPUs     string `x:"cpus"`
	Memory   string `x:"memory"`
	Groups   string `x:"groups"`
	State    string `x:"VMState"`
	Networks []VMNETSpec
}

type VMNETSpec struct {
	NIC            string `x:"nic" numbered:"true"` //type ? nat, intnet, ...
	MAC            string `x:"macaddress" numbered:"true"`
	Name           string `x:"natnet,intnet" numbered:"true"` // identify other types and append list
	CableConnected string `x:"cableconnected" numbered:"true"`
}

var vmsExp = regexp.MustCompile(`"(?P<name>.+)" {(?P<uuid>.+)}`)

func (vbmng VBoxManager) ListVMs() (vms []VM) {
	vmList := vbmng.Command("list", "vms")
	scanner := bufio.NewScanner(strings.NewReader(vmList))
	for scanner.Scan() {
		vmMatch := reSubMatchMap(vmsExp, scanner.Text())
		vms = append(vms, VM{
			Name: vmMatch["name"],
			UUID: vmMatch["uuid"],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return vms
}
