package cmd

import (
	"fmt"
	"os"
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

func Example_getInputTable() {
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

func TestActionGetPath(t *testing.T) {
	a := &Action{}
	a.getPath("../test-action/action.yml")
	if a.ActionDir == "" {
		t.Error()
	}
	notWant := "flynshue/gha-docs/test-action@VERSION"
	if a.Use == notWant {
		t.Error()
	}
	fmt.Println(a.Use)
}

func ExampleGenUsage_WithDefault() {
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

func ExampleGenUsage_WithOutDefault() {
	a := Action{
		Name: "Test Action",
		Use:  "test-action",
		Inputs: map[string]Input{
			"input-one": {Default: ""},
		},
	}
	fmt.Println(genUsage(a))

	// Output:
	// - name: Test Action
	//   uses: test-action
	//   with:
	//     input-one: <CHANGEME>
}

func Test_parseGitUrl(t *testing.T) {
	testCases := []struct {
		name   string
		gitUrl string
		owner  string
		repo   string
	}{
		{"ssh", "git@github.com:flynshue/gha-docs.git", "flynshue", "gha-docs"},
		{"https", "https://github.com/flynshue/gha-docs.git", "flynshue", "gha-docs"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			owner, repo := parseGitUrl(tc.gitUrl)
			if tc.owner != owner {
				t.Errorf("got: %s; wanted: %s", owner, tc.owner)
			}
			if tc.repo != repo {
				t.Errorf("got: %s; wanted: %s", repo, tc.repo)
			}
		})
	}
}

func Test_GitRepo(t *testing.T) {
	t.Run("validGitRepo", func(t *testing.T) {
		_, err := gitRepoUrl(".")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("invalidGitRepo", func(t *testing.T) {
		repoDir := "/tmp/fake-action"
		if err := os.Mkdir(repoDir, 0775); err != nil {
			t.Error(err)
		}
		defer os.RemoveAll(repoDir)
		b, err := os.ReadFile("../test-action/action.yml")
		if err != nil {
			t.Error(err)
		}
		if err := os.WriteFile(repoDir+"/action.yml", b, 0644); err != nil {
			t.Error(err)
		}
		_, err = gitRepoUrl(repoDir)
		if err == nil {
			t.Error("valid git repo found; should not be a git repository")
		}
	})
}

func Test_gitTag(t *testing.T) {
	tag := gitTag("./")
	if tag == "VERSION" {
		t.Error()
	}
	fmt.Println(tag)
}
