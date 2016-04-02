package currency

import (
	"bytes"
	"encoding/json"
)

var currencies []Currency
var currenciesByCode map[Code]Currency

func init() {
	if err := json.NewDecoder(bytes.NewReader([]byte(currenciesJSON))).Decode(&currencies); err != nil {
		panic(err)
	}

	currenciesByCode = make(map[Code]Currency)
	for _, v := range currencies {
		currenciesByCode[v.Code] = v
	}
}

const currenciesJSON = `[
  {
    "name": "US Dollar",
    "symbol": "$",
    "native_symbol": "$",
    "plural_name": "US dollars",
    "decimal_digits": 2,
    "code": "USD"
  },
  {
    "name": "Canadian Dollar",
    "symbol": "CA$",
    "native_symbol": "$",
    "plural_name": "Canadian dollars",
    "decimal_digits": 2,
    "code": "CAD"
  },
  {
    "name": "Euro",
    "symbol": "€",
    "native_symbol": "€",
    "plural_name": "euros",
    "decimal_digits": 2,
    "code": "EUR"
  },
  {
    "name": "United Arab Emirates Dirham",
    "symbol": "AED",
    "native_symbol": "د.إ.‏",
    "plural_name": "UAE dirhams",
    "decimal_digits": 2,
    "code": "AED"
  },
  {
    "name": "Afghan Afghani",
    "symbol": "Af",
    "native_symbol": "؋",
    "plural_name": "Afghan Afghanis",
    "decimal_digits": 0,
    "code": "AFN"
  },
  {
    "name": "Albanian Lek",
    "symbol": "ALL",
    "native_symbol": "Lek",
    "plural_name": "Albanian lekë",
    "decimal_digits": 0,
    "code": "ALL"
  },
  {
    "name": "Armenian Dram",
    "symbol": "AMD",
    "native_symbol": "դր.",
    "plural_name": "Armenian drams",
    "decimal_digits": 0,
    "code": "AMD"
  },
  {
    "name": "Argentine Peso",
    "symbol": "AR$",
    "native_symbol": "$",
    "plural_name": "Argentine pesos",
    "decimal_digits": 2,
    "code": "ARS"
  },
  {
    "name": "Australian Dollar",
    "symbol": "AU$",
    "native_symbol": "$",
    "plural_name": "Australian dollars",
    "decimal_digits": 2,
    "code": "AUD"
  },
  {
    "name": "Azerbaijani Manat",
    "symbol": "man.",
    "native_symbol": "ман.",
    "plural_name": "Azerbaijani manats",
    "decimal_digits": 2,
    "code": "AZN"
  },
  {
    "name": "Bosnia-Herzegovina Convertible Mark",
    "symbol": "KM",
    "native_symbol": "KM",
    "plural_name": "Bosnia-Herzegovina convertible marks",
    "decimal_digits": 2,
    "code": "BAM"
  },
  {
    "name": "Bangladeshi Taka",
    "symbol": "Tk",
    "native_symbol": "৳",
    "plural_name": "Bangladeshi takas",
    "decimal_digits": 2,
    "code": "BDT"
  },
  {
    "name": "Bulgarian Lev",
    "symbol": "BGN",
    "native_symbol": "лв.",
    "plural_name": "Bulgarian leva",
    "decimal_digits": 2,
    "code": "BGN"
  },
  {
    "name": "Bahraini Dinar",
    "symbol": "BD",
    "native_symbol": "د.ب.‏",
    "plural_name": "Bahraini dinars",
    "decimal_digits": 3,
    "code": "BHD"
  },
  {
    "name": "Burundian Franc",
    "symbol": "FBu",
    "native_symbol": "FBu",
    "plural_name": "Burundian francs",
    "decimal_digits": 0,
    "code": "BIF"
  },
  {
    "name": "Brunei Dollar",
    "symbol": "BN$",
    "native_symbol": "$",
    "plural_name": "Brunei dollars",
    "decimal_digits": 2,
    "code": "BND"
  },
  {
    "name": "Bolivian Boliviano",
    "symbol": "Bs",
    "native_symbol": "Bs",
    "plural_name": "Bolivian bolivianos",
    "decimal_digits": 2,
    "code": "BOB"
  },
  {
    "name": "Brazilian Real",
    "symbol": "R$",
    "native_symbol": "R$",
    "plural_name": "Brazilian reals",
    "decimal_digits": 2,
    "code": "BRL"
  },
  {
    "name": "Botswanan Pula",
    "symbol": "BWP",
    "native_symbol": "P",
    "plural_name": "Botswanan pulas",
    "decimal_digits": 2,
    "code": "BWP"
  },
  {
    "name": "Belarusian Ruble",
    "symbol": "BYR",
    "native_symbol": "BYR",
    "plural_name": "Belarusian rubles",
    "decimal_digits": 0,
    "code": "BYR"
  },
  {
    "name": "Belize Dollar",
    "symbol": "BZ$",
    "native_symbol": "$",
    "plural_name": "Belize dollars",
    "decimal_digits": 2,
    "code": "BZD"
  },
  {
    "name": "Congolese Franc",
    "symbol": "CDF",
    "native_symbol": "FrCD",
    "plural_name": "Congolese francs",
    "decimal_digits": 2,
    "code": "CDF"
  },
  {
    "name": "Swiss Franc",
    "symbol": "CHF",
    "native_symbol": "CHF",
    "plural_name": "Swiss francs",
    "decimal_digits": 2,
    "code": "CHF"
  },
  {
    "name": "Chilean Peso",
    "symbol": "CL$",
    "native_symbol": "$",
    "plural_name": "Chilean pesos",
    "decimal_digits": 0,
    "code": "CLP"
  },
  {
    "name": "Chinese Yuan",
    "symbol": "CN¥",
    "native_symbol": "CN¥",
    "plural_name": "Chinese yuan",
    "decimal_digits": 2,
    "code": "CNY"
  },
  {
    "name": "Colombian Peso",
    "symbol": "CO$",
    "native_symbol": "$",
    "plural_name": "Colombian pesos",
    "decimal_digits": 0,
    "code": "COP"
  },
  {
    "name": "Costa Rican Colón",
    "symbol": "₡",
    "native_symbol": "₡",
    "plural_name": "Costa Rican colóns",
    "decimal_digits": 0,
    "code": "CRC"
  },
  {
    "name": "Cape Verdean Escudo",
    "symbol": "CV$",
    "native_symbol": "CV$",
    "plural_name": "Cape Verdean escudos",
    "decimal_digits": 2,
    "code": "CVE"
  },
  {
    "name": "Czech Republic Koruna",
    "symbol": "Kč",
    "native_symbol": "Kč",
    "plural_name": "Czech Republic korunas",
    "decimal_digits": 2,
    "code": "CZK"
  },
  {
    "name": "Djiboutian Franc",
    "symbol": "Fdj",
    "native_symbol": "Fdj",
    "plural_name": "Djiboutian francs",
    "decimal_digits": 0,
    "code": "DJF"
  },
  {
    "name": "Danish Krone",
    "symbol": "Dkr",
    "native_symbol": "kr",
    "plural_name": "Danish kroner",
    "decimal_digits": 2,
    "code": "DKK"
  },
  {
    "name": "Dominican Peso",
    "symbol": "RD$",
    "native_symbol": "RD$",
    "plural_name": "Dominican pesos",
    "decimal_digits": 2,
    "code": "DOP"
  },
  {
    "name": "Algerian Dinar",
    "symbol": "DA",
    "native_symbol": "د.ج.‏",
    "plural_name": "Algerian dinars",
    "decimal_digits": 2,
    "code": "DZD"
  },
  {
    "name": "Estonian Kroon",
    "symbol": "Ekr",
    "native_symbol": "kr",
    "plural_name": "Estonian kroons",
    "decimal_digits": 2,
    "code": "EEK"
  },
  {
    "name": "Egyptian Pound",
    "symbol": "EGP",
    "native_symbol": "ج.م.‏",
    "plural_name": "Egyptian pounds",
    "decimal_digits": 2,
    "code": "EGP"
  },
  {
    "name": "Eritrean Nakfa",
    "symbol": "Nfk",
    "native_symbol": "Nfk",
    "plural_name": "Eritrean nakfas",
    "decimal_digits": 2,
    "code": "ERN"
  },
  {
    "name": "Ethiopian Birr",
    "symbol": "Br",
    "native_symbol": "Br",
    "plural_name": "Ethiopian birrs",
    "decimal_digits": 2,
    "code": "ETB"
  },
  {
    "name": "British Pound Sterling",
    "symbol": "£",
    "native_symbol": "£",
    "plural_name": "British pounds sterling",
    "decimal_digits": 2,
    "code": "GBP"
  },
  {
    "name": "Georgian Lari",
    "symbol": "GEL",
    "native_symbol": "GEL",
    "plural_name": "Georgian laris",
    "decimal_digits": 2,
    "code": "GEL"
  },
  {
    "name": "Ghanaian Cedi",
    "symbol": "GH₵",
    "native_symbol": "GH₵",
    "plural_name": "Ghanaian cedis",
    "decimal_digits": 2,
    "code": "GHS"
  },
  {
    "name": "Guinean Franc",
    "symbol": "FG",
    "native_symbol": "FG",
    "plural_name": "Guinean francs",
    "decimal_digits": 0,
    "code": "GNF"
  },
  {
    "name": "Guatemalan Quetzal",
    "symbol": "GTQ",
    "native_symbol": "Q",
    "plural_name": "Guatemalan quetzals",
    "decimal_digits": 2,
    "code": "GTQ"
  },
  {
    "name": "Hong Kong Dollar",
    "symbol": "HK$",
    "native_symbol": "$",
    "plural_name": "Hong Kong dollars",
    "decimal_digits": 2,
    "code": "HKD"
  },
  {
    "name": "Honduran Lempira",
    "symbol": "HNL",
    "native_symbol": "L",
    "plural_name": "Honduran lempiras",
    "decimal_digits": 2,
    "code": "HNL"
  },
  {
    "name": "Croatian Kuna",
    "symbol": "kn",
    "native_symbol": "kn",
    "plural_name": "Croatian kunas",
    "decimal_digits": 2,
    "code": "HRK"
  },
  {
    "name": "Hungarian Forint",
    "symbol": "Ft",
    "native_symbol": "Ft",
    "plural_name": "Hungarian forints",
    "decimal_digits": 0,
    "code": "HUF"
  },
  {
    "name": "Indonesian Rupiah",
    "symbol": "Rp",
    "native_symbol": "Rp",
    "plural_name": "Indonesian rupiahs",
    "decimal_digits": 0,
    "code": "IDR"
  },
  {
    "name": "Israeli New Sheqel",
    "symbol": "₪",
    "native_symbol": "₪",
    "plural_name": "Israeli new sheqels",
    "decimal_digits": 2,
    "code": "ILS"
  },
  {
    "name": "Indian Rupee",
    "symbol": "Rs",
    "native_symbol": "টকা",
    "plural_name": "Indian rupees",
    "decimal_digits": 2,
    "code": "INR"
  },
  {
    "name": "Iraqi Dinar",
    "symbol": "IQD",
    "native_symbol": "د.ع.‏",
    "plural_name": "Iraqi dinars",
    "decimal_digits": 0,
    "code": "IQD"
  },
  {
    "name": "Iranian Rial",
    "symbol": "IRR",
    "native_symbol": "﷼",
    "plural_name": "Iranian rials",
    "decimal_digits": 0,
    "code": "IRR"
  },
  {
    "name": "Icelandic Króna",
    "symbol": "Ikr",
    "native_symbol": "kr",
    "plural_name": "Icelandic krónur",
    "decimal_digits": 0,
    "code": "ISK"
  },
  {
    "name": "Jamaican Dollar",
    "symbol": "J$",
    "native_symbol": "$",
    "plural_name": "Jamaican dollars",
    "decimal_digits": 2,
    "code": "JMD"
  },
  {
    "name": "Jordanian Dinar",
    "symbol": "JD",
    "native_symbol": "د.أ.‏",
    "plural_name": "Jordanian dinars",
    "decimal_digits": 3,
    "code": "JOD"
  },
  {
    "name": "Japanese Yen",
    "symbol": "¥",
    "native_symbol": "￥",
    "plural_name": "Japanese yen",
    "decimal_digits": 0,
    "code": "JPY"
  },
  {
    "name": "Kenyan Shilling",
    "symbol": "Ksh",
    "native_symbol": "Ksh",
    "plural_name": "Kenyan shillings",
    "decimal_digits": 2,
    "code": "KES"
  },
  {
    "name": "Cambodian Riel",
    "symbol": "KHR",
    "native_symbol": "៛",
    "plural_name": "Cambodian riels",
    "decimal_digits": 2,
    "code": "KHR"
  },
  {
    "name": "Comorian Franc",
    "symbol": "CF",
    "native_symbol": "FC",
    "plural_name": "Comorian francs",
    "decimal_digits": 0,
    "code": "KMF"
  },
  {
    "name": "South Korean Won",
    "symbol": "₩",
    "native_symbol": "₩",
    "plural_name": "South Korean won",
    "decimal_digits": 0,
    "code": "KRW"
  },
  {
    "name": "Kuwaiti Dinar",
    "symbol": "KD",
    "native_symbol": "د.ك.‏",
    "plural_name": "Kuwaiti dinars",
    "decimal_digits": 3,
    "code": "KWD"
  },
  {
    "name": "Kazakhstani Tenge",
    "symbol": "KZT",
    "native_symbol": "тңг.",
    "plural_name": "Kazakhstani tenges",
    "decimal_digits": 2,
    "code": "KZT"
  },
  {
    "name": "Lebanese Pound",
    "symbol": "LB£",
    "native_symbol": "ل.ل.‏",
    "plural_name": "Lebanese pounds",
    "decimal_digits": 0,
    "code": "LBP"
  },
  {
    "name": "Sri Lankan Rupee",
    "symbol": "SLRs",
    "native_symbol": "SL Re",
    "plural_name": "Sri Lankan rupees",
    "decimal_digits": 2,
    "code": "LKR"
  },
  {
    "name": "Lithuanian Litas",
    "symbol": "Lt",
    "native_symbol": "Lt",
    "plural_name": "Lithuanian litai",
    "decimal_digits": 2,
    "code": "LTL"
  },
  {
    "name": "Latvian Lats",
    "symbol": "Ls",
    "native_symbol": "Ls",
    "plural_name": "Latvian lati",
    "decimal_digits": 2,
    "code": "LVL"
  },
  {
    "name": "Libyan Dinar",
    "symbol": "LD",
    "native_symbol": "د.ل.‏",
    "plural_name": "Libyan dinars",
    "decimal_digits": 3,
    "code": "LYD"
  },
  {
    "name": "Moroccan Dirham",
    "symbol": "MAD",
    "native_symbol": "د.م.‏",
    "plural_name": "Moroccan dirhams",
    "decimal_digits": 2,
    "code": "MAD"
  },
  {
    "name": "Moldovan Leu",
    "symbol": "MDL",
    "native_symbol": "MDL",
    "plural_name": "Moldovan lei",
    "decimal_digits": 2,
    "code": "MDL"
  },
  {
    "name": "Malagasy Ariary",
    "symbol": "MGA",
    "native_symbol": "MGA",
    "plural_name": "Malagasy Ariaries",
    "decimal_digits": 0,
    "code": "MGA"
  },
  {
    "name": "Macedonian Denar",
    "symbol": "MKD",
    "native_symbol": "MKD",
    "plural_name": "Macedonian denari",
    "decimal_digits": 2,
    "code": "MKD"
  },
  {
    "name": "Myanma Kyat",
    "symbol": "MMK",
    "native_symbol": "K",
    "plural_name": "Myanma kyats",
    "decimal_digits": 0,
    "code": "MMK"
  },
  {
    "name": "Macanese Pataca",
    "symbol": "MOP$",
    "native_symbol": "MOP$",
    "plural_name": "Macanese patacas",
    "decimal_digits": 2,
    "code": "MOP"
  },
  {
    "name": "Mauritian Rupee",
    "symbol": "MURs",
    "native_symbol": "MURs",
    "plural_name": "Mauritian rupees",
    "decimal_digits": 0,
    "code": "MUR"
  },
  {
    "name": "Mexican Peso",
    "symbol": "MX$",
    "native_symbol": "$",
    "plural_name": "Mexican pesos",
    "decimal_digits": 2,
    "code": "MXN"
  },
  {
    "name": "Malaysian Ringgit",
    "symbol": "RM",
    "native_symbol": "RM",
    "plural_name": "Malaysian ringgits",
    "decimal_digits": 2,
    "code": "MYR"
  },
  {
    "name": "Mozambican Metical",
    "symbol": "MTn",
    "native_symbol": "MTn",
    "plural_name": "Mozambican meticals",
    "decimal_digits": 2,
    "code": "MZN"
  },
  {
    "name": "Namibian Dollar",
    "symbol": "N$",
    "native_symbol": "N$",
    "plural_name": "Namibian dollars",
    "decimal_digits": 2,
    "code": "NAD"
  },
  {
    "name": "Nigerian Naira",
    "symbol": "₦",
    "native_symbol": "₦",
    "plural_name": "Nigerian nairas",
    "decimal_digits": 2,
    "code": "NGN"
  },
  {
    "name": "Nicaraguan Córdoba",
    "symbol": "C$",
    "native_symbol": "C$",
    "plural_name": "Nicaraguan córdobas",
    "decimal_digits": 2,
    "code": "NIO"
  },
  {
    "name": "Norwegian Krone",
    "symbol": "Nkr",
    "native_symbol": "kr",
    "plural_name": "Norwegian kroner",
    "decimal_digits": 2,
    "code": "NOK"
  },
  {
    "name": "Nepalese Rupee",
    "symbol": "NPRs",
    "native_symbol": "नेरू",
    "plural_name": "Nepalese rupees",
    "decimal_digits": 2,
    "code": "NPR"
  },
  {
    "name": "New Zealand Dollar",
    "symbol": "NZ$",
    "native_symbol": "$",
    "plural_name": "New Zealand dollars",
    "decimal_digits": 2,
    "code": "NZD"
  },
  {
    "name": "Omani Rial",
    "symbol": "OMR",
    "native_symbol": "ر.ع.‏",
    "plural_name": "Omani rials",
    "decimal_digits": 3,
    "code": "OMR"
  },
  {
    "name": "Panamanian Balboa",
    "symbol": "B/.",
    "native_symbol": "B/.",
    "plural_name": "Panamanian balboas",
    "decimal_digits": 2,
    "code": "PAB"
  },
  {
    "name": "Peruvian Nuevo Sol",
    "symbol": "S/.",
    "native_symbol": "S/.",
    "plural_name": "Peruvian nuevos soles",
    "decimal_digits": 2,
    "code": "PEN"
  },
  {
    "name": "Philippine Peso",
    "symbol": "₱",
    "native_symbol": "₱",
    "plural_name": "Philippine pesos",
    "decimal_digits": 2,
    "code": "PHP"
  },
  {
    "name": "Pakistani Rupee",
    "symbol": "PKRs",
    "native_symbol": "₨",
    "plural_name": "Pakistani rupees",
    "decimal_digits": 0,
    "code": "PKR"
  },
  {
    "name": "Polish Zloty",
    "symbol": "zł",
    "native_symbol": "zł",
    "plural_name": "Polish zlotys",
    "decimal_digits": 2,
    "code": "PLN"
  },
  {
    "name": "Paraguayan Guarani",
    "symbol": "₲",
    "native_symbol": "₲",
    "plural_name": "Paraguayan guaranis",
    "decimal_digits": 0,
    "code": "PYG"
  },
  {
    "name": "Qatari Rial",
    "symbol": "QR",
    "native_symbol": "ر.ق.‏",
    "plural_name": "Qatari rials",
    "decimal_digits": 2,
    "code": "QAR"
  },
  {
    "name": "Romanian Leu",
    "symbol": "RON",
    "native_symbol": "RON",
    "plural_name": "Romanian lei",
    "decimal_digits": 2,
    "code": "RON"
  },
  {
    "name": "Serbian Dinar",
    "symbol": "din.",
    "native_symbol": "дин.",
    "plural_name": "Serbian dinars",
    "decimal_digits": 0,
    "code": "RSD"
  },
  {
    "name": "Russian Ruble",
    "symbol": "RUB",
    "native_symbol": "руб.",
    "plural_name": "Russian rubles",
    "decimal_digits": 2,
    "code": "RUB"
  },
  {
    "name": "Rwandan Franc",
    "symbol": "RWF",
    "native_symbol": "FR",
    "plural_name": "Rwandan francs",
    "decimal_digits": 0,
    "code": "RWF"
  },
  {
    "name": "Saudi Riyal",
    "symbol": "SR",
    "native_symbol": "ر.س.‏",
    "plural_name": "Saudi riyals",
    "decimal_digits": 2,
    "code": "SAR"
  },
  {
    "name": "Sudanese Pound",
    "symbol": "SDG",
    "native_symbol": "SDG",
    "plural_name": "Sudanese pounds",
    "decimal_digits": 2,
    "code": "SDG"
  },
  {
    "name": "Swedish Krona",
    "symbol": "Skr",
    "native_symbol": "kr",
    "plural_name": "Swedish kronor",
    "decimal_digits": 2,
    "code": "SEK"
  },
  {
    "name": "Singapore Dollar",
    "symbol": "S$",
    "native_symbol": "$",
    "plural_name": "Singapore dollars",
    "decimal_digits": 2,
    "code": "SGD"
  },
  {
    "name": "Somali Shilling",
    "symbol": "Ssh",
    "native_symbol": "Ssh",
    "plural_name": "Somali shillings",
    "decimal_digits": 0,
    "code": "SOS"
  },
  {
    "name": "Syrian Pound",
    "symbol": "SY£",
    "native_symbol": "ل.س.‏",
    "plural_name": "Syrian pounds",
    "decimal_digits": 0,
    "code": "SYP"
  },
  {
    "name": "Thai Baht",
    "symbol": "฿",
    "native_symbol": "฿",
    "plural_name": "Thai baht",
    "decimal_digits": 2,
    "code": "THB"
  },
  {
    "name": "Tunisian Dinar",
    "symbol": "DT",
    "native_symbol": "د.ت.‏",
    "plural_name": "Tunisian dinars",
    "decimal_digits": 3,
    "code": "TND"
  },
  {
    "name": "Tongan Paʻanga",
    "symbol": "T$",
    "native_symbol": "T$",
    "plural_name": "Tongan paʻanga",
    "decimal_digits": 2,
    "code": "TOP"
  },
  {
    "name": "Turkish Lira",
    "symbol": "TL",
    "native_symbol": "TL",
    "plural_name": "Turkish Lira",
    "decimal_digits": 2,
    "code": "TRY"
  },
  {
    "name": "Trinidad and Tobago Dollar",
    "symbol": "TT$",
    "native_symbol": "$",
    "plural_name": "Trinidad and Tobago dollars",
    "decimal_digits": 2,
    "code": "TTD"
  },
  {
    "name": "New Taiwan Dollar",
    "symbol": "NT$",
    "native_symbol": "NT$",
    "plural_name": "New Taiwan dollars",
    "decimal_digits": 2,
    "code": "TWD"
  },
  {
    "name": "Tanzanian Shilling",
    "symbol": "TSh",
    "native_symbol": "TSh",
    "plural_name": "Tanzanian shillings",
    "decimal_digits": 0,
    "code": "TZS"
  },
  {
    "name": "Ukrainian Hryvnia",
    "symbol": "₴",
    "native_symbol": "₴",
    "plural_name": "Ukrainian hryvnias",
    "decimal_digits": 2,
    "code": "UAH"
  },
  {
    "name": "Ugandan Shilling",
    "symbol": "USh",
    "native_symbol": "USh",
    "plural_name": "Ugandan shillings",
    "decimal_digits": 0,
    "code": "UGX"
  },
  {
    "name": "Uruguayan Peso",
    "symbol": "$U",
    "native_symbol": "$",
    "plural_name": "Uruguayan pesos",
    "decimal_digits": 2,
    "code": "UYU"
  },
  {
    "name": "Uzbekistan Som",
    "symbol": "UZS",
    "native_symbol": "UZS",
    "plural_name": "Uzbekistan som",
    "decimal_digits": 0,
    "code": "UZS"
  },
  {
    "name": "Venezuelan Bolívar",
    "symbol": "Bs.F.",
    "native_symbol": "Bs.F.",
    "plural_name": "Venezuelan bolívars",
    "decimal_digits": 2,
    "code": "VEF"
  },
  {
    "name": "Vietnamese Dong",
    "symbol": "₫",
    "native_symbol": "₫",
    "plural_name": "Vietnamese dong",
    "decimal_digits": 0,
    "code": "VND"
  },
  {
    "name": "CFA Franc BEAC",
    "symbol": "FCFA",
    "native_symbol": "FCFA",
    "plural_name": "CFA francs BEAC",
    "decimal_digits": 0,
    "code": "XAF"
  },
  {
    "name": "CFA Franc BCEAO",
    "symbol": "CFA",
    "native_symbol": "CFA",
    "plural_name": "CFA francs BCEAO",
    "decimal_digits": 0,
    "code": "XOF"
  },
  {
    "name": "Yemeni Rial",
    "symbol": "YR",
    "native_symbol": "ر.ي.‏",
    "plural_name": "Yemeni rials",
    "decimal_digits": 0,
    "code": "YER"
  },
  {
    "name": "South African Rand",
    "symbol": "R",
    "native_symbol": "R",
    "plural_name": "South African rand",
    "decimal_digits": 2,
    "code": "ZAR"
  },
  {
    "name": "Zambian Kwacha",
    "symbol": "ZK",
    "native_symbol": "ZK",
    "plural_name": "Zambian kwachas",
    "decimal_digits": 0,
    "code": "ZMK"
  }
]`
