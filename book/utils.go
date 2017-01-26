package book

import (
	"regexp"
	"strconv"
)

const defaultRating = 0.0

var ratingClassStringPattern = regexp.MustCompile("[0-5]")

func parseRating(classString string) float64 {
	s := ratingClassStringPattern.FindAllString(classString, -1)
	if len(s) == 0 {
		return defaultRating
	}

	rating, err := strconv.ParseFloat(s[0], 32)
	if err != nil {
		return defaultRating
	}

	return rating
}
