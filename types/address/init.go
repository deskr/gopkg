package address

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/deskr/gopkg/types/country"
	"github.com/deskr/gopkg/types/geo"
)

var postalCodeInfosByCountryCode map[country.Code]map[string]PostalCodeInfo

func loadPostalCodeInfoByCountryCode(cc country.Code, code string) (PostalCodeInfo, bool) {
	if v, ok := postalCodeInfosByCountryCode[cc]; ok {
		return v[code], true
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	csvFile, err := os.Open(path.Join(dir,
		fmt.Sprintf("data/p%s.csv", cc.String())))
	if err != nil {
		csvFile, err = os.Open(path.Join("",
			fmt.Sprintf("data/p%s.csv", cc.String())))
		if err != nil {
			panic(err)
		}
	}

	r := csv.NewReader(csvFile)

	postalCodeInfosByCountryCode = make(map[country.Code]map[string]PostalCodeInfo)

	var infos = make(map[string]PostalCodeInfo)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		postalCodeInfo := parsePostalCodeInfo(record)
		postalCodeInfo.CountryCode = cc

		infos[postalCodeInfo.Code] = postalCodeInfo
	}

	postalCodeInfosByCountryCode[cc] = infos

	csvFile.Close()

	if v, ok := infos[code]; ok {
		return v, true
	}
	return PostalCodeInfo{}, false
}

func parsePostalCodeInfo(record []string) PostalCodeInfo {

	lat, err := strconv.ParseFloat(record[4], 32)
	if err != nil {
		lat = 0
	}
	lon, err := strconv.ParseFloat(record[5], 32)
	if err != nil {
		lon = 0
	}

	return PostalCodeInfo{
		Code: record[0],
		F1:   record[1],
		F2:   record[2],
		F3:   record[3],
		Geo: geo.Geo{
			Latitude:  float32(lat),
			Longitude: float32(lon),
		},
	}
}
