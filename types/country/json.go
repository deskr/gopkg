package country

import (
	"bytes"
	"encoding/json"
)

var countries []Country
var countriesByCode map[Code]Country
var countriesByContinent map[Continent][]Country
var countriesByCurrency map[Currency][]Country

func init() {
	// countries
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

	countriesByCurrency = make(map[Currency][]Country)
	for _, v := range countries {
		countriesByCurrency[v.Currency] = append(countriesByCurrency[v.Currency], v)
	}
}

const countriesJSON = `[
  {
    "Code": "AD",
    "Name": "Andorra",
    "Capital": "Andorra la Vella",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "ca",
        "Name": "Catalan; Valencian",
        "NativeName": "Català"
      }
    ]
  },
  {
    "Code": "AE",
    "Name": "United Arab Emirates",
    "Capital": "Abu Dhabi",
    "Currency": "AED",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "fa",
        "Name": "Persian",
        "NativeName": "فارسی"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "hi",
        "Name": "Hindi",
        "NativeName": "हिन्दी, हिंदी"
      },
      {
        "Code": "ur",
        "Name": "Urdu",
        "NativeName": "اردو"
      }
    ]
  },
  {
    "Code": "AF",
    "Name": "Afghanistan",
    "Capital": "Kabul",
    "Currency": "AFN",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "fa",
        "Name": "Persian",
        "NativeName": "فارسی"
      },
      {
        "Code": "ps",
        "Name": "Pashto, Pushto",
        "NativeName": "پښتو"
      },
      {
        "Code": "uz",
        "Name": "Uzbek",
        "NativeName": "Zbek, Ўзбек, أۇزبېك‎"
      },
      {
        "Code": "tk",
        "Name": "Turkmen",
        "NativeName": "Türkmen, Түркмен"
      }
    ]
  },
  {
    "Code": "AG",
    "Name": "Antigua and Barbuda",
    "Capital": "St. John's",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "AI",
    "Name": "Anguilla",
    "Capital": "The Valley",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "AL",
    "Name": "Albania",
    "Capital": "Tirana",
    "Currency": "ALL",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sq",
        "Name": "Albanian",
        "NativeName": "Shqip"
      },
      {
        "Code": "el",
        "Name": "Greek, Modern",
        "NativeName": "Ελληνικά"
      }
    ]
  },
  {
    "Code": "AM",
    "Name": "Armenia",
    "Capital": "Yerevan",
    "Currency": "AMD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "hy",
        "Name": "Armenian",
        "NativeName": "Հայերեն"
      }
    ]
  },
  {
    "Code": "AO",
    "Name": "Angola",
    "Capital": "Luanda",
    "Currency": "AOA",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "AQ",
    "Name": "Antarctica",
    "Capital": "",
    "Currency": "",
    "Continent": "AN",
    "ContientName": "Antarctica",
    "Languages": null
  },
  {
    "Code": "AR",
    "Name": "Argentina",
    "Capital": "Buenos Aires",
    "Currency": "ARS",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      },
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "gn",
        "Name": "Guaraní",
        "NativeName": "Avañeẽ"
      }
    ]
  },
  {
    "Code": "AS",
    "Name": "American Samoa",
    "Capital": "Pago Pago",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "sm",
        "Name": "Samoan",
        "NativeName": "Gagana Faa Samoa"
      },
      {
        "Code": "to",
        "Name": "Tonga (Tonga Islands)",
        "NativeName": "Faka Tonga"
      }
    ]
  },
  {
    "Code": "AT",
    "Name": "Austria",
    "Capital": "Vienna",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      },
      {
        "Code": "hr",
        "Name": "Croatian",
        "NativeName": "Hrvatski"
      },
      {
        "Code": "hu",
        "Name": "Hungarian",
        "NativeName": "Magyar"
      },
      {
        "Code": "sl",
        "Name": "Slovene",
        "NativeName": "Slovenščina"
      }
    ]
  },
  {
    "Code": "AU",
    "Name": "Australia",
    "Capital": "Canberra",
    "Currency": "AUD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "AW",
    "Name": "Aruba",
    "Capital": "Oranjestad",
    "Currency": "AWG",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "AX",
    "Name": "Åland",
    "Capital": "Mariehamn",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sv",
        "Name": "Swedish",
        "NativeName": "Svenska"
      }
    ]
  },
  {
    "Code": "AZ",
    "Name": "Azerbaijan",
    "Capital": "Baku",
    "Currency": "AZN",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "az",
        "Name": "Azerbaijani",
        "NativeName": "Azərbaycan Dili"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "hy",
        "Name": "Armenian",
        "NativeName": "Հայերեն"
      }
    ]
  },
  {
    "Code": "BA",
    "Name": "Bosnia and Herzegovina",
    "Capital": "Sarajevo",
    "Currency": "BAM",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "bs",
        "Name": "Bosnian",
        "NativeName": "Bosanski Jezik"
      },
      {
        "Code": "hr",
        "Name": "Croatian",
        "NativeName": "Hrvatski"
      },
      {
        "Code": "sr",
        "Name": "Serbian",
        "NativeName": "Српски Језик"
      }
    ]
  },
  {
    "Code": "BB",
    "Name": "Barbados",
    "Capital": "Bridgetown",
    "Currency": "BBD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "BD",
    "Name": "Bangladesh",
    "Capital": "Dhaka",
    "Currency": "BDT",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "bn",
        "Name": "Bengali",
        "NativeName": "বাংলা"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "BE",
    "Name": "Belgium",
    "Capital": "Brussels",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      }
    ]
  },
  {
    "Code": "BF",
    "Name": "Burkina Faso",
    "Capital": "Ouagadougou",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "BG",
    "Name": "Bulgaria",
    "Capital": "Sofia",
    "Currency": "BGN",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "bg",
        "Name": "Bulgarian",
        "NativeName": "Български Език"
      },
      {
        "Code": "tr",
        "Name": "Turkish",
        "NativeName": "Türkçe"
      }
    ]
  },
  {
    "Code": "BH",
    "Name": "Bahrain",
    "Capital": "Manama",
    "Currency": "BHD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fa",
        "Name": "Persian",
        "NativeName": "فارسی"
      },
      {
        "Code": "ur",
        "Name": "Urdu",
        "NativeName": "اردو"
      }
    ]
  },
  {
    "Code": "BI",
    "Name": "Burundi",
    "Capital": "Bujumbura",
    "Currency": "BIF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "rn",
        "Name": "Kirundi",
        "NativeName": "KiRundi"
      }
    ]
  },
  {
    "Code": "BJ",
    "Name": "Benin",
    "Capital": "Porto-Novo",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "BL",
    "Name": "Saint Barthélemy",
    "Capital": "Gustavia",
    "Currency": "EUR",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "BM",
    "Name": "Bermuda",
    "Capital": "Hamilton",
    "Currency": "BMD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "BN",
    "Name": "Brunei",
    "Capital": "Bandar Seri Begawan",
    "Currency": "BND",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ms",
        "Name": "Malay",
        "NativeName": "Bahasa Melayu, بهاس ملايو‎"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "BO",
    "Name": "Bolivia",
    "Capital": "Sucre",
    "Currency": "BOB",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "qu",
        "Name": "Quechua",
        "NativeName": "Runa Simi, Kichwa"
      },
      {
        "Code": "ay",
        "Name": "Aymara",
        "NativeName": "Aymar Aru"
      }
    ]
  },
  {
    "Code": "BQ",
    "Name": "Bonaire",
    "Capital": "Kralendijk",
    "Currency": "USD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "BR",
    "Name": "Brazil",
    "Capital": "Brasília",
    "Currency": "BRL",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "BS",
    "Name": "Bahamas",
    "Capital": "Nassau",
    "Currency": "BSD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "BT",
    "Name": "Bhutan",
    "Capital": "Thimphu",
    "Currency": "BTN",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": null
  },
  {
    "Code": "BV",
    "Name": "Bouvet Island",
    "Capital": "",
    "Currency": "NOK",
    "Continent": "AN",
    "ContientName": "Antarctica",
    "Languages": null
  },
  {
    "Code": "BW",
    "Name": "Botswana",
    "Capital": "Gaborone",
    "Currency": "BWP",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "tn",
        "Name": "Tswana",
        "NativeName": "Setswana"
      }
    ]
  },
  {
    "Code": "BY",
    "Name": "Belarus",
    "Capital": "Minsk",
    "Currency": "BYR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "be",
        "Name": "Belarusian",
        "NativeName": "Беларуская"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      }
    ]
  },
  {
    "Code": "BZ",
    "Name": "Belize",
    "Capital": "Belmopan",
    "Currency": "BZD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "CA",
    "Name": "Canada",
    "Capital": "Ottawa",
    "Currency": "CAD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "iu",
        "Name": "Inuktitut",
        "NativeName": "ᐃᓄᒃᑎᑐᑦ"
      }
    ]
  },
  {
    "Code": "CC",
    "Name": "Cocos [Keeling] Islands",
    "Capital": "West Island",
    "Currency": "AUD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ms",
        "Name": "Malay",
        "NativeName": "Bahasa Melayu, بهاس ملايو‎"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "CD",
    "Name": "Democratic Republic of the Congo",
    "Capital": "Kinshasa",
    "Currency": "CDF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "ln",
        "Name": "Lingala",
        "NativeName": "Lingála"
      },
      {
        "Code": "kg",
        "Name": "Kongo",
        "NativeName": "KiKongo"
      }
    ]
  },
  {
    "Code": "CF",
    "Name": "Central African Republic",
    "Capital": "Bangui",
    "Currency": "XAF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "sg",
        "Name": "Sango",
        "NativeName": "Yângâ Tî Sängö"
      },
      {
        "Code": "ln",
        "Name": "Lingala",
        "NativeName": "Lingála"
      },
      {
        "Code": "kg",
        "Name": "Kongo",
        "NativeName": "KiKongo"
      }
    ]
  },
  {
    "Code": "CG",
    "Name": "Republic of the Congo",
    "Capital": "Brazzaville",
    "Currency": "XAF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "kg",
        "Name": "Kongo",
        "NativeName": "KiKongo"
      },
      {
        "Code": "ln",
        "Name": "Lingala",
        "NativeName": "Lingála"
      }
    ]
  },
  {
    "Code": "CH",
    "Name": "Switzerland",
    "Capital": "Bern",
    "Currency": "CHF",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      },
      {
        "Code": "rm",
        "Name": "Romansh",
        "NativeName": "Rumantsch Grischun"
      }
    ]
  },
  {
    "Code": "CI",
    "Name": "Ivory Coast",
    "Capital": "Yamoussoukro",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "CK",
    "Name": "Cook Islands",
    "Capital": "Avarua",
    "Currency": "NZD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "mi",
        "Name": "Māori",
        "NativeName": "Te Reo Māori"
      }
    ]
  },
  {
    "Code": "CL",
    "Name": "Chile",
    "Capital": "Santiago",
    "Currency": "CLP",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "CM",
    "Name": "Cameroon",
    "Capital": "Yaoundé",
    "Currency": "XAF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "CN",
    "Name": "China",
    "Capital": "Beijing",
    "Currency": "CNY",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "ug",
        "Name": "Uighur, Uyghur",
        "NativeName": "Uyƣurqə, ئۇيغۇرچە‎"
      },
      {
        "Code": "za",
        "Name": "Zhuang, Chuang",
        "NativeName": "Saɯ Cueŋƅ, Saw Cuengh"
      }
    ]
  },
  {
    "Code": "CO",
    "Name": "Colombia",
    "Capital": "Bogotá",
    "Currency": "COP",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "CR",
    "Name": "Costa Rica",
    "Capital": "San José",
    "Currency": "CRC",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "CU",
    "Name": "Cuba",
    "Capital": "Havana",
    "Currency": "CUP",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "CV",
    "Name": "Cape Verde",
    "Capital": "Praia",
    "Currency": "CVE",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "CW",
    "Name": "Curacao",
    "Capital": "Willemstad",
    "Currency": "ANG",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      }
    ]
  },
  {
    "Code": "CX",
    "Name": "Christmas Island",
    "Capital": "Flying Fish Cove",
    "Currency": "AUD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "ms",
        "Name": "Malay",
        "NativeName": "Bahasa Melayu, بهاس ملايو‎"
      }
    ]
  },
  {
    "Code": "CY",
    "Name": "Cyprus",
    "Capital": "Nicosia",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "el",
        "Name": "Greek, Modern",
        "NativeName": "Ελληνικά"
      },
      {
        "Code": "tr",
        "Name": "Turkish",
        "NativeName": "Türkçe"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "CZ",
    "Name": "Czech Republic",
    "Capital": "Prague",
    "Currency": "CZK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "cs",
        "Name": "Czech",
        "NativeName": "Česky, Čeština"
      },
      {
        "Code": "sk",
        "Name": "Slovak",
        "NativeName": "Slovenčina"
      }
    ]
  },
  {
    "Code": "DE",
    "Name": "Germany",
    "Capital": "Berlin",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      }
    ]
  },
  {
    "Code": "DJ",
    "Name": "Djibouti",
    "Capital": "Djibouti",
    "Currency": "DJF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "so",
        "Name": "Somali",
        "NativeName": "Soomaaliga, Af Soomaali"
      },
      {
        "Code": "aa",
        "Name": "Afar",
        "NativeName": "Afaraf"
      }
    ]
  },
  {
    "Code": "DK",
    "Name": "Denmark",
    "Capital": "Copenhagen",
    "Currency": "DKK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "da",
        "Name": "Danish",
        "NativeName": "Dansk"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fo",
        "Name": "Faroese",
        "NativeName": "Føroyskt"
      },
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      }
    ]
  },
  {
    "Code": "DM",
    "Name": "Dominica",
    "Capital": "Roseau",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "DO",
    "Name": "Dominican Republic",
    "Capital": "Santo Domingo",
    "Currency": "DOP",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "DZ",
    "Name": "Algeria",
    "Capital": "Algiers",
    "Currency": "DZD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "EC",
    "Name": "Ecuador",
    "Capital": "Quito",
    "Currency": "USD",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "EE",
    "Name": "Estonia",
    "Capital": "Tallinn",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "et",
        "Name": "Estonian",
        "NativeName": "Eesti, Eesti Keel"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      }
    ]
  },
  {
    "Code": "EG",
    "Name": "Egypt",
    "Capital": "Cairo",
    "Currency": "EGP",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "EH",
    "Name": "Western Sahara",
    "Capital": "Laâyoune / El Aaiún",
    "Currency": "MAD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "ER",
    "Name": "Eritrea",
    "Capital": "Asmara",
    "Currency": "ERN",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "aa",
        "Name": "Afar",
        "NativeName": "Afaraf"
      },
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "ti",
        "Name": "Tigrinya",
        "NativeName": "ትግርኛ"
      }
    ]
  },
  {
    "Code": "ES",
    "Name": "Spain",
    "Capital": "Madrid",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "ca",
        "Name": "Catalan; Valencian",
        "NativeName": "Català"
      },
      {
        "Code": "gl",
        "Name": "Galician",
        "NativeName": "Galego"
      },
      {
        "Code": "eu",
        "Name": "Basque",
        "NativeName": "Euskara, Euskera"
      },
      {
        "Code": "oc",
        "Name": "Occitan",
        "NativeName": "Occitan"
      }
    ]
  },
  {
    "Code": "ET",
    "Name": "Ethiopia",
    "Capital": "Addis Ababa",
    "Currency": "ETB",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "am",
        "Name": "Amharic",
        "NativeName": "አማርኛ"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "om",
        "Name": "Oromo",
        "NativeName": "Afaan Oromoo"
      },
      {
        "Code": "ti",
        "Name": "Tigrinya",
        "NativeName": "ትግርኛ"
      },
      {
        "Code": "so",
        "Name": "Somali",
        "NativeName": "Soomaaliga, Af Soomaali"
      }
    ]
  },
  {
    "Code": "FI",
    "Name": "Finland",
    "Capital": "Helsinki",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "fi",
        "Name": "Finnish",
        "NativeName": "Suomi, Suomen Kieli"
      },
      {
        "Code": "sv",
        "Name": "Swedish",
        "NativeName": "Svenska"
      }
    ]
  },
  {
    "Code": "FJ",
    "Name": "Fiji",
    "Capital": "Suva",
    "Currency": "FJD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fj",
        "Name": "Fijian",
        "NativeName": "Vosa Vakaviti"
      }
    ]
  },
  {
    "Code": "FK",
    "Name": "Falkland Islands",
    "Capital": "Stanley",
    "Currency": "FKP",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "FM",
    "Name": "Micronesia",
    "Capital": "Palikir",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "FO",
    "Name": "Faroe Islands",
    "Capital": "Tórshavn",
    "Currency": "DKK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "fo",
        "Name": "Faroese",
        "NativeName": "Føroyskt"
      },
      {
        "Code": "da",
        "Name": "Danish",
        "NativeName": "Dansk"
      }
    ]
  },
  {
    "Code": "FR",
    "Name": "France",
    "Capital": "Paris",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "br",
        "Name": "Breton",
        "NativeName": "Brezhoneg"
      },
      {
        "Code": "co",
        "Name": "Corsican",
        "NativeName": "Corsu, Lingua Corsa"
      },
      {
        "Code": "ca",
        "Name": "Catalan; Valencian",
        "NativeName": "Català"
      },
      {
        "Code": "eu",
        "Name": "Basque",
        "NativeName": "Euskara, Euskera"
      },
      {
        "Code": "oc",
        "Name": "Occitan",
        "NativeName": "Occitan"
      }
    ]
  },
  {
    "Code": "GA",
    "Name": "Gabon",
    "Capital": "Libreville",
    "Currency": "XAF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "GB",
    "Name": "United Kingdom",
    "Capital": "London",
    "Currency": "GBP",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "cy",
        "Name": "Welsh",
        "NativeName": "Cymraeg"
      },
      {
        "Code": "gd",
        "Name": "Scottish Gaelic; Gaelic",
        "NativeName": "Gàidhlig"
      }
    ]
  },
  {
    "Code": "GD",
    "Name": "Grenada",
    "Capital": "St. George's",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "GE",
    "Name": "Georgia",
    "Capital": "Tbilisi",
    "Currency": "GEL",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ka",
        "Name": "Georgian",
        "NativeName": "ქართული"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "hy",
        "Name": "Armenian",
        "NativeName": "Հայերեն"
      },
      {
        "Code": "az",
        "Name": "Azerbaijani",
        "NativeName": "Azərbaycan Dili"
      }
    ]
  },
  {
    "Code": "GF",
    "Name": "French Guiana",
    "Capital": "Cayenne",
    "Currency": "EUR",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "GG",
    "Name": "Guernsey",
    "Capital": "St Peter Port",
    "Currency": "GBP",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "GH",
    "Name": "Ghana",
    "Capital": "Accra",
    "Currency": "GHS",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ak",
        "Name": "Akan",
        "NativeName": "Akan"
      },
      {
        "Code": "ee",
        "Name": "Ewe",
        "NativeName": "Eʋegbe"
      },
      {
        "Code": "tw",
        "Name": "Twi",
        "NativeName": "Twi"
      }
    ]
  },
  {
    "Code": "GI",
    "Name": "Gibraltar",
    "Capital": "Gibraltar",
    "Currency": "GIP",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      },
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "GL",
    "Name": "Greenland",
    "Capital": "Nuuk",
    "Currency": "DKK",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "kl",
        "Name": "Kalaallisut, Greenlandic",
        "NativeName": "Kalaallisut, Kalaallit Oqaasii"
      },
      {
        "Code": "da",
        "Name": "Danish",
        "NativeName": "Dansk"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "GM",
    "Name": "Gambia",
    "Capital": "Bathurst",
    "Currency": "GMD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "wo",
        "Name": "Wolof",
        "NativeName": "Wollof"
      },
      {
        "Code": "ff",
        "Name": "Fula; Fulah; Pulaar; Pular",
        "NativeName": "Fulfulde, Pulaar, Pular"
      }
    ]
  },
  {
    "Code": "GN",
    "Name": "Guinea",
    "Capital": "Conakry",
    "Currency": "GNF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "GP",
    "Name": "Guadeloupe",
    "Capital": "Basse-Terre",
    "Currency": "EUR",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "GQ",
    "Name": "Equatorial Guinea",
    "Capital": "Malabo",
    "Currency": "XAF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "GR",
    "Name": "Greece",
    "Capital": "Athens",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "el",
        "Name": "Greek, Modern",
        "NativeName": "Ελληνικά"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "GS",
    "Name": "South Georgia and the South Sandwich Islands",
    "Capital": "Grytviken",
    "Currency": "GBP",
    "Continent": "AN",
    "ContientName": "Antarctica",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "GT",
    "Name": "Guatemala",
    "Capital": "Guatemala City",
    "Currency": "GTQ",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "GU",
    "Name": "Guam",
    "Capital": "Hagåtña",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ch",
        "Name": "Chamorro",
        "NativeName": "Chamoru"
      }
    ]
  },
  {
    "Code": "GW",
    "Name": "Guinea-Bissau",
    "Capital": "Bissau",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "GY",
    "Name": "Guyana",
    "Capital": "Georgetown",
    "Currency": "GYD",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "HK",
    "Name": "Hong Kong",
    "Capital": "Hong Kong",
    "Currency": "HKD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "HM",
    "Name": "Heard Island and McDonald Islands",
    "Capital": "",
    "Currency": "AUD",
    "Continent": "AN",
    "ContientName": "Antarctica",
    "Languages": null
  },
  {
    "Code": "HN",
    "Name": "Honduras",
    "Capital": "Tegucigalpa",
    "Currency": "HNL",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "HR",
    "Name": "Croatia",
    "Capital": "Zagreb",
    "Currency": "HRK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "hr",
        "Name": "Croatian",
        "NativeName": "Hrvatski"
      },
      {
        "Code": "sr",
        "Name": "Serbian",
        "NativeName": "Српски Језик"
      }
    ]
  },
  {
    "Code": "HT",
    "Name": "Haiti",
    "Capital": "Port-au-Prince",
    "Currency": "HTG",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "ht",
        "Name": "Haitian; Haitian Creole",
        "NativeName": "Kreyòl Ayisyen"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "HU",
    "Name": "Hungary",
    "Capital": "Budapest",
    "Currency": "HUF",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "hu",
        "Name": "Hungarian",
        "NativeName": "Magyar"
      }
    ]
  },
  {
    "Code": "ID",
    "Name": "Indonesia",
    "Capital": "Jakarta",
    "Currency": "IDR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "id",
        "Name": "Indonesian",
        "NativeName": "Bahasa Indonesia"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      },
      {
        "Code": "jv",
        "Name": "Javanese",
        "NativeName": "Basa Jawa"
      }
    ]
  },
  {
    "Code": "IE",
    "Name": "Ireland",
    "Capital": "Dublin",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ga",
        "Name": "Irish",
        "NativeName": "Gaeilge"
      }
    ]
  },
  {
    "Code": "IL",
    "Name": "Israel",
    "Capital": "",
    "Currency": "ILS",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "he",
        "Name": "Hebrew (modern)",
        "NativeName": "עברית"
      },
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "IM",
    "Name": "Isle of Man",
    "Capital": "Douglas",
    "Currency": "GBP",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "gv",
        "Name": "Manx",
        "NativeName": "Gaelg, Gailck"
      }
    ]
  },
  {
    "Code": "IN",
    "Name": "India",
    "Capital": "New Delhi",
    "Currency": "INR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "hi",
        "Name": "Hindi",
        "NativeName": "हिन्दी, हिंदी"
      },
      {
        "Code": "bn",
        "Name": "Bengali",
        "NativeName": "বাংলা"
      },
      {
        "Code": "te",
        "Name": "Telugu",
        "NativeName": "తెలుగు"
      },
      {
        "Code": "mr",
        "Name": "Marathi (Marāṭhī)",
        "NativeName": "मराठी"
      },
      {
        "Code": "ta",
        "Name": "Tamil",
        "NativeName": "தமிழ்"
      },
      {
        "Code": "ur",
        "Name": "Urdu",
        "NativeName": "اردو"
      },
      {
        "Code": "gu",
        "Name": "Gujarati",
        "NativeName": "ગુજરાતી"
      },
      {
        "Code": "kn",
        "Name": "Kannada",
        "NativeName": "ಕನ್ನಡ"
      },
      {
        "Code": "ml",
        "Name": "Malayalam",
        "NativeName": "മലയാളം"
      },
      {
        "Code": "or",
        "Name": "Oriya",
        "NativeName": "ଓଡ଼ିଆ"
      },
      {
        "Code": "pa",
        "Name": "Panjabi, Punjabi",
        "NativeName": "ਪੰਜਾਬੀ, پنجابی‎"
      },
      {
        "Code": "as",
        "Name": "Assamese",
        "NativeName": "অসমীয়া"
      },
      {
        "Code": "bh",
        "Name": "Bihari",
        "NativeName": "भोजपुरी"
      },
      {
        "Code": "ks",
        "Name": "Kashmiri",
        "NativeName": "कश्मीरी, كشميري‎"
      },
      {
        "Code": "ne",
        "Name": "Nepali",
        "NativeName": "नेपाली"
      },
      {
        "Code": "sd",
        "Name": "Sindhi",
        "NativeName": "सिन्धी, سنڌي، سندھی‎"
      },
      {
        "Code": "sa",
        "Name": "Sanskrit (Saṁskṛta)",
        "NativeName": "संस्कृतम्"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "IO",
    "Name": "British Indian Ocean Territory",
    "Capital": "",
    "Currency": "USD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "IQ",
    "Name": "Iraq",
    "Capital": "Baghdad",
    "Currency": "IQD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "ku",
        "Name": "Kurdish",
        "NativeName": "Kurdî, كوردی‎"
      },
      {
        "Code": "hy",
        "Name": "Armenian",
        "NativeName": "Հայերեն"
      }
    ]
  },
  {
    "Code": "IR",
    "Name": "Iran",
    "Capital": "Tehran",
    "Currency": "IRR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "fa",
        "Name": "Persian",
        "NativeName": "فارسی"
      },
      {
        "Code": "ku",
        "Name": "Kurdish",
        "NativeName": "Kurdî, كوردی‎"
      }
    ]
  },
  {
    "Code": "IS",
    "Name": "Iceland",
    "Capital": "Reykjavik",
    "Currency": "ISK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "is",
        "Name": "Icelandic",
        "NativeName": "Íslenska"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      },
      {
        "Code": "da",
        "Name": "Danish",
        "NativeName": "Dansk"
      },
      {
        "Code": "sv",
        "Name": "Swedish",
        "NativeName": "Svenska"
      },
      {
        "Code": "no",
        "Name": "Norwegian",
        "NativeName": "Norsk"
      }
    ]
  },
  {
    "Code": "IT",
    "Name": "Italy",
    "Capital": "Rome",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      },
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "sc",
        "Name": "Sardinian",
        "NativeName": "Sardu"
      },
      {
        "Code": "ca",
        "Name": "Catalan; Valencian",
        "NativeName": "Català"
      },
      {
        "Code": "co",
        "Name": "Corsican",
        "NativeName": "Corsu, Lingua Corsa"
      },
      {
        "Code": "sl",
        "Name": "Slovene",
        "NativeName": "Slovenščina"
      }
    ]
  },
  {
    "Code": "JE",
    "Name": "Jersey",
    "Capital": "Saint Helier",
    "Currency": "GBP",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "JM",
    "Name": "Jamaica",
    "Capital": "Kingston",
    "Currency": "JMD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "JO",
    "Name": "Jordan",
    "Capital": "Amman",
    "Currency": "JOD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "JP",
    "Name": "Japan",
    "Capital": "Tokyo",
    "Currency": "JPY",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ja",
        "Name": "Japanese",
        "NativeName": "日本語 (にほんご／にっぽんご)"
      }
    ]
  },
  {
    "Code": "KE",
    "Name": "Kenya",
    "Capital": "Nairobi",
    "Currency": "KES",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "sw",
        "Name": "Swahili",
        "NativeName": "Kiswahili"
      }
    ]
  },
  {
    "Code": "KG",
    "Name": "Kyrgyzstan",
    "Capital": "Bishkek",
    "Currency": "KGS",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ky",
        "Name": "Kirghiz, Kyrgyz",
        "NativeName": "Кыргыз Тили"
      },
      {
        "Code": "uz",
        "Name": "Uzbek",
        "NativeName": "Zbek, Ўзбек, أۇزبېك‎"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      }
    ]
  },
  {
    "Code": "KH",
    "Name": "Cambodia",
    "Capital": "Phnom Penh",
    "Currency": "KHR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "km",
        "Name": "Khmer",
        "NativeName": "ភាសាខ្មែរ"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "KI",
    "Name": "Kiribati",
    "Capital": "Tarawa",
    "Currency": "AUD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "KM",
    "Name": "Comoros",
    "Capital": "Moroni",
    "Currency": "KMF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "KN",
    "Name": "Saint Kitts and Nevis",
    "Capital": "Basseterre",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "KP",
    "Name": "North Korea",
    "Capital": "Pyongyang",
    "Currency": "KPW",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ko",
        "Name": "Korean",
        "NativeName": "한국어 (韓國語), 조선말 (朝鮮語)"
      }
    ]
  },
  {
    "Code": "KR",
    "Name": "South Korea",
    "Capital": "Seoul",
    "Currency": "KRW",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ko",
        "Name": "Korean",
        "NativeName": "한국어 (韓國語), 조선말 (朝鮮語)"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "KW",
    "Name": "Kuwait",
    "Capital": "Kuwait City",
    "Currency": "KWD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "KY",
    "Name": "Cayman Islands",
    "Capital": "George Town",
    "Currency": "KYD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "KZ",
    "Name": "Kazakhstan",
    "Capital": "Astana",
    "Currency": "KZT",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "kk",
        "Name": "Kazakh",
        "NativeName": "Қазақ Тілі"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      }
    ]
  },
  {
    "Code": "LA",
    "Name": "Laos",
    "Capital": "Vientiane",
    "Currency": "LAK",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "lo",
        "Name": "Lao",
        "NativeName": "ພາສາລາວ"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "LB",
    "Name": "Lebanon",
    "Capital": "Beirut",
    "Currency": "LBP",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "hy",
        "Name": "Armenian",
        "NativeName": "Հայերեն"
      }
    ]
  },
  {
    "Code": "LC",
    "Name": "Saint Lucia",
    "Capital": "Castries",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "LI",
    "Name": "Liechtenstein",
    "Capital": "Vaduz",
    "Currency": "CHF",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      }
    ]
  },
  {
    "Code": "LK",
    "Name": "Sri Lanka",
    "Capital": "Colombo",
    "Currency": "LKR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "si",
        "Name": "Sinhala, Sinhalese",
        "NativeName": "සිංහල"
      },
      {
        "Code": "ta",
        "Name": "Tamil",
        "NativeName": "தமிழ்"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "LR",
    "Name": "Liberia",
    "Capital": "Monrovia",
    "Currency": "LRD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "LS",
    "Name": "Lesotho",
    "Capital": "Maseru",
    "Currency": "LSL",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "st",
        "Name": "Southern Sotho",
        "NativeName": "Sesotho"
      },
      {
        "Code": "xh",
        "Name": "Xhosa",
        "NativeName": "IsiXhosa"
      }
    ]
  },
  {
    "Code": "LT",
    "Name": "Lithuania",
    "Capital": "Vilnius",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "lt",
        "Name": "Lithuanian",
        "NativeName": "Lietuvių Kalba"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "pl",
        "Name": "Polish",
        "NativeName": "Polski"
      }
    ]
  },
  {
    "Code": "LU",
    "Name": "Luxembourg",
    "Capital": "Luxembourg",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "lb",
        "Name": "Luxembourgish, Letzeburgesch",
        "NativeName": "Lëtzebuergesch"
      },
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "LV",
    "Name": "Latvia",
    "Capital": "Riga",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "lv",
        "Name": "Latvian",
        "NativeName": "Latviešu Valoda"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "lt",
        "Name": "Lithuanian",
        "NativeName": "Lietuvių Kalba"
      }
    ]
  },
  {
    "Code": "LY",
    "Name": "Libya",
    "Capital": "Tripoli",
    "Currency": "LYD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "MA",
    "Name": "Morocco",
    "Capital": "Rabat",
    "Currency": "MAD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "MC",
    "Name": "Monaco",
    "Capital": "Monaco",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      }
    ]
  },
  {
    "Code": "MD",
    "Name": "Moldova",
    "Capital": "Chişinău",
    "Currency": "MDL",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "ro",
        "Name": "Romanian, Moldavian, Moldovan",
        "NativeName": "Română"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "tr",
        "Name": "Turkish",
        "NativeName": "Türkçe"
      }
    ]
  },
  {
    "Code": "ME",
    "Name": "Montenegro",
    "Capital": "Podgorica",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sr",
        "Name": "Serbian",
        "NativeName": "Српски Језик"
      },
      {
        "Code": "hu",
        "Name": "Hungarian",
        "NativeName": "Magyar"
      },
      {
        "Code": "bs",
        "Name": "Bosnian",
        "NativeName": "Bosanski Jezik"
      },
      {
        "Code": "sq",
        "Name": "Albanian",
        "NativeName": "Shqip"
      },
      {
        "Code": "hr",
        "Name": "Croatian",
        "NativeName": "Hrvatski"
      }
    ]
  },
  {
    "Code": "MF",
    "Name": "Saint Martin",
    "Capital": "Marigot",
    "Currency": "EUR",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "MG",
    "Name": "Madagascar",
    "Capital": "Antananarivo",
    "Currency": "MGA",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "mg",
        "Name": "Malagasy",
        "NativeName": "Malagasy Fiteny"
      }
    ]
  },
  {
    "Code": "MH",
    "Name": "Marshall Islands",
    "Capital": "Majuro",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "mh",
        "Name": "Marshallese",
        "NativeName": "Kajin M̧ajeļ"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "MK",
    "Name": "Macedonia",
    "Capital": "Skopje",
    "Currency": "MKD",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "mk",
        "Name": "Macedonian",
        "NativeName": "Македонски Јазик"
      },
      {
        "Code": "sq",
        "Name": "Albanian",
        "NativeName": "Shqip"
      },
      {
        "Code": "tr",
        "Name": "Turkish",
        "NativeName": "Türkçe"
      },
      {
        "Code": "sr",
        "Name": "Serbian",
        "NativeName": "Српски Језик"
      }
    ]
  },
  {
    "Code": "ML",
    "Name": "Mali",
    "Capital": "Bamako",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "bm",
        "Name": "Bambara",
        "NativeName": "Bamanankan"
      }
    ]
  },
  {
    "Code": "MM",
    "Name": "Myanmar [Burma]",
    "Capital": "Naypyitaw",
    "Currency": "MMK",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "my",
        "Name": "Burmese",
        "NativeName": "ဗမာစာ"
      }
    ]
  },
  {
    "Code": "MN",
    "Name": "Mongolia",
    "Capital": "Ulan Bator",
    "Currency": "MNT",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "mn",
        "Name": "Mongolian",
        "NativeName": "Монгол"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      }
    ]
  },
  {
    "Code": "MO",
    "Name": "Macao",
    "Capital": "Macao",
    "Currency": "MOP",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "MP",
    "Name": "Northern Mariana Islands",
    "Capital": "Saipan",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "tl",
        "Name": "Tagalog",
        "NativeName": "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "ch",
        "Name": "Chamorro",
        "NativeName": "Chamoru"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "MQ",
    "Name": "Martinique",
    "Capital": "Fort-de-France",
    "Currency": "EUR",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "MR",
    "Name": "Mauritania",
    "Capital": "Nouakchott",
    "Currency": "MRO",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "wo",
        "Name": "Wolof",
        "NativeName": "Wollof"
      }
    ]
  },
  {
    "Code": "MS",
    "Name": "Montserrat",
    "Capital": "Plymouth",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "MT",
    "Name": "Malta",
    "Capital": "Valletta",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "mt",
        "Name": "Maltese",
        "NativeName": "Malti"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "MU",
    "Name": "Mauritius",
    "Capital": "Port Louis",
    "Currency": "MUR",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "MV",
    "Name": "Maldives",
    "Capital": "Malé",
    "Currency": "MVR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "dv",
        "Name": "Divehi; Dhivehi; Maldivian;",
        "NativeName": "ދިވެހި"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "MW",
    "Name": "Malawi",
    "Capital": "Lilongwe",
    "Currency": "MWK",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ny",
        "Name": "Chichewa; Chewa; Nyanja",
        "NativeName": "ChiCheŵa, Chinyanja"
      }
    ]
  },
  {
    "Code": "MX",
    "Name": "Mexico",
    "Capital": "Mexico City",
    "Currency": "MXN",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "MY",
    "Name": "Malaysia",
    "Capital": "Kuala Lumpur",
    "Currency": "MYR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ms",
        "Name": "Malay",
        "NativeName": "Bahasa Melayu, بهاس ملايو‎"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "ta",
        "Name": "Tamil",
        "NativeName": "தமிழ்"
      },
      {
        "Code": "te",
        "Name": "Telugu",
        "NativeName": "తెలుగు"
      },
      {
        "Code": "ml",
        "Name": "Malayalam",
        "NativeName": "മലയാളം"
      },
      {
        "Code": "pa",
        "Name": "Panjabi, Punjabi",
        "NativeName": "ਪੰਜਾਬੀ, پنجابی‎"
      },
      {
        "Code": "th",
        "Name": "Thai",
        "NativeName": "ไทย"
      }
    ]
  },
  {
    "Code": "MZ",
    "Name": "Mozambique",
    "Capital": "Maputo",
    "Currency": "MZN",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "NA",
    "Name": "Namibia",
    "Capital": "Windhoek",
    "Currency": "NAD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "af",
        "Name": "Afrikaans",
        "NativeName": "Afrikaans"
      },
      {
        "Code": "de",
        "Name": "German",
        "NativeName": "Deutsch"
      },
      {
        "Code": "hz",
        "Name": "Herero",
        "NativeName": "Otjiherero"
      }
    ]
  },
  {
    "Code": "NC",
    "Name": "New Caledonia",
    "Capital": "Noumea",
    "Currency": "XPF",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "NE",
    "Name": "Niger",
    "Capital": "Niamey",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "ha",
        "Name": "Hausa",
        "NativeName": "Hausa, هَوُسَ"
      },
      {
        "Code": "kr",
        "Name": "Kanuri",
        "NativeName": "Kanuri"
      }
    ]
  },
  {
    "Code": "NF",
    "Name": "Norfolk Island",
    "Capital": "Kingston",
    "Currency": "AUD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "NG",
    "Name": "Nigeria",
    "Capital": "Abuja",
    "Currency": "NGN",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ha",
        "Name": "Hausa",
        "NativeName": "Hausa, هَوُسَ"
      },
      {
        "Code": "yo",
        "Name": "Yoruba",
        "NativeName": "Yorùbá"
      },
      {
        "Code": "ig",
        "Name": "Igbo",
        "NativeName": "Asụsụ Igbo"
      },
      {
        "Code": "ff",
        "Name": "Fula; Fulah; Pulaar; Pular",
        "NativeName": "Fulfulde, Pulaar, Pular"
      }
    ]
  },
  {
    "Code": "NI",
    "Name": "Nicaragua",
    "Capital": "Managua",
    "Currency": "NIO",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "NL",
    "Name": "Netherlands",
    "Capital": "Amsterdam",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      },
      {
        "Code": "fy",
        "Name": "Western Frisian",
        "NativeName": "Frysk"
      }
    ]
  },
  {
    "Code": "NO",
    "Name": "Norway",
    "Capital": "Oslo",
    "Currency": "NOK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "no",
        "Name": "Norwegian",
        "NativeName": "Norsk"
      },
      {
        "Code": "nb",
        "Name": "Norwegian Bokmål",
        "NativeName": "Norsk Bokmål"
      },
      {
        "Code": "nn",
        "Name": "Norwegian Nynorsk",
        "NativeName": "Norsk Nynorsk"
      },
      {
        "Code": "se",
        "Name": "Northern Sami",
        "NativeName": "Davvisámegiella"
      },
      {
        "Code": "fi",
        "Name": "Finnish",
        "NativeName": "Suomi, Suomen Kieli"
      }
    ]
  },
  {
    "Code": "NP",
    "Name": "Nepal",
    "Capital": "Kathmandu",
    "Currency": "NPR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ne",
        "Name": "Nepali",
        "NativeName": "नेपाली"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "NR",
    "Name": "Nauru",
    "Capital": "Yaren",
    "Currency": "AUD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "na",
        "Name": "Nauru",
        "NativeName": "Ekakairũ Naoero"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "NU",
    "Name": "Niue",
    "Capital": "Alofi",
    "Currency": "NZD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "NZ",
    "Name": "New Zealand",
    "Capital": "Wellington",
    "Currency": "NZD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "mi",
        "Name": "Māori",
        "NativeName": "Te Reo Māori"
      }
    ]
  },
  {
    "Code": "OM",
    "Name": "Oman",
    "Capital": "Muscat",
    "Currency": "OMR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ur",
        "Name": "Urdu",
        "NativeName": "اردو"
      }
    ]
  },
  {
    "Code": "PA",
    "Name": "Panama",
    "Capital": "Panama City",
    "Currency": "PAB",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "PE",
    "Name": "Peru",
    "Capital": "Lima",
    "Currency": "PEN",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "qu",
        "Name": "Quechua",
        "NativeName": "Runa Simi, Kichwa"
      },
      {
        "Code": "ay",
        "Name": "Aymara",
        "NativeName": "Aymar Aru"
      }
    ]
  },
  {
    "Code": "PF",
    "Name": "French Polynesia",
    "Capital": "Papeete",
    "Currency": "XPF",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "ty",
        "Name": "Tahitian",
        "NativeName": "Reo Tahiti"
      }
    ]
  },
  {
    "Code": "PG",
    "Name": "Papua New Guinea",
    "Capital": "Port Moresby",
    "Currency": "PGK",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ho",
        "Name": "Hiri Motu",
        "NativeName": "Hiri Motu"
      }
    ]
  },
  {
    "Code": "PH",
    "Name": "Philippines",
    "Capital": "Manila",
    "Currency": "PHP",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "tl",
        "Name": "Tagalog",
        "NativeName": "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "PK",
    "Name": "Pakistan",
    "Capital": "Islamabad",
    "Currency": "PKR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ur",
        "Name": "Urdu",
        "NativeName": "اردو"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "pa",
        "Name": "Panjabi, Punjabi",
        "NativeName": "ਪੰਜਾਬੀ, پنجابی‎"
      },
      {
        "Code": "sd",
        "Name": "Sindhi",
        "NativeName": "सिन्धी, سنڌي، سندھی‎"
      },
      {
        "Code": "ps",
        "Name": "Pashto, Pushto",
        "NativeName": "پښتو"
      }
    ]
  },
  {
    "Code": "PL",
    "Name": "Poland",
    "Capital": "Warsaw",
    "Currency": "PLN",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "pl",
        "Name": "Polish",
        "NativeName": "Polski"
      }
    ]
  },
  {
    "Code": "PM",
    "Name": "Saint Pierre and Miquelon",
    "Capital": "Saint-Pierre",
    "Currency": "EUR",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "PN",
    "Name": "Pitcairn Islands",
    "Capital": "Adamstown",
    "Currency": "NZD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "PR",
    "Name": "Puerto Rico",
    "Capital": "San Juan",
    "Currency": "USD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "PS",
    "Name": "Palestine",
    "Capital": "",
    "Currency": "ILS",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "PT",
    "Name": "Portugal",
    "Capital": "Lisbon",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "PW",
    "Name": "Palau",
    "Capital": "Melekeok",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ja",
        "Name": "Japanese",
        "NativeName": "日本語 (にほんご／にっぽんご)"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      }
    ]
  },
  {
    "Code": "PY",
    "Name": "Paraguay",
    "Capital": "Asunción",
    "Currency": "PYG",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "gn",
        "Name": "Guaraní",
        "NativeName": "Avañeẽ"
      }
    ]
  },
  {
    "Code": "QA",
    "Name": "Qatar",
    "Capital": "Doha",
    "Currency": "QAR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "RE",
    "Name": "Réunion",
    "Capital": "Saint-Denis",
    "Currency": "EUR",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "RO",
    "Name": "Romania",
    "Capital": "Bucharest",
    "Currency": "RON",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "ro",
        "Name": "Romanian, Moldavian, Moldovan",
        "NativeName": "Română"
      },
      {
        "Code": "hu",
        "Name": "Hungarian",
        "NativeName": "Magyar"
      }
    ]
  },
  {
    "Code": "RS",
    "Name": "Serbia",
    "Capital": "Belgrade",
    "Currency": "RSD",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sr",
        "Name": "Serbian",
        "NativeName": "Српски Језик"
      },
      {
        "Code": "hu",
        "Name": "Hungarian",
        "NativeName": "Magyar"
      },
      {
        "Code": "bs",
        "Name": "Bosnian",
        "NativeName": "Bosanski Jezik"
      }
    ]
  },
  {
    "Code": "RU",
    "Name": "Russia",
    "Capital": "Moscow",
    "Currency": "RUB",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "tt",
        "Name": "Tatar",
        "NativeName": "Татарча, Tatarça, تاتارچا‎"
      },
      {
        "Code": "kv",
        "Name": "Komi",
        "NativeName": "Коми Кыв"
      },
      {
        "Code": "ce",
        "Name": "Chechen",
        "NativeName": "Нохчийн Мотт"
      },
      {
        "Code": "cv",
        "Name": "Chuvash",
        "NativeName": "Чӑваш Чӗлхи"
      },
      {
        "Code": "ba",
        "Name": "Bashkir",
        "NativeName": "Башҡорт Теле"
      }
    ]
  },
  {
    "Code": "RW",
    "Name": "Rwanda",
    "Capital": "Kigali",
    "Currency": "RWF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "rw",
        "Name": "Kinyarwanda",
        "NativeName": "Ikinyarwanda"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "sw",
        "Name": "Swahili",
        "NativeName": "Kiswahili"
      }
    ]
  },
  {
    "Code": "SA",
    "Name": "Saudi Arabia",
    "Capital": "Riyadh",
    "Currency": "SAR",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "SB",
    "Name": "Solomon Islands",
    "Capital": "Honiara",
    "Currency": "SBD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "SC",
    "Name": "Seychelles",
    "Capital": "Victoria",
    "Currency": "SCR",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "SD",
    "Name": "Sudan",
    "Capital": "Khartoum",
    "Currency": "SDG",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "SE",
    "Name": "Sweden",
    "Capital": "Stockholm",
    "Currency": "SEK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sv",
        "Name": "Swedish",
        "NativeName": "Svenska"
      },
      {
        "Code": "se",
        "Name": "Northern Sami",
        "NativeName": "Davvisámegiella"
      },
      {
        "Code": "fi",
        "Name": "Finnish",
        "NativeName": "Suomi, Suomen Kieli"
      }
    ]
  },
  {
    "Code": "SG",
    "Name": "Singapore",
    "Capital": "Singapore",
    "Currency": "SGD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ms",
        "Name": "Malay",
        "NativeName": "Bahasa Melayu, بهاس ملايو‎"
      },
      {
        "Code": "ta",
        "Name": "Tamil",
        "NativeName": "தமிழ்"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      }
    ]
  },
  {
    "Code": "SH",
    "Name": "Saint Helena",
    "Capital": "Jamestown",
    "Currency": "SHP",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "SI",
    "Name": "Slovenia",
    "Capital": "Ljubljana",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sl",
        "Name": "Slovene",
        "NativeName": "Slovenščina"
      }
    ]
  },
  {
    "Code": "SJ",
    "Name": "Svalbard and Jan Mayen",
    "Capital": "Longyearbyen",
    "Currency": "NOK",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "no",
        "Name": "Norwegian",
        "NativeName": "Norsk"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      }
    ]
  },
  {
    "Code": "SK",
    "Name": "Slovakia",
    "Capital": "Bratislava",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sk",
        "Name": "Slovak",
        "NativeName": "Slovenčina"
      },
      {
        "Code": "hu",
        "Name": "Hungarian",
        "NativeName": "Magyar"
      }
    ]
  },
  {
    "Code": "SL",
    "Name": "Sierra Leone",
    "Capital": "Freetown",
    "Currency": "SLL",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "SM",
    "Name": "San Marino",
    "Capital": "San Marino",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      }
    ]
  },
  {
    "Code": "SN",
    "Name": "Senegal",
    "Capital": "Dakar",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "wo",
        "Name": "Wolof",
        "NativeName": "Wollof"
      }
    ]
  },
  {
    "Code": "SO",
    "Name": "Somalia",
    "Capital": "Mogadishu",
    "Currency": "SOS",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "so",
        "Name": "Somali",
        "NativeName": "Soomaaliga, Af Soomaali"
      },
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "SR",
    "Name": "Suriname",
    "Capital": "Paramaribo",
    "Currency": "SRD",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "jv",
        "Name": "Javanese",
        "NativeName": "Basa Jawa"
      }
    ]
  },
  {
    "Code": "SS",
    "Name": "South Sudan",
    "Capital": "Juba",
    "Currency": "SSP",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "ST",
    "Name": "São Tomé and Príncipe",
    "Capital": "São Tomé",
    "Currency": "STD",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      }
    ]
  },
  {
    "Code": "SV",
    "Name": "El Salvador",
    "Capital": "San Salvador",
    "Currency": "USD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "SX",
    "Name": "Sint Maarten",
    "Capital": "Philipsburg",
    "Currency": "ANG",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "nl",
        "Name": "Dutch",
        "NativeName": "Nederlands, Vlaams"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "SY",
    "Name": "Syria",
    "Capital": "Damascus",
    "Currency": "SYP",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "ku",
        "Name": "Kurdish",
        "NativeName": "Kurdî, كوردی‎"
      },
      {
        "Code": "hy",
        "Name": "Armenian",
        "NativeName": "Հայերեն"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "SZ",
    "Name": "Swaziland",
    "Capital": "Mbabane",
    "Currency": "SZL",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ss",
        "Name": "Swati",
        "NativeName": "SiSwati"
      }
    ]
  },
  {
    "Code": "TC",
    "Name": "Turks and Caicos Islands",
    "Capital": "Cockburn Town",
    "Currency": "USD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "TD",
    "Name": "Chad",
    "Capital": "N'Djamena",
    "Currency": "XAF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "TF",
    "Name": "French Southern Territories",
    "Capital": "Port-aux-Français",
    "Currency": "EUR",
    "Continent": "AN",
    "ContientName": "Antarctica",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "TG",
    "Name": "Togo",
    "Capital": "Lomé",
    "Currency": "XOF",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "ee",
        "Name": "Ewe",
        "NativeName": "Eʋegbe"
      },
      {
        "Code": "ha",
        "Name": "Hausa",
        "NativeName": "Hausa, هَوُسَ"
      }
    ]
  },
  {
    "Code": "TH",
    "Name": "Thailand",
    "Capital": "Bangkok",
    "Currency": "THB",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "th",
        "Name": "Thai",
        "NativeName": "ไทย"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "TJ",
    "Name": "Tajikistan",
    "Capital": "Dushanbe",
    "Currency": "TJS",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "tg",
        "Name": "Tajik",
        "NativeName": "Тоҷикӣ, Toğikī, تاجیکی‎"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      }
    ]
  },
  {
    "Code": "TK",
    "Name": "Tokelau",
    "Capital": "",
    "Currency": "NZD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "TL",
    "Name": "East Timor",
    "Capital": "Dili",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "pt",
        "Name": "Portuguese",
        "NativeName": "Português"
      },
      {
        "Code": "id",
        "Name": "Indonesian",
        "NativeName": "Bahasa Indonesia"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "TM",
    "Name": "Turkmenistan",
    "Capital": "Ashgabat",
    "Currency": "TMT",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "tk",
        "Name": "Turkmen",
        "NativeName": "Türkmen, Түркмен"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "uz",
        "Name": "Uzbek",
        "NativeName": "Zbek, Ўзбек, أۇزبېك‎"
      }
    ]
  },
  {
    "Code": "TN",
    "Name": "Tunisia",
    "Capital": "Tunis",
    "Currency": "TND",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "TO",
    "Name": "Tonga",
    "Capital": "Nuku'alofa",
    "Currency": "TOP",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "to",
        "Name": "Tonga (Tonga Islands)",
        "NativeName": "Faka Tonga"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "TR",
    "Name": "Turkey",
    "Capital": "Ankara",
    "Currency": "TRY",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "tr",
        "Name": "Turkish",
        "NativeName": "Türkçe"
      },
      {
        "Code": "ku",
        "Name": "Kurdish",
        "NativeName": "Kurdî, كوردی‎"
      },
      {
        "Code": "az",
        "Name": "Azerbaijani",
        "NativeName": "Azərbaycan Dili"
      },
      {
        "Code": "av",
        "Name": "Avaric",
        "NativeName": "Авар МацӀ, МагӀарул МацӀ"
      }
    ]
  },
  {
    "Code": "TT",
    "Name": "Trinidad and Tobago",
    "Capital": "Port of Spain",
    "Currency": "TTD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      }
    ]
  },
  {
    "Code": "TV",
    "Name": "Tuvalu",
    "Capital": "Funafuti",
    "Currency": "AUD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "sm",
        "Name": "Samoan",
        "NativeName": "Gagana Faa Samoa"
      }
    ]
  },
  {
    "Code": "TW",
    "Name": "Taiwan",
    "Capital": "Taipei",
    "Currency": "TWD",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      }
    ]
  },
  {
    "Code": "TZ",
    "Name": "Tanzania",
    "Capital": "Dodoma",
    "Currency": "TZS",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "sw",
        "Name": "Swahili",
        "NativeName": "Kiswahili"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "UA",
    "Name": "Ukraine",
    "Capital": "Kiev",
    "Currency": "UAH",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "uk",
        "Name": "Ukrainian",
        "NativeName": "Українська"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "pl",
        "Name": "Polish",
        "NativeName": "Polski"
      },
      {
        "Code": "hu",
        "Name": "Hungarian",
        "NativeName": "Magyar"
      }
    ]
  },
  {
    "Code": "UG",
    "Name": "Uganda",
    "Capital": "Kampala",
    "Currency": "UGX",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "lg",
        "Name": "Luganda",
        "NativeName": "Luganda"
      },
      {
        "Code": "sw",
        "Name": "Swahili",
        "NativeName": "Kiswahili"
      },
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "UM",
    "Name": "U.S. Minor Outlying Islands",
    "Capital": "",
    "Currency": "USD",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "US",
    "Name": "United States",
    "Capital": "Washington",
    "Currency": "USD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "UY",
    "Name": "Uruguay",
    "Capital": "Montevideo",
    "Currency": "UYU",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "UZ",
    "Name": "Uzbekistan",
    "Capital": "Tashkent",
    "Currency": "UZS",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "uz",
        "Name": "Uzbek",
        "NativeName": "Zbek, Ўзбек, أۇزبېك‎"
      },
      {
        "Code": "ru",
        "Name": "Russian",
        "NativeName": "Русский Язык"
      },
      {
        "Code": "tg",
        "Name": "Tajik",
        "NativeName": "Тоҷикӣ, Toğikī, تاجیکی‎"
      }
    ]
  },
  {
    "Code": "VA",
    "Name": "Vatican City",
    "Capital": "Vatican City",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "la",
        "Name": "Latin",
        "NativeName": "Latine, Lingua Latina"
      },
      {
        "Code": "it",
        "Name": "Italian",
        "NativeName": "Italiano"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "VC",
    "Name": "Saint Vincent and the Grenadines",
    "Capital": "Kingstown",
    "Currency": "XCD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "VE",
    "Name": "Venezuela",
    "Capital": "Caracas",
    "Currency": "VEF",
    "Continent": "SA",
    "ContientName": "South America",
    "Languages": [
      {
        "Code": "es",
        "Name": "Spanish; Castilian",
        "NativeName": "Español, Castellano"
      }
    ]
  },
  {
    "Code": "VG",
    "Name": "British Virgin Islands",
    "Capital": "Road Town",
    "Currency": "USD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "VI",
    "Name": "U.S. Virgin Islands",
    "Capital": "Charlotte Amalie",
    "Currency": "USD",
    "Continent": "NA",
    "ContientName": "North America",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "VN",
    "Name": "Vietnam",
    "Capital": "Hanoi",
    "Currency": "VND",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "vi",
        "Name": "Vietnamese",
        "NativeName": "Tiếng Việt"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      },
      {
        "Code": "zh",
        "Name": "Chinese",
        "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
      },
      {
        "Code": "km",
        "Name": "Khmer",
        "NativeName": "ភាសាខ្មែរ"
      }
    ]
  },
  {
    "Code": "VU",
    "Name": "Vanuatu",
    "Capital": "Port Vila",
    "Currency": "VUV",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "bi",
        "Name": "Bislama",
        "NativeName": "Bislama"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "WF",
    "Name": "Wallis and Futuna",
    "Capital": "Mata-Utu",
    "Currency": "XPF",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "WS",
    "Name": "Samoa",
    "Capital": "Apia",
    "Currency": "WST",
    "Continent": "OC",
    "ContientName": "Oceania",
    "Languages": [
      {
        "Code": "sm",
        "Name": "Samoan",
        "NativeName": "Gagana Faa Samoa"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      }
    ]
  },
  {
    "Code": "XK",
    "Name": "Kosovo",
    "Capital": "Pristina",
    "Currency": "EUR",
    "Continent": "EU",
    "ContientName": "Europe",
    "Languages": [
      {
        "Code": "sq",
        "Name": "Albanian",
        "NativeName": "Shqip"
      },
      {
        "Code": "sr",
        "Name": "Serbian",
        "NativeName": "Српски Језик"
      }
    ]
  },
  {
    "Code": "YE",
    "Name": "Yemen",
    "Capital": "Sanaa",
    "Currency": "YER",
    "Continent": "AS",
    "ContientName": "Asia",
    "Languages": [
      {
        "Code": "ar",
        "Name": "Arabic",
        "NativeName": "العربية"
      }
    ]
  },
  {
    "Code": "YT",
    "Name": "Mayotte",
    "Capital": "Mamoudzou",
    "Currency": "EUR",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "fr",
        "Name": "French",
        "NativeName": "Français, Langue Française"
      }
    ]
  },
  {
    "Code": "ZA",
    "Name": "South Africa",
    "Capital": "Pretoria",
    "Currency": "ZAR",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "xh",
        "Name": "Xhosa",
        "NativeName": "IsiXhosa"
      },
      {
        "Code": "af",
        "Name": "Afrikaans",
        "NativeName": "Afrikaans"
      },
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "tn",
        "Name": "Tswana",
        "NativeName": "Setswana"
      },
      {
        "Code": "st",
        "Name": "Southern Sotho",
        "NativeName": "Sesotho"
      },
      {
        "Code": "ts",
        "Name": "Tsonga",
        "NativeName": "Xitsonga"
      },
      {
        "Code": "ss",
        "Name": "Swati",
        "NativeName": "SiSwati"
      },
      {
        "Code": "ve",
        "Name": "Venda",
        "NativeName": "Tshivenḓa"
      },
      {
        "Code": "nr",
        "Name": "South Ndebele",
        "NativeName": "IsiNdebele"
      }
    ]
  },
  {
    "Code": "ZM",
    "Name": "Zambia",
    "Capital": "Lusaka",
    "Currency": "ZMW",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "ny",
        "Name": "Chichewa; Chewa; Nyanja",
        "NativeName": "ChiCheŵa, Chinyanja"
      }
    ]
  },
  {
    "Code": "ZW",
    "Name": "Zimbabwe",
    "Capital": "Harare",
    "Currency": "ZWL",
    "Continent": "AF",
    "ContientName": "Africa",
    "Languages": [
      {
        "Code": "en",
        "Name": "English",
        "NativeName": "English"
      },
      {
        "Code": "sn",
        "Name": "Shona",
        "NativeName": "ChiShona"
      },
      {
        "Code": "nr",
        "Name": "South Ndebele",
        "NativeName": "IsiNdebele"
      },
      {
        "Code": "nd",
        "Name": "North Ndebele",
        "NativeName": "IsiNdebele"
      }
    ]
  }
]`
