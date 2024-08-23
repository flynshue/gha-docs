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
  uses: flynshue/gha-docs/test-action@VERSION
  with:
    thing-on: false
    input-one: one
    input-two: two
```