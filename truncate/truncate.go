package truncate

import (
	"encoding/xml"
	"log"
	"math"
	"strconv"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Trk     Trk      `xml:"trk"`
}

type Trk struct {
	Trkseg Trkseg `xml:"trkseg"`
}

type Trkseg struct {
	Trkpt []Trkpt `xml:"trkpt"`
}

type Trkpt struct {
	Lat  string `xml:"lat,attr"`
	Lon  string `xml:"lon,attr"`
	Ele  string `xml:"ele"`
	Time string `xml:"time"`
}

func Truncate(data []byte, precision int) (string, error) {
	var gpx GPX
  var err error

	err = xml.Unmarshal(data, &gpx)
	if err != nil {
		log.Fatal(err)
	}

	for i, trkpt := range gpx.Trk.Trkseg.Trkpt {
		lat, err := strconv.ParseFloat(trkpt.Lat, 64)
		if err != nil {
			log.Fatal(err)
		}

		lon, err := strconv.ParseFloat(trkpt.Lon, 64)
		if err != nil {
			log.Fatal(err)
		}

		lat = math.Round(lat*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
		lon = math.Round(lon*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))

		gpx.Trk.Trkseg.Trkpt[i].Lat = strconv.FormatFloat(lat, 'f', -1, 64)
		gpx.Trk.Trkseg.Trkpt[i].Lon = strconv.FormatFloat(lon, 'f', -1, 64)
	}

	output, err := xml.MarshalIndent(gpx, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(output), err
}
