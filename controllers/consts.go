// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package controllers

import "time"

// SHARED CONSTANTS
const (
	requeueInterval = 15 * time.Second
)

// AGENT POOL CONTROLLER'S CONSTANTS
const (
	agentPoolFinalizer = "agentpool.app.terraform.io/finalizer"
)

// MODULE CONTROLLER'S CONSTANTS
const (
	requeueConfigurationUploadInterval = 10 * time.Second
	requeueNewRunInterval              = 10 * time.Second
	requeueRunStatusInterval           = 30 * time.Second
	moduleFinalizer                    = "module.app.terraform.io/finalizer"

	moduleTemplate = `
{{- if .Variables }}
  {{ range $v := .Variables }}
variable "{{ $v.Name }}" {}
  {{- end}}
{{- end }}

module "this" {
  source = "{{ .Module.Source }}"
{{- if .Module.Version }}
  version = "{{ .Module.Version }}"
{{- end }}

{{- if .Variables }}
  {{ range $v := .Variables }}
    {{ $v.Name }} = var.{{ $v.Name }}
  {{- end}}
{{- end }}
}

{{- if .Outputs }}
  {{ range $o := .Outputs }}
output "{{ $o.Name }}" {
  value = module.this.{{ $o.Name }}
  sensitive = {{ $o.Sensitive }}
}
  {{- end}}
{{- end }}
`
)

// WORKSPACE CONTROLLER'S CONSTANTS
const (
	workspaceFinalizer = "workspace.app.terraform.io/finalizer"
)
