package address

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/deskr/gopkg/types/country"
	"github.com/deskr/gopkg/types/geo"
)

var postalCodeInfos []PostalCodeInfo
var postalCodeInfosByCountryCode map[country.Code]map[string]PostalCodeInfo

func init() {
	csvFile, err := os.Open("post.csv.gz")
	if err != nil {
		panic(err)
	}

	gr, err := gzip.NewReader(csvFile)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(gr)

	postalCodeInfos = make([]PostalCodeInfo, 0)
	postalCodeInfosByCountryCode = make(map[country.Code]map[string]PostalCodeInfo)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		postalCodeInfo := parsePostalCodeInfo(record)

		postalCodeInfos = append(postalCodeInfos, postalCodeInfo)
		if postalCodeInfosByCountryCode[postalCodeInfo.CountryCode] == nil {
			postalCodeInfosByCountryCode[postalCodeInfo.CountryCode] = make(map[string]PostalCodeInfo)
		}
		postalCodeInfosByCountryCode[postalCodeInfo.CountryCode][postalCodeInfo.Code] = postalCodeInfo
	}
}

func parsePostalCodeInfo(record []string) PostalCodeInfo {

	lat, err := strconv.ParseFloat(record[5], 32)
	if err != nil {
		lat = 0
	}
	lon, err := strconv.ParseFloat(record[6], 32)
	if err != nil {
		lon = 0
	}

	return PostalCodeInfo{
		CountryCode: country.Code(strings.ToUpper(record[0])),
		Code:        record[1],
		F1:          record[2],
		F2:          record[3],
		F3:          record[4],
		Geo: geo.Geo{
			Latitude:  float32(lat),
			Longitude: float32(lon),
		},
	}
}
