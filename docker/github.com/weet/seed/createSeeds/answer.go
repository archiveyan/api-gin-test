package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedAnswers() {

	answers_infos := []map[string]string{
		map[string]string{
			"QuestionID": "1",
			"Name":       "A型",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "B型",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "O型",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "AB型",
		},
		map[string]string{ //5
			"QuestionID": "2",
			"Name":       "インドア派",
		},
		map[string]string{
			"QuestionID": "2",
			"Name":       "アウトドア派",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "毎日",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "2~3日に1回",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "4~6日に1回",
		},
		map[string]string{ //10
			"QuestionID": "3",
			"Name":       "1週間に1回",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "2週間に1回",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "月に1回以下",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "欲しい",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "欲しくない",
		},
		map[string]string{ //15
			"QuestionID": "5",
			"Name":       "生活費の節約",
		},
		map[string]string{
			"QuestionID": "5",
			"Name":       "誰かがいる安心感を得たい",
		},
		map[string]string{
			"QuestionID": "5",
			"Name":       "気まぐれ",
		},
	}

	for _, info := range answers_infos {
		questionID, _ := strconv.Atoi(info["QuestionID"])
		createAnswer(model.Answer{
			QuestionID: uint(questionID),
			Name:       info["Name"],
		})
	}
}

func createAnswer(answer model.Answer) {
	answer, err := model.CreateAnswer(answer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("answer created\n")
}
