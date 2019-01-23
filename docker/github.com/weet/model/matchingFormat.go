package model

import (
	"fmt"
)

// formatTableにInsertする
func CreateMatchingFormat(matchingFormat MatchingFormat) (MatchingFormat, error) {
	err := db.Create(&matchingFormat).Error
	if err != nil {
		return MatchingFormat{}, err
	}
	return matchingFormat, nil
}

func GetMatchingFormatNameByID(matchingFormatId uint) string {
	var matchingFormatName string

	db.Raw("select name AS matching_format_name from matching_formats where id = ?", matchingFormatId).Scan(&matchingFormatName)

	fmt.Printf(matchingFormatName)
	return matchingFormatName
}
