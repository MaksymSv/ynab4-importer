package ynab4

import "strings"

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

	categoriesMap := map[string]string{
		"LIDL":        "Everyday Expenses:Продукты",
		"KAUFLAND":    "Everyday Expenses:Продукты",
		"TESCO":       "Everyday Expenses:Продукты",
		"BILLA":       "Everyday Expenses:Продукты",
		"HOFER":       "Everyday Expenses:Продукты",
		"RADATZ":      "Everyday Expenses:Продукты",
		"TERNO":       "Everyday Expenses:Продукты",
		"VINOTEKA":    "Everyday Expenses:Продукты",
		"SHELL":       "Машина:Бензин",
		"BID":         "Everyday Expenses:Мелкие расходы",
		"KFC":         "Everyday Expenses:Развлечения",
		"MCDONALD":    "Everyday Expenses:Развлечения",
		"SIDE KEBAB":  "Everyday Expenses:Развлечения",
		"GATTO MATTO": "Everyday Expenses:Развлечения",

		"PEPCO":    "Everyday Expenses:Одежда/Обувь",
		"RESERVED": "Everyday Expenses:Одежда/Обувь",

		"DM":           "Everyday Expenses:Промтовары",
		"ALIEXPRESS":   "Everyday Expenses:Промтовары",
		"ACTION":       "Everyday Expenses:Промтовары",
		"JYSK":         "Everyday Expenses:Промтовары",
		"IKEA":         "Everyday Expenses:Промтовары",
		"ALZA.SK":      "Everyday Expenses:Промтовары",
		"4KA.SK":       "Дом:Телефон",
		"TELEKOMSZAML": "Дом:Телефон",

		"BINANCE": "Savings Goals / Rainy Day:Инвестиции",
	}

	for key := range categoriesMap {
		if strings.Contains(inputUpperCase, key) {
			return categoriesMap[key]
		}
	}
	return ""
}

func ParsePayee(input string) string {
	return ""
}
