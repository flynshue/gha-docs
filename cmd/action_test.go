package cmd

import (
	"fmt"
	"reflect"
	"testing"
)

var testAction = &Action{
	Name:        "Test Action",
	Description: "Description for fake Test Action",
	Inputs: map[string]Input{
		"input-one": {"Description for input-one", true, "one", "string"},
		"input-two": {"Description for input-two", false, "two", "string"},
		"thing-on":  {"Turn thing on", false, "false", "boolean"},
	},
	Outputs: map[string]Output{
		"thing-one": {"Description for output thing-one"},
	},
}

func Test_readActionFile(t *testing.T) {
	got, err := readActionFile("../test-action/action.yml")
	if !reflect.DeepEqual(testAction, got) {
		fmt.Println("got: ", got)
		fmt.Println("want: ", testAction)
		t.Error(err)
	}
}

func TestAction_setInputType(t *testing.T) {
	testCases := []struct {
		name      string
		inputType string
		want      string
	}{
		{"empty", "", "string"},
		{"bool", "boolean", "boolean"},
		{"int", "int", "int"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			action := &Action{
				Inputs: map[string]Input{
					tc.name: {InputType: tc.inputType},
				},
			}
			action.setInputType()
			got := action.Inputs[tc.name].InputType
			if got != tc.want {
				t.Errorf("got %s, want %s\n", got, tc.want)
			}
		})
	}
}

func ExampleGetInputTable() {
	inputs := map[string]Input{
		"input-one": {"Description for input-one", true, "one", "string"},
	}
	fmt.Println(getInputTable(inputs))
	// Output:
	// | Name | Type | Default | Description | Required |
	// | --- | --- | --- | --- | --- |
	// | input-one | string | one | Description for input-one | true |
}

func ExampleGetOutputTable() {
	outputs := map[string]Output{
		"thing-one": {"Description for output thing-one"},
	}
	fmt.Println(getOutputTable(outputs))

	// Output:
	// | Name | Description |
	// | --- | --- |
	// | thing-one | Description for output thing-one |
}

func Test_getBasePath(t *testing.T) {
	want := "test-action"
	got := getBasePath("../test-action/action.yml")
	if got != want {
		t.Errorf("got %s; wanted %s", got, want)
	}
}

func TestActionGetPath(t *testing.T) {
	a := &Action{}
	a.getPath("../test-action/action.yml")
	if a.ActionDir == "" {
		t.Error()
	}
	if a.Use != "test-action" {
		t.Error()
	}
}

func ExampleGenUsage() {
	a := Action{
		Name: "Test Action",
		Use:  "test-action",
		Inputs: map[string]Input{
			"input-one": {Default: "one"},
		},
	}
	fmt.Println(genUsage(a))

	// Output:
	// - name: Test Action
	//   uses: test-action
	//   with:
	//     input-one: one
}
