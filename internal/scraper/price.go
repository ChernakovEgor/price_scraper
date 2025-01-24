package scraper

import (
	"fmt"
	"strconv"
	"strings"
)

const digits = "0123456789."

func getPrice(priceString string) (float64, error) {
	stripped := strings.ReplaceAll(priceString, " ", "")

	var cleanedString string

	for _, r := range stripped {
		if strings.ContainsRune(digits, r) {
			cleanedString += string(r)
		}
	}

	price, err := strconv.ParseFloat(cleanedString, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing float: %v", err)
	}

	return price, nil
}
