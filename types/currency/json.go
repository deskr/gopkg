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
    "Name": "US Dollar",
    "Symbol": "$",
    "NativeSymbol": "$",
    "PluralName": "US dollars",
    "DecimalDigits": 2,
    "Code": "USD"
  },
  {
    "Name": "Canadian Dollar",
    "Symbol": "CA$",
    "NativeSymbol": "$",
    "PluralName": "Canadian dollars",
    "DecimalDigits": 2,
    "Code": "CAD"
  },
  {
    "Name": "Euro",
    "Symbol": "€",
    "NativeSymbol": "€",
    "PluralName": "euros",
    "DecimalDigits": 2,
    "Code": "EUR"
  },
  {
    "Name": "United Arab Emirates Dirham",
    "Symbol": "AED",
    "NativeSymbol": "د.إ.‏",
    "PluralName": "UAE dirhams",
    "DecimalDigits": 2,
    "Code": "AED"
  },
  {
    "Name": "Afghan Afghani",
    "Symbol": "Af",
    "NativeSymbol": "؋",
    "PluralName": "Afghan Afghanis",
    "DecimalDigits": 0,
    "Code": "AFN"
  },
  {
    "Name": "Albanian Lek",
    "Symbol": "ALL",
    "NativeSymbol": "Lek",
    "PluralName": "Albanian lekë",
    "DecimalDigits": 0,
    "Code": "ALL"
  },
  {
    "Name": "Armenian Dram",
    "Symbol": "AMD",
    "NativeSymbol": "դր.",
    "PluralName": "Armenian drams",
    "DecimalDigits": 0,
    "Code": "AMD"
  },
  {
    "Name": "Argentine Peso",
    "Symbol": "AR$",
    "NativeSymbol": "$",
    "PluralName": "Argentine pesos",
    "DecimalDigits": 2,
    "Code": "ARS"
  },
  {
    "Name": "Australian Dollar",
    "Symbol": "AU$",
    "NativeSymbol": "$",
    "PluralName": "Australian dollars",
    "DecimalDigits": 2,
    "Code": "AUD"
  },
  {
    "Name": "Azerbaijani Manat",
    "Symbol": "man.",
    "NativeSymbol": "ман.",
    "PluralName": "Azerbaijani manats",
    "DecimalDigits": 2,
    "Code": "AZN"
  },
  {
    "Name": "Bosnia-Herzegovina Convertible Mark",
    "Symbol": "KM",
    "NativeSymbol": "KM",
    "PluralName": "Bosnia-Herzegovina convertible marks",
    "DecimalDigits": 2,
    "Code": "BAM"
  },
  {
    "Name": "Bangladeshi Taka",
    "Symbol": "Tk",
    "NativeSymbol": "৳",
    "PluralName": "Bangladeshi takas",
    "DecimalDigits": 2,
    "Code": "BDT"
  },
  {
    "Name": "Bulgarian Lev",
    "Symbol": "BGN",
    "NativeSymbol": "лв.",
    "PluralName": "Bulgarian leva",
    "DecimalDigits": 2,
    "Code": "BGN"
  },
  {
    "Name": "Bahraini Dinar",
    "Symbol": "BD",
    "NativeSymbol": "د.ب.‏",
    "PluralName": "Bahraini dinars",
    "DecimalDigits": 3,
    "Code": "BHD"
  },
  {
    "Name": "Burundian Franc",
    "Symbol": "FBu",
    "NativeSymbol": "FBu",
    "PluralName": "Burundian francs",
    "DecimalDigits": 0,
    "Code": "BIF"
  },
  {
    "Name": "Brunei Dollar",
    "Symbol": "BN$",
    "NativeSymbol": "$",
    "PluralName": "Brunei dollars",
    "DecimalDigits": 2,
    "Code": "BND"
  },
  {
    "Name": "Bolivian Boliviano",
    "Symbol": "Bs",
    "NativeSymbol": "Bs",
    "PluralName": "Bolivian bolivianos",
    "DecimalDigits": 2,
    "Code": "BOB"
  },
  {
    "Name": "Brazilian Real",
    "Symbol": "R$",
    "NativeSymbol": "R$",
    "PluralName": "Brazilian reals",
    "DecimalDigits": 2,
    "Code": "BRL"
  },
  {
    "Name": "Botswanan Pula",
    "Symbol": "BWP",
    "NativeSymbol": "P",
    "PluralName": "Botswanan pulas",
    "DecimalDigits": 2,
    "Code": "BWP"
  },
  {
    "Name": "Belarusian Ruble",
    "Symbol": "BYR",
    "NativeSymbol": "BYR",
    "PluralName": "Belarusian rubles",
    "DecimalDigits": 0,
    "Code": "BYR"
  },
  {
    "Name": "Belize Dollar",
    "Symbol": "BZ$",
    "NativeSymbol": "$",
    "PluralName": "Belize dollars",
    "DecimalDigits": 2,
    "Code": "BZD"
  },
  {
    "Name": "Congolese Franc",
    "Symbol": "CDF",
    "NativeSymbol": "FrCD",
    "PluralName": "Congolese francs",
    "DecimalDigits": 2,
    "Code": "CDF"
  },
  {
    "Name": "Swiss Franc",
    "Symbol": "CHF",
    "NativeSymbol": "CHF",
    "PluralName": "Swiss francs",
    "DecimalDigits": 2,
    "Code": "CHF"
  },
  {
    "Name": "Chilean Peso",
    "Symbol": "CL$",
    "NativeSymbol": "$",
    "PluralName": "Chilean pesos",
    "DecimalDigits": 0,
    "Code": "CLP"
  },
  {
    "Name": "Chinese Yuan",
    "Symbol": "CN¥",
    "NativeSymbol": "CN¥",
    "PluralName": "Chinese yuan",
    "DecimalDigits": 2,
    "Code": "CNY"
  },
  {
    "Name": "Colombian Peso",
    "Symbol": "CO$",
    "NativeSymbol": "$",
    "PluralName": "Colombian pesos",
    "DecimalDigits": 0,
    "Code": "COP"
  },
  {
    "Name": "Costa Rican Colón",
    "Symbol": "₡",
    "NativeSymbol": "₡",
    "PluralName": "Costa Rican colóns",
    "DecimalDigits": 0,
    "Code": "CRC"
  },
  {
    "Name": "Cape Verdean Escudo",
    "Symbol": "CV$",
    "NativeSymbol": "CV$",
    "PluralName": "Cape Verdean escudos",
    "DecimalDigits": 2,
    "Code": "CVE"
  },
  {
    "Name": "Czech Republic Koruna",
    "Symbol": "Kč",
    "NativeSymbol": "Kč",
    "PluralName": "Czech Republic korunas",
    "DecimalDigits": 2,
    "Code": "CZK"
  },
  {
    "Name": "Djiboutian Franc",
    "Symbol": "Fdj",
    "NativeSymbol": "Fdj",
    "PluralName": "Djiboutian francs",
    "DecimalDigits": 0,
    "Code": "DJF"
  },
  {
    "Name": "Danish Krone",
    "Symbol": "Dkr",
    "NativeSymbol": "kr",
    "PluralName": "Danish kroner",
    "DecimalDigits": 2,
    "Code": "DKK"
  },
  {
    "Name": "Dominican Peso",
    "Symbol": "RD$",
    "NativeSymbol": "RD$",
    "PluralName": "Dominican pesos",
    "DecimalDigits": 2,
    "Code": "DOP"
  },
  {
    "Name": "Algerian Dinar",
    "Symbol": "DA",
    "NativeSymbol": "د.ج.‏",
    "PluralName": "Algerian dinars",
    "DecimalDigits": 2,
    "Code": "DZD"
  },
  {
    "Name": "Estonian Kroon",
    "Symbol": "Ekr",
    "NativeSymbol": "kr",
    "PluralName": "Estonian kroons",
    "DecimalDigits": 2,
    "Code": "EEK"
  },
  {
    "Name": "Egyptian Pound",
    "Symbol": "EGP",
    "NativeSymbol": "ج.م.‏",
    "PluralName": "Egyptian pounds",
    "DecimalDigits": 2,
    "Code": "EGP"
  },
  {
    "Name": "Eritrean Nakfa",
    "Symbol": "Nfk",
    "NativeSymbol": "Nfk",
    "PluralName": "Eritrean nakfas",
    "DecimalDigits": 2,
    "Code": "ERN"
  },
  {
    "Name": "Ethiopian Birr",
    "Symbol": "Br",
    "NativeSymbol": "Br",
    "PluralName": "Ethiopian birrs",
    "DecimalDigits": 2,
    "Code": "ETB"
  },
  {
    "Name": "British Pound Sterling",
    "Symbol": "£",
    "NativeSymbol": "£",
    "PluralName": "British pounds sterling",
    "DecimalDigits": 2,
    "Code": "GBP"
  },
  {
    "Name": "Georgian Lari",
    "Symbol": "GEL",
    "NativeSymbol": "GEL",
    "PluralName": "Georgian laris",
    "DecimalDigits": 2,
    "Code": "GEL"
  },
  {
    "Name": "Ghanaian Cedi",
    "Symbol": "GH₵",
    "NativeSymbol": "GH₵",
    "PluralName": "Ghanaian cedis",
    "DecimalDigits": 2,
    "Code": "GHS"
  },
  {
    "Name": "Guinean Franc",
    "Symbol": "FG",
    "NativeSymbol": "FG",
    "PluralName": "Guinean francs",
    "DecimalDigits": 0,
    "Code": "GNF"
  },
  {
    "Name": "Guatemalan Quetzal",
    "Symbol": "GTQ",
    "NativeSymbol": "Q",
    "PluralName": "Guatemalan quetzals",
    "DecimalDigits": 2,
    "Code": "GTQ"
  },
  {
    "Name": "Hong Kong Dollar",
    "Symbol": "HK$",
    "NativeSymbol": "$",
    "PluralName": "Hong Kong dollars",
    "DecimalDigits": 2,
    "Code": "HKD"
  },
  {
    "Name": "Honduran Lempira",
    "Symbol": "HNL",
    "NativeSymbol": "L",
    "PluralName": "Honduran lempiras",
    "DecimalDigits": 2,
    "Code": "HNL"
  },
  {
    "Name": "Croatian Kuna",
    "Symbol": "kn",
    "NativeSymbol": "kn",
    "PluralName": "Croatian kunas",
    "DecimalDigits": 2,
    "Code": "HRK"
  },
  {
    "Name": "Hungarian Forint",
    "Symbol": "Ft",
    "NativeSymbol": "Ft",
    "PluralName": "Hungarian forints",
    "DecimalDigits": 0,
    "Code": "HUF"
  },
  {
    "Name": "Indonesian Rupiah",
    "Symbol": "Rp",
    "NativeSymbol": "Rp",
    "PluralName": "Indonesian rupiahs",
    "DecimalDigits": 0,
    "Code": "IDR"
  },
  {
    "Name": "Israeli New Sheqel",
    "Symbol": "₪",
    "NativeSymbol": "₪",
    "PluralName": "Israeli new sheqels",
    "DecimalDigits": 2,
    "Code": "ILS"
  },
  {
    "Name": "Indian Rupee",
    "Symbol": "Rs",
    "NativeSymbol": "টকা",
    "PluralName": "Indian rupees",
    "DecimalDigits": 2,
    "Code": "INR"
  },
  {
    "Name": "Iraqi Dinar",
    "Symbol": "IQD",
    "NativeSymbol": "د.ع.‏",
    "PluralName": "Iraqi dinars",
    "DecimalDigits": 0,
    "Code": "IQD"
  },
  {
    "Name": "Iranian Rial",
    "Symbol": "IRR",
    "NativeSymbol": "﷼",
    "PluralName": "Iranian rials",
    "DecimalDigits": 0,
    "Code": "IRR"
  },
  {
    "Name": "Icelandic Króna",
    "Symbol": "Ikr",
    "NativeSymbol": "kr",
    "PluralName": "Icelandic krónur",
    "DecimalDigits": 0,
    "Code": "ISK"
  },
  {
    "Name": "Jamaican Dollar",
    "Symbol": "J$",
    "NativeSymbol": "$",
    "PluralName": "Jamaican dollars",
    "DecimalDigits": 2,
    "Code": "JMD"
  },
  {
    "Name": "Jordanian Dinar",
    "Symbol": "JD",
    "NativeSymbol": "د.أ.‏",
    "PluralName": "Jordanian dinars",
    "DecimalDigits": 3,
    "Code": "JOD"
  },
  {
    "Name": "Japanese Yen",
    "Symbol": "¥",
    "NativeSymbol": "￥",
    "PluralName": "Japanese yen",
    "DecimalDigits": 0,
    "Code": "JPY"
  },
  {
    "Name": "Kenyan Shilling",
    "Symbol": "Ksh",
    "NativeSymbol": "Ksh",
    "PluralName": "Kenyan shillings",
    "DecimalDigits": 2,
    "Code": "KES"
  },
  {
    "Name": "Cambodian Riel",
    "Symbol": "KHR",
    "NativeSymbol": "៛",
    "PluralName": "Cambodian riels",
    "DecimalDigits": 2,
    "Code": "KHR"
  },
  {
    "Name": "Comorian Franc",
    "Symbol": "CF",
    "NativeSymbol": "FC",
    "PluralName": "Comorian francs",
    "DecimalDigits": 0,
    "Code": "KMF"
  },
  {
    "Name": "South Korean Won",
    "Symbol": "₩",
    "NativeSymbol": "₩",
    "PluralName": "South Korean won",
    "DecimalDigits": 0,
    "Code": "KRW"
  },
  {
    "Name": "Kuwaiti Dinar",
    "Symbol": "KD",
    "NativeSymbol": "د.ك.‏",
    "PluralName": "Kuwaiti dinars",
    "DecimalDigits": 3,
    "Code": "KWD"
  },
  {
    "Name": "Kazakhstani Tenge",
    "Symbol": "KZT",
    "NativeSymbol": "тңг.",
    "PluralName": "Kazakhstani tenges",
    "DecimalDigits": 2,
    "Code": "KZT"
  },
  {
    "Name": "Lebanese Pound",
    "Symbol": "LB£",
    "NativeSymbol": "ل.ل.‏",
    "PluralName": "Lebanese pounds",
    "DecimalDigits": 0,
    "Code": "LBP"
  },
  {
    "Name": "Sri Lankan Rupee",
    "Symbol": "SLRs",
    "NativeSymbol": "SL Re",
    "PluralName": "Sri Lankan rupees",
    "DecimalDigits": 2,
    "Code": "LKR"
  },
  {
    "Name": "Lithuanian Litas",
    "Symbol": "Lt",
    "NativeSymbol": "Lt",
    "PluralName": "Lithuanian litai",
    "DecimalDigits": 2,
    "Code": "LTL"
  },
  {
    "Name": "Latvian Lats",
    "Symbol": "Ls",
    "NativeSymbol": "Ls",
    "PluralName": "Latvian lati",
    "DecimalDigits": 2,
    "Code": "LVL"
  },
  {
    "Name": "Libyan Dinar",
    "Symbol": "LD",
    "NativeSymbol": "د.ل.‏",
    "PluralName": "Libyan dinars",
    "DecimalDigits": 3,
    "Code": "LYD"
  },
  {
    "Name": "Moroccan Dirham",
    "Symbol": "MAD",
    "NativeSymbol": "د.م.‏",
    "PluralName": "Moroccan dirhams",
    "DecimalDigits": 2,
    "Code": "MAD"
  },
  {
    "Name": "Moldovan Leu",
    "Symbol": "MDL",
    "NativeSymbol": "MDL",
    "PluralName": "Moldovan lei",
    "DecimalDigits": 2,
    "Code": "MDL"
  },
  {
    "Name": "Malagasy Ariary",
    "Symbol": "MGA",
    "NativeSymbol": "MGA",
    "PluralName": "Malagasy Ariaries",
    "DecimalDigits": 0,
    "Code": "MGA"
  },
  {
    "Name": "Macedonian Denar",
    "Symbol": "MKD",
    "NativeSymbol": "MKD",
    "PluralName": "Macedonian denari",
    "DecimalDigits": 2,
    "Code": "MKD"
  },
  {
    "Name": "Myanma Kyat",
    "Symbol": "MMK",
    "NativeSymbol": "K",
    "PluralName": "Myanma kyats",
    "DecimalDigits": 0,
    "Code": "MMK"
  },
  {
    "Name": "Macanese Pataca",
    "Symbol": "MOP$",
    "NativeSymbol": "MOP$",
    "PluralName": "Macanese patacas",
    "DecimalDigits": 2,
    "Code": "MOP"
  },
  {
    "Name": "Mauritian Rupee",
    "Symbol": "MURs",
    "NativeSymbol": "MURs",
    "PluralName": "Mauritian rupees",
    "DecimalDigits": 0,
    "Code": "MUR"
  },
  {
    "Name": "Mexican Peso",
    "Symbol": "MX$",
    "NativeSymbol": "$",
    "PluralName": "Mexican pesos",
    "DecimalDigits": 2,
    "Code": "MXN"
  },
  {
    "Name": "Malaysian Ringgit",
    "Symbol": "RM",
    "NativeSymbol": "RM",
    "PluralName": "Malaysian ringgits",
    "DecimalDigits": 2,
    "Code": "MYR"
  },
  {
    "Name": "Mozambican Metical",
    "Symbol": "MTn",
    "NativeSymbol": "MTn",
    "PluralName": "Mozambican meticals",
    "DecimalDigits": 2,
    "Code": "MZN"
  },
  {
    "Name": "Namibian Dollar",
    "Symbol": "N$",
    "NativeSymbol": "N$",
    "PluralName": "Namibian dollars",
    "DecimalDigits": 2,
    "Code": "NAD"
  },
  {
    "Name": "Nigerian Naira",
    "Symbol": "₦",
    "NativeSymbol": "₦",
    "PluralName": "Nigerian nairas",
    "DecimalDigits": 2,
    "Code": "NGN"
  },
  {
    "Name": "Nicaraguan Córdoba",
    "Symbol": "C$",
    "NativeSymbol": "C$",
    "PluralName": "Nicaraguan córdobas",
    "DecimalDigits": 2,
    "Code": "NIO"
  },
  {
    "Name": "Norwegian Krone",
    "Symbol": "Nkr",
    "NativeSymbol": "kr",
    "PluralName": "Norwegian kroner",
    "DecimalDigits": 2,
    "Code": "NOK"
  },
  {
    "Name": "Nepalese Rupee",
    "Symbol": "NPRs",
    "NativeSymbol": "नेरू",
    "PluralName": "Nepalese rupees",
    "DecimalDigits": 2,
    "Code": "NPR"
  },
  {
    "Name": "New Zealand Dollar",
    "Symbol": "NZ$",
    "NativeSymbol": "$",
    "PluralName": "New Zealand dollars",
    "DecimalDigits": 2,
    "Code": "NZD"
  },
  {
    "Name": "Omani Rial",
    "Symbol": "OMR",
    "NativeSymbol": "ر.ع.‏",
    "PluralName": "Omani rials",
    "DecimalDigits": 3,
    "Code": "OMR"
  },
  {
    "Name": "Panamanian Balboa",
    "Symbol": "B/.",
    "NativeSymbol": "B/.",
    "PluralName": "Panamanian balboas",
    "DecimalDigits": 2,
    "Code": "PAB"
  },
  {
    "Name": "Peruvian Nuevo Sol",
    "Symbol": "S/.",
    "NativeSymbol": "S/.",
    "PluralName": "Peruvian nuevos soles",
    "DecimalDigits": 2,
    "Code": "PEN"
  },
  {
    "Name": "Philippine Peso",
    "Symbol": "₱",
    "NativeSymbol": "₱",
    "PluralName": "Philippine pesos",
    "DecimalDigits": 2,
    "Code": "PHP"
  },
  {
    "Name": "Pakistani Rupee",
    "Symbol": "PKRs",
    "NativeSymbol": "₨",
    "PluralName": "Pakistani rupees",
    "DecimalDigits": 0,
    "Code": "PKR"
  },
  {
    "Name": "Polish Zloty",
    "Symbol": "zł",
    "NativeSymbol": "zł",
    "PluralName": "Polish zlotys",
    "DecimalDigits": 2,
    "Code": "PLN"
  },
  {
    "Name": "Paraguayan Guarani",
    "Symbol": "₲",
    "NativeSymbol": "₲",
    "PluralName": "Paraguayan guaranis",
    "DecimalDigits": 0,
    "Code": "PYG"
  },
  {
    "Name": "Qatari Rial",
    "Symbol": "QR",
    "NativeSymbol": "ر.ق.‏",
    "PluralName": "Qatari rials",
    "DecimalDigits": 2,
    "Code": "QAR"
  },
  {
    "Name": "Romanian Leu",
    "Symbol": "RON",
    "NativeSymbol": "RON",
    "PluralName": "Romanian lei",
    "DecimalDigits": 2,
    "Code": "RON"
  },
  {
    "Name": "Serbian Dinar",
    "Symbol": "din.",
    "NativeSymbol": "дин.",
    "PluralName": "Serbian dinars",
    "DecimalDigits": 0,
    "Code": "RSD"
  },
  {
    "Name": "Russian Ruble",
    "Symbol": "RUB",
    "NativeSymbol": "руб.",
    "PluralName": "Russian rubles",
    "DecimalDigits": 2,
    "Code": "RUB"
  },
  {
    "Name": "Rwandan Franc",
    "Symbol": "RWF",
    "NativeSymbol": "FR",
    "PluralName": "Rwandan francs",
    "DecimalDigits": 0,
    "Code": "RWF"
  },
  {
    "Name": "Saudi Riyal",
    "Symbol": "SR",
    "NativeSymbol": "ر.س.‏",
    "PluralName": "Saudi riyals",
    "DecimalDigits": 2,
    "Code": "SAR"
  },
  {
    "Name": "Sudanese Pound",
    "Symbol": "SDG",
    "NativeSymbol": "SDG",
    "PluralName": "Sudanese pounds",
    "DecimalDigits": 2,
    "Code": "SDG"
  },
  {
    "Name": "Swedish Krona",
    "Symbol": "Skr",
    "NativeSymbol": "kr",
    "PluralName": "Swedish kronor",
    "DecimalDigits": 2,
    "Code": "SEK"
  },
  {
    "Name": "Singapore Dollar",
    "Symbol": "S$",
    "NativeSymbol": "$",
    "PluralName": "Singapore dollars",
    "DecimalDigits": 2,
    "Code": "SGD"
  },
  {
    "Name": "Somali Shilling",
    "Symbol": "Ssh",
    "NativeSymbol": "Ssh",
    "PluralName": "Somali shillings",
    "DecimalDigits": 0,
    "Code": "SOS"
  },
  {
    "Name": "Syrian Pound",
    "Symbol": "SY£",
    "NativeSymbol": "ل.س.‏",
    "PluralName": "Syrian pounds",
    "DecimalDigits": 0,
    "Code": "SYP"
  },
  {
    "Name": "Thai Baht",
    "Symbol": "฿",
    "NativeSymbol": "฿",
    "PluralName": "Thai baht",
    "DecimalDigits": 2,
    "Code": "THB"
  },
  {
    "Name": "Tunisian Dinar",
    "Symbol": "DT",
    "NativeSymbol": "د.ت.‏",
    "PluralName": "Tunisian dinars",
    "DecimalDigits": 3,
    "Code": "TND"
  },
  {
    "Name": "Tongan Paʻanga",
    "Symbol": "T$",
    "NativeSymbol": "T$",
    "PluralName": "Tongan paʻanga",
    "DecimalDigits": 2,
    "Code": "TOP"
  },
  {
    "Name": "Turkish Lira",
    "Symbol": "TL",
    "NativeSymbol": "TL",
    "PluralName": "Turkish Lira",
    "DecimalDigits": 2,
    "Code": "TRY"
  },
  {
    "Name": "Trinidad and Tobago Dollar",
    "Symbol": "TT$",
    "NativeSymbol": "$",
    "PluralName": "Trinidad and Tobago dollars",
    "DecimalDigits": 2,
    "Code": "TTD"
  },
  {
    "Name": "New Taiwan Dollar",
    "Symbol": "NT$",
    "NativeSymbol": "NT$",
    "PluralName": "New Taiwan dollars",
    "DecimalDigits": 2,
    "Code": "TWD"
  },
  {
    "Name": "Tanzanian Shilling",
    "Symbol": "TSh",
    "NativeSymbol": "TSh",
    "PluralName": "Tanzanian shillings",
    "DecimalDigits": 0,
    "Code": "TZS"
  },
  {
    "Name": "Ukrainian Hryvnia",
    "Symbol": "₴",
    "NativeSymbol": "₴",
    "PluralName": "Ukrainian hryvnias",
    "DecimalDigits": 2,
    "Code": "UAH"
  },
  {
    "Name": "Ugandan Shilling",
    "Symbol": "USh",
    "NativeSymbol": "USh",
    "PluralName": "Ugandan shillings",
    "DecimalDigits": 0,
    "Code": "UGX"
  },
  {
    "Name": "Uruguayan Peso",
    "Symbol": "$U",
    "NativeSymbol": "$",
    "PluralName": "Uruguayan pesos",
    "DecimalDigits": 2,
    "Code": "UYU"
  },
  {
    "Name": "Uzbekistan Som",
    "Symbol": "UZS",
    "NativeSymbol": "UZS",
    "PluralName": "Uzbekistan som",
    "DecimalDigits": 0,
    "Code": "UZS"
  },
  {
    "Name": "Venezuelan Bolívar",
    "Symbol": "Bs.F.",
    "NativeSymbol": "Bs.F.",
    "PluralName": "Venezuelan bolívars",
    "DecimalDigits": 2,
    "Code": "VEF"
  },
  {
    "Name": "Vietnamese Dong",
    "Symbol": "₫",
    "NativeSymbol": "₫",
    "PluralName": "Vietnamese dong",
    "DecimalDigits": 0,
    "Code": "VND"
  },
  {
    "Name": "CFA Franc BEAC",
    "Symbol": "FCFA",
    "NativeSymbol": "FCFA",
    "PluralName": "CFA francs BEAC",
    "DecimalDigits": 0,
    "Code": "XAF"
  },
  {
    "Name": "CFA Franc BCEAO",
    "Symbol": "CFA",
    "NativeSymbol": "CFA",
    "PluralName": "CFA francs BCEAO",
    "DecimalDigits": 0,
    "Code": "XOF"
  },
  {
    "Name": "Yemeni Rial",
    "Symbol": "YR",
    "NativeSymbol": "ر.ي.‏",
    "PluralName": "Yemeni rials",
    "DecimalDigits": 0,
    "Code": "YER"
  },
  {
    "Name": "South African Rand",
    "Symbol": "R",
    "NativeSymbol": "R",
    "PluralName": "South African rand",
    "DecimalDigits": 2,
    "Code": "ZAR"
  },
  {
    "Name": "Zambian Kwacha",
    "Symbol": "ZK",
    "NativeSymbol": "ZK",
    "PluralName": "Zambian kwachas",
    "DecimalDigits": 0,
    "Code": "ZMK"
  }
]`