package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedQuestions() {

	question_infos := []map[string]string{
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "血液型",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "休日の過ごし方",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"Name":             "デートの頻度",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "子供の願望",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"Name":             "ルームシェアの目的",
		},
	}

	for _, info := range question_infos {
		matchingFormatID, _ := strconv.Atoi(info["MatchingFormatID"])
		createQuestion(model.Question{
			MatchingFormatID: uint(matchingFormatID),
			Name:             info["Name"],
		})
	}
}

func createQuestion(question model.Question) {
	question, err := model.CreateQuestion(question)
	if err != nil {
		panic(err)
	}
	fmt.Printf("question created\n")
}
