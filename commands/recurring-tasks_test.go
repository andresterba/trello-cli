package commands

import "testing"

func TestRecurringTasksIsForCommand(t *testing.T) {
	rc := recurringCommand{}
	rc.subCommands = make(map[string]subCommandFunction)
	rc.registerSubCommands()

	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{[]string{"test"}, false},
		{[]string{"recurring"}, true},
		{[]string{"r"}, true},
	}

	for _, test := range tests {
		got := rc.IsForCommand(test.input)

		if got != test.want {
			t.Errorf("Expected %t, but got %t for params %s", test.want, got, test.input)
		}
	}
}
