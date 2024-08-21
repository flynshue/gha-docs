# gha-docs
Auto generate documentation for GitHub Actions

## Usage
Will generate docs based off an action.yml file found in the current directory
```bash
gha-docs
```

Generate docs dry run mode
```bash
gha-docs --dry-run
# Overview
Test Action Description for fake Test Action

## Inputs
| Name | Type | Default | Description | Required |
| --- | --- | --- | --- | --- |
| input-one | string | one | Description for input-one | true |
| input-two | string | two | Description for input-two | false |
| thing-on | boolean | false | Turn thing on | false |

## Outputs
| Name | Description |
| --- | --- |
| thing-one | Description for output thing-one |

## Example Usage
```
- name: Test Action
  uses: test-action
  with:
    input-one: one
    input-two: two
    thing-on: false
```
```


Generate docs specify action file and write to README.md in the same directory as the action file
```bash
gha-docs -f test-action/action.yml
```

Generate docs specify action file and output file
```bash
gha-docs -f test-action/action.yml -o /tmp/docs.md
```