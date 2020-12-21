package vbox

import (
	"testing"
)

func TestExec(t *testing.T) {
	var testCases = []struct {
		in  string
		out string
	}{
		{
			in:  "Write-Host hello\n",
			out: "hello",
		},
		{
			in: `$world = "world"
Write-Host hello $world
`,
			out: "hello world",
		},
	}
	for _, testCase := range testCases {
		cmd := NewCmd("Powershell.exe", "-NoProfile")
		cmd.Write(testCase.in)
		// got := cmd.Read()
		cmd.Close()
		// if got != testCase.out {
		// 	t.Errorf("Expected %v, but got %v", testCase.out, got)
		// }
	}
}
