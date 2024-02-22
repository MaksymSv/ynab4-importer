package ynab4

import "strings"

var categoriesMap = map[string]string{
	"LIDL":          "Everyday Expenses:Продукты|Cупермаркет",
	"KAUFLAND":      "Everyday Expenses:Продукты|Cупермаркет",
	"TESCO":         "Everyday Expenses:Продукты|Cупермаркет",
	"BILLA":         "Everyday Expenses:Продукты|Cупермаркет",
	"HOFER":         "Everyday Expenses:Продукты|Cупермаркет",
	"RADATZ":        "Everyday Expenses:Продукты|Cупермаркет",
	"TERNO":         "Everyday Expenses:Продукты|Cупермаркет",
	"VINOTEKA":      "Everyday Expenses:Продукты|Cупермаркет",
	"CISAROV PEKAR": "Everyday Expenses:Продукты|Cупермаркет",

	"LEKAREN": "Everyday Expenses:Медицина|Аптека",
	"BENU SK": "Everyday Expenses:Медицина|Аптека",

	"SHELL":    "Машина:Бензин|Заправка",
	"SLOVNAFT": "Машина:Бензин|Заправка",
	"OMV":      "Машина:Бензин|Заправка",

	//"BID": "Everyday Expenses:Мелкие расходы",

	"BID":  "Everyday Expenses:Транспорт|Автобус",
	"FLIX": "Everyday Expenses:Транспорт|Автобус",

	"KFC":         "Everyday Expenses:Развлечения|Кафетерий",
	"MCDONALD":    "Everyday Expenses:Развлечения|Кафетерий",
	"SIDE KEBAB":  "Everyday Expenses:Развлечения|Кафетерий",
	"GATTO MATTO": "Everyday Expenses:Развлечения|Кафетерий",
	"KONDITOREI":  "Everyday Expenses:Развлечения|Кафетерий",
	"NETFLIX":     "Everyday Expenses:Развлечения|Netflix",

	"PEPCO":               "Everyday Expenses:Одежда/Обувь|Pepco",
	"RESERVED":            "Everyday Expenses:Одежда/Обувь|Reserved",
	"MUSTANG SHOP":        "Everyday Expenses:Одежда/Обувь|Mustang",
	"COLUMBIA SPORTSWEAR": "Everyday Expenses:Одежда/Обувь|Columbia Sportswear",
	"MOUNTAIN WAREHOUSE":  "Everyday Expenses:Одежда/Обувь|Mountain Warehouse",

	"PRISPEVOK NA CINNOST": "Дети:Школа / Занятия|Бассейн",

	"DM":             "Everyday Expenses:Промтовары|DM",
	"ALIEXPRESS":     "Everyday Expenses:Промтовары|ALIEXPRESS",
	"ACTION":         "Everyday Expenses:Промтовары|ACTION",
	"JYSK":           "Everyday Expenses:Промтовары|JYSK",
	"IKEA":           "Everyday Expenses:Промтовары|IKEA",
	"ALZA.SK":        "Everyday Expenses:Промтовары|ALZA",
	"PRIMARK":        "Everyday Expenses:Промтовары|PRIMARK",
	"OBI BRATISLAVA": "Everyday Expenses:Промтовары|OBI",

	"4KA.SK":       "Дом:Телефон|4KA",
	"TELEKOMSZAML": "Дом:Телефон|Telekom",

	"BINANCE": "Savings Goals / Rainy Day:Инвестиции|BINANCE",

	"TATRALANDIA":             "Savings Goals / Rainy Day:Отпуск|Аквапарк",
	"HAPPY END LM DEMANOVSKA": "Savings Goals / Rainy Day:Отпуск|Кафетерий",
	"GOPASS.TRAVEL":           "Savings Goals / Rainy Day:Отпуск|Лыжный Центр",
	"TMR  BERNARDINO BURGER":  "Savings Goals / Rainy Day:Отпуск|Кафетерий",
	"TATRYMOTION BIELA PUT":   "Savings Goals / Rainy Day:Отпуск|Инструктор",
}

type FormatService struct {
}

// NewFormatService creates instance of FormatService
func NewFormatService() *FormatService {
	s := new(FormatService)
	return s
}

func (s *FormatService) FormatSomwthig() []string {

	return nil
}

// ParseCategory privides the name of Cathegory for given spending code name
func ParseCategory(input string) string {
	inputUpperCase := strings.ToUpper(input)
	for key := range categoriesMap {
		if strings.Contains(inputUpperCase, key) {
			return parseCategoryPayee(categoriesMap[key]).Category
		}
	}
	return ""
}

func ParsePayee(input string) string {
	inputUpperCase := strings.ToUpper(input)
	for key := range categoriesMap {
		if strings.Contains(inputUpperCase, key) {
			return parseCategoryPayee(categoriesMap[key]).Payee
		}
	}
	return ""
}

func parseCategoryPayee(arg string) CategoryPayee {
	result := CategoryPayee{}

	parsed := strings.Split(arg, "|")
	if len(parsed) > 0 {
		result.Category = parsed[0]
	}
	if len(parsed) > 1 {
		result.Payee = parsed[1]
	}

	return result
}

type CategoryPayee struct {
	Category string
	Payee    string
}
