package commands

import "testing"

func TestListIsForCommand(t *testing.T) {
	lc := listCommand{}
	lc.subCommands = make(map[string]subCommandFunction)
	lc.registerSubCommands()

	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{[]string{"test"}, false},
		{[]string{"l"}, true},
		{[]string{"list"}, true},
	}

	for _, test := range tests {
		got := lc.IsForCommand(test.input)

		if got != test.want {
			t.Errorf("Expected %t, but got %t for params %s", test.want, got, test.input)
		}
	}
}
