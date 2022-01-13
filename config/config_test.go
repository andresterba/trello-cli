package config

import (
	"fmt"
	"testing"
)

func testForString(test *testing.T, expectedOutput string, output string) {
	if output != expectedOutput {
		fmt.Printf("Expected %s, got %s\n", expectedOutput, output)
		test.FailNow()
	}
}

func TestGetConfig(test *testing.T) {
	config, err := LoadConfig("./test-resources/trello-cli.json.test")
	if err != nil {
		test.Errorf("%s", err)
		test.FailNow()
	}

	testForString(test, "my-app-key-1337", config.AppKey)
	testForString(test, "my-token-1337", config.Token)
	testForString(test, "board-id-personal", config.PersonalConfig.BoardID)
	testForString(test, "Recurring 1", config.PersonalConfig.RecurringTasks[0].Name)
	testForString(test, "Recurring 2", config.PersonalConfig.RecurringTasks[1].Name)
	testForString(test, "board-id-shopping", config.ShoppingConfig.BoardID)
	testForString(test, "shopping-list", config.ShoppingConfig.ListCardName)
	testForString(test, "board-id-work", config.WorkConfig.BoardID)
}

func TestPersonalConfig(test *testing.T) {
	config, err := LoadConfig("./test-resources/trello-cli.json.test")
	if err != nil {
		test.Errorf("%s", err)
		test.FailNow()
	}

	testForString(test, "my-app-key-1337", config.AppKey)
	testForString(test, "my-token-1337", config.Token)
	testForString(test, "board-id-personal", config.PersonalConfig.BoardID)
	testForString(test, "Recurring 1", config.PersonalConfig.RecurringTasks[0].Name)
	testForString(test, "some-list-id", config.PersonalConfig.RecurringTasks[0].ListID)
	testForString(test, "Test-Label-1", config.PersonalConfig.RecurringTasks[0].Labels[0])
	testForString(test, "Test-Label-2", config.PersonalConfig.RecurringTasks[0].Labels[1])
	testForString(test, "Recurring 2", config.PersonalConfig.RecurringTasks[1].Name)
	testForString(test, "some-list-id", config.PersonalConfig.RecurringTasks[0].ListID)
	testForString(test, "Test-Label-1", config.PersonalConfig.RecurringTasks[1].Labels[0])
	testForString(test, "Test-Label-2", config.PersonalConfig.RecurringTasks[1].Labels[1])
	testForString(test, "Recurring 3", config.PersonalConfig.RecurringTasks[2].Name)
	testForString(test, "some-list-id", config.PersonalConfig.RecurringTasks[0].ListID)
}
