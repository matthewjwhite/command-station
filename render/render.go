package render

import (
	"bytes"
	"github.com/matthewjwhite/command-station/asset"
	"github.com/matthewjwhite/command-station/config"
	"text/template"
)

type renderData struct {
	Endpoint string
	Config   config.Config
}

var stationHTML = string(asset.MustAsset("template/station.html"))
var stationTemplate = template.Must(template.New("station").Parse(stationHTML))

// Station evaluates the final station code and writes it to the provided writer.
func Station(config config.Config, endpoint string) ([]byte, error) {
	var buffer bytes.Buffer
	if err := stationTemplate.Execute(&buffer, renderData{endpoint, config}); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
