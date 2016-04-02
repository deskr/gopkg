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
    "code": "ab",
    "name": "Abkhaz",
    "native_name": "Аҧсуа"
  },
  {
    "code": "aa",
    "name": "Afar",
    "native_name": "Afaraf"
  },
  {
    "code": "af",
    "name": "Afrikaans",
    "native_name": "Afrikaans"
  },
  {
    "code": "ak",
    "name": "Akan",
    "native_name": "Akan"
  },
  {
    "code": "sq",
    "name": "Albanian",
    "native_name": "Shqip"
  },
  {
    "code": "am",
    "name": "Amharic",
    "native_name": "አማርኛ"
  },
  {
    "code": "ar",
    "name": "Arabic",
    "native_name": "العربية"
  },
  {
    "code": "an",
    "name": "Aragonese",
    "native_name": "Aragonés"
  },
  {
    "code": "hy",
    "name": "Armenian",
    "native_name": "Հայերեն"
  },
  {
    "code": "as",
    "name": "Assamese",
    "native_name": "অসমীয়া"
  },
  {
    "code": "av",
    "name": "Avaric",
    "native_name": "Авар МацӀ, МагӀарул МацӀ"
  },
  {
    "code": "ae",
    "name": "Avestan",
    "native_name": "Avesta"
  },
  {
    "code": "ay",
    "name": "Aymara",
    "native_name": "Aymar Aru"
  },
  {
    "code": "az",
    "name": "Azerbaijani",
    "native_name": "Azərbaycan Dili"
  },
  {
    "code": "bm",
    "name": "Bambara",
    "native_name": "Bamanankan"
  },
  {
    "code": "ba",
    "name": "Bashkir",
    "native_name": "Башҡорт Теле"
  },
  {
    "code": "eu",
    "name": "Basque",
    "native_name": "Euskara, Euskera"
  },
  {
    "code": "be",
    "name": "Belarusian",
    "native_name": "Беларуская"
  },
  {
    "code": "bn",
    "name": "Bengali",
    "native_name": "বাংলা"
  },
  {
    "code": "bh",
    "name": "Bihari",
    "native_name": "भोजपुरी"
  },
  {
    "code": "bi",
    "name": "Bislama",
    "native_name": "Bislama"
  },
  {
    "code": "bs",
    "name": "Bosnian",
    "native_name": "Bosanski Jezik"
  },
  {
    "code": "br",
    "name": "Breton",
    "native_name": "Brezhoneg"
  },
  {
    "code": "bg",
    "name": "Bulgarian",
    "native_name": "Български Език"
  },
  {
    "code": "my",
    "name": "Burmese",
    "native_name": "ဗမာစာ"
  },
  {
    "code": "ca",
    "name": "Catalan",
    "native_name": "Català"
  },
  {
    "code": "ch",
    "name": "Chamorro",
    "native_name": "Chamoru"
  },
  {
    "code": "ce",
    "name": "Chechen",
    "native_name": "Нохчийн Мотт"
  },
  {
    "code": "ny",
    "name": "Chichewa",
    "native_name": "ChiCheŵa"
  },
  {
    "code": "zh",
    "name": "Chinese",
    "native_name": "中文 (Zhōngwén), 汉语, 漢語"
  },
  {
    "code": "cv",
    "name": "Chuvash",
    "native_name": "Чӑваш Чӗлхи"
  },
  {
    "code": "kw",
    "name": "Cornish",
    "native_name": "Kernewek"
  },
  {
    "code": "co",
    "name": "Corsican",
    "native_name": "Corsu, Lingua Corsa"
  },
  {
    "code": "cr",
    "name": "Cree",
    "native_name": "ᓀᐦᐃᔭᐍᐏᐣ"
  },
  {
    "code": "hr",
    "name": "Croatian",
    "native_name": "Hrvatski"
  },
  {
    "code": "cs",
    "name": "Czech",
    "native_name": "Česky, Čeština"
  },
  {
    "code": "da",
    "name": "Danish",
    "native_name": "Dansk"
  },
  {
    "code": "dv",
    "name": "Maldivian",
    "native_name": ""
  },
  {
    "code": "nl",
    "name": "Dutch",
    "native_name": "Nederlands, Vlaams"
  },
  {
    "code": "en",
    "name": "English",
    "native_name": "English"
  },
  {
    "code": "eo",
    "name": "Esperanto",
    "native_name": "Esperanto"
  },
  {
    "code": "et",
    "name": "Estonian",
    "native_name": "Eesti, Eesti Keel"
  },
  {
    "code": "ee",
    "name": "Ewe",
    "native_name": "Eʋegbe"
  },
  {
    "code": "fo",
    "name": "Faroese",
    "native_name": "Føroyskt"
  },
  {
    "code": "fj",
    "name": "Fijian",
    "native_name": "Vosa Vakaviti"
  },
  {
    "code": "fi",
    "name": "Finnish",
    "native_name": "Suomi, Suomen Kieli"
  },
  {
    "code": "fr",
    "name": "French",
    "native_name": "Français, Langue Française"
  },
  {
    "code": "ff",
    "name": "Pular",
    "native_name": "Pular"
  },
  {
    "code": "gl",
    "name": "Galician",
    "native_name": "Galego"
  },
  {
    "code": "ka",
    "name": "Georgian",
    "native_name": "ქართული"
  },
  {
    "code": "de",
    "name": "German",
    "native_name": "Deutsch"
  },
  {
    "code": "el",
    "name": "Greek, Modern",
    "native_name": "Ελληνικά"
  },
  {
    "code": "gn",
    "name": "Guaraní",
    "native_name": "Avañeẽ"
  },
  {
    "code": "gu",
    "name": "Gujarati",
    "native_name": "ગુજરાતી"
  },
  {
    "code": "ht",
    "name": "Haitian",
    "native_name": "Kreyòl Ayisyen"
  },
  {
    "code": "ha",
    "name": "Hausa",
    "native_name": "Hausa, هَوُسَ"
  },
  {
    "code": "he",
    "name": "Hebrew (modern)",
    "native_name": "עברית"
  },
  {
    "code": "hz",
    "name": "Herero",
    "native_name": "Otjiherero"
  },
  {
    "code": "hi",
    "name": "Hindi",
    "native_name": "हिन्दी, हिंदी"
  },
  {
    "code": "ho",
    "name": "Hiri Motu",
    "native_name": "Hiri Motu"
  },
  {
    "code": "hu",
    "name": "Hungarian",
    "native_name": "Magyar"
  },
  {
    "code": "ia",
    "name": "Interlingua",
    "native_name": "Interlingua"
  },
  {
    "code": "id",
    "name": "Indonesian",
    "native_name": "Bahasa Indonesia"
  },
  {
    "code": "ie",
    "name": "Interlingue",
    "native_name": ""
  },
  {
    "code": "ga",
    "name": "Irish",
    "native_name": "Gaeilge"
  },
  {
    "code": "ig",
    "name": "Igbo",
    "native_name": "Asụsụ Igbo"
  },
  {
    "code": "ik",
    "name": "Inupiaq",
    "native_name": "Iñupiaq, Iñupiatun"
  },
  {
    "code": "io",
    "name": "Ido",
    "native_name": "Ido"
  },
  {
    "code": "is",
    "name": "Icelandic",
    "native_name": "Íslenska"
  },
  {
    "code": "it",
    "name": "Italian",
    "native_name": "Italiano"
  },
  {
    "code": "iu",
    "name": "Inuktitut",
    "native_name": "ᐃᓄᒃᑎᑐᑦ"
  },
  {
    "code": "ja",
    "name": "Japanese",
    "native_name": "日本語 (にほんご／にっぽんご)"
  },
  {
    "code": "jv",
    "name": "Javanese",
    "native_name": "Basa Jawa"
  },
  {
    "code": "kl",
    "name": "Kalaallisut, Greenlandic",
    "native_name": "Kalaallisut, Kalaallit Oqaasii"
  },
  {
    "code": "kn",
    "name": "Kannada",
    "native_name": "ಕನ್ನಡ"
  },
  {
    "code": "kr",
    "name": "Kanuri",
    "native_name": "Kanuri"
  },
  {
    "code": "ks",
    "name": "Kashmiri",
    "native_name": "कश्मीरी, كشميري‎"
  },
  {
    "code": "kk",
    "name": "Kazakh",
    "native_name": "Қазақ Тілі"
  },
  {
    "code": "km",
    "name": "Khmer",
    "native_name": "ភាសាខ្មែរ"
  },
  {
    "code": "ki",
    "name": "Kikuyu, Gikuyu",
    "native_name": "Gĩkũyũ"
  },
  {
    "code": "rw",
    "name": "Kinyarwanda",
    "native_name": "Ikinyarwanda"
  },
  {
    "code": "ky",
    "name": "Kirghiz, Kyrgyz",
    "native_name": "Кыргыз Тили"
  },
  {
    "code": "kv",
    "name": "Komi",
    "native_name": "Коми Кыв"
  },
  {
    "code": "kg",
    "name": "Kongo",
    "native_name": "KiKongo"
  },
  {
    "code": "ko",
    "name": "Korean",
    "native_name": "한국어 (韓國語), 조선말 (朝鮮語)"
  },
  {
    "code": "ku",
    "name": "Kurdish",
    "native_name": "Kurdî, كوردی‎"
  },
  {
    "code": "kj",
    "name": "Kwanyama, Kuanyama",
    "native_name": "Kuanyama"
  },
  {
    "code": "la",
    "name": "Latin",
    "native_name": "Latine, Lingua Latina"
  },
  {
    "code": "lb",
    "name": "Luxembourgish, Letzeburgesch",
    "native_name": "Lëtzebuergesch"
  },
  {
    "code": "lg",
    "name": "Luganda",
    "native_name": "Luganda"
  },
  {
    "code": "li",
    "name": "Limburgish, Limburgan, Limburger",
    "native_name": "Limburgs"
  },
  {
    "code": "ln",
    "name": "Lingala",
    "native_name": "Lingála"
  },
  {
    "code": "lo",
    "name": "Lao",
    "native_name": "ພາສາລາວ"
  },
  {
    "code": "lt",
    "name": "Lithuanian",
    "native_name": "Lietuvių Kalba"
  },
  {
    "code": "lu",
    "name": "Luba-Katanga",
    "native_name": ""
  },
  {
    "code": "lv",
    "name": "Latvian",
    "native_name": "Latviešu Valoda"
  },
  {
    "code": "gv",
    "name": "Manx",
    "native_name": "Gaelg, Gailck"
  },
  {
    "code": "mk",
    "name": "Macedonian",
    "native_name": "Македонски Јазик"
  },
  {
    "code": "mg",
    "name": "Malagasy",
    "native_name": "Malagasy Fiteny"
  },
  {
    "code": "ms",
    "name": "Malay",
    "native_name": "Bahasa Melayu, بهاس ملايو‎"
  },
  {
    "code": "ml",
    "name": "Malayalam",
    "native_name": "മലയാളം"
  },
  {
    "code": "mt",
    "name": "Maltese",
    "native_name": "Malti"
  },
  {
    "code": "mi",
    "name": "Māori",
    "native_name": "Te Reo Māori"
  },
  {
    "code": "mr",
    "name": "Marathi (Marāṭhī)",
    "native_name": "मराठी"
  },
  {
    "code": "mh",
    "name": "Marshallese",
    "native_name": "Kajin M̧ajeļ"
  },
  {
    "code": "mn",
    "name": "Mongolian",
    "native_name": "Монгол"
  },
  {
    "code": "na",
    "name": "Nauru",
    "native_name": "Ekakairũ Naoero"
  },
  {
    "code": "nv",
    "name": "Navajo, Navaho",
    "native_name": "Diné Bizaad, Dinékʼehǰí"
  },
  {
    "code": "nb",
    "name": "Norwegian Bokmål",
    "native_name": "Norsk Bokmål"
  },
  {
    "code": "nd",
    "name": "North Ndebele",
    "native_name": "IsiNdebele"
  },
  {
    "code": "ne",
    "name": "Nepali",
    "native_name": "नेपाली"
  },
  {
    "code": "ng",
    "name": "Ndonga",
    "native_name": "Owambo"
  },
  {
    "code": "nn",
    "name": "Norwegian Nynorsk",
    "native_name": "Norsk Nynorsk"
  },
  {
    "code": "no",
    "name": "Norwegian",
    "native_name": "Norsk"
  },
  {
    "code": "ii",
    "name": "Nuosu",
    "native_name": "ꆈꌠ꒿ Nuosuhxop"
  },
  {
    "code": "nr",
    "name": "South Ndebele",
    "native_name": "IsiNdebele"
  },
  {
    "code": "oc",
    "name": "Occitan",
    "native_name": "Occitan"
  },
  {
    "code": "oj",
    "name": "Ojibwe, Ojibwa",
    "native_name": "ᐊᓂᔑᓈᐯᒧᐎᓐ"
  },
  {
    "code": "cu",
    "name": "Old Church Slavonic, Church Slavic, Church Slavonic, Old Bulgarian, Old Slavonic",
    "native_name": "Ѩзыкъ Словѣньскъ"
  },
  {
    "code": "om",
    "name": "Oromo",
    "native_name": "Afaan Oromoo"
  },
  {
    "code": "or",
    "name": "Oriya",
    "native_name": "ଓଡ଼ିଆ"
  },
  {
    "code": "os",
    "name": "Ossetian, Ossetic",
    "native_name": "Ирон Æвзаг"
  },
  {
    "code": "pa",
    "name": "Panjabi, Punjabi",
    "native_name": "ਪੰਜਾਬੀ, پنجابی‎"
  },
  {
    "code": "pi",
    "name": "Pāli",
    "native_name": "पाऴि"
  },
  {
    "code": "fa",
    "name": "Persian",
    "native_name": "فارسی"
  },
  {
    "code": "pl",
    "name": "Polish",
    "native_name": "Polski"
  },
  {
    "code": "ps",
    "name": "Pashto, Pushto",
    "native_name": "پښتو"
  },
  {
    "code": "pt",
    "name": "Portuguese",
    "native_name": "Português"
  },
  {
    "code": "qu",
    "name": "Quechua",
    "native_name": "Runa Simi, Kichwa"
  },
  {
    "code": "rm",
    "name": "Romansh",
    "native_name": "Rumantsch Grischun"
  },
  {
    "code": "rn",
    "name": "Kirundi",
    "native_name": "KiRundi"
  },
  {
    "code": "ro",
    "name": "Romanian, Moldavian, Moldovan",
    "native_name": "Română"
  },
  {
    "code": "ru",
    "name": "Russian",
    "native_name": "Русский Язык"
  },
  {
    "code": "sa",
    "name": "Sanskrit (Saṁskṛta)",
    "native_name": "संस्कृतम्"
  },
  {
    "code": "sc",
    "name": "Sardinian",
    "native_name": "Sardu"
  },
  {
    "code": "sd",
    "name": "Sindhi",
    "native_name": "सिन्धी, سنڌي، سندھی‎"
  },
  {
    "code": "se",
    "name": "Northern Sami",
    "native_name": "Davvisámegiella"
  },
  {
    "code": "sm",
    "name": "Samoan",
    "native_name": "Gagana Faa Samoa"
  },
  {
    "code": "sg",
    "name": "Sango",
    "native_name": "Yângâ Tî Sängö"
  },
  {
    "code": "sr",
    "name": "Serbian",
    "native_name": "Српски Језик"
  },
  {
    "code": "gd",
    "name": "Gaelic",
    "native_name": "Gàidhlig"
  },
  {
    "code": "sn",
    "name": "Shona",
    "native_name": "ChiShona"
  },
  {
    "code": "si",
    "name": "Sinhala, Sinhalese",
    "native_name": "සිංහල"
  },
  {
    "code": "sk",
    "name": "Slovak",
    "native_name": "Slovenčina"
  },
  {
    "code": "sl",
    "name": "Slovene",
    "native_name": "Slovenščina"
  },
  {
    "code": "so",
    "name": "Somali",
    "native_name": "Soomaaliga, Af Soomaali"
  },
  {
    "code": "st",
    "name": "Southern Sotho",
    "native_name": "Sesotho"
  },
  {
    "code": "es",
    "name": "Spanish",
    "native_name": "Español"
  },
  {
    "code": "su",
    "name": "Sundanese",
    "native_name": "Basa Sunda"
  },
  {
    "code": "sw",
    "name": "Swahili",
    "native_name": "Kiswahili"
  },
  {
    "code": "ss",
    "name": "Swati",
    "native_name": "SiSwati"
  },
  {
    "code": "sv",
    "name": "Swedish",
    "native_name": "Svenska"
  },
  {
    "code": "ta",
    "name": "Tamil",
    "native_name": "தமிழ்"
  },
  {
    "code": "te",
    "name": "Telugu",
    "native_name": "తెలుగు"
  },
  {
    "code": "tg",
    "name": "Tajik",
    "native_name": "Тоҷикӣ, Toğikī, تاجیکی‎"
  },
  {
    "code": "th",
    "name": "Thai",
    "native_name": "ไทย"
  },
  {
    "code": "ti",
    "name": "Tigrinya",
    "native_name": "ትግርኛ"
  },
  {
    "code": "bo",
    "name": "Tibetan Standard, Tibetan, Central",
    "native_name": "བོད་ཡིག"
  },
  {
    "code": "tk",
    "name": "Turkmen",
    "native_name": "Türkmen, Түркмен"
  },
  {
    "code": "tl",
    "name": "Tagalog",
    "native_name": "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔"
  },
  {
    "code": "tn",
    "name": "Tswana",
    "native_name": "Setswana"
  },
  {
    "code": "to",
    "name": "Tonga (Tonga Islands)",
    "native_name": "Faka Tonga"
  },
  {
    "code": "tr",
    "name": "Turkish",
    "native_name": "Türkçe"
  },
  {
    "code": "ts",
    "name": "Tsonga",
    "native_name": "Xitsonga"
  },
  {
    "code": "tt",
    "name": "Tatar",
    "native_name": "Татарча, Tatarça, تاتارچا‎"
  },
  {
    "code": "tw",
    "name": "Twi",
    "native_name": "Twi"
  },
  {
    "code": "ty",
    "name": "Tahitian",
    "native_name": "Reo Tahiti"
  },
  {
    "code": "ug",
    "name": "Uighur, Uyghur",
    "native_name": "Uyƣurqə, ئۇيغۇرچە‎"
  },
  {
    "code": "uk",
    "name": "Ukrainian",
    "native_name": "Українська"
  },
  {
    "code": "ur",
    "name": "Urdu",
    "native_name": "اردو"
  },
  {
    "code": "uz",
    "name": "Uzbek",
    "native_name": "Zbek, Ўзбек, أۇزبېك‎"
  },
  {
    "code": "ve",
    "name": "Venda",
    "native_name": "Tshivenḓa"
  },
  {
    "code": "vi",
    "name": "Vietnamese",
    "native_name": "Tiếng Việt"
  },
  {
    "code": "vo",
    "name": "Volapük",
    "native_name": "Volapük"
  },
  {
    "code": "wa",
    "name": "Walloon",
    "native_name": "Walon"
  },
  {
    "code": "cy",
    "name": "Welsh",
    "native_name": "Cymraeg"
  },
  {
    "code": "wo",
    "name": "Wolof",
    "native_name": "Wollof"
  },
  {
    "code": "fy",
    "name": "Western Frisian",
    "native_name": "Frysk"
  },
  {
    "code": "xh",
    "name": "Xhosa",
    "native_name": "IsiXhosa"
  },
  {
    "code": "yi",
    "name": "Yiddish",
    "native_name": "ייִדיש"
  },
  {
    "code": "yo",
    "name": "Yoruba",
    "native_name": "Yorùbá"
  },
  {
    "code": "za",
    "name": "Zhuang, Chuang",
    "native_name": "Saɯ Cueŋƅ, Saw Cuengh"
  }
]`
