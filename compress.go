package main

import (
	"encoding/xml"
	"log"
	"math"
)

func Compress(data []byte, count int) (string, error) {
	var gpx GPX
	var err error

	err = xml.Unmarshal(data, &gpx)
	if err != nil {
		log.Fatal(err)
	}

	seg := gpx.Track.Segment
	if len(seg.Points) > count {
		seg.Points = compress(seg.Points, count)
		gpx.Track.Segment = seg
	}

	output, err := xml.MarshalIndent(gpx, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(output), err
}

func compress(points []GPXTrackPoint, count int) []GPXTrackPoint {
	newPoints := make([]GPXTrackPoint, count)

	// always include first and last point
	newPoints[0] = points[0]
	newPoints[count-1] = points[len(points)-1]

	// compute number of points to skip
	skip := float64(len(points)-2) / float64(count-2)
	skipCount := 0.0

	// copy remaining points uniformly
	for i := 1; i < count-1; i++ {
		skipCount += skip
		skipWhole := math.Floor(skipCount)
		skipRemainder := skipCount - skipWhole

		newPoints[i].Lat = points[int(skipWhole)+1].Lat*skipRemainder + points[int(skipWhole)].Lat*(1-skipRemainder)
		newPoints[i].Lon = points[int(skipWhole)+1].Lon*skipRemainder + points[int(skipWhole)].Lon*(1-skipRemainder)
		newPoints[i].Ele = points[int(skipWhole)+1].Ele
	}

	return newPoints
}
