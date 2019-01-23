package createSeeds

import (
	"fmt"

	"github.com/weet/model"
)

func CreateSeedMatchingFormats() {

	matchingFormat_infos := []map[string]string{
		map[string]string{
			"Name": "友達",
		},
		map[string]string{
			"Name": "恋愛",
		},
		map[string]string{
			"Name": "婚活",
		},
		map[string]string{
			"Name": "ルームメイト",
		},
	}

	for _, info := range matchingFormat_infos {
		createMatchingFormat(model.MatchingFormat{
			Name: info["Name"],
		})
	}
}

func createMatchingFormat(matchingFormat model.MatchingFormat) {
	matchingFormat, err := model.CreateMatchingFormat(matchingFormat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("matchingFormat created\n")
}
