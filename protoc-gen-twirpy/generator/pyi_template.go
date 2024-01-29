package generator

import "text/template"

// TwirpTemplate - Template for twirp server and client
var PyiTwirpTemplate = template.Must(template.New("PyiTwirpTemplate").Parse(`# -*- coding: utf-8 -*-
# Generated by https://github.com/verloop/twirpy/protoc-gen-twirpy.  DO NOT EDIT!
# source: {{.FileName}}

from typing import Protocol
from twirp.server import TwirpServer

{{- range .Imports}}
{{if .From}}from {{.From}} {{end}}import {{.Import}} as {{.Alias}}
{{- end}}
{{range .Services}}
class {{.Name}}ServiceProtocol(Protocol):
	{{- range .Methods }}
	def {{.Name}}(self, request: {{.QualifiedInput}}) -> {{.QualifiedOutput}}: ...
	{{- end }}

class {{.Name}}Server(TwirpServer):
	def __init__(self, *args, service: DataPlatformServiceProtocol, server_path_prefix: str = "/twirp"): ...
{{end}}`))