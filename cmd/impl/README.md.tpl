# {{ .Name }}

{{ .Description }}
## Overview
{{ .Documentation }}
{{ if ne (len .Inputs) 0 }}
## Inputs

| Input | Description | Default | Required |
| -------- | -------- | -------- | ------- |
{{ range $k, $v := .Inputs -}}
| {{ $k }} | {{ $v.Description }} | {{ $v.Default }} | {{ $v.Required }} |
{{ end }}
{{- end }}
{{ if ne (len .Outputs) 0 -}}
## Outputs

| Output | Description |
| --- | --- |
{{ range $k, $v := .Outputs -}}
| {{ $k }} | {{ $v.Description }} |
{{ end }}
{{- end -}}