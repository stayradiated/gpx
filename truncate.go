package main

import (
	"encoding/xml"
	"log"
	"math"
)

func Truncate(data []byte, precision int) (string, error) {
	var gpx GPX
	var err error

	err = xml.Unmarshal(data, &gpx)
	if err != nil {
		log.Fatal(err)
	}

	for i, trkpt := range gpx.Track.Segment.Points {
		lat := math.Round(trkpt.Lat*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
		lon := math.Round(trkpt.Lon*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))

		gpx.Track.Segment.Points[i].Lat = lat
		gpx.Track.Segment.Points[i].Lon = lon
	}

	output, err := xml.MarshalIndent(gpx, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(output), err
}
