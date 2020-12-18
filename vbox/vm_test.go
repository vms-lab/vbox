package vbox

import (
	"testing"
)

func (vm VM) Equal(vmToCmp VM) bool {
	if vm.Name != vmToCmp.Name {
		return false
	}
	if vm.UUID != vmToCmp.UUID {
		return false
	}
	return true
}

func TestParseVMsList(t *testing.T) {
	var testCases = []struct {
		vmList string
		vms    []VM
	}{
		{
			vmList: "\"base\" {b405ac98-8946-4c8f-aafb-620738641db2}",
			vms: []VM{
				VM{
					Name: "base",
					UUID: "b405ac98-8946-4c8f-aafb-620738641db2",
				},
			},
		},
		{
			vmList: `"k8s_node_1" {c67502ef-f0e1-4d2d-909a-6a75bf1f26e5}
"k8s_node_2" {2233b86a-7935-48a7-8aa3-ded1e3dbaafe}
"k8s_node_3" {628d6877-6ab1-4d41-b605-aeee24e62670}
`,

			vms: []VM{
				VM{
					Name: "k8s_node_1",
					UUID: "c67502ef-f0e1-4d2d-909a-6a75bf1f26e5",
				},
				VM{
					Name: "k8s_node_2",
					UUID: "2233b86a-7935-48a7-8aa3-ded1e3dbaafe",
				},
				VM{
					Name: "k8s_node_3",
					UUID: "628d6877-6ab1-4d41-b605-aeee24e62670",
				},
			},
		},
	}
	for _, testCase := range testCases {
		got := ParseVMsList(testCase.vmList)
		for vmIndex, expectedVM := range testCase.vms {
			gotVM := got[vmIndex]
			if !gotVM.Equal(expectedVM) {
				t.Errorf("Expected %v, but got %v", expectedVM, gotVM)
			}
		}
	}
}
