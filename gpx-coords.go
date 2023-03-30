package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
)

type TomlData struct {
	Type        string
	Start       string
	End         string
	Date        string
	Mode        string
	Distance    int
	Duration    int
	Coordinates [][]float64
}

func ConvertToGpx(data []byte) (string, error) {
	data = bytes.Trim(data, "-\n")

	// Decode TOML data into struct
	var tomlData TomlData
	if err := toml.Unmarshal(data, &tomlData); err != nil {
		return "", err
	}

	// Create GPX struct
	gpx := GPX{
		Track: GPXTrack{
			Name: fmt.Sprintf("%s to %s via %s", tomlData.Start, tomlData.End, tomlData.Mode),
			Segment: GPXTrackSegment{
				Points: make([]GPXTrackPoint, len(tomlData.Coordinates)),
			},
		},
	}

	// Copy coordinates to GPX struct
	for i, coord := range tomlData.Coordinates {
		gpx.Track.Segment.Points[i].Lat = coord[0]
		gpx.Track.Segment.Points[i].Lon = coord[1]
	}

	// Encode GPX struct into XML
	var buf bytes.Buffer
	if err := xml.NewEncoder(&buf).Encode(gpx); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}
