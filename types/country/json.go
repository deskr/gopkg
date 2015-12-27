package country

import (
	"bytes"
	"encoding/json"
)

var countries []Country
var countriesByCode map[Code]Country
var countriesByContinent map[Continent][]Country

func init() {
	if err := json.NewDecoder(bytes.NewReader([]byte(countriesJSON))).Decode(&countries); err != nil {
		panic(err)
	}

	countriesByCode = make(map[Code]Country)
	for _, v := range countries {
		countriesByCode[v.Code] = v
	}

	countriesByContinent = make(map[Continent][]Country)
	for _, v := range countries {
		countriesByContinent[v.Continent] = append(countriesByContinent[v.Continent], v)
	}
}

const countriesJSON = `[
  {
    "Code": "AD",
    "Name": "Andorra",
    "Capital": "Andorra la Vella",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "ca"
    ]
  },
  {
    "Code": "AE",
    "Name": "United Arab Emirates",
    "Capital": "Abu Dhabi",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "AED",
    "LanguageCodes": [
      "ar",
      "fa",
      "en",
      "hi",
      "ur"
    ]
  },
  {
    "Code": "AF",
    "Name": "Afghanistan",
    "Capital": "Kabul",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "AFN",
    "LanguageCodes": [
      "fa",
      "ps",
      "uz",
      "tk"
    ]
  },
  {
    "Code": "AG",
    "Name": "Antigua and Barbuda",
    "Capital": "St. John's",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "AI",
    "Name": "Anguilla",
    "Capital": "The Valley",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "AL",
    "Name": "Albania",
    "Capital": "Tirana",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "ALL",
    "LanguageCodes": [
      "sq",
      "el"
    ]
  },
  {
    "Code": "AM",
    "Name": "Armenia",
    "Capital": "Yerevan",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "AMD",
    "LanguageCodes": [
      "hy"
    ]
  },
  {
    "Code": "AO",
    "Name": "Angola",
    "Capital": "Luanda",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "AOA",
    "LanguageCodes": [
      "pt"
    ]
  },
  {
    "Code": "AQ",
    "Name": "Antarctica",
    "Capital": "",
    "Continent": "AN",
    "ContinentName": "Antarctica",
    "CurrencyCode": "",
    "LanguageCodes": null
  },
  {
    "Code": "AR",
    "Name": "Argentina",
    "Capital": "Buenos Aires",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "ARS",
    "LanguageCodes": [
      "es",
      "en",
      "it",
      "de",
      "fr",
      "gn"
    ]
  },
  {
    "Code": "AS",
    "Name": "American Samoa",
    "Capital": "Pago Pago",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en",
      "sm",
      "to"
    ]
  },
  {
    "Code": "AT",
    "Name": "Austria",
    "Capital": "Vienna",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "de",
      "hr",
      "hu",
      "sl"
    ]
  },
  {
    "Code": "AU",
    "Name": "Australia",
    "Capital": "Canberra",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "AUD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "AW",
    "Name": "Aruba",
    "Capital": "Oranjestad",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "AWG",
    "LanguageCodes": [
      "nl",
      "es",
      "en"
    ]
  },
  {
    "Code": "AX",
    "Name": "Åland",
    "Capital": "Mariehamn",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "sv"
    ]
  },
  {
    "Code": "AZ",
    "Name": "Azerbaijan",
    "Capital": "Baku",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "AZN",
    "LanguageCodes": [
      "az",
      "ru",
      "hy"
    ]
  },
  {
    "Code": "BA",
    "Name": "Bosnia and Herzegovina",
    "Capital": "Sarajevo",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "BAM",
    "LanguageCodes": [
      "bs",
      "hr",
      "sr"
    ]
  },
  {
    "Code": "BB",
    "Name": "Barbados",
    "Capital": "Bridgetown",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "BBD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "BD",
    "Name": "Bangladesh",
    "Capital": "Dhaka",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "BDT",
    "LanguageCodes": [
      "bn",
      "en"
    ]
  },
  {
    "Code": "BE",
    "Name": "Belgium",
    "Capital": "Brussels",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "nl",
      "fr",
      "de"
    ]
  },
  {
    "Code": "BF",
    "Name": "Burkina Faso",
    "Capital": "Ouagadougou",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "BG",
    "Name": "Bulgaria",
    "Capital": "Sofia",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "BGN",
    "LanguageCodes": [
      "bg",
      "tr"
    ]
  },
  {
    "Code": "BH",
    "Name": "Bahrain",
    "Capital": "Manama",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "BHD",
    "LanguageCodes": [
      "ar",
      "en",
      "fa",
      "ur"
    ]
  },
  {
    "Code": "BI",
    "Name": "Burundi",
    "Capital": "Bujumbura",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "BIF",
    "LanguageCodes": [
      "fr",
      "rn"
    ]
  },
  {
    "Code": "BJ",
    "Name": "Benin",
    "Capital": "Porto-Novo",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "BL",
    "Name": "Saint Barthélemy",
    "Capital": "Gustavia",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "BM",
    "Name": "Bermuda",
    "Capital": "Hamilton",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "BMD",
    "LanguageCodes": [
      "en",
      "pt"
    ]
  },
  {
    "Code": "BN",
    "Name": "Brunei",
    "Capital": "Bandar Seri Begawan",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "BND",
    "LanguageCodes": [
      "ms",
      "en"
    ]
  },
  {
    "Code": "BO",
    "Name": "Bolivia",
    "Capital": "Sucre",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "BOB",
    "LanguageCodes": [
      "es",
      "qu",
      "ay"
    ]
  },
  {
    "Code": "BQ",
    "Name": "Bonaire",
    "Capital": "Kralendijk",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "nl",
      "en"
    ]
  },
  {
    "Code": "BR",
    "Name": "Brazil",
    "Capital": "Brasília",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "BRL",
    "LanguageCodes": [
      "pt",
      "es",
      "en",
      "fr"
    ]
  },
  {
    "Code": "BS",
    "Name": "Bahamas",
    "Capital": "Nassau",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "BSD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "BT",
    "Name": "Bhutan",
    "Capital": "Thimphu",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "BTN",
    "LanguageCodes": null
  },
  {
    "Code": "BV",
    "Name": "Bouvet Island",
    "Capital": "",
    "Continent": "AN",
    "ContinentName": "Antarctica",
    "CurrencyCode": "NOK",
    "LanguageCodes": null
  },
  {
    "Code": "BW",
    "Name": "Botswana",
    "Capital": "Gaborone",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "BWP",
    "LanguageCodes": [
      "en",
      "tn"
    ]
  },
  {
    "Code": "BY",
    "Name": "Belarus",
    "Capital": "Minsk",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "BYR",
    "LanguageCodes": [
      "be",
      "ru"
    ]
  },
  {
    "Code": "BZ",
    "Name": "Belize",
    "Capital": "Belmopan",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "BZD",
    "LanguageCodes": [
      "en",
      "es"
    ]
  },
  {
    "Code": "CA",
    "Name": "Canada",
    "Capital": "Ottawa",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "CAD",
    "LanguageCodes": [
      "en",
      "fr",
      "iu"
    ]
  },
  {
    "Code": "CC",
    "Name": "Cocos [Keeling] Islands",
    "Capital": "West Island",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "AUD",
    "LanguageCodes": [
      "ms",
      "en"
    ]
  },
  {
    "Code": "CD",
    "Name": "Democratic Republic of the Congo",
    "Capital": "Kinshasa",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "CDF",
    "LanguageCodes": [
      "fr",
      "ln",
      "kg"
    ]
  },
  {
    "Code": "CF",
    "Name": "Central African Republic",
    "Capital": "Bangui",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XAF",
    "LanguageCodes": [
      "fr",
      "sg",
      "ln",
      "kg"
    ]
  },
  {
    "Code": "CG",
    "Name": "Republic of the Congo",
    "Capital": "Brazzaville",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XAF",
    "LanguageCodes": [
      "fr",
      "kg",
      "ln"
    ]
  },
  {
    "Code": "CH",
    "Name": "Switzerland",
    "Capital": "Bern",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "CHF",
    "LanguageCodes": [
      "de",
      "fr",
      "it",
      "rm"
    ]
  },
  {
    "Code": "CI",
    "Name": "Ivory Coast",
    "Capital": "Yamoussoukro",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "CK",
    "Name": "Cook Islands",
    "Capital": "Avarua",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "NZD",
    "LanguageCodes": [
      "en",
      "mi"
    ]
  },
  {
    "Code": "CL",
    "Name": "Chile",
    "Capital": "Santiago",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "CLP",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "CM",
    "Name": "Cameroon",
    "Capital": "Yaoundé",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XAF",
    "LanguageCodes": [
      "en",
      "fr"
    ]
  },
  {
    "Code": "CN",
    "Name": "China",
    "Capital": "Beijing",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "CNY",
    "LanguageCodes": [
      "zh",
      "ug",
      "za"
    ]
  },
  {
    "Code": "CO",
    "Name": "Colombia",
    "Capital": "Bogotá",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "COP",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "CR",
    "Name": "Costa Rica",
    "Capital": "San José",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "CRC",
    "LanguageCodes": [
      "es",
      "en"
    ]
  },
  {
    "Code": "CU",
    "Name": "Cuba",
    "Capital": "Havana",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "CUP",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "CV",
    "Name": "Cape Verde",
    "Capital": "Praia",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "CVE",
    "LanguageCodes": [
      "pt"
    ]
  },
  {
    "Code": "CW",
    "Name": "Curacao",
    "Capital": "Willemstad",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "ANG",
    "LanguageCodes": [
      "nl"
    ]
  },
  {
    "Code": "CX",
    "Name": "Christmas Island",
    "Capital": "Flying Fish Cove",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "AUD",
    "LanguageCodes": [
      "en",
      "zh",
      "ms"
    ]
  },
  {
    "Code": "CY",
    "Name": "Cyprus",
    "Capital": "Nicosia",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "el",
      "tr",
      "en"
    ]
  },
  {
    "Code": "CZ",
    "Name": "Czech Republic",
    "Capital": "Prague",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "CZK",
    "LanguageCodes": [
      "cs",
      "sk"
    ]
  },
  {
    "Code": "DE",
    "Name": "Germany",
    "Capital": "Berlin",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "de"
    ]
  },
  {
    "Code": "DJ",
    "Name": "Djibouti",
    "Capital": "Djibouti",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "DJF",
    "LanguageCodes": [
      "fr",
      "ar",
      "so",
      "aa"
    ]
  },
  {
    "Code": "DK",
    "Name": "Denmark",
    "Capital": "Copenhagen",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "DKK",
    "LanguageCodes": [
      "da",
      "en",
      "fo",
      "de"
    ]
  },
  {
    "Code": "DM",
    "Name": "Dominica",
    "Capital": "Roseau",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "DO",
    "Name": "Dominican Republic",
    "Capital": "Santo Domingo",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "DOP",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "DZ",
    "Name": "Algeria",
    "Capital": "Algiers",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "DZD",
    "LanguageCodes": [
      "ar"
    ]
  },
  {
    "Code": "EC",
    "Name": "Ecuador",
    "Capital": "Quito",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "EE",
    "Name": "Estonia",
    "Capital": "Tallinn",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "et",
      "ru"
    ]
  },
  {
    "Code": "EG",
    "Name": "Egypt",
    "Capital": "Cairo",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "EGP",
    "LanguageCodes": [
      "ar",
      "en",
      "fr"
    ]
  },
  {
    "Code": "EH",
    "Name": "Western Sahara",
    "Capital": "Laâyoune / El Aaiún",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "MAD",
    "LanguageCodes": [
      "ar"
    ]
  },
  {
    "Code": "ER",
    "Name": "Eritrea",
    "Capital": "Asmara",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "ERN",
    "LanguageCodes": [
      "aa",
      "ar",
      "ti"
    ]
  },
  {
    "Code": "ES",
    "Name": "Spain",
    "Capital": "Madrid",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "es",
      "ca",
      "gl",
      "eu",
      "oc"
    ]
  },
  {
    "Code": "ET",
    "Name": "Ethiopia",
    "Capital": "Addis Ababa",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "ETB",
    "LanguageCodes": [
      "am",
      "en",
      "om",
      "ti",
      "so"
    ]
  },
  {
    "Code": "FI",
    "Name": "Finland",
    "Capital": "Helsinki",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fi",
      "sv"
    ]
  },
  {
    "Code": "FJ",
    "Name": "Fiji",
    "Capital": "Suva",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "FJD",
    "LanguageCodes": [
      "en",
      "fj"
    ]
  },
  {
    "Code": "FK",
    "Name": "Falkland Islands",
    "Capital": "Stanley",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "FKP",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "FM",
    "Name": "Micronesia",
    "Capital": "Palikir",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "FO",
    "Name": "Faroe Islands",
    "Capital": "Tórshavn",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "DKK",
    "LanguageCodes": [
      "fo",
      "da"
    ]
  },
  {
    "Code": "FR",
    "Name": "France",
    "Capital": "Paris",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr",
      "br",
      "co",
      "ca",
      "eu",
      "oc"
    ]
  },
  {
    "Code": "GA",
    "Name": "Gabon",
    "Capital": "Libreville",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XAF",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "GB",
    "Name": "United Kingdom",
    "Capital": "London",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "GBP",
    "LanguageCodes": [
      "en",
      "cy",
      "gd"
    ]
  },
  {
    "Code": "GD",
    "Name": "Grenada",
    "Capital": "St. George's",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "GE",
    "Name": "Georgia",
    "Capital": "Tbilisi",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "GEL",
    "LanguageCodes": [
      "ka",
      "ru",
      "hy",
      "az"
    ]
  },
  {
    "Code": "GF",
    "Name": "French Guiana",
    "Capital": "Cayenne",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "GG",
    "Name": "Guernsey",
    "Capital": "St Peter Port",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "GBP",
    "LanguageCodes": [
      "en",
      "fr"
    ]
  },
  {
    "Code": "GH",
    "Name": "Ghana",
    "Capital": "Accra",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "GHS",
    "LanguageCodes": [
      "en",
      "ak",
      "ee",
      "tw"
    ]
  },
  {
    "Code": "GI",
    "Name": "Gibraltar",
    "Capital": "Gibraltar",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "GIP",
    "LanguageCodes": [
      "en",
      "es",
      "it",
      "pt"
    ]
  },
  {
    "Code": "GL",
    "Name": "Greenland",
    "Capital": "Nuuk",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "DKK",
    "LanguageCodes": [
      "kl",
      "da",
      "en"
    ]
  },
  {
    "Code": "GM",
    "Name": "Gambia",
    "Capital": "Bathurst",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "GMD",
    "LanguageCodes": [
      "en",
      "wo",
      "ff"
    ]
  },
  {
    "Code": "GN",
    "Name": "Guinea",
    "Capital": "Conakry",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "GNF",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "GP",
    "Name": "Guadeloupe",
    "Capital": "Basse-Terre",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "GQ",
    "Name": "Equatorial Guinea",
    "Capital": "Malabo",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XAF",
    "LanguageCodes": [
      "es",
      "fr"
    ]
  },
  {
    "Code": "GR",
    "Name": "Greece",
    "Capital": "Athens",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "el",
      "en",
      "fr"
    ]
  },
  {
    "Code": "GS",
    "Name": "South Georgia and the South Sandwich Islands",
    "Capital": "Grytviken",
    "Continent": "AN",
    "ContinentName": "Antarctica",
    "CurrencyCode": "GBP",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "GT",
    "Name": "Guatemala",
    "Capital": "Guatemala City",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "GTQ",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "GU",
    "Name": "Guam",
    "Capital": "Hagåtña",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en",
      "ch"
    ]
  },
  {
    "Code": "GW",
    "Name": "Guinea-Bissau",
    "Capital": "Bissau",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "pt"
    ]
  },
  {
    "Code": "GY",
    "Name": "Guyana",
    "Capital": "Georgetown",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "GYD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "HK",
    "Name": "Hong Kong",
    "Capital": "Hong Kong",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "HKD",
    "LanguageCodes": [
      "zh",
      "zh",
      "en"
    ]
  },
  {
    "Code": "HM",
    "Name": "Heard Island and McDonald Islands",
    "Capital": "",
    "Continent": "AN",
    "ContinentName": "Antarctica",
    "CurrencyCode": "AUD",
    "LanguageCodes": null
  },
  {
    "Code": "HN",
    "Name": "Honduras",
    "Capital": "Tegucigalpa",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "HNL",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "HR",
    "Name": "Croatia",
    "Capital": "Zagreb",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "HRK",
    "LanguageCodes": [
      "hr",
      "sr"
    ]
  },
  {
    "Code": "HT",
    "Name": "Haiti",
    "Capital": "Port-au-Prince",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "HTG",
    "LanguageCodes": [
      "ht",
      "fr"
    ]
  },
  {
    "Code": "HU",
    "Name": "Hungary",
    "Capital": "Budapest",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "HUF",
    "LanguageCodes": [
      "hu"
    ]
  },
  {
    "Code": "ID",
    "Name": "Indonesia",
    "Capital": "Jakarta",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "IDR",
    "LanguageCodes": [
      "id",
      "en",
      "nl",
      "jv"
    ]
  },
  {
    "Code": "IE",
    "Name": "Ireland",
    "Capital": "Dublin",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "en",
      "ga"
    ]
  },
  {
    "Code": "IL",
    "Name": "Israel",
    "Capital": "",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "ILS",
    "LanguageCodes": [
      "he",
      "ar",
      "en"
    ]
  },
  {
    "Code": "IM",
    "Name": "Isle of Man",
    "Capital": "Douglas",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "GBP",
    "LanguageCodes": [
      "en",
      "gv"
    ]
  },
  {
    "Code": "IN",
    "Name": "India",
    "Capital": "New Delhi",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "INR",
    "LanguageCodes": [
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
    ]
  },
  {
    "Code": "IO",
    "Name": "British Indian Ocean Territory",
    "Capital": "",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "IQ",
    "Name": "Iraq",
    "Capital": "Baghdad",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "IQD",
    "LanguageCodes": [
      "ar",
      "ku",
      "hy"
    ]
  },
  {
    "Code": "IR",
    "Name": "Iran",
    "Capital": "Tehran",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "IRR",
    "LanguageCodes": [
      "fa",
      "ku"
    ]
  },
  {
    "Code": "IS",
    "Name": "Iceland",
    "Capital": "Reykjavik",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "ISK",
    "LanguageCodes": [
      "is",
      "en",
      "de",
      "da",
      "sv",
      "no"
    ]
  },
  {
    "Code": "IT",
    "Name": "Italy",
    "Capital": "Rome",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "it",
      "de",
      "fr",
      "sc",
      "ca",
      "co",
      "sl"
    ]
  },
  {
    "Code": "JE",
    "Name": "Jersey",
    "Capital": "Saint Helier",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "GBP",
    "LanguageCodes": [
      "en",
      "pt"
    ]
  },
  {
    "Code": "JM",
    "Name": "Jamaica",
    "Capital": "Kingston",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "JMD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "JO",
    "Name": "Jordan",
    "Capital": "Amman",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "JOD",
    "LanguageCodes": [
      "ar",
      "en"
    ]
  },
  {
    "Code": "JP",
    "Name": "Japan",
    "Capital": "Tokyo",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "JPY",
    "LanguageCodes": [
      "ja"
    ]
  },
  {
    "Code": "KE",
    "Name": "Kenya",
    "Capital": "Nairobi",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "KES",
    "LanguageCodes": [
      "en",
      "sw"
    ]
  },
  {
    "Code": "KG",
    "Name": "Kyrgyzstan",
    "Capital": "Bishkek",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "KGS",
    "LanguageCodes": [
      "ky",
      "uz",
      "ru"
    ]
  },
  {
    "Code": "KH",
    "Name": "Cambodia",
    "Capital": "Phnom Penh",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "KHR",
    "LanguageCodes": [
      "km",
      "fr",
      "en"
    ]
  },
  {
    "Code": "KI",
    "Name": "Kiribati",
    "Capital": "Tarawa",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "AUD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "KM",
    "Name": "Comoros",
    "Capital": "Moroni",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "KMF",
    "LanguageCodes": [
      "ar",
      "fr"
    ]
  },
  {
    "Code": "KN",
    "Name": "Saint Kitts and Nevis",
    "Capital": "Basseterre",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "KP",
    "Name": "North Korea",
    "Capital": "Pyongyang",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "KPW",
    "LanguageCodes": [
      "ko"
    ]
  },
  {
    "Code": "KR",
    "Name": "South Korea",
    "Capital": "Seoul",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "KRW",
    "LanguageCodes": [
      "ko",
      "en"
    ]
  },
  {
    "Code": "KW",
    "Name": "Kuwait",
    "Capital": "Kuwait City",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "KWD",
    "LanguageCodes": [
      "ar",
      "en"
    ]
  },
  {
    "Code": "KY",
    "Name": "Cayman Islands",
    "Capital": "George Town",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "KYD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "KZ",
    "Name": "Kazakhstan",
    "Capital": "Astana",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "KZT",
    "LanguageCodes": [
      "kk",
      "ru"
    ]
  },
  {
    "Code": "LA",
    "Name": "Laos",
    "Capital": "Vientiane",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "LAK",
    "LanguageCodes": [
      "lo",
      "fr",
      "en"
    ]
  },
  {
    "Code": "LB",
    "Name": "Lebanon",
    "Capital": "Beirut",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "LBP",
    "LanguageCodes": [
      "ar",
      "fr",
      "en",
      "hy"
    ]
  },
  {
    "Code": "LC",
    "Name": "Saint Lucia",
    "Capital": "Castries",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "LI",
    "Name": "Liechtenstein",
    "Capital": "Vaduz",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "CHF",
    "LanguageCodes": [
      "de"
    ]
  },
  {
    "Code": "LK",
    "Name": "Sri Lanka",
    "Capital": "Colombo",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "LKR",
    "LanguageCodes": [
      "si",
      "ta",
      "en"
    ]
  },
  {
    "Code": "LR",
    "Name": "Liberia",
    "Capital": "Monrovia",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "LRD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "LS",
    "Name": "Lesotho",
    "Capital": "Maseru",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "LSL",
    "LanguageCodes": [
      "en",
      "st",
      "xh"
    ]
  },
  {
    "Code": "LT",
    "Name": "Lithuania",
    "Capital": "Vilnius",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "lt",
      "ru",
      "pl"
    ]
  },
  {
    "Code": "LU",
    "Name": "Luxembourg",
    "Capital": "Luxembourg",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "lb",
      "de",
      "fr"
    ]
  },
  {
    "Code": "LV",
    "Name": "Latvia",
    "Capital": "Riga",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "lv",
      "ru",
      "lt"
    ]
  },
  {
    "Code": "LY",
    "Name": "Libya",
    "Capital": "Tripoli",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "LYD",
    "LanguageCodes": [
      "ar",
      "it",
      "en"
    ]
  },
  {
    "Code": "MA",
    "Name": "Morocco",
    "Capital": "Rabat",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "MAD",
    "LanguageCodes": [
      "ar",
      "fr"
    ]
  },
  {
    "Code": "MC",
    "Name": "Monaco",
    "Capital": "Monaco",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr",
      "en",
      "it"
    ]
  },
  {
    "Code": "MD",
    "Name": "Moldova",
    "Capital": "Chişinău",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "MDL",
    "LanguageCodes": [
      "ro",
      "ru",
      "tr"
    ]
  },
  {
    "Code": "ME",
    "Name": "Montenegro",
    "Capital": "Podgorica",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "sr",
      "hu",
      "bs",
      "sq",
      "hr"
    ]
  },
  {
    "Code": "MF",
    "Name": "Saint Martin",
    "Capital": "Marigot",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "MG",
    "Name": "Madagascar",
    "Capital": "Antananarivo",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "MGA",
    "LanguageCodes": [
      "fr",
      "mg"
    ]
  },
  {
    "Code": "MH",
    "Name": "Marshall Islands",
    "Capital": "Majuro",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "mh",
      "en"
    ]
  },
  {
    "Code": "MK",
    "Name": "Macedonia",
    "Capital": "Skopje",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "MKD",
    "LanguageCodes": [
      "mk",
      "sq",
      "tr",
      "sr"
    ]
  },
  {
    "Code": "ML",
    "Name": "Mali",
    "Capital": "Bamako",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "fr",
      "bm"
    ]
  },
  {
    "Code": "MM",
    "Name": "Myanmar [Burma]",
    "Capital": "Naypyitaw",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "MMK",
    "LanguageCodes": [
      "my"
    ]
  },
  {
    "Code": "MN",
    "Name": "Mongolia",
    "Capital": "Ulan Bator",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "MNT",
    "LanguageCodes": [
      "mn",
      "ru"
    ]
  },
  {
    "Code": "MO",
    "Name": "Macao",
    "Capital": "Macao",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "MOP",
    "LanguageCodes": [
      "zh",
      "zh",
      "pt"
    ]
  },
  {
    "Code": "MP",
    "Name": "Northern Mariana Islands",
    "Capital": "Saipan",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "tl",
      "zh",
      "ch",
      "en"
    ]
  },
  {
    "Code": "MQ",
    "Name": "Martinique",
    "Capital": "Fort-de-France",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "MR",
    "Name": "Mauritania",
    "Capital": "Nouakchott",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "MRO",
    "LanguageCodes": [
      "ar",
      "fr",
      "wo"
    ]
  },
  {
    "Code": "MS",
    "Name": "Montserrat",
    "Capital": "Plymouth",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "MT",
    "Name": "Malta",
    "Capital": "Valletta",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "mt",
      "en"
    ]
  },
  {
    "Code": "MU",
    "Name": "Mauritius",
    "Capital": "Port Louis",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "MUR",
    "LanguageCodes": [
      "en",
      "fr"
    ]
  },
  {
    "Code": "MV",
    "Name": "Maldives",
    "Capital": "Malé",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "MVR",
    "LanguageCodes": [
      "dv",
      "en"
    ]
  },
  {
    "Code": "MW",
    "Name": "Malawi",
    "Capital": "Lilongwe",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "MWK",
    "LanguageCodes": [
      "ny"
    ]
  },
  {
    "Code": "MX",
    "Name": "Mexico",
    "Capital": "Mexico City",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "MXN",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "MY",
    "Name": "Malaysia",
    "Capital": "Kuala Lumpur",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "MYR",
    "LanguageCodes": [
      "ms",
      "en",
      "zh",
      "ta",
      "te",
      "ml",
      "pa",
      "th"
    ]
  },
  {
    "Code": "MZ",
    "Name": "Mozambique",
    "Capital": "Maputo",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "MZN",
    "LanguageCodes": [
      "pt"
    ]
  },
  {
    "Code": "NA",
    "Name": "Namibia",
    "Capital": "Windhoek",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "NAD",
    "LanguageCodes": [
      "en",
      "af",
      "de",
      "hz"
    ]
  },
  {
    "Code": "NC",
    "Name": "New Caledonia",
    "Capital": "Noumea",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "XPF",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "NE",
    "Name": "Niger",
    "Capital": "Niamey",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "fr",
      "ha",
      "kr"
    ]
  },
  {
    "Code": "NF",
    "Name": "Norfolk Island",
    "Capital": "Kingston",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "AUD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "NG",
    "Name": "Nigeria",
    "Capital": "Abuja",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "NGN",
    "LanguageCodes": [
      "en",
      "ha",
      "yo",
      "ig",
      "ff"
    ]
  },
  {
    "Code": "NI",
    "Name": "Nicaragua",
    "Capital": "Managua",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "NIO",
    "LanguageCodes": [
      "es",
      "en"
    ]
  },
  {
    "Code": "NL",
    "Name": "Netherlands",
    "Capital": "Amsterdam",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "nl",
      "fy"
    ]
  },
  {
    "Code": "NO",
    "Name": "Norway",
    "Capital": "Oslo",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "NOK",
    "LanguageCodes": [
      "no",
      "nb",
      "nn",
      "se",
      "fi"
    ]
  },
  {
    "Code": "NP",
    "Name": "Nepal",
    "Capital": "Kathmandu",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "NPR",
    "LanguageCodes": [
      "ne",
      "en"
    ]
  },
  {
    "Code": "NR",
    "Name": "Nauru",
    "Capital": "Yaren",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "AUD",
    "LanguageCodes": [
      "na",
      "en"
    ]
  },
  {
    "Code": "NU",
    "Name": "Niue",
    "Capital": "Alofi",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "NZD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "NZ",
    "Name": "New Zealand",
    "Capital": "Wellington",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "NZD",
    "LanguageCodes": [
      "en",
      "mi"
    ]
  },
  {
    "Code": "OM",
    "Name": "Oman",
    "Capital": "Muscat",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "OMR",
    "LanguageCodes": [
      "ar",
      "en",
      "ur"
    ]
  },
  {
    "Code": "PA",
    "Name": "Panama",
    "Capital": "Panama City",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "PAB",
    "LanguageCodes": [
      "es",
      "en"
    ]
  },
  {
    "Code": "PE",
    "Name": "Peru",
    "Capital": "Lima",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "PEN",
    "LanguageCodes": [
      "es",
      "qu",
      "ay"
    ]
  },
  {
    "Code": "PF",
    "Name": "French Polynesia",
    "Capital": "Papeete",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "XPF",
    "LanguageCodes": [
      "fr",
      "ty"
    ]
  },
  {
    "Code": "PG",
    "Name": "Papua New Guinea",
    "Capital": "Port Moresby",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "PGK",
    "LanguageCodes": [
      "en",
      "ho"
    ]
  },
  {
    "Code": "PH",
    "Name": "Philippines",
    "Capital": "Manila",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "PHP",
    "LanguageCodes": [
      "tl",
      "en"
    ]
  },
  {
    "Code": "PK",
    "Name": "Pakistan",
    "Capital": "Islamabad",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "PKR",
    "LanguageCodes": [
      "ur",
      "en",
      "pa",
      "sd",
      "ps"
    ]
  },
  {
    "Code": "PL",
    "Name": "Poland",
    "Capital": "Warsaw",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "PLN",
    "LanguageCodes": [
      "pl"
    ]
  },
  {
    "Code": "PM",
    "Name": "Saint Pierre and Miquelon",
    "Capital": "Saint-Pierre",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "PN",
    "Name": "Pitcairn Islands",
    "Capital": "Adamstown",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "NZD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "PR",
    "Name": "Puerto Rico",
    "Capital": "San Juan",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en",
      "es"
    ]
  },
  {
    "Code": "PS",
    "Name": "Palestine",
    "Capital": "",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "ILS",
    "LanguageCodes": [
      "ar"
    ]
  },
  {
    "Code": "PT",
    "Name": "Portugal",
    "Capital": "Lisbon",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "pt"
    ]
  },
  {
    "Code": "PW",
    "Name": "Palau",
    "Capital": "Melekeok",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en",
      "ja",
      "zh"
    ]
  },
  {
    "Code": "PY",
    "Name": "Paraguay",
    "Capital": "Asunción",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "PYG",
    "LanguageCodes": [
      "es",
      "gn"
    ]
  },
  {
    "Code": "QA",
    "Name": "Qatar",
    "Capital": "Doha",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "QAR",
    "LanguageCodes": [
      "ar",
      "es"
    ]
  },
  {
    "Code": "RE",
    "Name": "Réunion",
    "Capital": "Saint-Denis",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "RO",
    "Name": "Romania",
    "Capital": "Bucharest",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "RON",
    "LanguageCodes": [
      "ro",
      "hu"
    ]
  },
  {
    "Code": "RS",
    "Name": "Serbia",
    "Capital": "Belgrade",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "RSD",
    "LanguageCodes": [
      "sr",
      "hu",
      "bs"
    ]
  },
  {
    "Code": "RU",
    "Name": "Russia",
    "Capital": "Moscow",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "RUB",
    "LanguageCodes": [
      "ru",
      "tt",
      "kv",
      "ce",
      "cv",
      "ba"
    ]
  },
  {
    "Code": "RW",
    "Name": "Rwanda",
    "Capital": "Kigali",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "RWF",
    "LanguageCodes": [
      "rw",
      "en",
      "fr",
      "sw"
    ]
  },
  {
    "Code": "SA",
    "Name": "Saudi Arabia",
    "Capital": "Riyadh",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "SAR",
    "LanguageCodes": [
      "ar"
    ]
  },
  {
    "Code": "SB",
    "Name": "Solomon Islands",
    "Capital": "Honiara",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "SBD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "SC",
    "Name": "Seychelles",
    "Capital": "Victoria",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "SCR",
    "LanguageCodes": [
      "en",
      "fr"
    ]
  },
  {
    "Code": "SD",
    "Name": "Sudan",
    "Capital": "Khartoum",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "SDG",
    "LanguageCodes": [
      "ar",
      "en"
    ]
  },
  {
    "Code": "SE",
    "Name": "Sweden",
    "Capital": "Stockholm",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "SEK",
    "LanguageCodes": [
      "sv",
      "se",
      "fi"
    ]
  },
  {
    "Code": "SG",
    "Name": "Singapore",
    "Capital": "Singapore",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "SGD",
    "LanguageCodes": [
      "en",
      "ms",
      "ta",
      "zh"
    ]
  },
  {
    "Code": "SH",
    "Name": "Saint Helena",
    "Capital": "Jamestown",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "SHP",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "SI",
    "Name": "Slovenia",
    "Capital": "Ljubljana",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "sl"
    ]
  },
  {
    "Code": "SJ",
    "Name": "Svalbard and Jan Mayen",
    "Capital": "Longyearbyen",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "NOK",
    "LanguageCodes": [
      "no",
      "ru"
    ]
  },
  {
    "Code": "SK",
    "Name": "Slovakia",
    "Capital": "Bratislava",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "sk",
      "hu"
    ]
  },
  {
    "Code": "SL",
    "Name": "Sierra Leone",
    "Capital": "Freetown",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "SLL",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "SM",
    "Name": "San Marino",
    "Capital": "San Marino",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "it"
    ]
  },
  {
    "Code": "SN",
    "Name": "Senegal",
    "Capital": "Dakar",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "fr",
      "wo"
    ]
  },
  {
    "Code": "SO",
    "Name": "Somalia",
    "Capital": "Mogadishu",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "SOS",
    "LanguageCodes": [
      "so",
      "ar",
      "it",
      "en"
    ]
  },
  {
    "Code": "SR",
    "Name": "Suriname",
    "Capital": "Paramaribo",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "SRD",
    "LanguageCodes": [
      "nl",
      "en",
      "jv"
    ]
  },
  {
    "Code": "SS",
    "Name": "South Sudan",
    "Capital": "Juba",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "SSP",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "ST",
    "Name": "São Tomé and Príncipe",
    "Capital": "São Tomé",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "STD",
    "LanguageCodes": [
      "pt"
    ]
  },
  {
    "Code": "SV",
    "Name": "El Salvador",
    "Capital": "San Salvador",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "SX",
    "Name": "Sint Maarten",
    "Capital": "Philipsburg",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "ANG",
    "LanguageCodes": [
      "nl",
      "en"
    ]
  },
  {
    "Code": "SY",
    "Name": "Syria",
    "Capital": "Damascus",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "SYP",
    "LanguageCodes": [
      "ar",
      "ku",
      "hy",
      "fr",
      "en"
    ]
  },
  {
    "Code": "SZ",
    "Name": "Swaziland",
    "Capital": "Mbabane",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "SZL",
    "LanguageCodes": [
      "en",
      "ss"
    ]
  },
  {
    "Code": "TC",
    "Name": "Turks and Caicos Islands",
    "Capital": "Cockburn Town",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "TD",
    "Name": "Chad",
    "Capital": "N'Djamena",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XAF",
    "LanguageCodes": [
      "fr",
      "ar"
    ]
  },
  {
    "Code": "TF",
    "Name": "French Southern Territories",
    "Capital": "Port-aux-Français",
    "Continent": "AN",
    "ContinentName": "Antarctica",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "TG",
    "Name": "Togo",
    "Capital": "Lomé",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "XOF",
    "LanguageCodes": [
      "fr",
      "ee",
      "ha"
    ]
  },
  {
    "Code": "TH",
    "Name": "Thailand",
    "Capital": "Bangkok",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "THB",
    "LanguageCodes": [
      "th",
      "en"
    ]
  },
  {
    "Code": "TJ",
    "Name": "Tajikistan",
    "Capital": "Dushanbe",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "TJS",
    "LanguageCodes": [
      "tg",
      "ru"
    ]
  },
  {
    "Code": "TK",
    "Name": "Tokelau",
    "Capital": "",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "NZD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "TL",
    "Name": "East Timor",
    "Capital": "Dili",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "pt",
      "id",
      "en"
    ]
  },
  {
    "Code": "TM",
    "Name": "Turkmenistan",
    "Capital": "Ashgabat",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "TMT",
    "LanguageCodes": [
      "tk",
      "ru",
      "uz"
    ]
  },
  {
    "Code": "TN",
    "Name": "Tunisia",
    "Capital": "Tunis",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "TND",
    "LanguageCodes": [
      "ar",
      "fr"
    ]
  },
  {
    "Code": "TO",
    "Name": "Tonga",
    "Capital": "Nuku'alofa",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "TOP",
    "LanguageCodes": [
      "to",
      "en"
    ]
  },
  {
    "Code": "TR",
    "Name": "Turkey",
    "Capital": "Ankara",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "TRY",
    "LanguageCodes": [
      "tr",
      "ku",
      "az",
      "av"
    ]
  },
  {
    "Code": "TT",
    "Name": "Trinidad and Tobago",
    "Capital": "Port of Spain",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "TTD",
    "LanguageCodes": [
      "en",
      "fr",
      "es",
      "zh"
    ]
  },
  {
    "Code": "TV",
    "Name": "Tuvalu",
    "Capital": "Funafuti",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "AUD",
    "LanguageCodes": [
      "en",
      "sm"
    ]
  },
  {
    "Code": "TW",
    "Name": "Taiwan",
    "Capital": "Taipei",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "TWD",
    "LanguageCodes": [
      "zh",
      "zh"
    ]
  },
  {
    "Code": "TZ",
    "Name": "Tanzania",
    "Capital": "Dodoma",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "TZS",
    "LanguageCodes": [
      "sw",
      "en",
      "ar"
    ]
  },
  {
    "Code": "UA",
    "Name": "Ukraine",
    "Capital": "Kiev",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "UAH",
    "LanguageCodes": [
      "uk",
      "ru",
      "pl",
      "hu"
    ]
  },
  {
    "Code": "UG",
    "Name": "Uganda",
    "Capital": "Kampala",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "UGX",
    "LanguageCodes": [
      "en",
      "lg",
      "sw",
      "ar"
    ]
  },
  {
    "Code": "UM",
    "Name": "U.S. Minor Outlying Islands",
    "Capital": "",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "US",
    "Name": "United States",
    "Capital": "Washington",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en",
      "es",
      "fr"
    ]
  },
  {
    "Code": "UY",
    "Name": "Uruguay",
    "Capital": "Montevideo",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "UYU",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "UZ",
    "Name": "Uzbekistan",
    "Capital": "Tashkent",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "UZS",
    "LanguageCodes": [
      "uz",
      "ru",
      "tg"
    ]
  },
  {
    "Code": "VA",
    "Name": "Vatican City",
    "Capital": "Vatican City",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "la",
      "it",
      "fr"
    ]
  },
  {
    "Code": "VC",
    "Name": "Saint Vincent and the Grenadines",
    "Capital": "Kingstown",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "XCD",
    "LanguageCodes": [
      "en",
      "fr"
    ]
  },
  {
    "Code": "VE",
    "Name": "Venezuela",
    "Capital": "Caracas",
    "Continent": "SA",
    "ContinentName": "South America",
    "CurrencyCode": "VEF",
    "LanguageCodes": [
      "es"
    ]
  },
  {
    "Code": "VG",
    "Name": "British Virgin Islands",
    "Capital": "Road Town",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "VI",
    "Name": "U.S. Virgin Islands",
    "Capital": "Charlotte Amalie",
    "Continent": "NA",
    "ContinentName": "North America",
    "CurrencyCode": "USD",
    "LanguageCodes": [
      "en"
    ]
  },
  {
    "Code": "VN",
    "Name": "Vietnam",
    "Capital": "Hanoi",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "VND",
    "LanguageCodes": [
      "vi",
      "en",
      "fr",
      "zh",
      "km"
    ]
  },
  {
    "Code": "VU",
    "Name": "Vanuatu",
    "Capital": "Port Vila",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "VUV",
    "LanguageCodes": [
      "bi",
      "en",
      "fr"
    ]
  },
  {
    "Code": "WF",
    "Name": "Wallis and Futuna",
    "Capital": "Mata-Utu",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "XPF",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "WS",
    "Name": "Samoa",
    "Capital": "Apia",
    "Continent": "OC",
    "ContinentName": "Oceania",
    "CurrencyCode": "WST",
    "LanguageCodes": [
      "sm",
      "en"
    ]
  },
  {
    "Code": "XK",
    "Name": "Kosovo",
    "Capital": "Pristina",
    "Continent": "EU",
    "ContinentName": "Europe",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "sq",
      "sr"
    ]
  },
  {
    "Code": "YE",
    "Name": "Yemen",
    "Capital": "Sanaa",
    "Continent": "AS",
    "ContinentName": "Asia",
    "CurrencyCode": "YER",
    "LanguageCodes": [
      "ar"
    ]
  },
  {
    "Code": "YT",
    "Name": "Mayotte",
    "Capital": "Mamoudzou",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "EUR",
    "LanguageCodes": [
      "fr"
    ]
  },
  {
    "Code": "ZA",
    "Name": "South Africa",
    "Capital": "Pretoria",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "ZAR",
    "LanguageCodes": [
      "xh",
      "af",
      "en",
      "tn",
      "st",
      "ts",
      "ss",
      "ve",
      "nr"
    ]
  },
  {
    "Code": "ZM",
    "Name": "Zambia",
    "Capital": "Lusaka",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "ZMW",
    "LanguageCodes": [
      "en",
      "ny"
    ]
  },
  {
    "Code": "ZW",
    "Name": "Zimbabwe",
    "Capital": "Harare",
    "Continent": "AF",
    "ContinentName": "Africa",
    "CurrencyCode": "ZWL",
    "LanguageCodes": [
      "en",
      "sn",
      "nr",
      "nd"
    ]
  }
]`
