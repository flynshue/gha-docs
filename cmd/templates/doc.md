# Overview
{{ .Name}} {{ .Description}}

{{if .Inputs}}
## Inputs
{{getInputs .Inputs}}
{{end}}

{{if .Outputs}}
## Outputs
{{getOutputs .Outputs}}
{{end}}