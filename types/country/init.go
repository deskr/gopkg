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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "ca"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "AED",
    "language_codes": [
      "ar",
      "fa",
      "en",
      "hi",
      "ur"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "AFN",
    "language_codes": [
      "fa",
      "ps",
      "uz",
      "tk"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "ALL",
    "language_codes": [
      "sq",
      "el"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "AMD",
    "language_codes": [
      "hy"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "AOA",
    "language_codes": [
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "Antarctica",
    "currency_code": "",
    "language_codes": null,
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "ARS",
    "language_codes": [
      "es",
      "en",
      "it",
      "de",
      "fr",
      "gn"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "en",
      "sm",
      "to"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "de",
      "hr",
      "hu",
      "sl"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "AUD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "AWG",
    "language_codes": [
      "nl",
      "es",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "sv"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "AZN",
    "language_codes": [
      "az",
      "ru",
      "hy"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "BAM",
    "language_codes": [
      "bs",
      "hr",
      "sr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "BBD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "BDT",
    "language_codes": [
      "bn",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "nl",
      "fr",
      "de"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "BGN",
    "language_codes": [
      "bg",
      "tr"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "BHD",
    "language_codes": [
      "ar",
      "en",
      "fa",
      "ur"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "BIF",
    "language_codes": [
      "fr",
      "rn"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "BMD",
    "language_codes": [
      "en",
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "BND",
    "language_codes": [
      "ms",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "BOB",
    "language_codes": [
      "es",
      "qu",
      "ay"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "USD",
    "language_codes": [
      "nl",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "BRL",
    "language_codes": [
      "pt",
      "es",
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "BSD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "BTN",
    "language_codes": null,
    "vat_rates": {
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
    "continent_name": "Antarctica",
    "currency_code": "NOK",
    "language_codes": null,
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "BWP",
    "language_codes": [
      "en",
      "tn"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "BYR",
    "language_codes": [
      "be",
      "ru"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "BZD",
    "language_codes": [
      "en",
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "CAD",
    "language_codes": [
      "en",
      "fr",
      "iu"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "AUD",
    "language_codes": [
      "ms",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "CDF",
    "language_codes": [
      "fr",
      "ln",
      "kg"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XAF",
    "language_codes": [
      "fr",
      "sg",
      "ln",
      "kg"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XAF",
    "language_codes": [
      "fr",
      "kg",
      "ln"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "CHF",
    "language_codes": [
      "de",
      "fr",
      "it",
      "rm"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "NZD",
    "language_codes": [
      "en",
      "mi"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "CLP",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XAF",
    "language_codes": [
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "CNY",
    "language_codes": [
      "zh",
      "ug",
      "za"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "COP",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "CRC",
    "language_codes": [
      "es",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "CUP",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "CVE",
    "language_codes": [
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "ANG",
    "language_codes": [
      "nl"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "AUD",
    "language_codes": [
      "en",
      "zh",
      "ms"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "el",
      "tr",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "CZK",
    "language_codes": [
      "cs",
      "sk"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "de"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "DJF",
    "language_codes": [
      "fr",
      "ar",
      "so",
      "aa"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "DKK",
    "language_codes": [
      "da",
      "en",
      "fo",
      "de"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "DOP",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "DZD",
    "language_codes": [
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "USD",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "et",
      "ru"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "EGP",
    "language_codes": [
      "ar",
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "MAD",
    "language_codes": [
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "ERN",
    "language_codes": [
      "aa",
      "ar",
      "ti"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "es",
      "ca",
      "gl",
      "eu",
      "oc"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "ETB",
    "language_codes": [
      "am",
      "en",
      "om",
      "ti",
      "so"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "fi",
      "sv"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "FJD",
    "language_codes": [
      "en",
      "fj"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "FKP",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "DKK",
    "language_codes": [
      "fo",
      "da"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "fr",
      "br",
      "co",
      "ca",
      "eu",
      "oc"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XAF",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "GBP",
    "language_codes": [
      "en",
      "cy",
      "gd"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "GEL",
    "language_codes": [
      "ka",
      "ru",
      "hy",
      "az"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "GBP",
    "language_codes": [
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "GHS",
    "language_codes": [
      "en",
      "ak",
      "ee",
      "tw"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "GIP",
    "language_codes": [
      "en",
      "es",
      "it",
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "DKK",
    "language_codes": [
      "kl",
      "da",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "GMD",
    "language_codes": [
      "en",
      "wo",
      "ff"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "GNF",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XAF",
    "language_codes": [
      "es",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "el",
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Antarctica",
    "currency_code": "GBP",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "GTQ",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "en",
      "ch"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "GYD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "HKD",
    "language_codes": [
      "zh",
      "zh",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Antarctica",
    "currency_code": "AUD",
    "language_codes": null,
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "HNL",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "HRK",
    "language_codes": [
      "hr",
      "sr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "HTG",
    "language_codes": [
      "ht",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "HUF",
    "language_codes": [
      "hu"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "IDR",
    "language_codes": [
      "id",
      "en",
      "nl",
      "jv"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "en",
      "ga"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "ILS",
    "language_codes": [
      "he",
      "ar",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "GBP",
    "language_codes": [
      "en",
      "gv"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "INR",
    "language_codes": [
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
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "USD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "IQD",
    "language_codes": [
      "ar",
      "ku",
      "hy"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "IRR",
    "language_codes": [
      "fa",
      "ku"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "ISK",
    "language_codes": [
      "is",
      "en",
      "de",
      "da",
      "sv",
      "no"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "it",
      "de",
      "fr",
      "sc",
      "ca",
      "co",
      "sl"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "GBP",
    "language_codes": [
      "en",
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "JMD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "JOD",
    "language_codes": [
      "ar",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "JPY",
    "language_codes": [
      "ja"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "KES",
    "language_codes": [
      "en",
      "sw"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "KGS",
    "language_codes": [
      "ky",
      "uz",
      "ru"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "KHR",
    "language_codes": [
      "km",
      "fr",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "AUD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "KMF",
    "language_codes": [
      "ar",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "KPW",
    "language_codes": [
      "ko"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "KRW",
    "language_codes": [
      "ko",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "KWD",
    "language_codes": [
      "ar",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "KYD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "KZT",
    "language_codes": [
      "kk",
      "ru"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "LAK",
    "language_codes": [
      "lo",
      "fr",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "LBP",
    "language_codes": [
      "ar",
      "fr",
      "en",
      "hy"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "CHF",
    "language_codes": [
      "de"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "LKR",
    "language_codes": [
      "si",
      "ta",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "LRD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "LSL",
    "language_codes": [
      "en",
      "st",
      "xh"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "lt",
      "ru",
      "pl"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "lb",
      "de",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "lv",
      "ru",
      "lt"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "LYD",
    "language_codes": [
      "ar",
      "it",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "MAD",
    "language_codes": [
      "ar",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "fr",
      "en",
      "it"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "MDL",
    "language_codes": [
      "ro",
      "ru",
      "tr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "sr",
      "hu",
      "bs",
      "sq",
      "hr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "MGA",
    "language_codes": [
      "fr",
      "mg"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "mh",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "MKD",
    "language_codes": [
      "mk",
      "sq",
      "tr",
      "sr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "fr",
      "bm"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "MMK",
    "language_codes": [
      "my"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "MNT",
    "language_codes": [
      "mn",
      "ru"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "MOP",
    "language_codes": [
      "zh",
      "zh",
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "tl",
      "zh",
      "ch",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "MRO",
    "language_codes": [
      "ar",
      "fr",
      "wo"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "mt",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "MUR",
    "language_codes": [
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "MVR",
    "language_codes": [
      "dv",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "MWK",
    "language_codes": [
      "ny"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "MXN",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "MYR",
    "language_codes": [
      "ms",
      "en",
      "zh",
      "ta",
      "te",
      "ml",
      "pa",
      "th"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "MZN",
    "language_codes": [
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "NAD",
    "language_codes": [
      "en",
      "af",
      "de",
      "hz"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "XPF",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "fr",
      "ha",
      "kr"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "AUD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "NGN",
    "language_codes": [
      "en",
      "ha",
      "yo",
      "ig",
      "ff"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "NIO",
    "language_codes": [
      "es",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "nl",
      "fy"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "NOK",
    "language_codes": [
      "no",
      "nb",
      "nn",
      "se",
      "fi"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "NPR",
    "language_codes": [
      "ne",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "AUD",
    "language_codes": [
      "na",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "NZD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "NZD",
    "language_codes": [
      "en",
      "mi"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "OMR",
    "language_codes": [
      "ar",
      "en",
      "ur"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "PAB",
    "language_codes": [
      "es",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "PEN",
    "language_codes": [
      "es",
      "qu",
      "ay"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "XPF",
    "language_codes": [
      "fr",
      "ty"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "PGK",
    "language_codes": [
      "en",
      "ho"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "PHP",
    "language_codes": [
      "tl",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "PKR",
    "language_codes": [
      "ur",
      "en",
      "pa",
      "sd",
      "ps"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "PLN",
    "language_codes": [
      "pl"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "NZD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "USD",
    "language_codes": [
      "en",
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "ILS",
    "language_codes": [
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "en",
      "ja",
      "zh"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "PYG",
    "language_codes": [
      "es",
      "gn"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "QAR",
    "language_codes": [
      "ar",
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "RON",
    "language_codes": [
      "ro",
      "hu"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "RSD",
    "language_codes": [
      "sr",
      "hu",
      "bs"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "RUB",
    "language_codes": [
      "ru",
      "tt",
      "kv",
      "ce",
      "cv",
      "ba"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "RWF",
    "language_codes": [
      "rw",
      "en",
      "fr",
      "sw"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "SAR",
    "language_codes": [
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "SBD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "SCR",
    "language_codes": [
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "SDG",
    "language_codes": [
      "ar",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "SEK",
    "language_codes": [
      "sv",
      "se",
      "fi"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "SGD",
    "language_codes": [
      "en",
      "ms",
      "ta",
      "zh"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "SHP",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "sl"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "NOK",
    "language_codes": [
      "no",
      "ru"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "sk",
      "hu"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "SLL",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "it"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "fr",
      "wo"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "SOS",
    "language_codes": [
      "so",
      "ar",
      "it",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "SRD",
    "language_codes": [
      "nl",
      "en",
      "jv"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "SSP",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "STD",
    "language_codes": [
      "pt"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "USD",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "ANG",
    "language_codes": [
      "nl",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "SYP",
    "language_codes": [
      "ar",
      "ku",
      "hy",
      "fr",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "SZL",
    "language_codes": [
      "en",
      "ss"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "USD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XAF",
    "language_codes": [
      "fr",
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "Antarctica",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "XOF",
    "language_codes": [
      "fr",
      "ee",
      "ha"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "THB",
    "language_codes": [
      "th",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "TJS",
    "language_codes": [
      "tg",
      "ru"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "NZD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "pt",
      "id",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "TMT",
    "language_codes": [
      "tk",
      "ru",
      "uz"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "TND",
    "language_codes": [
      "ar",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "TOP",
    "language_codes": [
      "to",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "TRY",
    "language_codes": [
      "tr",
      "ku",
      "az",
      "av"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "TTD",
    "language_codes": [
      "en",
      "fr",
      "es",
      "zh"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "AUD",
    "language_codes": [
      "en",
      "sm"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "TWD",
    "language_codes": [
      "zh",
      "zh"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "TZS",
    "language_codes": [
      "sw",
      "en",
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "UAH",
    "language_codes": [
      "uk",
      "ru",
      "pl",
      "hu"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "UGX",
    "language_codes": [
      "en",
      "lg",
      "sw",
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "USD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "USD",
    "language_codes": [
      "en",
      "es",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "UYU",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "UZS",
    "language_codes": [
      "uz",
      "ru",
      "tg"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "la",
      "it",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "XCD",
    "language_codes": [
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "South America",
    "currency_code": "VEF",
    "language_codes": [
      "es"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "USD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "North America",
    "currency_code": "USD",
    "language_codes": [
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "VND",
    "language_codes": [
      "vi",
      "en",
      "fr",
      "zh",
      "km"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "VUV",
    "language_codes": [
      "bi",
      "en",
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "XPF",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Oceania",
    "currency_code": "WST",
    "language_codes": [
      "sm",
      "en"
    ],
    "vat_rates": {
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
    "continent_name": "Europe",
    "currency_code": "EUR",
    "language_codes": [
      "sq",
      "sr"
    ],
    "vat_rates": {
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
    "continent_name": "Asia",
    "currency_code": "YER",
    "language_codes": [
      "ar"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "EUR",
    "language_codes": [
      "fr"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "ZAR",
    "language_codes": [
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
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "ZMW",
    "language_codes": [
      "en",
      "ny"
    ],
    "vat_rates": {
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
    "continent_name": "Africa",
    "currency_code": "ZWL",
    "language_codes": [
      "en",
      "sn",
      "nr",
      "nd"
    ],
    "vat_rates": {
      "standard": 0,
      "reduced1": 0,
      "reduced2": 0
    }
  }
}`
