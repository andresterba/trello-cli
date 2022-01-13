package commands

import "testing"

func TestShoppingListIsForCommand(t *testing.T) {
	slc := shoppingListCommand{}
	slc.subCommands = make(map[string]subCommandFunction)
	slc.registerSubCommands()

	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{[]string{"test"}, false},
		{[]string{"shopping-list"}, true},
		{[]string{"sl"}, true},
	}

	for _, test := range tests {
		got := slc.IsForCommand(test.input)

		if got != test.want {
			t.Errorf("Expected %t, but got %t for params %s", test.want, got, test.input)
		}
	}
}
