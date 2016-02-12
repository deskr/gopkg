package language

import (
	"bytes"
	"encoding/json"
)

var languages []Language
var languagesByCode map[Code]Language

func init() {
	if err := json.NewDecoder(bytes.NewReader([]byte(languagesJSON))).Decode(&languages); err != nil {
		panic(err)
	}

	languagesByCode = make(map[Code]Language)
	for _, v := range languages {
		languagesByCode[v.Code] = v
	}
}

const languagesJSON = `[
  {
    "Code": "ab",
    "Name": "Abkhaz",
    "NativeName": "Аҧсуа"
  },
  {
    "Code": "aa",
    "Name": "Afar",
    "NativeName": "Afaraf"
  },
  {
    "Code": "af",
    "Name": "Afrikaans",
    "NativeName": "Afrikaans"
  },
  {
    "Code": "ak",
    "Name": "Akan",
    "NativeName": "Akan"
  },
  {
    "Code": "sq",
    "Name": "Albanian",
    "NativeName": "Shqip"
  },
  {
    "Code": "am",
    "Name": "Amharic",
    "NativeName": "አማርኛ"
  },
  {
    "Code": "ar",
    "Name": "Arabic",
    "NativeName": "العربية"
  },
  {
    "Code": "an",
    "Name": "Aragonese",
    "NativeName": "Aragonés"
  },
  {
    "Code": "hy",
    "Name": "Armenian",
    "NativeName": "Հայերեն"
  },
  {
    "Code": "as",
    "Name": "Assamese",
    "NativeName": "অসমীয়া"
  },
  {
    "Code": "av",
    "Name": "Avaric",
    "NativeName": "Авар МацӀ, МагӀарул МацӀ"
  },
  {
    "Code": "ae",
    "Name": "Avestan",
    "NativeName": "Avesta"
  },
  {
    "Code": "ay",
    "Name": "Aymara",
    "NativeName": "Aymar Aru"
  },
  {
    "Code": "az",
    "Name": "Azerbaijani",
    "NativeName": "Azərbaycan Dili"
  },
  {
    "Code": "bm",
    "Name": "Bambara",
    "NativeName": "Bamanankan"
  },
  {
    "Code": "ba",
    "Name": "Bashkir",
    "NativeName": "Башҡорт Теле"
  },
  {
    "Code": "eu",
    "Name": "Basque",
    "NativeName": "Euskara, Euskera"
  },
  {
    "Code": "be",
    "Name": "Belarusian",
    "NativeName": "Беларуская"
  },
  {
    "Code": "bn",
    "Name": "Bengali",
    "NativeName": "বাংলা"
  },
  {
    "Code": "bh",
    "Name": "Bihari",
    "NativeName": "भोजपुरी"
  },
  {
    "Code": "bi",
    "Name": "Bislama",
    "NativeName": "Bislama"
  },
  {
    "Code": "bs",
    "Name": "Bosnian",
    "NativeName": "Bosanski Jezik"
  },
  {
    "Code": "br",
    "Name": "Breton",
    "NativeName": "Brezhoneg"
  },
  {
    "Code": "bg",
    "Name": "Bulgarian",
    "NativeName": "Български Език"
  },
  {
    "Code": "my",
    "Name": "Burmese",
    "NativeName": "ဗမာစာ"
  },
  {
    "Code": "ca",
    "Name": "Catalan",
    "NativeName": "Català"
  },
  {
    "Code": "ch",
    "Name": "Chamorro",
    "NativeName": "Chamoru"
  },
  {
    "Code": "ce",
    "Name": "Chechen",
    "NativeName": "Нохчийн Мотт"
  },
  {
    "Code": "ny",
    "Name": "Chichewa",
    "NativeName": "ChiCheŵa"
  },
  {
    "Code": "zh",
    "Name": "Chinese",
    "NativeName": "中文 (Zhōngwén), 汉语, 漢語"
  },
  {
    "Code": "cv",
    "Name": "Chuvash",
    "NativeName": "Чӑваш Чӗлхи"
  },
  {
    "Code": "kw",
    "Name": "Cornish",
    "NativeName": "Kernewek"
  },
  {
    "Code": "co",
    "Name": "Corsican",
    "NativeName": "Corsu, Lingua Corsa"
  },
  {
    "Code": "cr",
    "Name": "Cree",
    "NativeName": "ᓀᐦᐃᔭᐍᐏᐣ"
  },
  {
    "Code": "hr",
    "Name": "Croatian",
    "NativeName": "Hrvatski"
  },
  {
    "Code": "cs",
    "Name": "Czech",
    "NativeName": "Česky, Čeština"
  },
  {
    "Code": "da",
    "Name": "Danish",
    "NativeName": "Dansk"
  },
  {
    "Code": "dv",
    "Name": "Maldivian",
    "NativeName": ""
  },
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
    "Code": "eo",
    "Name": "Esperanto",
    "NativeName": "Esperanto"
  },
  {
    "Code": "et",
    "Name": "Estonian",
    "NativeName": "Eesti, Eesti Keel"
  },
  {
    "Code": "ee",
    "Name": "Ewe",
    "NativeName": "Eʋegbe"
  },
  {
    "Code": "fo",
    "Name": "Faroese",
    "NativeName": "Føroyskt"
  },
  {
    "Code": "fj",
    "Name": "Fijian",
    "NativeName": "Vosa Vakaviti"
  },
  {
    "Code": "fi",
    "Name": "Finnish",
    "NativeName": "Suomi, Suomen Kieli"
  },
  {
    "Code": "fr",
    "Name": "French",
    "NativeName": "Français, Langue Française"
  },
  {
    "Code": "ff",
    "Name": "Pular",
    "NativeName": "Pular"
  },
  {
    "Code": "gl",
    "Name": "Galician",
    "NativeName": "Galego"
  },
  {
    "Code": "ka",
    "Name": "Georgian",
    "NativeName": "ქართული"
  },
  {
    "Code": "de",
    "Name": "German",
    "NativeName": "Deutsch"
  },
  {
    "Code": "el",
    "Name": "Greek, Modern",
    "NativeName": "Ελληνικά"
  },
  {
    "Code": "gn",
    "Name": "Guaraní",
    "NativeName": "Avañeẽ"
  },
  {
    "Code": "gu",
    "Name": "Gujarati",
    "NativeName": "ગુજરાતી"
  },
  {
    "Code": "ht",
    "Name": "Haitian",
    "NativeName": "Kreyòl Ayisyen"
  },
  {
    "Code": "ha",
    "Name": "Hausa",
    "NativeName": "Hausa, هَوُسَ"
  },
  {
    "Code": "he",
    "Name": "Hebrew (modern)",
    "NativeName": "עברית"
  },
  {
    "Code": "hz",
    "Name": "Herero",
    "NativeName": "Otjiherero"
  },
  {
    "Code": "hi",
    "Name": "Hindi",
    "NativeName": "हिन्दी, हिंदी"
  },
  {
    "Code": "ho",
    "Name": "Hiri Motu",
    "NativeName": "Hiri Motu"
  },
  {
    "Code": "hu",
    "Name": "Hungarian",
    "NativeName": "Magyar"
  },
  {
    "Code": "ia",
    "Name": "Interlingua",
    "NativeName": "Interlingua"
  },
  {
    "Code": "id",
    "Name": "Indonesian",
    "NativeName": "Bahasa Indonesia"
  },
  {
    "Code": "ie",
    "Name": "Interlingue",
    "NativeName": ""
  },
  {
    "Code": "ga",
    "Name": "Irish",
    "NativeName": "Gaeilge"
  },
  {
    "Code": "ig",
    "Name": "Igbo",
    "NativeName": "Asụsụ Igbo"
  },
  {
    "Code": "ik",
    "Name": "Inupiaq",
    "NativeName": "Iñupiaq, Iñupiatun"
  },
  {
    "Code": "io",
    "Name": "Ido",
    "NativeName": "Ido"
  },
  {
    "Code": "is",
    "Name": "Icelandic",
    "NativeName": "Íslenska"
  },
  {
    "Code": "it",
    "Name": "Italian",
    "NativeName": "Italiano"
  },
  {
    "Code": "iu",
    "Name": "Inuktitut",
    "NativeName": "ᐃᓄᒃᑎᑐᑦ"
  },
  {
    "Code": "ja",
    "Name": "Japanese",
    "NativeName": "日本語 (にほんご／にっぽんご)"
  },
  {
    "Code": "jv",
    "Name": "Javanese",
    "NativeName": "Basa Jawa"
  },
  {
    "Code": "kl",
    "Name": "Kalaallisut, Greenlandic",
    "NativeName": "Kalaallisut, Kalaallit Oqaasii"
  },
  {
    "Code": "kn",
    "Name": "Kannada",
    "NativeName": "ಕನ್ನಡ"
  },
  {
    "Code": "kr",
    "Name": "Kanuri",
    "NativeName": "Kanuri"
  },
  {
    "Code": "ks",
    "Name": "Kashmiri",
    "NativeName": "कश्मीरी, كشميري‎"
  },
  {
    "Code": "kk",
    "Name": "Kazakh",
    "NativeName": "Қазақ Тілі"
  },
  {
    "Code": "km",
    "Name": "Khmer",
    "NativeName": "ភាសាខ្មែរ"
  },
  {
    "Code": "ki",
    "Name": "Kikuyu, Gikuyu",
    "NativeName": "Gĩkũyũ"
  },
  {
    "Code": "rw",
    "Name": "Kinyarwanda",
    "NativeName": "Ikinyarwanda"
  },
  {
    "Code": "ky",
    "Name": "Kirghiz, Kyrgyz",
    "NativeName": "Кыргыз Тили"
  },
  {
    "Code": "kv",
    "Name": "Komi",
    "NativeName": "Коми Кыв"
  },
  {
    "Code": "kg",
    "Name": "Kongo",
    "NativeName": "KiKongo"
  },
  {
    "Code": "ko",
    "Name": "Korean",
    "NativeName": "한국어 (韓國語), 조선말 (朝鮮語)"
  },
  {
    "Code": "ku",
    "Name": "Kurdish",
    "NativeName": "Kurdî, كوردی‎"
  },
  {
    "Code": "kj",
    "Name": "Kwanyama, Kuanyama",
    "NativeName": "Kuanyama"
  },
  {
    "Code": "la",
    "Name": "Latin",
    "NativeName": "Latine, Lingua Latina"
  },
  {
    "Code": "lb",
    "Name": "Luxembourgish, Letzeburgesch",
    "NativeName": "Lëtzebuergesch"
  },
  {
    "Code": "lg",
    "Name": "Luganda",
    "NativeName": "Luganda"
  },
  {
    "Code": "li",
    "Name": "Limburgish, Limburgan, Limburger",
    "NativeName": "Limburgs"
  },
  {
    "Code": "ln",
    "Name": "Lingala",
    "NativeName": "Lingála"
  },
  {
    "Code": "lo",
    "Name": "Lao",
    "NativeName": "ພາສາລາວ"
  },
  {
    "Code": "lt",
    "Name": "Lithuanian",
    "NativeName": "Lietuvių Kalba"
  },
  {
    "Code": "lu",
    "Name": "Luba-Katanga",
    "NativeName": ""
  },
  {
    "Code": "lv",
    "Name": "Latvian",
    "NativeName": "Latviešu Valoda"
  },
  {
    "Code": "gv",
    "Name": "Manx",
    "NativeName": "Gaelg, Gailck"
  },
  {
    "Code": "mk",
    "Name": "Macedonian",
    "NativeName": "Македонски Јазик"
  },
  {
    "Code": "mg",
    "Name": "Malagasy",
    "NativeName": "Malagasy Fiteny"
  },
  {
    "Code": "ms",
    "Name": "Malay",
    "NativeName": "Bahasa Melayu, بهاس ملايو‎"
  },
  {
    "Code": "ml",
    "Name": "Malayalam",
    "NativeName": "മലയാളം"
  },
  {
    "Code": "mt",
    "Name": "Maltese",
    "NativeName": "Malti"
  },
  {
    "Code": "mi",
    "Name": "Māori",
    "NativeName": "Te Reo Māori"
  },
  {
    "Code": "mr",
    "Name": "Marathi (Marāṭhī)",
    "NativeName": "मराठी"
  },
  {
    "Code": "mh",
    "Name": "Marshallese",
    "NativeName": "Kajin M̧ajeļ"
  },
  {
    "Code": "mn",
    "Name": "Mongolian",
    "NativeName": "Монгол"
  },
  {
    "Code": "na",
    "Name": "Nauru",
    "NativeName": "Ekakairũ Naoero"
  },
  {
    "Code": "nv",
    "Name": "Navajo, Navaho",
    "NativeName": "Diné Bizaad, Dinékʼehǰí"
  },
  {
    "Code": "nb",
    "Name": "Norwegian Bokmål",
    "NativeName": "Norsk Bokmål"
  },
  {
    "Code": "nd",
    "Name": "North Ndebele",
    "NativeName": "IsiNdebele"
  },
  {
    "Code": "ne",
    "Name": "Nepali",
    "NativeName": "नेपाली"
  },
  {
    "Code": "ng",
    "Name": "Ndonga",
    "NativeName": "Owambo"
  },
  {
    "Code": "nn",
    "Name": "Norwegian Nynorsk",
    "NativeName": "Norsk Nynorsk"
  },
  {
    "Code": "no",
    "Name": "Norwegian",
    "NativeName": "Norsk"
  },
  {
    "Code": "ii",
    "Name": "Nuosu",
    "NativeName": "ꆈꌠ꒿ Nuosuhxop"
  },
  {
    "Code": "nr",
    "Name": "South Ndebele",
    "NativeName": "IsiNdebele"
  },
  {
    "Code": "oc",
    "Name": "Occitan",
    "NativeName": "Occitan"
  },
  {
    "Code": "oj",
    "Name": "Ojibwe, Ojibwa",
    "NativeName": "ᐊᓂᔑᓈᐯᒧᐎᓐ"
  },
  {
    "Code": "cu",
    "Name": "Old Church Slavonic, Church Slavic, Church Slavonic, Old Bulgarian, Old Slavonic",
    "NativeName": "Ѩзыкъ Словѣньскъ"
  },
  {
    "Code": "om",
    "Name": "Oromo",
    "NativeName": "Afaan Oromoo"
  },
  {
    "Code": "or",
    "Name": "Oriya",
    "NativeName": "ଓଡ଼ିଆ"
  },
  {
    "Code": "os",
    "Name": "Ossetian, Ossetic",
    "NativeName": "Ирон Æвзаг"
  },
  {
    "Code": "pa",
    "Name": "Panjabi, Punjabi",
    "NativeName": "ਪੰਜਾਬੀ, پنجابی‎"
  },
  {
    "Code": "pi",
    "Name": "Pāli",
    "NativeName": "पाऴि"
  },
  {
    "Code": "fa",
    "Name": "Persian",
    "NativeName": "فارسی"
  },
  {
    "Code": "pl",
    "Name": "Polish",
    "NativeName": "Polski"
  },
  {
    "Code": "ps",
    "Name": "Pashto, Pushto",
    "NativeName": "پښتو"
  },
  {
    "Code": "pt",
    "Name": "Portuguese",
    "NativeName": "Português"
  },
  {
    "Code": "qu",
    "Name": "Quechua",
    "NativeName": "Runa Simi, Kichwa"
  },
  {
    "Code": "rm",
    "Name": "Romansh",
    "NativeName": "Rumantsch Grischun"
  },
  {
    "Code": "rn",
    "Name": "Kirundi",
    "NativeName": "KiRundi"
  },
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
    "Code": "sa",
    "Name": "Sanskrit (Saṁskṛta)",
    "NativeName": "संस्कृतम्"
  },
  {
    "Code": "sc",
    "Name": "Sardinian",
    "NativeName": "Sardu"
  },
  {
    "Code": "sd",
    "Name": "Sindhi",
    "NativeName": "सिन्धी, سنڌي، سندھی‎"
  },
  {
    "Code": "se",
    "Name": "Northern Sami",
    "NativeName": "Davvisámegiella"
  },
  {
    "Code": "sm",
    "Name": "Samoan",
    "NativeName": "Gagana Faa Samoa"
  },
  {
    "Code": "sg",
    "Name": "Sango",
    "NativeName": "Yângâ Tî Sängö"
  },
  {
    "Code": "sr",
    "Name": "Serbian",
    "NativeName": "Српски Језик"
  },
  {
    "Code": "gd",
    "Name": "Gaelic",
    "NativeName": "Gàidhlig"
  },
  {
    "Code": "sn",
    "Name": "Shona",
    "NativeName": "ChiShona"
  },
  {
    "Code": "si",
    "Name": "Sinhala, Sinhalese",
    "NativeName": "සිංහල"
  },
  {
    "Code": "sk",
    "Name": "Slovak",
    "NativeName": "Slovenčina"
  },
  {
    "Code": "sl",
    "Name": "Slovene",
    "NativeName": "Slovenščina"
  },
  {
    "Code": "so",
    "Name": "Somali",
    "NativeName": "Soomaaliga, Af Soomaali"
  },
  {
    "Code": "st",
    "Name": "Southern Sotho",
    "NativeName": "Sesotho"
  },
  {
    "Code": "es",
    "Name": "Spanish",
    "NativeName": "Español"
  },
  {
    "Code": "su",
    "Name": "Sundanese",
    "NativeName": "Basa Sunda"
  },
  {
    "Code": "sw",
    "Name": "Swahili",
    "NativeName": "Kiswahili"
  },
  {
    "Code": "ss",
    "Name": "Swati",
    "NativeName": "SiSwati"
  },
  {
    "Code": "sv",
    "Name": "Swedish",
    "NativeName": "Svenska"
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
    "Code": "tg",
    "Name": "Tajik",
    "NativeName": "Тоҷикӣ, Toğikī, تاجیکی‎"
  },
  {
    "Code": "th",
    "Name": "Thai",
    "NativeName": "ไทย"
  },
  {
    "Code": "ti",
    "Name": "Tigrinya",
    "NativeName": "ትግርኛ"
  },
  {
    "Code": "bo",
    "Name": "Tibetan Standard, Tibetan, Central",
    "NativeName": "བོད་ཡིག"
  },
  {
    "Code": "tk",
    "Name": "Turkmen",
    "NativeName": "Türkmen, Түркмен"
  },
  {
    "Code": "tl",
    "Name": "Tagalog",
    "NativeName": "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔"
  },
  {
    "Code": "tn",
    "Name": "Tswana",
    "NativeName": "Setswana"
  },
  {
    "Code": "to",
    "Name": "Tonga (Tonga Islands)",
    "NativeName": "Faka Tonga"
  },
  {
    "Code": "tr",
    "Name": "Turkish",
    "NativeName": "Türkçe"
  },
  {
    "Code": "ts",
    "Name": "Tsonga",
    "NativeName": "Xitsonga"
  },
  {
    "Code": "tt",
    "Name": "Tatar",
    "NativeName": "Татарча, Tatarça, تاتارچا‎"
  },
  {
    "Code": "tw",
    "Name": "Twi",
    "NativeName": "Twi"
  },
  {
    "Code": "ty",
    "Name": "Tahitian",
    "NativeName": "Reo Tahiti"
  },
  {
    "Code": "ug",
    "Name": "Uighur, Uyghur",
    "NativeName": "Uyƣurqə, ئۇيغۇرچە‎"
  },
  {
    "Code": "uk",
    "Name": "Ukrainian",
    "NativeName": "Українська"
  },
  {
    "Code": "ur",
    "Name": "Urdu",
    "NativeName": "اردو"
  },
  {
    "Code": "uz",
    "Name": "Uzbek",
    "NativeName": "Zbek, Ўзбек, أۇزبېك‎"
  },
  {
    "Code": "ve",
    "Name": "Venda",
    "NativeName": "Tshivenḓa"
  },
  {
    "Code": "vi",
    "Name": "Vietnamese",
    "NativeName": "Tiếng Việt"
  },
  {
    "Code": "vo",
    "Name": "Volapük",
    "NativeName": "Volapük"
  },
  {
    "Code": "wa",
    "Name": "Walloon",
    "NativeName": "Walon"
  },
  {
    "Code": "cy",
    "Name": "Welsh",
    "NativeName": "Cymraeg"
  },
  {
    "Code": "wo",
    "Name": "Wolof",
    "NativeName": "Wollof"
  },
  {
    "Code": "fy",
    "Name": "Western Frisian",
    "NativeName": "Frysk"
  },
  {
    "Code": "xh",
    "Name": "Xhosa",
    "NativeName": "IsiXhosa"
  },
  {
    "Code": "yi",
    "Name": "Yiddish",
    "NativeName": "ייִדיש"
  },
  {
    "Code": "yo",
    "Name": "Yoruba",
    "NativeName": "Yorùbá"
  },
  {
    "Code": "za",
    "Name": "Zhuang, Chuang",
    "NativeName": "Saɯ Cueŋƅ, Saw Cuengh"
  }
]`
