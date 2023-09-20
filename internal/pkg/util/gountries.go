package util

import (
	"encoding/json"
	"fmt"
	"github.com/pariz/gountries"
)

type Countries []struct {
	Iso2     string `json:"ISO2"`
	Iso3     string `json:"ISO3"`
	Digits   string `json:"DIGITS"`
	ISO31662 string `json:"ISO-3166-2"`
	English  string `json:"English"`
	China    string `json:"China"`
	Taiwan   string `json:"Taiwan"`
	Hongkong string `json:"Hongkong"`
	Memo     string `json:"Memo"`
}

func ConvertCountryNameToChinese(name string) string {
	query := gountries.New()

	/////////////////
	// Find country //
	/////////////////

	country, _ := query.FindCountryByName(name)

	if country.Alpha3 == "" {
		country, _ = query.FindCountryByAlpha(name)
	}

	// fmt.Println(country.Alpha3)
	// fmt.Println(country.Name.Official)

	if country.Alpha3 != "" {
		countryList := GetContries()
		for _, i := range countryList {
			if i.Iso3 == country.Alpha3 {
				return i.Taiwan
			}
		}
	}

	return name
}

func ConvertCountryChineseToName(name string) string {
	countryList := GetContries()
	for _, i := range countryList {
		if i.Taiwan == name {
			return i.English
		}
	}
	return name
}

func GetContries() (countries Countries) {
	//對照查表 https://gist.github.com/jacobbubu/060d84c2bdf005d412db
	str := []byte(`[
		{
		 "ISO2": "AD",
		 "ISO3": "AND",
		 "DIGITS": "20",
		 "ISO-3166-2": "ISO 3166-2:AD",
		 "English": " Andorra",
		 "China": "安道尔",
		 "Taiwan": "安道爾",
		 "Hongkong": "安道爾",
		 "Memo": ""
		},
		{
		 "ISO2": "AE",
		 "ISO3": "ARE",
		 "DIGITS": "784",
		 "ISO-3166-2": "ISO 3166-2:AE",
		 "English": " United Arab Emirates",
		 "China": "阿联酋",
		 "Taiwan": "杜拜",
		 "Hongkong": "阿聯酋",
		 "Memo": ""
		},
		{
		 "ISO2": "AF",
		 "ISO3": "AFG",
		 "DIGITS": "4",
		 "ISO-3166-2": "ISO 3166-2:AF",
		 "English": " Afghanistan",
		 "China": "阿富汗",
		 "Taiwan": "阿富汗",
		 "Hongkong": "阿富汗",
		 "Memo": ""
		},
		{
		 "ISO2": "AG",
		 "ISO3": "ATG",
		 "DIGITS": "28",
		 "ISO-3166-2": "ISO 3166-2:AG",
		 "English": " Antigua and Barbuda",
		 "China": "安提瓜和巴布达",
		 "Taiwan": "安提瓜和巴布達",
		 "Hongkong": "安提瓜和巴布達",
		 "Memo": ""
		},
		{
		 "ISO2": "AI",
		 "ISO3": "AIA",
		 "DIGITS": "660",
		 "ISO-3166-2": "ISO 3166-2:AI",
		 "English": " Anguilla",
		 "China": "安圭拉",
		 "Taiwan": "英屬安圭拉",
		 "Hongkong": "安圭拉島",
		 "Memo": ""
		},
		{
		 "ISO2": "AL",
		 "ISO3": "ALB",
		 "DIGITS": "8",
		 "ISO-3166-2": "ISO 3166-2:AL",
		 "English": " Albania",
		 "China": "阿尔巴尼亚",
		 "Taiwan": "阿爾巴尼亞",
		 "Hongkong": "阿爾巴尼亞",
		 "Memo": ""
		},
		{
		 "ISO2": "AM",
		 "ISO3": "ARM",
		 "DIGITS": "51",
		 "ISO-3166-2": "ISO 3166-2:AM",
		 "English": " Armenia",
		 "China": "亚美尼亚",
		 "Taiwan": "亞美尼亞",
		 "Hongkong": "亞美尼亞",
		 "Memo": ""
		},
		{
		 "ISO2": "AO",
		 "ISO3": "AGO",
		 "DIGITS": "24",
		 "ISO-3166-2": "ISO 3166-2:AO",
		 "English": " Angola",
		 "China": "安哥拉",
		 "Taiwan": "安哥拉",
		 "Hongkong": "安哥拉",
		 "Memo": ""
		},
		{
		 "ISO2": "AQ",
		 "ISO3": "ATA",
		 "DIGITS": "10",
		 "ISO-3166-2": "ISO 3166-2:AQ",
		 "English": " Antarctica",
		 "China": "南极洲",
		 "Taiwan": "南極洲",
		 "Hongkong": "南極洲",
		 "Memo": ""
		},
		{
		 "ISO2": "AR",
		 "ISO3": "ARG",
		 "DIGITS": "32",
		 "ISO-3166-2": "ISO 3166-2:AR",
		 "English": " Argentina",
		 "China": "阿根廷",
		 "Taiwan": "阿根廷",
		 "Hongkong": "阿根廷",
		 "Memo": ""
		},
		{
		 "ISO2": "AS",
		 "ISO3": "ASM",
		 "DIGITS": "16",
		 "ISO-3166-2": "ISO 3166-2:AS",
		 "English": " American Samoa",
		 "China": "美属萨摩亚",
		 "Taiwan": "美屬薩摩亞",
		 "Hongkong": "美屬薩摩亞",
		 "Memo": ""
		},
		{
		 "ISO2": "AT",
		 "ISO3": "AUT",
		 "DIGITS": "40",
		 "ISO-3166-2": "ISO 3166-2:AT",
		 "English": " Austria",
		 "China": "奥地利",
		 "Taiwan": "奧地利",
		 "Hongkong": "奧地利",
		 "Memo": ""
		},
		{
		 "ISO2": "AU",
		 "ISO3": "AUS",
		 "DIGITS": "36",
		 "ISO-3166-2": "ISO 3166-2:AU",
		 "English": " Australia",
		 "China": "澳大利亚",
		 "Taiwan": "澳洲",
		 "Hongkong": "澳洲",
		 "Memo": ""
		},
		{
		 "ISO2": "AW",
		 "ISO3": "ABW",
		 "DIGITS": "533",
		 "ISO-3166-2": "ISO 3166-2:AW",
		 "English": " Aruba",
		 "China": "阿鲁巴",
		 "Taiwan": "阿魯巴",
		 "Hongkong": "阿魯巴",
		 "Memo": ""
		},
		{
		 "ISO2": "AX",
		 "ISO3": "ALA",
		 "DIGITS": "248",
		 "ISO-3166-2": "ISO 3166-2:AX",
		 "English": " Åaland Islands",
		 "China": "奥兰群岛",
		 "Taiwan": "奧蘭群島",
		 "Hongkong": "亞蘭群島",
		 "Memo": ""
		},
		{
		 "ISO2": "AZ",
		 "ISO3": "AZE",
		 "DIGITS": "31",
		 "ISO-3166-2": "ISO 3166-2:AZ",
		 "English": " Azerbaijan",
		 "China": "阿塞拜疆",
		 "Taiwan": "亞賽拜然",
		 "Hongkong": "阿塞拜疆",
		 "Memo": ""
		},
		{
		 "ISO2": "BA",
		 "ISO3": "BIH",
		 "DIGITS": "70",
		 "ISO-3166-2": "ISO 3166-2:BA",
		 "English": " Bosnia and Herzegovina",
		 "China": "波黑",
		 "Taiwan": "波赫",
		 "Hongkong": "波黑",
		 "Memo": ""
		},
		{
		 "ISO2": "BB",
		 "ISO3": "BRB",
		 "DIGITS": "52",
		 "ISO-3166-2": "ISO 3166-2:BB",
		 "English": " Barbados",
		 "China": "巴巴多斯",
		 "Taiwan": "巴貝多",
		 "Hongkong": "巴巴多斯",
		 "Memo": ""
		},
		{
		 "ISO2": "BD",
		 "ISO3": "BGD",
		 "DIGITS": "50",
		 "ISO-3166-2": "ISO 3166-2:BD",
		 "English": " Bangladesh",
		 "China": "孟加拉国",
		 "Taiwan": "孟加拉",
		 "Hongkong": "孟加拉",
		 "Memo": ""
		},
		{
		 "ISO2": "BE",
		 "ISO3": "BEL",
		 "DIGITS": "56",
		 "ISO-3166-2": "ISO 3166-2:BE",
		 "English": " Belgium",
		 "China": "比利时",
		 "Taiwan": "比利時",
		 "Hongkong": "比利時",
		 "Memo": ""
		},
		{
		 "ISO2": "BF",
		 "ISO3": "BFA",
		 "DIGITS": "854",
		 "ISO-3166-2": "ISO 3166-2:BF",
		 "English": " Burkina Faso",
		 "China": "布基纳法索",
		 "Taiwan": "布吉納法索",
		 "Hongkong": "布基納法索",
		 "Memo": ""
		},
		{
		 "ISO2": "BG",
		 "ISO3": "BGR",
		 "DIGITS": "100",
		 "ISO-3166-2": "ISO 3166-2:BG",
		 "English": " Bulgaria",
		 "China": "保加利亚",
		 "Taiwan": "保加利亞",
		 "Hongkong": "保加利亞",
		 "Memo": ""
		},
		{
		 "ISO2": "BH",
		 "ISO3": "BHR",
		 "DIGITS": "48",
		 "ISO-3166-2": "ISO 3166-2:BH",
		 "English": " Bahrain",
		 "China": "巴林",
		 "Taiwan": "巴林",
		 "Hongkong": "巴林",
		 "Memo": ""
		},
		{
		 "ISO2": "BI",
		 "ISO3": "BDI",
		 "DIGITS": "108",
		 "ISO-3166-2": "ISO 3166-2:BI",
		 "English": " Burundi",
		 "China": "布隆迪",
		 "Taiwan": "蒲隆地",
		 "Hongkong": "布隆迪",
		 "Memo": ""
		},
		{
		 "ISO2": "BJ",
		 "ISO3": "BEN",
		 "DIGITS": "204",
		 "ISO-3166-2": "ISO 3166-2:BJ",
		 "English": " Benin",
		 "China": "贝宁",
		 "Taiwan": "貝南",
		 "Hongkong": "貝寧",
		 "Memo": ""
		},
		{
		 "ISO2": "BL",
		 "ISO3": "BLM",
		 "DIGITS": "652",
		 "ISO-3166-2": "ISO 3166-2:BL",
		 "English": " Saint Barthélemy",
		 "China": "圣巴泰勒米岛",
		 "Taiwan": "聖巴瑟米",
		 "Hongkong": "聖巴托洛繆島",
		 "Memo": ""
		},
		{
		 "ISO2": "BM",
		 "ISO3": "BMU",
		 "DIGITS": "60",
		 "ISO-3166-2": "ISO 3166-2:BM",
		 "English": " Bermuda",
		 "China": "百慕大",
		 "Taiwan": "百慕達",
		 "Hongkong": "百慕達",
		 "Memo": ""
		},
		{
		 "ISO2": "BN",
		 "ISO3": "BRN",
		 "DIGITS": "96",
		 "ISO-3166-2": "ISO 3166-2:BN",
		 "English": " Brunei Darussalam",
		 "China": "文莱",
		 "Taiwan": "汶萊",
		 "Hongkong": "汶萊",
		 "Memo": ""
		},
		{
		 "ISO2": "BO",
		 "ISO3": "BOL",
		 "DIGITS": "68",
		 "ISO-3166-2": "ISO 3166-2:BO",
		 "English": " Bolivia (Plurinational State of)",
		 "China": "玻利维亚",
		 "Taiwan": "波利維亞",
		 "Hongkong": "玻利維亞",
		 "Memo": ""
		},
		{
		 "ISO2": "BQ",
		 "ISO3": "BES",
		 "DIGITS": "535",
		 "ISO-3166-2": "ISO 3166-2:BQ",
		 "English": " Bonaire, Sint Eustatius and Saba",
		 "China": "荷兰加勒比区",
		 "Taiwan": "荷蘭加勒比區",
		 "Hongkong": "荷蘭加勒比區",
		 "Memo": ""
		},
		{
		 "ISO2": "BR",
		 "ISO3": "BRA",
		 "DIGITS": "76",
		 "ISO-3166-2": "ISO 3166-2:BR",
		 "English": " Brazil",
		 "China": "巴西",
		 "Taiwan": "巴西",
		 "Hongkong": "巴西",
		 "Memo": ""
		},
		{
		 "ISO2": "BS",
		 "ISO3": "BHS",
		 "DIGITS": "44",
		 "ISO-3166-2": "ISO 3166-2:BS",
		 "English": " Bahamas",
		 "China": "巴哈马",
		 "Taiwan": "巴哈馬",
		 "Hongkong": "巴哈馬",
		 "Memo": ""
		},
		{
		 "ISO2": "BT",
		 "ISO3": "BTN",
		 "DIGITS": "64",
		 "ISO-3166-2": "ISO 3166-2:BT",
		 "English": " Bhutan",
		 "China": "不丹",
		 "Taiwan": "不丹",
		 "Hongkong": "不丹",
		 "Memo": ""
		},
		{
		 "ISO2": "BV",
		 "ISO3": "BVT",
		 "DIGITS": "74",
		 "ISO-3166-2": "ISO 3166-2:BV",
		 "English": " Bouvet Island",
		 "China": "布韦岛",
		 "Taiwan": "布威島",
		 "Hongkong": "鮑威特島",
		 "Memo": ""
		},
		{
		 "ISO2": "BW",
		 "ISO3": "BWA",
		 "DIGITS": "72",
		 "ISO-3166-2": "ISO 3166-2:BW",
		 "English": " Botswana",
		 "China": "博茨瓦纳",
		 "Taiwan": "波札那",
		 "Hongkong": "博茨瓦納",
		 "Memo": ""
		},
		{
		 "ISO2": "BY",
		 "ISO3": "BLR",
		 "DIGITS": "112",
		 "ISO-3166-2": "ISO 3166-2:BY",
		 "English": " Belarus",
		 "China": "白俄罗斯",
		 "Taiwan": "白俄羅斯",
		 "Hongkong": "白俄羅斯",
		 "Memo": ""
		},
		{
		 "ISO2": "BZ",
		 "ISO3": "BLZ",
		 "DIGITS": "84",
		 "ISO-3166-2": "ISO 3166-2:BZ",
		 "English": " Belize",
		 "China": "伯利兹",
		 "Taiwan": "貝里斯",
		 "Hongkong": "伯利茲",
		 "Memo": ""
		},
		{
		 "ISO2": "CA",
		 "ISO3": "CAN",
		 "DIGITS": "124",
		 "ISO-3166-2": "ISO 3166-2:CA",
		 "English": " Canada",
		 "China": "加拿大",
		 "Taiwan": "加拿大",
		 "Hongkong": "加拿大",
		 "Memo": ""
		},
		{
		 "ISO2": "CC",
		 "ISO3": "CCK",
		 "DIGITS": "166",
		 "ISO-3166-2": "ISO 3166-2:CC",
		 "English": " Cocos (Keeling) Islands",
		 "China": "科科斯群岛",
		 "Taiwan": "可可斯群島",
		 "Hongkong": "科科斯群島",
		 "Memo": ""
		},
		{
		 "ISO2": "CD",
		 "ISO3": "COD",
		 "DIGITS": "180",
		 "ISO-3166-2": "ISO 3166-2:CD",
		 "English": " Congo (Democratic Republic of the)",
		 "China": "刚果（金）",
		 "Taiwan": "民主剛果",
		 "Hongkong": "民主剛果",
		 "Memo": "中國大陸主要使用「剛果（金）」一詞，意指「首都為金沙薩的剛果（共和國）」，而「民主剛果」一詞亦普遍為民間所用。"
		},
		{
		 "ISO2": "CF",
		 "ISO3": "CAF",
		 "DIGITS": "140",
		 "ISO-3166-2": "ISO 3166-2:CF",
		 "English": " Central African Republic",
		 "China": "中非",
		 "Taiwan": "中非",
		 "Hongkong": "中非",
		 "Memo": ""
		},
		{
		 "ISO2": "CG",
		 "ISO3": "COG",
		 "DIGITS": "178",
		 "ISO-3166-2": "ISO 3166-2:CG",
		 "English": " Congo",
		 "China": "刚果（布）",
		 "Taiwan": "剛果",
		 "Hongkong": "剛果",
		 "Memo": "中國大陸主要使用「剛果（布）」一詞，意指「首都為布拉柴維爾的剛果（共和國）」，而「剛果」一詞亦普遍為民間所用。"
		},
		{
		 "ISO2": "CH",
		 "ISO3": "CHE",
		 "DIGITS": "756",
		 "ISO-3166-2": "ISO 3166-2:CH",
		 "English": " Switzerland",
		 "China": "瑞士",
		 "Taiwan": "瑞士",
		 "Hongkong": "瑞士",
		 "Memo": ""
		},
		{
		 "ISO2": "CI",
		 "ISO3": "CIV",
		 "DIGITS": "384",
		 "ISO-3166-2": "ISO 3166-2:CI",
		 "English": " Côte d'Ivoire",
		 "China": "科特迪瓦",
		 "Taiwan": "象牙海岸",
		 "Hongkong": "科特迪瓦",
		 "Memo": "香港亦普遍採用「象牙海岸」一詞於其它場合（如香港郵政的郵政指南附錄表）"
		},
		{
		 "ISO2": "CK",
		 "ISO3": "COK",
		 "DIGITS": "184",
		 "ISO-3166-2": "ISO 3166-2:CK",
		 "English": " Cook Islands",
		 "China": "库克群岛",
		 "Taiwan": "庫克群島",
		 "Hongkong": "庫克群島",
		 "Memo": "香港亦普遍採用「[[科克群島]]」（CNS 12842譯名）一詞於其它場合"
		},
		{
		 "ISO2": "CL",
		 "ISO3": "CHL",
		 "DIGITS": "152",
		 "ISO-3166-2": "ISO 3166-2:CL",
		 "English": " Chile",
		 "China": "智利",
		 "Taiwan": "智利",
		 "Hongkong": "智利",
		 "Memo": ""
		},
		{
		 "ISO2": "CM",
		 "ISO3": "CMR",
		 "DIGITS": "120",
		 "ISO-3166-2": "ISO 3166-2:CM",
		 "English": " Cameroon",
		 "China": "喀麦隆",
		 "Taiwan": "喀麥隆",
		 "Hongkong": "喀麥隆",
		 "Memo": ""
		},
		{
		 "ISO2": "CN",
		 "ISO3": "CHN",
		 "DIGITS": "156",
		 "ISO-3166-2": "ISO 3166-2:CN",
		 "English": " People's Republic of China",
		 "China": "中国(內地)",
		 "Taiwan": "中國(大陸)",
		 "Hongkong": "大陸(內地)",
		 "Memo": "「GB/T 2659-2000」的「CN」適用於整個中華人民共和國轄區（包括中國大陸、香港、澳門）。而「ISO 3166-1」和「CNS 12842」的「CN」則僅適用於中國大陸，不包括港澳地區。"
		},
		{
		 "ISO2": "CO",
		 "ISO3": "COL",
		 "DIGITS": "170",
		 "ISO-3166-2": "ISO 3166-2:CO",
		 "English": " Colombia",
		 "China": "哥伦比亚",
		 "Taiwan": "哥倫比亞",
		 "Hongkong": "哥倫比亞",
		 "Memo": ""
		},
		{
		 "ISO2": "CR",
		 "ISO3": "CRI",
		 "DIGITS": "188",
		 "ISO-3166-2": "ISO 3166-2:CR",
		 "English": " Costa Rica",
		 "China": "哥斯达黎加",
		 "Taiwan": "哥斯大黎加",
		 "Hongkong": "哥斯達黎加",
		 "Memo": ""
		},
		{
		 "ISO2": "CU",
		 "ISO3": "CUB",
		 "DIGITS": "192",
		 "ISO-3166-2": "ISO 3166-2:CU",
		 "English": " Cuba",
		 "China": "古巴",
		 "Taiwan": "古巴",
		 "Hongkong": "古巴",
		 "Memo": ""
		},
		{
		 "ISO2": "CV",
		 "ISO3": "CPV",
		 "DIGITS": "132",
		 "ISO-3166-2": "ISO 3166-2:CV",
		 "English": " Cape Verde",
		 "China": "佛得角",
		 "Taiwan": "維德角",
		 "Hongkong": "佛得角",
		 "Memo": ""
		},
		{
		 "ISO2": "CW",
		 "ISO3": "CUW",
		 "DIGITS": "531",
		 "ISO-3166-2": "ISO 3166-2:CW",
		 "English": " Curaçao",
		 "China": "库拉索",
		 "Taiwan": "庫拉索",
		 "Hongkong": "庫拉索",
		 "Memo": ""
		},
		{
		 "ISO2": "CX",
		 "ISO3": "CXR",
		 "DIGITS": "162",
		 "ISO-3166-2": "ISO 3166-2:CX",
		 "English": " Christmas Island",
		 "China": "圣诞岛",
		 "Taiwan": "聖誕島",
		 "Hongkong": "聖誕島",
		 "Memo": ""
		},
		{
		 "ISO2": "CY",
		 "ISO3": "CYP",
		 "DIGITS": "196",
		 "ISO-3166-2": "ISO 3166-2:CY",
		 "English": " Cyprus",
		 "China": "塞浦路斯",
		 "Taiwan": "塞普勒斯",
		 "Hongkong": "塞浦路斯",
		 "Memo": ""
		},
		{
		 "ISO2": "CZ",
		 "ISO3": "CZE",
		 "DIGITS": "203",
		 "ISO-3166-2": "ISO 3166-2:CZ",
		 "English": " Czech Republic",
		 "China": "捷克",
		 "Taiwan": "捷克",
		 "Hongkong": "捷克",
		 "Memo": ""
		},
		{
		 "ISO2": "DE",
		 "ISO3": "DEU",
		 "DIGITS": "276",
		 "ISO-3166-2": "ISO 3166-2:DE",
		 "English": " Germany",
		 "China": "德国",
		 "Taiwan": "德國",
		 "Hongkong": "德國",
		 "Memo": ""
		},
		{
		 "ISO2": "DJ",
		 "ISO3": "DJI",
		 "DIGITS": "262",
		 "ISO-3166-2": "ISO 3166-2:DJ",
		 "English": " Djibouti",
		 "China": "吉布提",
		 "Taiwan": "吉布地",
		 "Hongkong": "吉布提",
		 "Memo": ""
		},
		{
		 "ISO2": "DK",
		 "ISO3": "DNK",
		 "DIGITS": "208",
		 "ISO-3166-2": "ISO 3166-2:DK",
		 "English": " Denmark",
		 "China": "丹麦",
		 "Taiwan": "丹麥",
		 "Hongkong": "丹麥",
		 "Memo": ""
		},
		{
		 "ISO2": "DM",
		 "ISO3": "DMA",
		 "DIGITS": "212",
		 "ISO-3166-2": "ISO 3166-2:DM",
		 "English": " Dominica",
		 "China": "多米尼克",
		 "Taiwan": "多米尼克",
		 "Hongkong": "多米尼克",
		 "Memo": ""
		},
		{
		 "ISO2": "DO",
		 "ISO3": "DOM",
		 "DIGITS": "214",
		 "ISO-3166-2": "ISO 3166-2:DO",
		 "English": " Dominican Republic",
		 "China": "多米尼加",
		 "Taiwan": "多明尼加",
		 "Hongkong": "多明尼加",
		 "Memo": ""
		},
		{
		 "ISO2": "DZ",
		 "ISO3": "DZA",
		 "DIGITS": "12",
		 "ISO-3166-2": "ISO 3166-2:DZ",
		 "English": " Algeria",
		 "China": "阿尔及利亚",
		 "Taiwan": "阿爾及利亞",
		 "Hongkong": "阿爾及利亞",
		 "Memo": ""
		},
		{
		 "ISO2": "EC",
		 "ISO3": "ECU",
		 "DIGITS": "218",
		 "ISO-3166-2": "ISO 3166-2:EC",
		 "English": " Ecuador",
		 "China": "厄瓜多尔",
		 "Taiwan": "厄瓜多",
		 "Hongkong": "厄瓜多爾",
		 "Memo": ""
		},
		{
		 "ISO2": "EE",
		 "ISO3": "EST",
		 "DIGITS": "233",
		 "ISO-3166-2": "ISO 3166-2:EE",
		 "English": " Estonia",
		 "China": "爱沙尼亚",
		 "Taiwan": "愛沙尼亞",
		 "Hongkong": "愛沙尼亞",
		 "Memo": ""
		},
		{
		 "ISO2": "EG",
		 "ISO3": "EGY",
		 "DIGITS": "818",
		 "ISO-3166-2": "ISO 3166-2:EG",
		 "English": " Egypt",
		 "China": "埃及",
		 "Taiwan": "埃及",
		 "Hongkong": "埃及",
		 "Memo": ""
		},
		{
		 "ISO2": "EH",
		 "ISO3": "ESH",
		 "DIGITS": "732",
		 "ISO-3166-2": "ISO 3166-2:EH",
		 "English": " Western Sahara",
		 "China": "西撒哈拉",
		 "Taiwan": "西撒哈拉",
		 "Hongkong": "西撒哈拉",
		 "Memo": ""
		},
		{
		 "ISO2": "ER",
		 "ISO3": "ERI",
		 "DIGITS": "232",
		 "ISO-3166-2": "ISO 3166-2:ER",
		 "English": " Eritrea",
		 "China": "厄立特里亚",
		 "Taiwan": "厄立垂亞",
		 "Hongkong": "厄立特里亞",
		 "Memo": ""
		},
		{
		 "ISO2": "ES",
		 "ISO3": "ESP",
		 "DIGITS": "724",
		 "ISO-3166-2": "ISO 3166-2:ES",
		 "English": " Spain",
		 "China": "西班牙",
		 "Taiwan": "西班牙",
		 "Hongkong": "西班牙",
		 "Memo": ""
		},
		{
		 "ISO2": "ET",
		 "ISO3": "ETH",
		 "DIGITS": "231",
		 "ISO-3166-2": "ISO 3166-2:ET",
		 "English": " Ethiopia",
		 "China": "埃塞俄比亚",
		 "Taiwan": "衣索比亞",
		 "Hongkong": "埃塞俄比亞",
		 "Memo": "亦有部份人士使用「衣索匹亞」一詞於台灣"
		},
		{
		 "ISO2": "FI",
		 "ISO3": "FIN",
		 "DIGITS": "246",
		 "ISO-3166-2": "ISO 3166-2:FI",
		 "English": " Finland",
		 "China": "芬兰",
		 "Taiwan": "芬蘭",
		 "Hongkong": "芬蘭",
		 "Memo": ""
		},
		{
		 "ISO2": "FJ",
		 "ISO3": "FJI",
		 "DIGITS": "242",
		 "ISO-3166-2": "ISO 3166-2:FJ",
		 "English": " Fiji",
		 "China": "斐济群岛",
		 "Taiwan": "斐濟",
		 "Hongkong": "斐濟",
		 "Memo": ""
		},
		{
		 "ISO2": "FK",
		 "ISO3": "FLK",
		 "DIGITS": "238",
		 "ISO-3166-2": "ISO 3166-2:FK",
		 "English": " Falkland Islands (Malvinas)",
		 "China": "马尔维纳斯群岛（福克兰）",
		 "Taiwan": "福克蘭群島",
		 "Hongkong": "福克蘭群島（馬爾維納斯）",
		 "Memo": ""
		},
		{
		 "ISO2": "FM",
		 "ISO3": "FSM",
		 "DIGITS": "583",
		 "ISO-3166-2": "ISO 3166-2:FM",
		 "English": " Micronesia (Federated States of)",
		 "China": "密克罗尼西亚联邦",
		 "Taiwan": "密克羅尼西亞聯邦",
		 "Hongkong": "密克羅尼西亞",
		 "Memo": ""
		},
		{
		 "ISO2": "FO",
		 "ISO3": "FRO",
		 "DIGITS": "234",
		 "ISO-3166-2": "ISO 3166-2:FO",
		 "English": " Faroe Islands",
		 "China": "法罗群岛",
		 "Taiwan": "法羅群島",
		 "Hongkong": "法羅群島",
		 "Memo": ""
		},
		{
		 "ISO2": "FR",
		 "ISO3": "FRA",
		 "DIGITS": "250",
		 "ISO-3166-2": "ISO 3166-2:FR",
		 "English": " France",
		 "China": "法国",
		 "Taiwan": "法國",
		 "Hongkong": "法國",
		 "Memo": ""
		},
		{
		 "ISO2": "GA",
		 "ISO3": "GAB",
		 "DIGITS": "266",
		 "ISO-3166-2": "ISO 3166-2:GA",
		 "English": " Gabon",
		 "China": "加蓬",
		 "Taiwan": "加彭",
		 "Hongkong": "加蓬",
		 "Memo": ""
		},
		{
		 "ISO2": "GB",
		 "ISO3": "GBR",
		 "DIGITS": "826",
		 "ISO-3166-2": "ISO 3166-2:GB",
		 "English": " United Kingdom",
		 "China": "英国",
		 "Taiwan": "英國",
		 "Hongkong": "英國",
		 "Memo": "台灣和香港亦普遍採用「聯合王國」一詞於其它場合"
		},
		{
		 "ISO2": "GD",
		 "ISO3": "GRD",
		 "DIGITS": "308",
		 "ISO-3166-2": "ISO 3166-2:GD",
		 "English": " Grenada",
		 "China": "格林纳达",
		 "Taiwan": "格瑞那達",
		 "Hongkong": "格林納達",
		 "Memo": ""
		},
		{
		 "ISO2": "GE",
		 "ISO3": "GEO",
		 "DIGITS": "268",
		 "ISO-3166-2": "ISO 3166-2:GE",
		 "English": " Georgia",
		 "China": "格鲁吉亚",
		 "Taiwan": "喬治亞",
		 "Hongkong": "格魯吉亞",
		 "Memo": ""
		},
		{
		 "ISO2": "GF",
		 "ISO3": "GUF",
		 "DIGITS": "254",
		 "ISO-3166-2": "ISO 3166-2:GF",
		 "English": " French Guiana",
		 "China": "法属圭亚那",
		 "Taiwan": "法屬圭亞那",
		 "Hongkong": "法屬圭亞那",
		 "Memo": ""
		},
		{
		 "ISO2": "GG",
		 "ISO3": "GGY",
		 "DIGITS": "831",
		 "ISO-3166-2": "ISO 3166-2:GG",
		 "English": " Guernsey",
		 "China": "根西岛",
		 "Taiwan": "根息島",
		 "Hongkong": "根西島",
		 "Memo": "中國大陸曾將之普遍譯作「格恩西岛」"
		},
		{
		 "ISO2": "GH",
		 "ISO3": "GHA",
		 "DIGITS": "288",
		 "ISO-3166-2": "ISO 3166-2:GH",
		 "English": " Ghana",
		 "China": "加纳",
		 "Taiwan": "迦納",
		 "Hongkong": "加納",
		 "Memo": ""
		},
		{
		 "ISO2": "GI",
		 "ISO3": "GIB",
		 "DIGITS": "292",
		 "ISO-3166-2": "ISO 3166-2:GI",
		 "English": " Gibraltar",
		 "China": "直布罗陀",
		 "Taiwan": "直布羅陀",
		 "Hongkong": "直布羅陀",
		 "Memo": ""
		},
		{
		 "ISO2": "GL",
		 "ISO3": "GRL",
		 "DIGITS": "304",
		 "ISO-3166-2": "ISO 3166-2:GL",
		 "English": " Greenland",
		 "China": "格陵兰",
		 "Taiwan": "格陵蘭",
		 "Hongkong": "格陵蘭",
		 "Memo": ""
		},
		{
		 "ISO2": "GM",
		 "ISO3": "GMB",
		 "DIGITS": "270",
		 "ISO-3166-2": "ISO 3166-2:GM",
		 "English": " Gambia",
		 "China": "冈比亚",
		 "Taiwan": "甘比亞",
		 "Hongkong": "岡比亞",
		 "Memo": "亦有部份人士使用「剛比亞」一詞於港澳地區"
		},
		{
		 "ISO2": "GN",
		 "ISO3": "GIN",
		 "DIGITS": "324",
		 "ISO-3166-2": "ISO 3166-2:GN",
		 "English": " Guinea",
		 "China": "几内亚",
		 "Taiwan": "幾內亞",
		 "Hongkong": "幾內亞",
		 "Memo": ""
		},
		{
		 "ISO2": "GP",
		 "ISO3": "GLP",
		 "DIGITS": "312",
		 "ISO-3166-2": "ISO 3166-2:GP",
		 "English": " Guadeloupe",
		 "China": "瓜德罗普",
		 "Taiwan": "瓜德魯普島",
		 "Hongkong": "瓜德魯普島",
		 "Memo": ""
		},
		{
		 "ISO2": "GQ",
		 "ISO3": "GNQ",
		 "DIGITS": "226",
		 "ISO-3166-2": "ISO 3166-2:GQ",
		 "English": " Equatorial Guinea",
		 "China": "赤道几内亚",
		 "Taiwan": "赤道幾內亞",
		 "Hongkong": "赤道幾內亞",
		 "Memo": ""
		},
		{
		 "ISO2": "GR",
		 "ISO3": "GRC",
		 "DIGITS": "300",
		 "ISO-3166-2": "ISO 3166-2:GR",
		 "English": " Greece",
		 "China": "希腊",
		 "Taiwan": "希臘",
		 "Hongkong": "希臘",
		 "Memo": ""
		},
		{
		 "ISO2": "GS",
		 "ISO3": "SGS",
		 "DIGITS": "239",
		 "ISO-3166-2": "ISO 3166-2:GS",
		 "English": " South Georgia and the South Sandwich Islands",
		 "China": "南乔治亚岛和南桑威奇群岛",
		 "Taiwan": "南喬治亞與南三明治群島",
		 "Hongkong": "南喬治亞島與南桑威奇群島",
		 "Memo": ""
		},
		{
		 "ISO2": "GT",
		 "ISO3": "GTM",
		 "DIGITS": "320",
		 "ISO-3166-2": "ISO 3166-2:GT",
		 "English": " Guatemala",
		 "China": "危地马拉",
		 "Taiwan": "瓜地馬拉",
		 "Hongkong": "危地馬拉",
		 "Memo": ""
		},
		{
		 "ISO2": "GU",
		 "ISO3": "GUM",
		 "DIGITS": "316",
		 "ISO-3166-2": "ISO 3166-2:GU",
		 "English": " Guam",
		 "China": "关岛",
		 "Taiwan": "關島",
		 "Hongkong": "關島",
		 "Memo": ""
		},
		{
		 "ISO2": "GW",
		 "ISO3": "GNB",
		 "DIGITS": "624",
		 "ISO-3166-2": "ISO 3166-2:GW",
		 "English": " Guinea-Bissau",
		 "China": "几内亚比绍",
		 "Taiwan": "幾內亞比索",
		 "Hongkong": "幾內亞比紹",
		 "Memo": ""
		},
		{
		 "ISO2": "GY",
		 "ISO3": "GUY",
		 "DIGITS": "328",
		 "ISO-3166-2": "ISO 3166-2:GY",
		 "English": " Guyana",
		 "China": "圭亚那",
		 "Taiwan": "蓋亞那",
		 "Hongkong": "圭亞那",
		 "Memo": ""
		},
		{
		 "ISO2": "HK",
		 "ISO3": "HKG",
		 "DIGITS": "344",
		 "ISO-3166-2": "ISO 3166-2:HK",
		 "English": " Hong Kong",
		 "China": "香港",
		 "Taiwan": "香港",
		 "Hongkong": "香港",
		 "Memo": ""
		},
		{
		 "ISO2": "HM",
		 "ISO3": "HMD",
		 "DIGITS": "334",
		 "ISO-3166-2": "ISO 3166-2:HM",
		 "English": " Heard Island and McDonald Islands",
		 "China": "赫德岛和麦克唐纳群岛",
		 "Taiwan": "赫德及麥當勞群島",
		 "Hongkong": "赫德島和麥克唐納群島",
		 "Memo": ""
		},
		{
		 "ISO2": "HN",
		 "ISO3": "HND",
		 "DIGITS": "340",
		 "ISO-3166-2": "ISO 3166-2:HN",
		 "English": " Honduras",
		 "China": "洪都拉斯",
		 "Taiwan": "宏都拉斯",
		 "Hongkong": "宏都拉斯",
		 "Memo": ""
		},
		{
		 "ISO2": "HR",
		 "ISO3": "HRV",
		 "DIGITS": "191",
		 "ISO-3166-2": "ISO 3166-2:HR",
		 "English": " Croatia",
		 "China": "克罗地亚",
		 "Taiwan": "克羅地亞",
		 "Hongkong": "克羅地亞",
		 "Memo": ""
		},
		{
		 "ISO2": "HT",
		 "ISO3": "HTI",
		 "DIGITS": "332",
		 "ISO-3166-2": "ISO 3166-2:HT",
		 "English": " Haiti",
		 "China": "海地",
		 "Taiwan": "海地",
		 "Hongkong": "海地",
		 "Memo": ""
		},
		{
		 "ISO2": "HU",
		 "ISO3": "HUN",
		 "DIGITS": "348",
		 "ISO-3166-2": "ISO 3166-2:HU",
		 "English": " Hungary",
		 "China": "匈牙利",
		 "Taiwan": "匈牙利",
		 "Hongkong": "匈牙利",
		 "Memo": ""
		},
		{
		 "ISO2": "ID",
		 "ISO3": "IDN",
		 "DIGITS": "360",
		 "ISO-3166-2": "ISO 3166-2:ID",
		 "English": " Indonesia",
		 "China": "印尼",
		 "Taiwan": "印尼",
		 "Hongkong": "印尼",
		 "Memo": ""
		},
		{
		 "ISO2": "IE",
		 "ISO3": "IRL",
		 "DIGITS": "372",
		 "ISO-3166-2": "ISO 3166-2:IE",
		 "English": " Ireland",
		 "China": "爱尔兰",
		 "Taiwan": "愛爾蘭",
		 "Hongkong": "愛爾蘭",
		 "Memo": ""
		},
		{
		 "ISO2": "IL",
		 "ISO3": "ISR",
		 "DIGITS": "376",
		 "ISO-3166-2": "ISO 3166-2:IL",
		 "English": " Israel",
		 "China": "以色列",
		 "Taiwan": "以色列",
		 "Hongkong": "以色列",
		 "Memo": ""
		},
		{
		 "ISO2": "IM",
		 "ISO3": "IMN",
		 "DIGITS": "833",
		 "ISO-3166-2": "ISO 3166-2:IM",
		 "English": " Isle of Man",
		 "China": "马恩岛",
		 "Taiwan": "曼島",
		 "Hongkong": "萌島",
		 "Memo": ""
		},
		{
		 "ISO2": "IN",
		 "ISO3": "IND",
		 "DIGITS": "356",
		 "ISO-3166-2": "ISO 3166-2:IN",
		 "English": " India",
		 "China": "印度",
		 "Taiwan": "印度",
		 "Hongkong": "印度",
		 "Memo": ""
		},
		{
		 "ISO2": "IO",
		 "ISO3": "IOT",
		 "DIGITS": "86",
		 "ISO-3166-2": "ISO 3166-2:IO",
		 "English": " British Indian Ocean Territory",
		 "China": "英属印度洋领地",
		 "Taiwan": "英屬印度洋地區",
		 "Hongkong": "英屬印度洋地區",
		 "Memo": ""
		},
		{
		 "ISO2": "IQ",
		 "ISO3": "IRQ",
		 "DIGITS": "368",
		 "ISO-3166-2": "ISO 3166-2:IQ",
		 "English": " Iraq",
		 "China": "伊拉克",
		 "Taiwan": "伊拉克",
		 "Hongkong": "伊拉克",
		 "Memo": ""
		},
		{
		 "ISO2": "IR",
		 "ISO3": "IRN",
		 "DIGITS": "364",
		 "ISO-3166-2": "ISO 3166-2:IR",
		 "English": " Iran (Islamic Republic of)",
		 "China": "伊朗",
		 "Taiwan": "伊朗",
		 "Hongkong": "伊朗",
		 "Memo": ""
		},
		{
		 "ISO2": "IS",
		 "ISO3": "ISL",
		 "DIGITS": "352",
		 "ISO-3166-2": "ISO 3166-2:IS",
		 "English": " Iceland",
		 "China": "冰岛",
		 "Taiwan": "冰島",
		 "Hongkong": "冰島",
		 "Memo": ""
		},
		{
		 "ISO2": "IT",
		 "ISO3": "ITA",
		 "DIGITS": "380",
		 "ISO-3166-2": "ISO 3166-2:IT",
		 "English": " Italy",
		 "China": "意大利",
		 "Taiwan": "義大利",
		 "Hongkong": "意大利",
		 "Memo": ""
		},
		{
		 "ISO2": "JE",
		 "ISO3": "JEY",
		 "DIGITS": "832",
		 "ISO-3166-2": "ISO 3166-2:JE",
		 "English": " Jersey",
		 "China": "泽西岛",
		 "Taiwan": "澤西島",
		 "Hongkong": "澤西",
		 "Memo": ""
		},
		{
		 "ISO2": "JM",
		 "ISO3": "JAM",
		 "DIGITS": "388",
		 "ISO-3166-2": "ISO 3166-2:JM",
		 "English": " Jamaica",
		 "China": "牙买加",
		 "Taiwan": "牙買加",
		 "Hongkong": "牙買加",
		 "Memo": ""
		},
		{
		 "ISO2": "JO",
		 "ISO3": "JOR",
		 "DIGITS": "400",
		 "ISO-3166-2": "ISO 3166-2:JO",
		 "English": " Jordan",
		 "China": "约旦",
		 "Taiwan": "約旦",
		 "Hongkong": "約旦",
		 "Memo": ""
		},
		{
		 "ISO2": "JP",
		 "ISO3": "JPN",
		 "DIGITS": "392",
		 "ISO-3166-2": "ISO 3166-2:JP",
		 "English": " Japan",
		 "China": "日本",
		 "Taiwan": "日本",
		 "Hongkong": "日本",
		 "Memo": ""
		},
		{
		 "ISO2": "KE",
		 "ISO3": "KEN",
		 "DIGITS": "404",
		 "ISO-3166-2": "ISO 3166-2:KE",
		 "English": " Kenya",
		 "China": "肯尼亚",
		 "Taiwan": "肯亞",
		 "Hongkong": "肯尼亞",
		 "Memo": "香港亦普遍採用「肯雅」一詞於其它場合"
		},
		{
		 "ISO2": "KG",
		 "ISO3": "KGZ",
		 "DIGITS": "417",
		 "ISO-3166-2": "ISO 3166-2:KG",
		 "English": " Kyrgyzstan",
		 "China": "吉尔吉斯斯坦",
		 "Taiwan": "吉爾吉斯",
		 "Hongkong": "吉爾吉斯",
		 "Memo": "香港習慣略去「斯坦」後綴，有必要會用全稱"
		},
		{
		 "ISO2": "KH",
		 "ISO3": "KHM",
		 "DIGITS": "116",
		 "ISO-3166-2": "ISO 3166-2:KH",
		 "English": " Cambodia",
		 "China": "柬埔寨",
		 "Taiwan": "柬埔寨",
		 "Hongkong": "柬埔寨",
		 "Memo": ""
		},
		{
		 "ISO2": "KI",
		 "ISO3": "KIR",
		 "DIGITS": "296",
		 "ISO-3166-2": "ISO 3166-2:KI",
		 "English": " Kiribati",
		 "China": "基里巴斯",
		 "Taiwan": "吉里巴斯",
		 "Hongkong": "基里巴斯",
		 "Memo": ""
		},
		{
		 "ISO2": "KM",
		 "ISO3": "COM",
		 "DIGITS": "174",
		 "ISO-3166-2": "ISO 3166-2:KM",
		 "English": " Comoros",
		 "China": "科摩罗",
		 "Taiwan": "葛摩",
		 "Hongkong": "科摩羅",
		 "Memo": ""
		},
		{
		 "ISO2": "KN",
		 "ISO3": "KNA",
		 "DIGITS": "659",
		 "ISO-3166-2": "ISO 3166-2:KN",
		 "English": " Saint Kitts and Nevis",
		 "China": "圣基茨和尼维斯",
		 "Taiwan": "聖克里斯多福及尼維斯",
		 "Hongkong": "聖基茨和尼維斯",
		 "Memo": "香港亦普遍採用「聖克里斯托佛島及尼維斯島」一詞於其它場合（如香港郵政的郵政指南附錄表）。亦有部份人士使用「聖吉斯納域斯」一詞於港澳地區"
		},
		{
		 "ISO2": "KP",
		 "ISO3": "PRK",
		 "DIGITS": "408",
		 "ISO-3166-2": "ISO 3166-2:KP",
		 "English": " Korea (Democratic People's Republic of)",
		 "China": "朝鲜(北朝鲜)",
		 "Taiwan": "北韓",
		 "Hongkong": "朝鲜(北韓)",
		 "Memo": "澳門習慣稱之為「北韓」"
		},
		{
		 "ISO2": "KR",
		 "ISO3": "KOR",
		 "DIGITS": "410",
		 "ISO-3166-2": "ISO 3166-2:KR",
		 "English": " Korea (Republic of)",
		 "China": "韩国(南朝鲜)",
		 "Taiwan": "韓國",
		 "Hongkong": "韓國(南韓)",
		 "Memo": "澳門習慣稱之為「南韓」"
		},
		{
		 "ISO2": "KW",
		 "ISO3": "KWT",
		 "DIGITS": "414",
		 "ISO-3166-2": "ISO 3166-2:KW",
		 "English": " Kuwait",
		 "China": "科威特",
		 "Taiwan": "科威特",
		 "Hongkong": "科威特",
		 "Memo": ""
		},
		{
		 "ISO2": "KY",
		 "ISO3": "CYM",
		 "DIGITS": "136",
		 "ISO-3166-2": "ISO 3166-2:KY",
		 "English": " Cayman Islands",
		 "China": "开曼群岛",
		 "Taiwan": "開曼群島",
		 "Hongkong": "開曼群島",
		 "Memo": ""
		},
		{
		 "ISO2": "KZ",
		 "ISO3": "KAZ",
		 "DIGITS": "398",
		 "ISO-3166-2": "ISO 3166-2:KZ",
		 "English": " Kazakhstan",
		 "China": "哈萨克斯坦",
		 "Taiwan": "哈薩克",
		 "Hongkong": "哈薩克",
		 "Memo": "香港習慣略去「斯坦」後綴，有必要會用全稱"
		},
		{
		 "ISO2": "LA",
		 "ISO3": "LAO",
		 "DIGITS": "418",
		 "ISO-3166-2": "ISO 3166-2:LA",
		 "English": " Lao People's Democratic Republic",
		 "China": "老挝",
		 "Taiwan": "寮國",
		 "Hongkong": "老挝",
		 "Memo": "新加坡與馬來西亞均將之譯作「寮国」"
		},
		{
		 "ISO2": "LB",
		 "ISO3": "LBN",
		 "DIGITS": "422",
		 "ISO-3166-2": "ISO 3166-2:LB",
		 "English": " Lebanon",
		 "China": "黎巴嫩",
		 "Taiwan": "黎巴嫩",
		 "Hongkong": "黎巴嫩",
		 "Memo": ""
		},
		{
		 "ISO2": "LC",
		 "ISO3": "LCA",
		 "DIGITS": "662",
		 "ISO-3166-2": "ISO 3166-2:LC",
		 "English": " Saint Lucia",
		 "China": "圣卢西亚",
		 "Taiwan": "聖露西亞",
		 "Hongkong": "聖盧西亞",
		 "Memo": "香港亦普遍採用「聖路西亞」一詞於其它場合"
		},
		{
		 "ISO2": "LI",
		 "ISO3": "LIE",
		 "DIGITS": "438",
		 "ISO-3166-2": "ISO 3166-2:LI",
		 "English": " Liechtenstein",
		 "China": "列支敦士登",
		 "Taiwan": "列支敦斯登",
		 "Hongkong": "列支敦士登",
		 "Memo": ""
		},
		{
		 "ISO2": "LK",
		 "ISO3": "LKA",
		 "DIGITS": "144",
		 "ISO-3166-2": "ISO 3166-2:LK",
		 "English": " Sri Lanka",
		 "China": "斯里兰卡",
		 "Taiwan": "斯里蘭卡",
		 "Hongkong": "斯里蘭卡",
		 "Memo": ""
		},
		{
		 "ISO2": "LR",
		 "ISO3": "LBR",
		 "DIGITS": "430",
		 "ISO-3166-2": "ISO 3166-2:LR",
		 "English": " Liberia",
		 "China": "利比里亚",
		 "Taiwan": "賴比瑞亞",
		 "Hongkong": "利比里亞",
		 "Memo": ""
		},
		{
		 "ISO2": "LS",
		 "ISO3": "LSO",
		 "DIGITS": "426",
		 "ISO-3166-2": "ISO 3166-2:LS",
		 "English": " Lesotho",
		 "China": "莱索托",
		 "Taiwan": "賴索托",
		 "Hongkong": "萊索托",
		 "Memo": ""
		},
		{
		 "ISO2": "LT",
		 "ISO3": "LTU",
		 "DIGITS": "440",
		 "ISO-3166-2": "ISO 3166-2:LT",
		 "English": " Lithuania",
		 "China": "立陶宛",
		 "Taiwan": "立陶宛",
		 "Hongkong": "立陶宛",
		 "Memo": ""
		},
		{
		 "ISO2": "LU",
		 "ISO3": "LUX",
		 "DIGITS": "442",
		 "ISO-3166-2": "ISO 3166-2:LU",
		 "English": " Luxembourg",
		 "China": "卢森堡",
		 "Taiwan": "盧森堡",
		 "Hongkong": "盧森堡",
		 "Memo": ""
		},
		{
		 "ISO2": "LV",
		 "ISO3": "LVA",
		 "DIGITS": "428",
		 "ISO-3166-2": "ISO 3166-2:LV",
		 "English": " Latvia",
		 "China": "拉脱维亚",
		 "Taiwan": "拉脫維亞",
		 "Hongkong": "拉脫維亞",
		 "Memo": ""
		},
		{
		 "ISO2": "LY",
		 "ISO3": "LBY",
		 "DIGITS": "434",
		 "ISO-3166-2": "ISO 3166-2:LY",
		 "English": " Libya",
		 "China": "利比亚",
		 "Taiwan": "利比亞",
		 "Hongkong": "利比亞",
		 "Memo": ""
		},
		{
		 "ISO2": "MA",
		 "ISO3": "MAR",
		 "DIGITS": "504",
		 "ISO-3166-2": "ISO 3166-2:MA",
		 "English": " Morocco",
		 "China": "摩洛哥",
		 "Taiwan": "摩洛哥",
		 "Hongkong": "摩洛哥",
		 "Memo": ""
		},
		{
		 "ISO2": "MC",
		 "ISO3": "MCO",
		 "DIGITS": "492",
		 "ISO-3166-2": "ISO 3166-2:MC",
		 "English": " Monaco",
		 "China": "摩纳哥",
		 "Taiwan": "摩納哥",
		 "Hongkong": "摩納哥",
		 "Memo": ""
		},
		{
		 "ISO2": "MD",
		 "ISO3": "MDA",
		 "DIGITS": "498",
		 "ISO-3166-2": "ISO 3166-2:MD",
		 "English": " Moldova (Republic of)",
		 "China": "摩尔多瓦",
		 "Taiwan": "摩爾多瓦",
		 "Hongkong": "摩爾多瓦",
		 "Memo": ""
		},
		{
		 "ISO2": "ME",
		 "ISO3": "MNE",
		 "DIGITS": "499",
		 "ISO-3166-2": "ISO 3166-2:ME",
		 "English": " Montenegro",
		 "China": "黑山",
		 "Taiwan": "蒙特內哥羅",
		 "Hongkong": "黑山",
		 "Memo": ""
		},
		{
		 "ISO2": "MF",
		 "ISO3": "MAF",
		 "DIGITS": "663",
		 "ISO-3166-2": "ISO 3166-2:MF",
		 "English": " Saint Martin (French part)",
		 "China": "法属圣马丁",
		 "Taiwan": "法屬聖馬丁　",
		 "Hongkong": "法屬聖馬丁",
		 "Memo": ""
		},
		{
		 "ISO2": "MG",
		 "ISO3": "MDG",
		 "DIGITS": "450",
		 "ISO-3166-2": "ISO 3166-2:MG",
		 "English": " Madagascar",
		 "China": "马达加斯加",
		 "Taiwan": "馬達加斯加",
		 "Hongkong": "馬達加斯加",
		 "Memo": ""
		},
		{
		 "ISO2": "MH",
		 "ISO3": "MHL",
		 "DIGITS": "584",
		 "ISO-3166-2": "ISO 3166-2:MH",
		 "English": " Marshall islands",
		 "China": "马绍尔群岛",
		 "Taiwan": "馬紹爾群島",
		 "Hongkong": "馬紹爾群島",
		 "Memo": ""
		},
		{
		 "ISO2": "MK",
		 "ISO3": "MKD",
		 "DIGITS": "807",
		 "ISO-3166-2": "ISO 3166-2:MK",
		 "English": " Macedonia (the former Yugoslav Republic of)",
		 "China": "马其顿",
		 "Taiwan": "馬其頓",
		 "Hongkong": "馬其頓",
		 "Memo": "因历史、政治原因，ISO官方称之为“Macedonia (the former Yugoslav Republic of)”，即“前南斯拉夫马其顿共和国”"
		},
		{
		 "ISO2": "ML",
		 "ISO3": "MLI",
		 "DIGITS": "466",
		 "ISO-3166-2": "ISO 3166-2:ML",
		 "English": " Mali",
		 "China": "马里",
		 "Taiwan": "馬利",
		 "Hongkong": "馬里",
		 "Memo": ""
		},
		{
		 "ISO2": "MM",
		 "ISO3": "MMR",
		 "DIGITS": "104",
		 "ISO-3166-2": "ISO 3166-2:MM",
		 "English": " Myanmar",
		 "China": "缅甸",
		 "Taiwan": "緬甸",
		 "Hongkong": "緬甸",
		 "Memo": ""
		},
		{
		 "ISO2": "MN",
		 "ISO3": "MNG",
		 "DIGITS": "496",
		 "ISO-3166-2": "ISO 3166-2:MN",
		 "English": " Mongolia",
		 "China": "蒙古国；蒙古",
		 "Taiwan": "蒙古",
		 "Hongkong": "蒙古國",
		 "Memo": "香港亦普遍採用「蒙古」一詞於其它場合"
		},
		{
		 "ISO2": "MO",
		 "ISO3": "MAC",
		 "DIGITS": "446",
		 "ISO-3166-2": "ISO 3166-2:MO",
		 "English": " Macao",
		 "China": "澳门",
		 "Taiwan": "澳門",
		 "Hongkong": "澳門",
		 "Memo": ""
		},
		{
		 "ISO2": "MP",
		 "ISO3": "MNP",
		 "DIGITS": "580",
		 "ISO-3166-2": "ISO 3166-2:MP",
		 "English": " Northern Mariana Islands",
		 "China": "北马里亚纳群岛",
		 "Taiwan": "北馬里亞納群島",
		 "Hongkong": "北馬里亞納群島",
		 "Memo": "亦有部份人士使用「北瑪利安娜群島」一詞於港澳地區"
		},
		{
		 "ISO2": "MQ",
		 "ISO3": "MTQ",
		 "DIGITS": "474",
		 "ISO-3166-2": "ISO 3166-2:MQ",
		 "English": " Martinique",
		 "China": "马提尼克",
		 "Taiwan": "法屬馬丁尼克",
		 "Hongkong": "馬提尼克",
		 "Memo": ""
		},
		{
		 "ISO2": "MR",
		 "ISO3": "MRT",
		 "DIGITS": "478",
		 "ISO-3166-2": "ISO 3166-2:MR",
		 "English": " Mauritania",
		 "China": "毛里塔尼亚",
		 "Taiwan": "茅利塔尼亞",
		 "Hongkong": "毛里塔尼亞",
		 "Memo": ""
		},
		{
		 "ISO2": "MS",
		 "ISO3": "MSR",
		 "DIGITS": "500",
		 "ISO-3166-2": "ISO 3166-2:MS",
		 "English": " Montserrat",
		 "China": "蒙塞拉特岛",
		 "Taiwan": "蒙塞拉特島",
		 "Hongkong": "蒙塞拉特島",
		 "Memo": ""
		},
		{
		 "ISO2": "MT",
		 "ISO3": "MLT",
		 "DIGITS": "470",
		 "ISO-3166-2": "ISO 3166-2:MT",
		 "English": " Malta",
		 "China": "马耳他",
		 "Taiwan": "馬耳他",
		 "Hongkong": "馬爾他",
		 "Memo": ""
		},
		{
		 "ISO2": "MU",
		 "ISO3": "MUS",
		 "DIGITS": "480",
		 "ISO-3166-2": "ISO 3166-2:MU",
		 "English": " Mauritius",
		 "China": "毛里求斯",
		 "Taiwan": "模里西斯",
		 "Hongkong": "毛里求斯",
		 "Memo": "香港亦普遍採用「毛里裘斯」一詞於其它場合"
		},
		{
		 "ISO2": "MV",
		 "ISO3": "MDV",
		 "DIGITS": "462",
		 "ISO-3166-2": "ISO 3166-2:MV",
		 "English": " Maldives",
		 "China": "马尔代夫",
		 "Taiwan": "馬爾地夫",
		 "Hongkong": "馬爾代夫",
		 "Memo": ""
		},
		{
		 "ISO2": "MW",
		 "ISO3": "MWI",
		 "DIGITS": "454",
		 "ISO-3166-2": "ISO 3166-2:MW",
		 "English": " Malawi",
		 "China": "马拉维",
		 "Taiwan": "馬拉威",
		 "Hongkong": "馬拉維",
		 "Memo": ""
		},
		{
		 "ISO2": "MX",
		 "ISO3": "MEX",
		 "DIGITS": "484",
		 "ISO-3166-2": "ISO 3166-2:MX",
		 "English": " Mexico",
		 "China": "墨西哥",
		 "Taiwan": "墨西哥",
		 "Hongkong": "墨西哥",
		 "Memo": ""
		},
		{
		 "ISO2": "MY",
		 "ISO3": "MYS",
		 "DIGITS": "458",
		 "ISO-3166-2": "ISO 3166-2:MY",
		 "English": " Malaysia",
		 "China": "马来西亚",
		 "Taiwan": "馬來西亞",
		 "Hongkong": "馬來西亞",
		 "Memo": ""
		},
		{
		 "ISO2": "MZ",
		 "ISO3": "MOZ",
		 "DIGITS": "508",
		 "ISO-3166-2": "ISO 3166-2:MZ",
		 "English": " Mozambique",
		 "China": "莫桑比克",
		 "Taiwan": "莫三比克",
		 "Hongkong": "莫桑比克",
		 "Memo": "亦有譯作「莫三鼻給」"
		},
		{
		 "ISO2": "NA",
		 "ISO3": "NAM",
		 "DIGITS": "516",
		 "ISO-3166-2": "ISO 3166-2:NA",
		 "English": " Namibia",
		 "China": "纳米比亚",
		 "Taiwan": "奈米比亞",
		 "Hongkong": "納米比亞",
		 "Memo": ""
		},
		{
		 "ISO2": "NC",
		 "ISO3": "NCL",
		 "DIGITS": "540",
		 "ISO-3166-2": "ISO 3166-2:NC",
		 "English": " New Caledonia",
		 "China": "新喀里多尼亚",
		 "Taiwan": "新喀里多尼亞島",
		 "Hongkong": "新喀里多尼亞",
		 "Memo": "亦有部份人士使用「新喀爾多尼亞」一詞於港澳地區"
		},
		{
		 "ISO2": "NE",
		 "ISO3": "NER",
		 "DIGITS": "562",
		 "ISO-3166-2": "ISO 3166-2:NE",
		 "English": " Niger",
		 "China": "尼日尔",
		 "Taiwan": "尼日",
		 "Hongkong": "尼日爾",
		 "Memo": ""
		},
		{
		 "ISO2": "NF",
		 "ISO3": "NFK",
		 "DIGITS": "574",
		 "ISO-3166-2": "ISO 3166-2:NF",
		 "English": " Norfolk Island",
		 "China": "诺福克岛",
		 "Taiwan": "諾福克島",
		 "Hongkong": "諾福克島",
		 "Memo": ""
		},
		{
		 "ISO2": "NG",
		 "ISO3": "NGA",
		 "DIGITS": "566",
		 "ISO-3166-2": "ISO 3166-2:NG",
		 "English": " Nigeria",
		 "China": "尼日利亚",
		 "Taiwan": "奈及利亞",
		 "Hongkong": "尼日利亞",
		 "Memo": ""
		},
		{
		 "ISO2": "NI",
		 "ISO3": "NIC",
		 "DIGITS": "558",
		 "ISO-3166-2": "ISO 3166-2:NI",
		 "English": " Nicaragua",
		 "China": "尼加拉瓜",
		 "Taiwan": "尼加拉瓜",
		 "Hongkong": "尼加拉瓜",
		 "Memo": ""
		},
		{
		 "ISO2": "NK",
		 "ISO3": "NKR",
		 "DIGITS": "296",
		 "ISO-3166-2": "ISO 3166-2:NK",
		 "English": " Nagorno-Karabakh",
		 "China": "纳戈尔诺-卡拉巴赫",
		 "Taiwan": "纳戈尔诺-卡拉巴赫",
		 "Hongkong": "纳戈尔诺-卡拉巴赫",
		 "Memo": "从未使用"
		},
		{
		 "ISO2": "NL",
		 "ISO3": "NLD",
		 "DIGITS": "528",
		 "ISO-3166-2": "ISO 3166-2:NL",
		 "English": " Netherlands",
		 "China": "荷兰",
		 "Taiwan": "荷蘭",
		 "Hongkong": "荷蘭",
		 "Memo": ""
		},
		{
		 "ISO2": "NO",
		 "ISO3": "NOR",
		 "DIGITS": "578",
		 "ISO-3166-2": "ISO 3166-2:NO",
		 "English": " Norway",
		 "China": "挪威",
		 "Taiwan": "挪威",
		 "Hongkong": "挪威",
		 "Memo": ""
		},
		{
		 "ISO2": "NP",
		 "ISO3": "NPL",
		 "DIGITS": "524",
		 "ISO-3166-2": "ISO 3166-2:NP",
		 "English": " Nepal",
		 "China": "尼泊尔",
		 "Taiwan": "尼泊爾",
		 "Hongkong": "尼泊爾",
		 "Memo": ""
		},
		{
		 "ISO2": "NR",
		 "ISO3": "NRU",
		 "DIGITS": "520",
		 "ISO-3166-2": "ISO 3166-2:NR",
		 "English": " Nauru",
		 "China": "瑙鲁",
		 "Taiwan": "諾魯",
		 "Hongkong": "瑙魯",
		 "Memo": ""
		},
		{
		 "ISO2": "NU",
		 "ISO3": "NIU",
		 "DIGITS": "570",
		 "ISO-3166-2": "ISO 3166-2:NU",
		 "English": " Niue",
		 "China": "纽埃",
		 "Taiwan": "紐埃",
		 "Hongkong": "紐埃",
		 "Memo": "台灣亦普遍採用「紐威島」（CNS 12842譯名）一詞於其它場合（如MSN台灣）"
		},
		{
		 "ISO2": "NZ",
		 "ISO3": "NZL",
		 "DIGITS": "554",
		 "ISO-3166-2": "ISO 3166-2:NZ",
		 "English": " New Zealand",
		 "China": "新西兰",
		 "Taiwan": "紐西蘭",
		 "Hongkong": "新西蘭",
		 "Memo": "新加坡與馬來西亞均將之譯作「纽西兰」。香港亦普遍採用「紐西蘭」一詞於其它場合"
		},
		{
		 "ISO2": "OM",
		 "ISO3": "OMN",
		 "DIGITS": "512",
		 "ISO-3166-2": "ISO 3166-2:OM",
		 "English": " Oman",
		 "China": "阿曼",
		 "Taiwan": "阿曼",
		 "Hongkong": "阿曼",
		 "Memo": ""
		},
		{
		 "ISO2": "PA",
		 "ISO3": "PAN",
		 "DIGITS": "591",
		 "ISO-3166-2": "ISO 3166-2:PA",
		 "English": " Panama",
		 "China": "巴拿马",
		 "Taiwan": "巴拿馬",
		 "Hongkong": "巴拿馬",
		 "Memo": ""
		},
		{
		 "ISO2": "PE",
		 "ISO3": "PER",
		 "DIGITS": "604",
		 "ISO-3166-2": "ISO 3166-2:PE",
		 "English": " Peru",
		 "China": "秘鲁",
		 "Taiwan": "秘魯",
		 "Hongkong": "秘魯",
		 "Memo": ""
		},
		{
		 "ISO2": "PF",
		 "ISO3": "PYF",
		 "DIGITS": "258",
		 "ISO-3166-2": "ISO 3166-2:PF",
		 "English": " French Polynesia",
		 "China": "法属波利尼西亚",
		 "Taiwan": "法屬玻里尼西亞",
		 "Hongkong": "法屬波利尼西亞",
		 "Memo": ""
		},
		{
		 "ISO2": "PG",
		 "ISO3": "PNG",
		 "DIGITS": "598",
		 "ISO-3166-2": "ISO 3166-2:PG",
		 "English": " Papua New Guinea",
		 "China": "巴布亚新几内亚",
		 "Taiwan": "巴布亞新幾內亞",
		 "Hongkong": "巴布亞新幾內亞",
		 "Memo": ""
		},
		{
		 "ISO2": "PH",
		 "ISO3": "PHL",
		 "DIGITS": "608",
		 "ISO-3166-2": "ISO 3166-2:PH",
		 "English": " Philippines",
		 "China": "菲律宾",
		 "Taiwan": "菲律賓",
		 "Hongkong": "菲律賓",
		 "Memo": ""
		},
		{
		 "ISO2": "PK",
		 "ISO3": "PAK",
		 "DIGITS": "586",
		 "ISO-3166-2": "ISO 3166-2:PK",
		 "English": " Pakistan",
		 "China": "巴基斯坦",
		 "Taiwan": "巴基斯坦",
		 "Hongkong": "巴基斯坦",
		 "Memo": ""
		},
		{
		 "ISO2": "PL",
		 "ISO3": "POL",
		 "DIGITS": "616",
		 "ISO-3166-2": "ISO 3166-2:PL",
		 "English": " Poland",
		 "China": "波兰",
		 "Taiwan": "波蘭",
		 "Hongkong": "波蘭",
		 "Memo": ""
		},
		{
		 "ISO2": "PM",
		 "ISO3": "SPM",
		 "DIGITS": "666",
		 "ISO-3166-2": "ISO 3166-2:PM",
		 "English": " Saint Pierre and Miquelon",
		 "China": "圣皮埃尔和密克隆",
		 "Taiwan": "聖皮耶與密克隆群島",
		 "Hongkong": "聖皮埃爾島和密克隆島",
		 "Memo": "香港亦普遍採用「聖皮埃蘭和密克隆群島」一詞於其它場合（如香港郵政的郵政指南附錄表）"
		},
		{
		 "ISO2": "PN",
		 "ISO3": "PCN",
		 "DIGITS": "612",
		 "ISO-3166-2": "ISO 3166-2:PN",
		 "English": " Pitcairn Islands",
		 "China": "皮特凯恩群岛",
		 "Taiwan": "皮特康島",
		 "Hongkong": "皮特凱恩群島",
		 "Memo": ""
		},
		{
		 "ISO2": "PR",
		 "ISO3": "PRI",
		 "DIGITS": "630",
		 "ISO-3166-2": "ISO 3166-2:PR",
		 "English": " Puerto Rico",
		 "China": "波多黎各",
		 "Taiwan": "波多黎各",
		 "Hongkong": "波多黎各",
		 "Memo": ""
		},
		{
		 "ISO2": "PS",
		 "ISO3": "PSE",
		 "DIGITS": "275",
		 "ISO-3166-2": "ISO 3166-2:PS",
		 "English": " Palestine, State of",
		 "China": "巴勒斯坦",
		 "Taiwan": "巴勒斯坦",
		 "Hongkong": "巴勒斯坦",
		 "Memo": ""
		},
		{
		 "ISO2": "PT",
		 "ISO3": "PRT",
		 "DIGITS": "620",
		 "ISO-3166-2": "ISO 3166-2:PT",
		 "English": " Portugal",
		 "China": "葡萄牙",
		 "Taiwan": "葡萄牙",
		 "Hongkong": "葡萄牙",
		 "Memo": "澳門民間亦普遍稱之為葡國"
		},
		{
		 "ISO2": "PW",
		 "ISO3": "PLW",
		 "DIGITS": "585",
		 "ISO-3166-2": "ISO 3166-2:PW",
		 "English": " Palau",
		 "China": "帕劳",
		 "Taiwan": "帛琉",
		 "Hongkong": "帕勞；帛琉",
		 "Memo": ""
		},
		{
		 "ISO2": "PY",
		 "ISO3": "PRY",
		 "DIGITS": "600",
		 "ISO-3166-2": "ISO 3166-2:PY",
		 "English": " Paraguay",
		 "China": "巴拉圭",
		 "Taiwan": "巴拉圭",
		 "Hongkong": "巴拉圭",
		 "Memo": ""
		},
		{
		 "ISO2": "QA",
		 "ISO3": "QAT",
		 "DIGITS": "634",
		 "ISO-3166-2": "ISO 3166-2:QA",
		 "English": " Qatar",
		 "China": "卡塔尔",
		 "Taiwan": "卡達",
		 "Hongkong": "卡塔爾",
		 "Memo": ""
		},
		{
		 "ISO2": "RE",
		 "ISO3": "REU",
		 "DIGITS": "638",
		 "ISO-3166-2": "ISO 3166-2:RE",
		 "English": " Réunion",
		 "China": "留尼汪",
		 "Taiwan": "留尼旺",
		 "Hongkong": "留尼汪",
		 "Memo": ""
		},
		{
		 "ISO2": "RO",
		 "ISO3": "ROU",
		 "DIGITS": "642",
		 "ISO-3166-2": "ISO 3166-2:RO",
		 "English": " Romania",
		 "China": "罗马尼亚",
		 "Taiwan": "羅馬尼亞",
		 "Hongkong": "羅馬尼亞",
		 "Memo": ""
		},
		{
		 "ISO2": "RS",
		 "ISO3": "SRB",
		 "DIGITS": "688",
		 "ISO-3166-2": "ISO 3166-2:RS",
		 "English": " Serbia",
		 "China": "塞尔维亚",
		 "Taiwan": "塞爾維亞",
		 "Hongkong": "塞爾維亞",
		 "Memo": ""
		},
		{
		 "ISO2": "RU",
		 "ISO3": "RUS",
		 "DIGITS": "643",
		 "ISO-3166-2": "ISO 3166-2:RU",
		 "English": " Russian Federation",
		 "China": "俄罗斯",
		 "Taiwan": "俄羅斯",
		 "Hongkong": "俄羅斯",
		 "Memo": ""
		},
		{
		 "ISO2": "RW",
		 "ISO3": "RWA",
		 "DIGITS": "646",
		 "ISO-3166-2": "ISO 3166-2:RW",
		 "English": " Rwanda",
		 "China": "卢旺达",
		 "Taiwan": "盧安達",
		 "Hongkong": "盧旺達",
		 "Memo": ""
		},
		{
		 "ISO2": "SA",
		 "ISO3": "SAU",
		 "DIGITS": "682",
		 "ISO-3166-2": "ISO 3166-2:SA",
		 "English": " Saudi Arabia",
		 "China": "沙特阿拉伯",
		 "Taiwan": "沙烏地阿拉伯",
		 "Hongkong": "沙特阿拉伯",
		 "Memo": "新加坡與馬來西亞均將之譯作「沙地阿拉伯」。香港亦普遍採用「沙地阿拉伯」一詞於其它場合"
		},
		{
		 "ISO2": "SB",
		 "ISO3": "SLB",
		 "DIGITS": "90",
		 "ISO-3166-2": "ISO 3166-2:SB",
		 "English": " Solomon Islands",
		 "China": "所罗门群岛",
		 "Taiwan": "所羅門群島",
		 "Hongkong": "所羅門群島",
		 "Memo": ""
		},
		{
		 "ISO2": "SC",
		 "ISO3": "SYC",
		 "DIGITS": "690",
		 "ISO-3166-2": "ISO 3166-2:SC",
		 "English": " Seychelles",
		 "China": "塞舌尔",
		 "Taiwan": "塞席爾",
		 "Hongkong": "塞舌爾",
		 "Memo": ""
		},
		{
		 "ISO2": "SD",
		 "ISO3": "SDN",
		 "DIGITS": "729",
		 "ISO-3166-2": "ISO 3166-2:SD",
		 "English": " Sudan",
		 "China": "苏丹",
		 "Taiwan": "蘇丹",
		 "Hongkong": "蘇丹",
		 "Memo": ""
		},
		{
		 "ISO2": "SE",
		 "ISO3": "SWE",
		 "DIGITS": "752",
		 "ISO-3166-2": "ISO 3166-2:SE",
		 "English": " Sweden",
		 "China": "瑞典",
		 "Taiwan": "瑞典",
		 "Hongkong": "瑞典",
		 "Memo": ""
		},
		{
		 "ISO2": "SG",
		 "ISO3": "SGP",
		 "DIGITS": "702",
		 "ISO-3166-2": "ISO 3166-2:SG",
		 "English": " Singapore",
		 "China": "新加坡",
		 "Taiwan": "新加坡",
		 "Hongkong": "新加坡(星加坡)",
		 "Memo": ""
		},
		{
		 "ISO2": "SH",
		 "ISO3": "SHN",
		 "DIGITS": "654",
		 "ISO-3166-2": "ISO 3166-2:SH",
		 "English": " Saint Helena, Ascension and Tristan da Cunha",
		 "China": "圣赫勒拿",
		 "Taiwan": "聖赫勒拿島",
		 "Hongkong": "聖赫勒拿",
		 "Memo": "香港亦普遍採用「聖赫勒拿島」一詞於其它場合（如香港郵政的郵政指南附錄表）。亦有部份人士使用「聖海倫娜島」一詞於港澳地區"
		},
		{
		 "ISO2": "SI",
		 "ISO3": "SVN",
		 "DIGITS": "705",
		 "ISO-3166-2": "ISO 3166-2:SI",
		 "English": " Slovenia",
		 "China": "斯洛文尼亚",
		 "Taiwan": "斯洛維尼亞",
		 "Hongkong": "斯洛文尼亞",
		 "Memo": ""
		},
		{
		 "ISO2": "SJ",
		 "ISO3": "SJM",
		 "DIGITS": "744",
		 "ISO-3166-2": "ISO 3166-2:SJ",
		 "English": " Svalbard and Jan Mayen",
		 "China": "斯瓦尔巴群岛和扬马延岛",
		 "Taiwan": "斯瓦巴及尖棉島",
		 "Hongkong": "斯瓦爾巴特群島",
		 "Memo": ""
		},
		{
		 "ISO2": "SK",
		 "ISO3": "SVK",
		 "DIGITS": "703",
		 "ISO-3166-2": "ISO 3166-2:SK",
		 "English": " Slovakia",
		 "China": "斯洛伐克",
		 "Taiwan": "斯洛伐克",
		 "Hongkong": "斯洛伐克",
		 "Memo": ""
		},
		{
		 "ISO2": "SL",
		 "ISO3": "SLE",
		 "DIGITS": "694",
		 "ISO-3166-2": "ISO 3166-2:SL",
		 "English": " Sierra Leone",
		 "China": "塞拉利昂",
		 "Taiwan": "獅子山",
		 "Hongkong": "塞拉利昂",
		 "Memo": ""
		},
		{
		 "ISO2": "SM",
		 "ISO3": "SMR",
		 "DIGITS": "674",
		 "ISO-3166-2": "ISO 3166-2:SM",
		 "English": " San Marino",
		 "China": "圣马力诺",
		 "Taiwan": "聖馬利諾",
		 "Hongkong": "聖馬力諾",
		 "Memo": ""
		},
		{
		 "ISO2": "SN",
		 "ISO3": "SEN",
		 "DIGITS": "686",
		 "ISO-3166-2": "ISO 3166-2:SN",
		 "English": " Senegal",
		 "China": "塞内加尔",
		 "Taiwan": "塞內加爾",
		 "Hongkong": "塞內加爾",
		 "Memo": ""
		},
		{
		 "ISO2": "SO",
		 "ISO3": "SOM",
		 "DIGITS": "706",
		 "ISO-3166-2": "ISO 3166-2:SO",
		 "English": " Somalia",
		 "China": "索马里",
		 "Taiwan": "索馬利亞",
		 "Hongkong": "索馬里",
		 "Memo": ""
		},
		{
		 "ISO2": "SR",
		 "ISO3": "SUR",
		 "DIGITS": "740",
		 "ISO-3166-2": "ISO 3166-2:SR",
		 "English": " Suriname",
		 "China": "苏里南",
		 "Taiwan": "蘇利南",
		 "Hongkong": "蘇里南",
		 "Memo": ""
		},
		{
		 "ISO2": "SS",
		 "ISO3": "SSD",
		 "DIGITS": "728",
		 "ISO-3166-2": "ISO 3166-2:SS",
		 "English": " South Sudan",
		 "China": "南苏丹",
		 "Taiwan": "南蘇丹",
		 "Hongkong": "南蘇丹",
		 "Memo": ""
		},
		{
		 "ISO2": "ST",
		 "ISO3": "STP",
		 "DIGITS": "678",
		 "ISO-3166-2": "ISO 3166-2:ST",
		 "English": " Sao Tome and Principe",
		 "China": "圣多美和普林西比",
		 "Taiwan": "聖多美普林西比",
		 "Hongkong": "聖多美及普林西比",
		 "Memo": ""
		},
		{
		 "ISO2": "SV",
		 "ISO3": "SLV",
		 "DIGITS": "222",
		 "ISO-3166-2": "ISO 3166-2:SV",
		 "English": " El Salvador",
		 "China": "萨尔瓦多",
		 "Taiwan": "薩爾瓦多",
		 "Hongkong": "薩爾瓦多",
		 "Memo": ""
		},
		{
		 "ISO2": "SX",
		 "ISO3": "SXM",
		 "DIGITS": "534",
		 "ISO-3166-2": "ISO 3166-2:SX",
		 "English": " Sint Maarten (Dutch part)",
		 "China": "荷属圣马丁",
		 "Taiwan": "荷屬聖馬丁",
		 "Hongkong": "荷屬聖馬丁",
		 "Memo": ""
		},
		{
		 "ISO2": "SY",
		 "ISO3": "SYR",
		 "DIGITS": "760",
		 "ISO-3166-2": "ISO 3166-2:SY",
		 "English": " Syrian Arab Republic",
		 "China": "叙利亚",
		 "Taiwan": "敘利亞",
		 "Hongkong": "敘利亞",
		 "Memo": ""
		},
		{
		 "ISO2": "SZ",
		 "ISO3": "SWZ",
		 "DIGITS": "748",
		 "ISO-3166-2": "ISO 3166-2:SZ",
		 "English": " Swaziland",
		 "China": "斯威士兰",
		 "Taiwan": "史瓦濟蘭",
		 "Hongkong": "斯威士蘭",
		 "Memo": ""
		},
		{
		 "ISO2": "TC",
		 "ISO3": "TCA",
		 "DIGITS": "796",
		 "ISO-3166-2": "ISO 3166-2:TC",
		 "English": " Turks and Caicos Islands",
		 "China": "特克斯和凯科斯群岛",
		 "Taiwan": "土克斯及開科斯群島",
		 "Hongkong": "特克斯和凱科斯群島",
		 "Memo": ""
		},
		{
		 "ISO2": "TD",
		 "ISO3": "TCD",
		 "DIGITS": "148",
		 "ISO-3166-2": "ISO 3166-2:TD",
		 "English": " Chad",
		 "China": "乍得",
		 "Taiwan": "查德",
		 "Hongkong": "乍得",
		 "Memo": ""
		},
		{
		 "ISO2": "TF",
		 "ISO3": "ATF",
		 "DIGITS": "260",
		 "ISO-3166-2": "ISO 3166-2:TF",
		 "English": " French Southern Territories",
		 "China": "法属南部领地",
		 "Taiwan": "法屬南部屬地",
		 "Hongkong": "法屬南部地區",
		 "Memo": "台灣亦普遍採用「法屬南方及南極陸地」一詞於其它場合（如MSN台灣）"
		},
		{
		 "ISO2": "TG",
		 "ISO3": "TGO",
		 "DIGITS": "768",
		 "ISO-3166-2": "ISO 3166-2:TG",
		 "English": " Togo",
		 "China": "多哥",
		 "Taiwan": "多哥",
		 "Hongkong": "多哥",
		 "Memo": ""
		},
		{
		 "ISO2": "TH",
		 "ISO3": "THA",
		 "DIGITS": "764",
		 "ISO-3166-2": "ISO 3166-2:TH",
		 "English": " Thailand",
		 "China": "泰国",
		 "Taiwan": "泰國",
		 "Hongkong": "泰國",
		 "Memo": ""
		},
		{
		 "ISO2": "TJ",
		 "ISO3": "TJK",
		 "DIGITS": "762",
		 "ISO-3166-2": "ISO 3166-2:TJ",
		 "English": " Tajikistan",
		 "China": "塔吉克斯坦",
		 "Taiwan": "塔吉克",
		 "Hongkong": "塔吉克",
		 "Memo": "香港習慣略去「斯坦」後綴，有必要會用全稱"
		},
		{
		 "ISO2": "TK",
		 "ISO3": "TKL",
		 "DIGITS": "772",
		 "ISO-3166-2": "ISO 3166-2:TK",
		 "English": " Tokelau",
		 "China": "托克劳",
		 "Taiwan": "托克勞群島",
		 "Hongkong": "托克勞群島",
		 "Memo": ""
		},
		{
		 "ISO2": "TL",
		 "ISO3": "TLS",
		 "DIGITS": "626",
		 "ISO-3166-2": "ISO 3166-2:TP",
		 "English": " Timor-Leste",
		 "China": "东帝汶",
		 "Taiwan": "東帝汶",
		 "Hongkong": "東帝汶",
		 "Memo": ""
		},
		{
		 "ISO2": "TM",
		 "ISO3": "TKM",
		 "DIGITS": "795",
		 "ISO-3166-2": "ISO 3166-2:TM",
		 "English": " Turkmenistan",
		 "China": "土库曼斯坦",
		 "Taiwan": "土庫曼",
		 "Hongkong": "土庫曼",
		 "Memo": "香港習慣略去「斯坦」後綴，有必要會用全稱"
		},
		{
		 "ISO2": "TN",
		 "ISO3": "TUN",
		 "DIGITS": "788",
		 "ISO-3166-2": "ISO 3166-2:TN",
		 "English": " Tunisia",
		 "China": "突尼斯",
		 "Taiwan": "突尼西亞",
		 "Hongkong": "突尼斯",
		 "Memo": ""
		},
		{
		 "ISO2": "TO",
		 "ISO3": "TON",
		 "DIGITS": "776",
		 "ISO-3166-2": "ISO 3166-2:TO",
		 "English": " Tonga",
		 "China": "汤加",
		 "Taiwan": "東加",
		 "Hongkong": "湯加",
		 "Memo": ""
		},
		{
		 "ISO2": "TR",
		 "ISO3": "TUR",
		 "DIGITS": "792",
		 "ISO-3166-2": "ISO 3166-2:TR",
		 "English": " Turkey",
		 "China": "土耳其",
		 "Taiwan": "土耳其",
		 "Hongkong": "土耳其",
		 "Memo": ""
		},
		{
		 "ISO2": "TT",
		 "ISO3": "TTO",
		 "DIGITS": "780",
		 "ISO-3166-2": "ISO 3166-2:TT",
		 "English": " Trinidad and Tobago",
		 "China": "特立尼达和多巴哥",
		 "Taiwan": "千里達及托巴哥",
		 "Hongkong": "千里達和多巴哥",
		 "Memo": "台灣和香港均將之簡稱為「千里達」"
		},
		{
		 "ISO2": "TV",
		 "ISO3": "TUV",
		 "DIGITS": "798",
		 "ISO-3166-2": "ISO 3166-2:TV",
		 "English": " Tuvalu",
		 "China": "图瓦卢",
		 "Taiwan": "吐瓦鲁",
		 "Hongkong": "圖瓦盧",
		 "Memo": ""
		},
		{
		 "ISO2": "TW",
		 "ISO3": "TWN",
		 "DIGITS": "158",
		 "ISO-3166-2": "ISO 3166-2:TW",
		 "English": " Taiwan, Province of China",
		 "China": "台湾",
		 "Taiwan": "台灣",
		 "Hongkong": "台湾(臺灣)",
		 "Memo": "因历史、政治原因，ISO官方称之为“Taiwan, Province of China”[1]。台灣的國際政治地位可參見未被國際普遍承認的國家列表、台海現狀以及舊金山條約，但臺灣官方仍稱為中華民國或中華民國（臺灣）。"
		},
		{
		 "ISO2": "TZ",
		 "ISO3": "TZA",
		 "DIGITS": "834",
		 "ISO-3166-2": "ISO 3166-2:TZ",
		 "English": " Tanzania, United Republic of",
		 "China": "坦桑尼亚",
		 "Taiwan": "坦尚尼亞",
		 "Hongkong": "坦桑尼亞",
		 "Memo": ""
		},
		{
		 "ISO2": "UA",
		 "ISO3": "UKR",
		 "DIGITS": "804",
		 "ISO-3166-2": "ISO 3166-2:UA",
		 "English": " Ukraine",
		 "China": "乌克兰",
		 "Taiwan": "烏克蘭",
		 "Hongkong": "烏克蘭",
		 "Memo": ""
		},
		{
		 "ISO2": "UG",
		 "ISO3": "UGA",
		 "DIGITS": "800",
		 "ISO-3166-2": "ISO 3166-2:UG",
		 "English": " Uganda",
		 "China": "乌干达",
		 "Taiwan": "烏干達",
		 "Hongkong": "烏干達",
		 "Memo": ""
		},
		{
		 "ISO2": "UM",
		 "ISO3": "UMI",
		 "DIGITS": "581",
		 "ISO-3166-2": "ISO 3166-2:UM",
		 "English": " United States Minor Outlying Islands",
		 "China": "美国本土外小岛屿",
		 "Taiwan": "美國邊疆小島",
		 "Hongkong": "美國海外小島",
		 "Memo": "台灣亦普遍採用「美國外島」一詞於其它場合（如MSN台灣）"
		},
		{
		 "ISO2": "US",
		 "ISO3": "USA",
		 "DIGITS": "840",
		 "ISO-3166-2": "ISO 3166-2:US",
		 "English": " United States",
		 "China": "美国",
		 "Taiwan": "美國",
		 "Hongkong": "美國",
		 "Memo": ""
		},
		{
		 "ISO2": "UY",
		 "ISO3": "URY",
		 "DIGITS": "858",
		 "ISO-3166-2": "ISO 3166-2:UY",
		 "English": " Uruguay",
		 "China": "乌拉圭",
		 "Taiwan": "烏拉圭",
		 "Hongkong": "烏拉圭",
		 "Memo": ""
		},
		{
		 "ISO2": "UZ",
		 "ISO3": "UZB",
		 "DIGITS": "860",
		 "ISO-3166-2": "ISO 3166-2:UZ",
		 "English": " Uzbekistan",
		 "China": "乌兹别克斯坦",
		 "Taiwan": "烏茲別克",
		 "Hongkong": "烏茲別克",
		 "Memo": "香港習慣略去「斯坦」後綴，有必要會用全稱"
		},
		{
		 "ISO2": "VA",
		 "ISO3": "VAT",
		 "DIGITS": "336",
		 "ISO-3166-2": "ISO 3166-2:VA",
		 "English": " Holy See (Vatican City State)",
		 "China": "梵蒂冈",
		 "Taiwan": "梵蒂岡",
		 "Hongkong": "梵蒂岡",
		 "Memo": ""
		},
		{
		 "ISO2": "VC",
		 "ISO3": "VCT",
		 "DIGITS": "670",
		 "ISO-3166-2": "ISO 3166-2:VC",
		 "English": " Saint Vincent and the Grenadines",
		 "China": "圣文森特和格林纳丁斯",
		 "Taiwan": "聖文森及格瑞那丁",
		 "Hongkong": "聖文森特和格林納丁斯",
		 "Memo": "台灣將之簡稱為「聖文森」"
		},
		{
		 "ISO2": "VE",
		 "ISO3": "VEN",
		 "DIGITS": "862",
		 "ISO-3166-2": "ISO 3166-2:VE",
		 "English": " Venezuela (Bolivarian Republic of)",
		 "China": "委内瑞拉",
		 "Taiwan": "委內瑞拉",
		 "Hongkong": "委內瑞拉",
		 "Memo": ""
		},
		{
		 "ISO2": "VG",
		 "ISO3": "VGB",
		 "DIGITS": "92",
		 "ISO-3166-2": "ISO 3166-2:VG",
		 "English": " Virgin Islands (British)",
		 "China": "英属维尔京群岛",
		 "Taiwan": "英屬維爾京群島",
		 "Hongkong": "英屬處女群島",
		 "Memo": ""
		},
		{
		 "ISO2": "VI",
		 "ISO3": "VIR",
		 "DIGITS": "850",
		 "ISO-3166-2": "ISO 3166-2:VI",
		 "English": " Virgin Islands (U.S.)",
		 "China": "美属维尔京群岛",
		 "Taiwan": "美屬維爾京群島",
		 "Hongkong": "美屬處女群島",
		 "Memo": ""
		},
		{
		 "ISO2": "VN",
		 "ISO3": "VNM",
		 "DIGITS": "704",
		 "ISO-3166-2": "ISO 3166-2:VN",
		 "English": " Vietnam",
		 "China": "越南",
		 "Taiwan": "越南",
		 "Hongkong": "越南",
		 "Memo": ""
		},
		{
		 "ISO2": "VU",
		 "ISO3": "VUT",
		 "DIGITS": "548",
		 "ISO-3166-2": "ISO 3166-2:VU",
		 "English": " Vanuatu",
		 "China": "瓦努阿图",
		 "Taiwan": "萬那杜",
		 "Hongkong": "瓦努阿圖",
		 "Memo": "亦有部份人士使用「溫納圖」一詞於港澳地區"
		},
		{
		 "ISO2": "WF",
		 "ISO3": "WLF",
		 "DIGITS": "876",
		 "ISO-3166-2": "ISO 3166-2:WF",
		 "English": " Wallis and Futuna",
		 "China": "瓦利斯和富图纳",
		 "Taiwan": "沃里斯與伏塔那島",
		 "Hongkong": "瓦利斯群島和富圖納群島",
		 "Memo": ""
		},
		{
		 "ISO2": "WS",
		 "ISO3": "WSM",
		 "DIGITS": "882",
		 "ISO-3166-2": "ISO 3166-2:WS",
		 "English": " Samoa",
		 "China": "萨摩亚",
		 "Taiwan": "薩摩亞",
		 "Hongkong": "薩摩亞",
		 "Memo": ""
		},
		{
		 "ISO2": "YE",
		 "ISO3": "YEM",
		 "DIGITS": "887",
		 "ISO-3166-2": "ISO 3166-2:YE",
		 "English": " Yemen",
		 "China": "也门",
		 "Taiwan": "葉門",
		 "Hongkong": "也門",
		 "Memo": ""
		},
		{
		 "ISO2": "YT",
		 "ISO3": "MYT",
		 "DIGITS": "175",
		 "ISO-3166-2": "ISO 3166-2:YT",
		 "English": " Mayotte",
		 "China": "马约特",
		 "Taiwan": "美亞特",
		 "Hongkong": "馬約特",
		 "Memo": ""
		},
		{
		 "ISO2": "ZA",
		 "ISO3": "ZAF",
		 "DIGITS": "710",
		 "ISO-3166-2": "ISO 3166-2:ZA",
		 "English": " South Africa",
		 "China": "南非",
		 "Taiwan": "南非",
		 "Hongkong": "南非",
		 "Memo": ""
		},
		{
		 "ISO2": "ZM",
		 "ISO3": "ZMB",
		 "DIGITS": "894",
		 "ISO-3166-2": "ISO 3166-2:ZM",
		 "English": " Zambia",
		 "China": "赞比亚",
		 "Taiwan": "尚比亞",
		 "Hongkong": "贊比亞",
		 "Memo": ""
		},
		{
		 "ISO2": "ZW",
		 "ISO3": "ZWE",
		 "DIGITS": "716",
		 "ISO-3166-2": "ISO 3166-2:ZW",
		 "English": " Zimbabwe",
		 "China": "津巴布韦",
		 "Taiwan": "辛巴威",
		 "Hongkong": "津巴布韋",
		 "Memo": ""
		}
	   ]`)

	err := json.Unmarshal([]byte(str), &countries)
	if err != nil {
		// ... handle error
		fmt.Println(err)
	}

	return countries
}
