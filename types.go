package main

import (
	"encoding/xml"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Track   GPXTrack `xml:"trk"`
}

type GPXTrack struct {
	Name    string          `xml:"name"`
	Type    string          `xml:"type"`
	Segment GPXTrackSegment `xml:"trkseg"`
}

type GPXTrackSegment struct {
	Points []GPXTrackPoint `xml:"trkpt"`
}

type GPXTrackPoint struct {
	Lat  float64 `xml:"lat,attr"`
	Lon  float64 `xml:"lon,attr"`
	Ele  float64 `xml:"ele"`
	Time string  `xml:"time"`
}
