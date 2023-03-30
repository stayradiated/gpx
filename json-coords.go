package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
)

func Coords(data []byte) (string, error) {
	var gpx GPX
	var err error

	err = xml.Unmarshal(data, &gpx)
	if err != nil {
		log.Fatal(err)
	}

	coords := [][]float64{}
	for _, trkpt := range gpx.Track.Segment.Points {
		coords = append(coords, []float64{trkpt.Lat, trkpt.Lon})
	}

	jsonCoords, err := json.Marshal(coords)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonCoords), err
}
