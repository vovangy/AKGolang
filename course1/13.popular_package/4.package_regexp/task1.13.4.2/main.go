package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Ad struct {
	Title       string
	Description string
}

func main() {
	ads := []Ad{
		{Title: "Куплю велосипед MeRiDa", Description: "Куплю велосипед meriDA в хорошем состоянии."},
		{Title: "Продам ВаЗ 2101", Description: "Продам ваз 2101 в хорошем состоянии."},
		{Title: "Продам БМВ", Description: "Продам бМв в хорошем состоянии."},
		{Title: "Продам macBook pro", Description: "Продам macBook PRO в хорошем состоянии."},
	}

	replacements := map[string]string{
		"велосипед meriDA": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macBook pro":      "Macbook Pro",
	}

	ads = censorAds(ads, replacements)

	for _, ad := range ads {
		fmt.Println(ad.Title)
		fmt.Println(ad.Description)
		fmt.Println()
	}
}

func censorAds(ads []Ad, replacements map[string]string) []Ad {
	for i, ad := range ads {
		lowerTitle := strings.ToLower(ad.Title)
		lowerDescription := strings.ToLower(ad.Description)

		newTitle := ad.Title
		newDescription := ad.Description

		for oldPhrase, newPhrase := range replacements {
			re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(oldPhrase))

			if re.MatchString(lowerTitle) {
				newTitle = re.ReplaceAllString(newTitle, newPhrase)
			}
			if re.MatchString(lowerDescription) {
				newDescription = re.ReplaceAllString(newDescription, newPhrase)
			}
		}

		ads[i].Title = newTitle
		ads[i].Description = newDescription
	}
	return ads
}
