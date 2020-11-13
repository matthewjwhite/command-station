package render

import (
	"bytes"
	"github.com/matthewjwhite/command-station/config"
	"text/template"
)

type renderData struct {
	Endpoint string
	Config   config.Config
}

const stationHTML = `
<link rel="preconnect" href="https://fonts.gstatic.com">
<link href="https://fonts.googleapis.com/css2?family=VT323&display=swap" rel="stylesheet"> 
<style>
* {
  font-family: vt323;
  text-transform: uppercase;
  font-size: 30px;
}

body {
  background: black;
}

h1 {
  color: #00ff1a;
  font-size: 50px;
}

a {
  background: grey;
  padding: 20px 20px;
  text-decoration: none;
  color: black;
  box-shadow: 5px 5px #3b453c;
  margin: 15px;
}

a:active {
  padding: 18px 18px;
  box-shadow: 3px 3px #3b453c;
}

a:hover {
  background:#34eb4c;
}
</style>
<center>
<h1>{{ .Config.Title }}</h1>
<br>
{{- $endpoint := .Endpoint -}}
{{ range .Config.Commands }}
    <a href="/{{ $endpoint }}/{{.Name}}">{{.Name}}</a>
{{ end }}
</center>
`

// Station evaluates the final station code and writes it to the provided writer.
func Station(config config.Config, endpoint string) ([]byte, error) {
	t, err := template.New("station").Parse(stationHTML)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	if err = t.Execute(&buffer, renderData{endpoint, config}); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
