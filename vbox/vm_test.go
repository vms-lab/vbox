package vbox

import "testing"

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
	}
	for _, testCase := range testCases {
		got := ParseVMsList(testCase.vmList)
		if got != testCase.vms {
			t.Errorf("Abs(-1) = %d; want 1", got)
		}
	}
}
