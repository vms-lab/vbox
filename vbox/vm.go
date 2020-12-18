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

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := map[string]string{}
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}

	return subMatchMap
}

var vmsExp = regexp.MustCompile(`"(?P<name>.+)" {(?P<uuid>.+)}`)

func ParseVMsList(vmList string) (vms []VM) {
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

func VMs() {
	listVMs := Manage("list", "vms")
	log.Println(listVMs)
}
