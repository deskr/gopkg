package country

import (
	"bytes"
	"encoding/json"
)

var countriesByCode map[Code]Country

func init() {
	if err := json.NewDecoder(bytes.NewReader([]byte(countriesJSON))).Decode(&countriesByCode); err != nil {
		panic(err)
	}
}

const countriesJSON = `{
  "AD": {
    "code": "AD",
    "name": "Andorra",
    "capital": "Andorra la Vella",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "ca"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AE": {
    "code": "AE",
    "name": "United Arab Emirates",
    "capital": "Abu Dhabi",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "AED",
    "languageCodes": [
      "ar",
      "fa",
      "en",
      "hi",
      "ur"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AF": {
    "code": "AF",
    "name": "Afghanistan",
    "capital": "Kabul",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "AFN",
    "languageCodes": [
      "fa",
      "ps",
      "uz",
      "tk"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AG": {
    "code": "AG",
    "name": "Antigua and Barbuda",
    "capital": "St. John's",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AI": {
    "code": "AI",
    "name": "Anguilla",
    "capital": "The Valley",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AL": {
    "code": "AL",
    "name": "Albania",
    "capital": "Tirana",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "ALL",
    "languageCodes": [
      "sq",
      "el"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AM": {
    "code": "AM",
    "name": "Armenia",
    "capital": "Yerevan",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "AMD",
    "languageCodes": [
      "hy"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AO": {
    "code": "AO",
    "name": "Angola",
    "capital": "Luanda",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "AOA",
    "languageCodes": [
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AQ": {
    "code": "AQ",
    "name": "Antarctica",
    "capital": "",
    "continent": "AN",
    "continentName": "Antarctica",
    "currencyCode": "",
    "languageCodes": null,
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AR": {
    "code": "AR",
    "name": "Argentina",
    "capital": "Buenos Aires",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "ARS",
    "languageCodes": [
      "es",
      "en",
      "it",
      "de",
      "fr",
      "gn"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AS": {
    "code": "AS",
    "name": "American Samoa",
    "capital": "Pago Pago",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "en",
      "sm",
      "to"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AT": {
    "code": "AT",
    "name": "Austria",
    "capital": "Vienna",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "de",
      "hr",
      "hu",
      "sl"
    ],
    "vatRates": {
      "standard": 20,
      "reduced1": 10,
      "reduced2": 13
    }
  },
  "AU": {
    "code": "AU",
    "name": "Australia",
    "capital": "Canberra",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "AUD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AW": {
    "code": "AW",
    "name": "Aruba",
    "capital": "Oranjestad",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "AWG",
    "languageCodes": [
      "nl",
      "es",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AX": {
    "code": "AX",
    "name": "\u00c5land",
    "capital": "Mariehamn",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "sv"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "AZ": {
    "code": "AZ",
    "name": "Azerbaijan",
    "capital": "Baku",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "AZN",
    "languageCodes": [
      "az",
      "ru",
      "hy"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BA": {
    "code": "BA",
    "name": "Bosnia and Herzegovina",
    "capital": "Sarajevo",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "BAM",
    "languageCodes": [
      "bs",
      "hr",
      "sr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BB": {
    "code": "BB",
    "name": "Barbados",
    "capital": "Bridgetown",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "BBD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BD": {
    "code": "BD",
    "name": "Bangladesh",
    "capital": "Dhaka",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "BDT",
    "languageCodes": [
      "bn",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BE": {
    "code": "BE",
    "name": "Belgium",
    "capital": "Brussels",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "nl",
      "fr",
      "de"
    ],
    "vatRates": {
      "standard": 21,
      "reduced1": 12,
      "reduced2": 6
    }
  },
  "BF": {
    "code": "BF",
    "name": "Burkina Faso",
    "capital": "Ouagadougou",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BG": {
    "code": "BG",
    "name": "Bulgaria",
    "capital": "Sofia",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "BGN",
    "languageCodes": [
      "bg",
      "tr"
    ],
    "vatRates": {
      "standard": 20,
      "reduced1": 9,
      "reduced2": 0
    }
  },
  "BH": {
    "code": "BH",
    "name": "Bahrain",
    "capital": "Manama",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "BHD",
    "languageCodes": [
      "ar",
      "en",
      "fa",
      "ur"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BI": {
    "code": "BI",
    "name": "Burundi",
    "capital": "Bujumbura",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "BIF",
    "languageCodes": [
      "fr",
      "rn"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BJ": {
    "code": "BJ",
    "name": "Benin",
    "capital": "Porto-Novo",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BL": {
    "code": "BL",
    "name": "Saint Barth\u00e9lemy",
    "capital": "Gustavia",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BM": {
    "code": "BM",
    "name": "Bermuda",
    "capital": "Hamilton",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "BMD",
    "languageCodes": [
      "en",
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BN": {
    "code": "BN",
    "name": "Brunei",
    "capital": "Bandar Seri Begawan",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "BND",
    "languageCodes": [
      "ms",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BO": {
    "code": "BO",
    "name": "Bolivia",
    "capital": "Sucre",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "BOB",
    "languageCodes": [
      "es",
      "qu",
      "ay"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BQ": {
    "code": "BQ",
    "name": "Bonaire",
    "capital": "Kralendijk",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "USD",
    "languageCodes": [
      "nl",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BR": {
    "code": "BR",
    "name": "Brazil",
    "capital": "Bras\u00edlia",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "BRL",
    "languageCodes": [
      "pt",
      "es",
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BS": {
    "code": "BS",
    "name": "Bahamas",
    "capital": "Nassau",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "BSD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BT": {
    "code": "BT",
    "name": "Bhutan",
    "capital": "Thimphu",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "BTN",
    "languageCodes": null,
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BV": {
    "code": "BV",
    "name": "Bouvet Island",
    "capital": "",
    "continent": "AN",
    "continentName": "Antarctica",
    "currencyCode": "NOK",
    "languageCodes": null,
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BW": {
    "code": "BW",
    "name": "Botswana",
    "capital": "Gaborone",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "BWP",
    "languageCodes": [
      "en",
      "tn"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BY": {
    "code": "BY",
    "name": "Belarus",
    "capital": "Minsk",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "BYR",
    "languageCodes": [
      "be",
      "ru"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "BZ": {
    "code": "BZ",
    "name": "Belize",
    "capital": "Belmopan",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "BZD",
    "languageCodes": [
      "en",
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CA": {
    "code": "CA",
    "name": "Canada",
    "capital": "Ottawa",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "CAD",
    "languageCodes": [
      "en",
      "fr",
      "iu"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CC": {
    "code": "CC",
    "name": "Cocos [Keeling] Islands",
    "capital": "West Island",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "AUD",
    "languageCodes": [
      "ms",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CD": {
    "code": "CD",
    "name": "Democratic Republic of the Congo",
    "capital": "Kinshasa",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "CDF",
    "languageCodes": [
      "fr",
      "ln",
      "kg"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CF": {
    "code": "CF",
    "name": "Central African Republic",
    "capital": "Bangui",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XAF",
    "languageCodes": [
      "fr",
      "sg",
      "ln",
      "kg"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CG": {
    "code": "CG",
    "name": "Republic of the Congo",
    "capital": "Brazzaville",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XAF",
    "languageCodes": [
      "fr",
      "kg",
      "ln"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CH": {
    "code": "CH",
    "name": "Switzerland",
    "capital": "Bern",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "CHF",
    "languageCodes": [
      "de",
      "fr",
      "it",
      "rm"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CI": {
    "code": "CI",
    "name": "Ivory Coast",
    "capital": "Yamoussoukro",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CK": {
    "code": "CK",
    "name": "Cook Islands",
    "capital": "Avarua",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "NZD",
    "languageCodes": [
      "en",
      "mi"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CL": {
    "code": "CL",
    "name": "Chile",
    "capital": "Santiago",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "CLP",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CM": {
    "code": "CM",
    "name": "Cameroon",
    "capital": "Yaound\u00e9",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XAF",
    "languageCodes": [
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CN": {
    "code": "CN",
    "name": "China",
    "capital": "Beijing",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "CNY",
    "languageCodes": [
      "zh",
      "ug",
      "za"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CO": {
    "code": "CO",
    "name": "Colombia",
    "capital": "Bogot\u00e1",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "COP",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CR": {
    "code": "CR",
    "name": "Costa Rica",
    "capital": "San Jos\u00e9",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "CRC",
    "languageCodes": [
      "es",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CU": {
    "code": "CU",
    "name": "Cuba",
    "capital": "Havana",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "CUP",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CV": {
    "code": "CV",
    "name": "Cape Verde",
    "capital": "Praia",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "CVE",
    "languageCodes": [
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CW": {
    "code": "CW",
    "name": "Curacao",
    "capital": "Willemstad",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "ANG",
    "languageCodes": [
      "nl"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CX": {
    "code": "CX",
    "name": "Christmas Island",
    "capital": "Flying Fish Cove",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "AUD",
    "languageCodes": [
      "en",
      "zh",
      "ms"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "CY": {
    "code": "CY",
    "name": "Cyprus",
    "capital": "Nicosia",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "el",
      "tr",
      "en"
    ],
    "vatRates": {
      "standard": 19,
      "reduced1": 9,
      "reduced2": 5
    }
  },
  "CZ": {
    "code": "CZ",
    "name": "Czech Republic",
    "capital": "Prague",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "CZK",
    "languageCodes": [
      "cs",
      "sk"
    ],
    "vatRates": {
      "standard": 21,
      "reduced1": 15,
      "reduced2": 10
    }
  },
  "DE": {
    "code": "DE",
    "name": "Germany",
    "capital": "Berlin",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "de"
    ],
    "vatRates": {
      "standard": 19,
      "reduced1": 7,
      "reduced2": 0
    }
  },
  "DJ": {
    "code": "DJ",
    "name": "Djibouti",
    "capital": "Djibouti",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "DJF",
    "languageCodes": [
      "fr",
      "ar",
      "so",
      "aa"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "DK": {
    "code": "DK",
    "name": "Denmark",
    "capital": "Copenhagen",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "DKK",
    "languageCodes": [
      "da",
      "en",
      "fo",
      "de"
    ],
    "vatRates": {
      "standard": 25,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "DM": {
    "code": "DM",
    "name": "Dominica",
    "capital": "Roseau",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "DO": {
    "code": "DO",
    "name": "Dominican Republic",
    "capital": "Santo Domingo",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "DOP",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "DZ": {
    "code": "DZ",
    "name": "Algeria",
    "capital": "Algiers",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "DZD",
    "languageCodes": [
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "EC": {
    "code": "EC",
    "name": "Ecuador",
    "capital": "Quito",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "USD",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "EE": {
    "code": "EE",
    "name": "Estonia",
    "capital": "Tallinn",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "et",
      "ru"
    ],
    "vatRates": {
      "standard": 20,
      "reduced1": 9,
      "reduced2": 0
    }
  },
  "EG": {
    "code": "EG",
    "name": "Egypt",
    "capital": "Cairo",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "EGP",
    "languageCodes": [
      "ar",
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "EH": {
    "code": "EH",
    "name": "Western Sahara",
    "capital": "La\u00e2youne \/ El Aai\u00fan",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "MAD",
    "languageCodes": [
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ER": {
    "code": "ER",
    "name": "Eritrea",
    "capital": "Asmara",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "ERN",
    "languageCodes": [
      "aa",
      "ar",
      "ti"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ES": {
    "code": "ES",
    "name": "Spain",
    "capital": "Madrid",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "es",
      "ca",
      "gl",
      "eu",
      "oc"
    ],
    "vatRates": {
      "standard": 21,
      "reduced1": 10,
      "reduced2": 0
    }
  },
  "ET": {
    "code": "ET",
    "name": "Ethiopia",
    "capital": "Addis Ababa",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "ETB",
    "languageCodes": [
      "am",
      "en",
      "om",
      "ti",
      "so"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "FI": {
    "code": "FI",
    "name": "Finland",
    "capital": "Helsinki",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "fi",
      "sv"
    ],
    "vatRates": {
      "standard": 24,
      "reduced1": 14,
      "reduced2": 10
    }
  },
  "FJ": {
    "code": "FJ",
    "name": "Fiji",
    "capital": "Suva",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "FJD",
    "languageCodes": [
      "en",
      "fj"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "FK": {
    "code": "FK",
    "name": "Falkland Islands",
    "capital": "Stanley",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "FKP",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "FM": {
    "code": "FM",
    "name": "Micronesia",
    "capital": "Palikir",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "FO": {
    "code": "FO",
    "name": "Faroe Islands",
    "capital": "T\u00f3rshavn",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "DKK",
    "languageCodes": [
      "fo",
      "da"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "FR": {
    "code": "FR",
    "name": "France",
    "capital": "Paris",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr",
      "br",
      "co",
      "ca",
      "eu",
      "oc"
    ],
    "vatRates": {
      "standard": 20,
      "reduced1": 10,
      "reduced2": 5.5
    }
  },
  "GA": {
    "code": "GA",
    "name": "Gabon",
    "capital": "Libreville",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XAF",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GB": {
    "code": "GB",
    "name": "United Kingdom",
    "capital": "London",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "GBP",
    "languageCodes": [
      "en",
      "cy",
      "gd"
    ],
    "vatRates": {
      "standard": 20,
      "reduced1": 5,
      "reduced2": 0
    }
  },
  "GD": {
    "code": "GD",
    "name": "Grenada",
    "capital": "St. George's",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GE": {
    "code": "GE",
    "name": "Georgia",
    "capital": "Tbilisi",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "GEL",
    "languageCodes": [
      "ka",
      "ru",
      "hy",
      "az"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GF": {
    "code": "GF",
    "name": "French Guiana",
    "capital": "Cayenne",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GG": {
    "code": "GG",
    "name": "Guernsey",
    "capital": "St Peter Port",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "GBP",
    "languageCodes": [
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GH": {
    "code": "GH",
    "name": "Ghana",
    "capital": "Accra",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "GHS",
    "languageCodes": [
      "en",
      "ak",
      "ee",
      "tw"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GI": {
    "code": "GI",
    "name": "Gibraltar",
    "capital": "Gibraltar",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "GIP",
    "languageCodes": [
      "en",
      "es",
      "it",
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GL": {
    "code": "GL",
    "name": "Greenland",
    "capital": "Nuuk",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "DKK",
    "languageCodes": [
      "kl",
      "da",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GM": {
    "code": "GM",
    "name": "Gambia",
    "capital": "Bathurst",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "GMD",
    "languageCodes": [
      "en",
      "wo",
      "ff"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GN": {
    "code": "GN",
    "name": "Guinea",
    "capital": "Conakry",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "GNF",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GP": {
    "code": "GP",
    "name": "Guadeloupe",
    "capital": "Basse-Terre",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GQ": {
    "code": "GQ",
    "name": "Equatorial Guinea",
    "capital": "Malabo",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XAF",
    "languageCodes": [
      "es",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GR": {
    "code": "GR",
    "name": "Greece",
    "capital": "Athens",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "el",
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 23,
      "reduced1": 13,
      "reduced2": 6
    }
  },
  "GS": {
    "code": "GS",
    "name": "South Georgia and the South Sandwich Islands",
    "capital": "Grytviken",
    "continent": "AN",
    "continentName": "Antarctica",
    "currencyCode": "GBP",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GT": {
    "code": "GT",
    "name": "Guatemala",
    "capital": "Guatemala City",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "GTQ",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GU": {
    "code": "GU",
    "name": "Guam",
    "capital": "Hag\u00e5t\u00f1a",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "en",
      "ch"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GW": {
    "code": "GW",
    "name": "Guinea-Bissau",
    "capital": "Bissau",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "GY": {
    "code": "GY",
    "name": "Guyana",
    "capital": "Georgetown",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "GYD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "HK": {
    "code": "HK",
    "name": "Hong Kong",
    "capital": "Hong Kong",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "HKD",
    "languageCodes": [
      "zh",
      "zh",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "HM": {
    "code": "HM",
    "name": "Heard Island and McDonald Islands",
    "capital": "",
    "continent": "AN",
    "continentName": "Antarctica",
    "currencyCode": "AUD",
    "languageCodes": null,
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "HN": {
    "code": "HN",
    "name": "Honduras",
    "capital": "Tegucigalpa",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "HNL",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "HR": {
    "code": "HR",
    "name": "Croatia",
    "capital": "Zagreb",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "HRK",
    "languageCodes": [
      "hr",
      "sr"
    ],
    "vatRates": {
      "standard": 25,
      "reduced1": 13,
      "reduced2": 5
    }
  },
  "HT": {
    "code": "HT",
    "name": "Haiti",
    "capital": "Port-au-Prince",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "HTG",
    "languageCodes": [
      "ht",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "HU": {
    "code": "HU",
    "name": "Hungary",
    "capital": "Budapest",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "HUF",
    "languageCodes": [
      "hu"
    ],
    "vatRates": {
      "standard": 27,
      "reduced1": 18,
      "reduced2": 5
    }
  },
  "ID": {
    "code": "ID",
    "name": "Indonesia",
    "capital": "Jakarta",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "IDR",
    "languageCodes": [
      "id",
      "en",
      "nl",
      "jv"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IE": {
    "code": "IE",
    "name": "Ireland",
    "capital": "Dublin",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "en",
      "ga"
    ],
    "vatRates": {
      "standard": 23,
      "reduced1": 13.5,
      "reduced2": 9
    }
  },
  "IL": {
    "code": "IL",
    "name": "Israel",
    "capital": "",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "ILS",
    "languageCodes": [
      "he",
      "ar",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IM": {
    "code": "IM",
    "name": "Isle of Man",
    "capital": "Douglas",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "GBP",
    "languageCodes": [
      "en",
      "gv"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IN": {
    "code": "IN",
    "name": "India",
    "capital": "New Delhi",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "INR",
    "languageCodes": [
      "en",
      "hi",
      "bn",
      "te",
      "mr",
      "ta",
      "ur",
      "gu",
      "kn",
      "ml",
      "or",
      "pa",
      "as",
      "bh",
      "ks",
      "ne",
      "sd",
      "sa",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IO": {
    "code": "IO",
    "name": "British Indian Ocean Territory",
    "capital": "",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "USD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IQ": {
    "code": "IQ",
    "name": "Iraq",
    "capital": "Baghdad",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "IQD",
    "languageCodes": [
      "ar",
      "ku",
      "hy"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IR": {
    "code": "IR",
    "name": "Iran",
    "capital": "Tehran",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "IRR",
    "languageCodes": [
      "fa",
      "ku"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IS": {
    "code": "IS",
    "name": "Iceland",
    "capital": "Reykjavik",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "ISK",
    "languageCodes": [
      "is",
      "en",
      "de",
      "da",
      "sv",
      "no"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "IT": {
    "code": "IT",
    "name": "Italy",
    "capital": "Rome",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "it",
      "de",
      "fr",
      "sc",
      "ca",
      "co",
      "sl"
    ],
    "vatRates": {
      "standard": 22,
      "reduced1": 10,
      "reduced2": 4
    }
  },
  "JE": {
    "code": "JE",
    "name": "Jersey",
    "capital": "Saint Helier",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "GBP",
    "languageCodes": [
      "en",
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "JM": {
    "code": "JM",
    "name": "Jamaica",
    "capital": "Kingston",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "JMD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "JO": {
    "code": "JO",
    "name": "Jordan",
    "capital": "Amman",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "JOD",
    "languageCodes": [
      "ar",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "JP": {
    "code": "JP",
    "name": "Japan",
    "capital": "Tokyo",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "JPY",
    "languageCodes": [
      "ja"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KE": {
    "code": "KE",
    "name": "Kenya",
    "capital": "Nairobi",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "KES",
    "languageCodes": [
      "en",
      "sw"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KG": {
    "code": "KG",
    "name": "Kyrgyzstan",
    "capital": "Bishkek",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "KGS",
    "languageCodes": [
      "ky",
      "uz",
      "ru"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KH": {
    "code": "KH",
    "name": "Cambodia",
    "capital": "Phnom Penh",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "KHR",
    "languageCodes": [
      "km",
      "fr",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KI": {
    "code": "KI",
    "name": "Kiribati",
    "capital": "Tarawa",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "AUD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KM": {
    "code": "KM",
    "name": "Comoros",
    "capital": "Moroni",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "KMF",
    "languageCodes": [
      "ar",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KN": {
    "code": "KN",
    "name": "Saint Kitts and Nevis",
    "capital": "Basseterre",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KP": {
    "code": "KP",
    "name": "North Korea",
    "capital": "Pyongyang",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "KPW",
    "languageCodes": [
      "ko"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KR": {
    "code": "KR",
    "name": "South Korea",
    "capital": "Seoul",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "KRW",
    "languageCodes": [
      "ko",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KW": {
    "code": "KW",
    "name": "Kuwait",
    "capital": "Kuwait City",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "KWD",
    "languageCodes": [
      "ar",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KY": {
    "code": "KY",
    "name": "Cayman Islands",
    "capital": "George Town",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "KYD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "KZ": {
    "code": "KZ",
    "name": "Kazakhstan",
    "capital": "Astana",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "KZT",
    "languageCodes": [
      "kk",
      "ru"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LA": {
    "code": "LA",
    "name": "Laos",
    "capital": "Vientiane",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "LAK",
    "languageCodes": [
      "lo",
      "fr",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LB": {
    "code": "LB",
    "name": "Lebanon",
    "capital": "Beirut",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "LBP",
    "languageCodes": [
      "ar",
      "fr",
      "en",
      "hy"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LC": {
    "code": "LC",
    "name": "Saint Lucia",
    "capital": "Castries",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LI": {
    "code": "LI",
    "name": "Liechtenstein",
    "capital": "Vaduz",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "CHF",
    "languageCodes": [
      "de"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LK": {
    "code": "LK",
    "name": "Sri Lanka",
    "capital": "Colombo",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "LKR",
    "languageCodes": [
      "si",
      "ta",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LR": {
    "code": "LR",
    "name": "Liberia",
    "capital": "Monrovia",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "LRD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LS": {
    "code": "LS",
    "name": "Lesotho",
    "capital": "Maseru",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "LSL",
    "languageCodes": [
      "en",
      "st",
      "xh"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "LT": {
    "code": "LT",
    "name": "Lithuania",
    "capital": "Vilnius",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "lt",
      "ru",
      "pl"
    ],
    "vatRates": {
      "standard": 21,
      "reduced1": 9,
      "reduced2": 5
    }
  },
  "LU": {
    "code": "LU",
    "name": "Luxembourg",
    "capital": "Luxembourg",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "lb",
      "de",
      "fr"
    ],
    "vatRates": {
      "standard": 17,
      "reduced1": 14,
      "reduced2": 8
    }
  },
  "LV": {
    "code": "LV",
    "name": "Latvia",
    "capital": "Riga",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "lv",
      "ru",
      "lt"
    ],
    "vatRates": {
      "standard": 21,
      "reduced1": 12,
      "reduced2": 0
    }
  },
  "LY": {
    "code": "LY",
    "name": "Libya",
    "capital": "Tripoli",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "LYD",
    "languageCodes": [
      "ar",
      "it",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MA": {
    "code": "MA",
    "name": "Morocco",
    "capital": "Rabat",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "MAD",
    "languageCodes": [
      "ar",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MC": {
    "code": "MC",
    "name": "Monaco",
    "capital": "Monaco",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr",
      "en",
      "it"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MD": {
    "code": "MD",
    "name": "Moldova",
    "capital": "Chi\u015fin\u0103u",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "MDL",
    "languageCodes": [
      "ro",
      "ru",
      "tr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ME": {
    "code": "ME",
    "name": "Montenegro",
    "capital": "Podgorica",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "sr",
      "hu",
      "bs",
      "sq",
      "hr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MF": {
    "code": "MF",
    "name": "Saint Martin",
    "capital": "Marigot",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MG": {
    "code": "MG",
    "name": "Madagascar",
    "capital": "Antananarivo",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "MGA",
    "languageCodes": [
      "fr",
      "mg"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MH": {
    "code": "MH",
    "name": "Marshall Islands",
    "capital": "Majuro",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "mh",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MK": {
    "code": "MK",
    "name": "Macedonia",
    "capital": "Skopje",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "MKD",
    "languageCodes": [
      "mk",
      "sq",
      "tr",
      "sr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ML": {
    "code": "ML",
    "name": "Mali",
    "capital": "Bamako",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "fr",
      "bm"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MM": {
    "code": "MM",
    "name": "Myanmar [Burma]",
    "capital": "Naypyitaw",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "MMK",
    "languageCodes": [
      "my"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MN": {
    "code": "MN",
    "name": "Mongolia",
    "capital": "Ulan Bator",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "MNT",
    "languageCodes": [
      "mn",
      "ru"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MO": {
    "code": "MO",
    "name": "Macao",
    "capital": "Macao",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "MOP",
    "languageCodes": [
      "zh",
      "zh",
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MP": {
    "code": "MP",
    "name": "Northern Mariana Islands",
    "capital": "Saipan",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "tl",
      "zh",
      "ch",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MQ": {
    "code": "MQ",
    "name": "Martinique",
    "capital": "Fort-de-France",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MR": {
    "code": "MR",
    "name": "Mauritania",
    "capital": "Nouakchott",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "MRO",
    "languageCodes": [
      "ar",
      "fr",
      "wo"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MS": {
    "code": "MS",
    "name": "Montserrat",
    "capital": "Plymouth",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MT": {
    "code": "MT",
    "name": "Malta",
    "capital": "Valletta",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "mt",
      "en"
    ],
    "vatRates": {
      "standard": 18,
      "reduced1": 7,
      "reduced2": 5
    }
  },
  "MU": {
    "code": "MU",
    "name": "Mauritius",
    "capital": "Port Louis",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "MUR",
    "languageCodes": [
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MV": {
    "code": "MV",
    "name": "Maldives",
    "capital": "Mal\u00e9",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "MVR",
    "languageCodes": [
      "dv",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MW": {
    "code": "MW",
    "name": "Malawi",
    "capital": "Lilongwe",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "MWK",
    "languageCodes": [
      "ny"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MX": {
    "code": "MX",
    "name": "Mexico",
    "capital": "Mexico City",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "MXN",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MY": {
    "code": "MY",
    "name": "Malaysia",
    "capital": "Kuala Lumpur",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "MYR",
    "languageCodes": [
      "ms",
      "en",
      "zh",
      "ta",
      "te",
      "ml",
      "pa",
      "th"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "MZ": {
    "code": "MZ",
    "name": "Mozambique",
    "capital": "Maputo",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "MZN",
    "languageCodes": [
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NA": {
    "code": "NA",
    "name": "Namibia",
    "capital": "Windhoek",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "NAD",
    "languageCodes": [
      "en",
      "af",
      "de",
      "hz"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NC": {
    "code": "NC",
    "name": "New Caledonia",
    "capital": "Noumea",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "XPF",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NE": {
    "code": "NE",
    "name": "Niger",
    "capital": "Niamey",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "fr",
      "ha",
      "kr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NF": {
    "code": "NF",
    "name": "Norfolk Island",
    "capital": "Kingston",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "AUD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NG": {
    "code": "NG",
    "name": "Nigeria",
    "capital": "Abuja",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "NGN",
    "languageCodes": [
      "en",
      "ha",
      "yo",
      "ig",
      "ff"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NI": {
    "code": "NI",
    "name": "Nicaragua",
    "capital": "Managua",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "NIO",
    "languageCodes": [
      "es",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NL": {
    "code": "NL",
    "name": "Netherlands",
    "capital": "Amsterdam",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "nl",
      "fy"
    ],
    "vatRates": {
      "standard": 21,
      "reduced1": 6,
      "reduced2": 0
    }
  },
  "NO": {
    "code": "NO",
    "name": "Norway",
    "capital": "Oslo",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "NOK",
    "languageCodes": [
      "no",
      "nb",
      "nn",
      "se",
      "fi"
    ],
    "vatRates": {
      "standard": 25,
      "reduced1": 15,
      "reduced2": 8
    }
  },
  "NP": {
    "code": "NP",
    "name": "Nepal",
    "capital": "Kathmandu",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "NPR",
    "languageCodes": [
      "ne",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NR": {
    "code": "NR",
    "name": "Nauru",
    "capital": "Yaren",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "AUD",
    "languageCodes": [
      "na",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NU": {
    "code": "NU",
    "name": "Niue",
    "capital": "Alofi",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "NZD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "NZ": {
    "code": "NZ",
    "name": "New Zealand",
    "capital": "Wellington",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "NZD",
    "languageCodes": [
      "en",
      "mi"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "OM": {
    "code": "OM",
    "name": "Oman",
    "capital": "Muscat",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "OMR",
    "languageCodes": [
      "ar",
      "en",
      "ur"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PA": {
    "code": "PA",
    "name": "Panama",
    "capital": "Panama City",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "PAB",
    "languageCodes": [
      "es",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PE": {
    "code": "PE",
    "name": "Peru",
    "capital": "Lima",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "PEN",
    "languageCodes": [
      "es",
      "qu",
      "ay"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PF": {
    "code": "PF",
    "name": "French Polynesia",
    "capital": "Papeete",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "XPF",
    "languageCodes": [
      "fr",
      "ty"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PG": {
    "code": "PG",
    "name": "Papua New Guinea",
    "capital": "Port Moresby",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "PGK",
    "languageCodes": [
      "en",
      "ho"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PH": {
    "code": "PH",
    "name": "Philippines",
    "capital": "Manila",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "PHP",
    "languageCodes": [
      "tl",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PK": {
    "code": "PK",
    "name": "Pakistan",
    "capital": "Islamabad",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "PKR",
    "languageCodes": [
      "ur",
      "en",
      "pa",
      "sd",
      "ps"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PL": {
    "code": "PL",
    "name": "Poland",
    "capital": "Warsaw",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "PLN",
    "languageCodes": [
      "pl"
    ],
    "vatRates": {
      "standard": 23,
      "reduced1": 8,
      "reduced2": 5
    }
  },
  "PM": {
    "code": "PM",
    "name": "Saint Pierre and Miquelon",
    "capital": "Saint-Pierre",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PN": {
    "code": "PN",
    "name": "Pitcairn Islands",
    "capital": "Adamstown",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "NZD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PR": {
    "code": "PR",
    "name": "Puerto Rico",
    "capital": "San Juan",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "USD",
    "languageCodes": [
      "en",
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PS": {
    "code": "PS",
    "name": "Palestine",
    "capital": "",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "ILS",
    "languageCodes": [
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PT": {
    "code": "PT",
    "name": "Portugal",
    "capital": "Lisbon",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "pt"
    ],
    "vatRates": {
      "standard": 23,
      "reduced1": 13,
      "reduced2": 6
    }
  },
  "PW": {
    "code": "PW",
    "name": "Palau",
    "capital": "Melekeok",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "en",
      "ja",
      "zh"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "PY": {
    "code": "PY",
    "name": "Paraguay",
    "capital": "Asunci\u00f3n",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "PYG",
    "languageCodes": [
      "es",
      "gn"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "QA": {
    "code": "QA",
    "name": "Qatar",
    "capital": "Doha",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "QAR",
    "languageCodes": [
      "ar",
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "RE": {
    "code": "RE",
    "name": "R\u00e9union",
    "capital": "Saint-Denis",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "RO": {
    "code": "RO",
    "name": "Romania",
    "capital": "Bucharest",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "RON",
    "languageCodes": [
      "ro",
      "hu"
    ],
    "vatRates": {
      "standard": 20,
      "reduced1": 9,
      "reduced2": 5
    }
  },
  "RS": {
    "code": "RS",
    "name": "Serbia",
    "capital": "Belgrade",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "RSD",
    "languageCodes": [
      "sr",
      "hu",
      "bs"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "RU": {
    "code": "RU",
    "name": "Russia",
    "capital": "Moscow",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "RUB",
    "languageCodes": [
      "ru",
      "tt",
      "kv",
      "ce",
      "cv",
      "ba"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "RW": {
    "code": "RW",
    "name": "Rwanda",
    "capital": "Kigali",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "RWF",
    "languageCodes": [
      "rw",
      "en",
      "fr",
      "sw"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SA": {
    "code": "SA",
    "name": "Saudi Arabia",
    "capital": "Riyadh",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "SAR",
    "languageCodes": [
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SB": {
    "code": "SB",
    "name": "Solomon Islands",
    "capital": "Honiara",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "SBD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SC": {
    "code": "SC",
    "name": "Seychelles",
    "capital": "Victoria",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "SCR",
    "languageCodes": [
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SD": {
    "code": "SD",
    "name": "Sudan",
    "capital": "Khartoum",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "SDG",
    "languageCodes": [
      "ar",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SE": {
    "code": "SE",
    "name": "Sweden",
    "capital": "Stockholm",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "SEK",
    "languageCodes": [
      "sv",
      "se",
      "fi"
    ],
    "vatRates": {
      "standard": 25,
      "reduced1": 12,
      "reduced2": 6
    }
  },
  "SG": {
    "code": "SG",
    "name": "Singapore",
    "capital": "Singapore",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "SGD",
    "languageCodes": [
      "en",
      "ms",
      "ta",
      "zh"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SH": {
    "code": "SH",
    "name": "Saint Helena",
    "capital": "Jamestown",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "SHP",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SI": {
    "code": "SI",
    "name": "Slovenia",
    "capital": "Ljubljana",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "sl"
    ],
    "vatRates": {
      "standard": 22,
      "reduced1": 9.5,
      "reduced2": 0
    }
  },
  "SJ": {
    "code": "SJ",
    "name": "Svalbard and Jan Mayen",
    "capital": "Longyearbyen",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "NOK",
    "languageCodes": [
      "no",
      "ru"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SK": {
    "code": "SK",
    "name": "Slovakia",
    "capital": "Bratislava",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "sk",
      "hu"
    ],
    "vatRates": {
      "standard": 20,
      "reduced1": 10,
      "reduced2": 0
    }
  },
  "SL": {
    "code": "SL",
    "name": "Sierra Leone",
    "capital": "Freetown",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "SLL",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SM": {
    "code": "SM",
    "name": "San Marino",
    "capital": "San Marino",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "it"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SN": {
    "code": "SN",
    "name": "Senegal",
    "capital": "Dakar",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "fr",
      "wo"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SO": {
    "code": "SO",
    "name": "Somalia",
    "capital": "Mogadishu",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "SOS",
    "languageCodes": [
      "so",
      "ar",
      "it",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SR": {
    "code": "SR",
    "name": "Suriname",
    "capital": "Paramaribo",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "SRD",
    "languageCodes": [
      "nl",
      "en",
      "jv"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SS": {
    "code": "SS",
    "name": "South Sudan",
    "capital": "Juba",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "SSP",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ST": {
    "code": "ST",
    "name": "S\u00e3o Tom\u00e9 and Pr\u00edncipe",
    "capital": "S\u00e3o Tom\u00e9",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "STD",
    "languageCodes": [
      "pt"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SV": {
    "code": "SV",
    "name": "El Salvador",
    "capital": "San Salvador",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "USD",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SX": {
    "code": "SX",
    "name": "Sint Maarten",
    "capital": "Philipsburg",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "ANG",
    "languageCodes": [
      "nl",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SY": {
    "code": "SY",
    "name": "Syria",
    "capital": "Damascus",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "SYP",
    "languageCodes": [
      "ar",
      "ku",
      "hy",
      "fr",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "SZ": {
    "code": "SZ",
    "name": "Swaziland",
    "capital": "Mbabane",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "SZL",
    "languageCodes": [
      "en",
      "ss"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TC": {
    "code": "TC",
    "name": "Turks and Caicos Islands",
    "capital": "Cockburn Town",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "USD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TD": {
    "code": "TD",
    "name": "Chad",
    "capital": "N'Djamena",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XAF",
    "languageCodes": [
      "fr",
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TF": {
    "code": "TF",
    "name": "French Southern Territories",
    "capital": "Port-aux-Fran\u00e7ais",
    "continent": "AN",
    "continentName": "Antarctica",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TG": {
    "code": "TG",
    "name": "Togo",
    "capital": "Lom\u00e9",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "XOF",
    "languageCodes": [
      "fr",
      "ee",
      "ha"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TH": {
    "code": "TH",
    "name": "Thailand",
    "capital": "Bangkok",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "THB",
    "languageCodes": [
      "th",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TJ": {
    "code": "TJ",
    "name": "Tajikistan",
    "capital": "Dushanbe",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "TJS",
    "languageCodes": [
      "tg",
      "ru"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TK": {
    "code": "TK",
    "name": "Tokelau",
    "capital": "",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "NZD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TL": {
    "code": "TL",
    "name": "East Timor",
    "capital": "Dili",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "pt",
      "id",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TM": {
    "code": "TM",
    "name": "Turkmenistan",
    "capital": "Ashgabat",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "TMT",
    "languageCodes": [
      "tk",
      "ru",
      "uz"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TN": {
    "code": "TN",
    "name": "Tunisia",
    "capital": "Tunis",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "TND",
    "languageCodes": [
      "ar",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TO": {
    "code": "TO",
    "name": "Tonga",
    "capital": "Nuku'alofa",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "TOP",
    "languageCodes": [
      "to",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TR": {
    "code": "TR",
    "name": "Turkey",
    "capital": "Ankara",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "TRY",
    "languageCodes": [
      "tr",
      "ku",
      "az",
      "av"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TT": {
    "code": "TT",
    "name": "Trinidad and Tobago",
    "capital": "Port of Spain",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "TTD",
    "languageCodes": [
      "en",
      "fr",
      "es",
      "zh"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TV": {
    "code": "TV",
    "name": "Tuvalu",
    "capital": "Funafuti",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "AUD",
    "languageCodes": [
      "en",
      "sm"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TW": {
    "code": "TW",
    "name": "Taiwan",
    "capital": "Taipei",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "TWD",
    "languageCodes": [
      "zh",
      "zh"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "TZ": {
    "code": "TZ",
    "name": "Tanzania",
    "capital": "Dodoma",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "TZS",
    "languageCodes": [
      "sw",
      "en",
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "UA": {
    "code": "UA",
    "name": "Ukraine",
    "capital": "Kiev",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "UAH",
    "languageCodes": [
      "uk",
      "ru",
      "pl",
      "hu"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "UG": {
    "code": "UG",
    "name": "Uganda",
    "capital": "Kampala",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "UGX",
    "languageCodes": [
      "en",
      "lg",
      "sw",
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "UM": {
    "code": "UM",
    "name": "U.S. Minor Outlying Islands",
    "capital": "",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "USD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "US": {
    "code": "US",
    "name": "United States",
    "capital": "Washington",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "USD",
    "languageCodes": [
      "en",
      "es",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "UY": {
    "code": "UY",
    "name": "Uruguay",
    "capital": "Montevideo",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "UYU",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "UZ": {
    "code": "UZ",
    "name": "Uzbekistan",
    "capital": "Tashkent",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "UZS",
    "languageCodes": [
      "uz",
      "ru",
      "tg"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "VA": {
    "code": "VA",
    "name": "Vatican City",
    "capital": "Vatican City",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "la",
      "it",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "VC": {
    "code": "VC",
    "name": "Saint Vincent and the Grenadines",
    "capital": "Kingstown",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "XCD",
    "languageCodes": [
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "VE": {
    "code": "VE",
    "name": "Venezuela",
    "capital": "Caracas",
    "continent": "SA",
    "continentName": "South America",
    "currencyCode": "VEF",
    "languageCodes": [
      "es"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "VG": {
    "code": "VG",
    "name": "British Virgin Islands",
    "capital": "Road Town",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "USD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "VI": {
    "code": "VI",
    "name": "U.S. Virgin Islands",
    "capital": "Charlotte Amalie",
    "continent": "NA",
    "continentName": "North America",
    "currencyCode": "USD",
    "languageCodes": [
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "VN": {
    "code": "VN",
    "name": "Vietnam",
    "capital": "Hanoi",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "VND",
    "languageCodes": [
      "vi",
      "en",
      "fr",
      "zh",
      "km"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "VU": {
    "code": "VU",
    "name": "Vanuatu",
    "capital": "Port Vila",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "VUV",
    "languageCodes": [
      "bi",
      "en",
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "WF": {
    "code": "WF",
    "name": "Wallis and Futuna",
    "capital": "Mata-Utu",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "XPF",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "WS": {
    "code": "WS",
    "name": "Samoa",
    "capital": "Apia",
    "continent": "OC",
    "continentName": "Oceania",
    "currencyCode": "WST",
    "languageCodes": [
      "sm",
      "en"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "XK": {
    "code": "XK",
    "name": "Kosovo",
    "capital": "Pristina",
    "continent": "EU",
    "continentName": "Europe",
    "currencyCode": "EUR",
    "languageCodes": [
      "sq",
      "sr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "YE": {
    "code": "YE",
    "name": "Yemen",
    "capital": "Sanaa",
    "continent": "AS",
    "continentName": "Asia",
    "currencyCode": "YER",
    "languageCodes": [
      "ar"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "YT": {
    "code": "YT",
    "name": "Mayotte",
    "capital": "Mamoudzou",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "EUR",
    "languageCodes": [
      "fr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ZA": {
    "code": "ZA",
    "name": "South Africa",
    "capital": "Pretoria",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "ZAR",
    "languageCodes": [
      "xh",
      "af",
      "en",
      "tn",
      "st",
      "ts",
      "ss",
      "ve",
      "nr"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ZM": {
    "code": "ZM",
    "name": "Zambia",
    "capital": "Lusaka",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "ZMW",
    "languageCodes": [
      "en",
      "ny"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  },
  "ZW": {
    "code": "ZW",
    "name": "Zimbabwe",
    "capital": "Harare",
    "continent": "AF",
    "continentName": "Africa",
    "currencyCode": "ZWL",
    "languageCodes": [
      "en",
      "sn",
      "nr",
      "nd"
    ],
    "vatRates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  }
}`
