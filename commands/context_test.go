package commands

import "testing"

func TestContextIsForCommand(t *testing.T) {
	cc := contextCommand{}
	cc.subCommands = make(map[string]subCommandFunction)
	cc.registerSubCommands()

	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{[]string{"test"}, false},
		{[]string{"context"}, true},
		{[]string{"c"}, true},
	}

	for _, test := range tests {
		got := cc.IsForCommand(test.input)

		if got != test.want {
			t.Errorf("Expected %t, but got %t for params %s", test.want, got, test.input)
		}
	}
}
