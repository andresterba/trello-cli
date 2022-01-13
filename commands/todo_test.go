package commands

import "testing"

func TestTodoIsForCommand(t *testing.T) {
	tc := todoCommand{}
	tc.subCommands = make(map[string]subCommandFunction)
	tc.registerSubCommands()

	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{[]string{"test"}, false},
		{[]string{"todo"}, true},
		{[]string{"t"}, true},
	}

	for _, test := range tests {
		got := tc.IsForCommand(test.input)

		if got != test.want {
			t.Errorf("Expected %t, but got %t for param %s", test.want, got, test.input)
		}
	}
}
