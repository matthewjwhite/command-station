package render

import (
	"bytes"
	"github.com/matthewjwhite/command-station/command"
	"text/template"
)

type renderData struct {
	Endpoint string
	Commands []command.Command
}

const stationHTML = `
<link href="https://kristopolous.github.io/BOOTSTRA.386/assets/css/bootstrap.css" rel="stylesheet">
<center>
<div class="page-header">ALL HANDS MAN YOUR BATTLE STATIONS</div>
{{- $endpoint := .Endpoint -}}
{{ range .Commands }}
    <a class="btn btn-primary btn-large" href="/{{ $endpoint }}/{{.Name}}">{{.Name}}</a>
{{ end }}
</center>
`

// Station evaluates the final station code and writes it to the provided writer.
func Station(commands []command.Command, endpoint string) ([]byte, error) {
	t, err := template.New("station").Parse(stationHTML)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	if err = t.Execute(&buffer, renderData{endpoint, commands}); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
