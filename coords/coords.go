package coords

import (
	"encoding/json"
	"encoding/xml"
	"log"
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

func Coords(data []byte) (string, error) {
	var gpx GPX
  var err error

  err = xml.Unmarshal(data, &gpx)
	if err != nil {
		log.Fatal(err)
	}

	coords := [][]float64{}
	for _, trkpt := range gpx.Trk.Trkseg.Trkpt {
		lat, err := strconv.ParseFloat(trkpt.Lat, 64)
		if err != nil {
			log.Fatal(err)
		}

		lon, err := strconv.ParseFloat(trkpt.Lon, 64)
		if err != nil {
			log.Fatal(err)
		}

		coords = append(coords, []float64{lat, lon})
	}

	jsonCoords, err := json.Marshal(coords)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonCoords), err
}
