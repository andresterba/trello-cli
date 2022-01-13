package commands

import "testing"

func TestVersionIsForCommand(t *testing.T) {
	cc := versionCommand{}
	cc.registerSubCommands()

	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{[]string{"test"}, false},
		{[]string{"version"}, true},
	}

	for _, test := range tests {
		got := cc.IsForCommand(test.input)

		if got != test.want {
			t.Errorf("Expected %t, but got %t for params %s", test.want, got, test.input)
		}
	}
}
