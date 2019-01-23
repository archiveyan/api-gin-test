package model

import (
	"github.com/weet/service"
)

func GetUserSpecialsById(userId uint) (service.UserSpecials, error) {
	userSpecial := service.UserSpecial{}
	userSpecials := service.UserSpecials{}
	var err error

	for i := 1; i < 5; i++ {

		//userSpecial.MatchingFormatName = GetMatchingFormatNameByID(uint(i))

		switch i {
		case 1:
			userSpecial.MatchingFormatName = "友達"
		case 2:
			userSpecial.MatchingFormatName = "恋愛"
		case 3:
			userSpecial.MatchingFormatName = "婚活"
		case 4:
			userSpecial.MatchingFormatName = "ルームメイト"
		}
		userSpecial.UserQuestionsAndAnswers, err = GetUserQuestionAndAnswerByUserIDAndFormatID(userId, uint(i))
		userSpecials = append(userSpecials, userSpecial)
	}
	return userSpecials, err
}

/*
//マッチング形式ごとのタイトルと質疑応答
type UserSpecial struct {
	[
		MatchingFormatName
		[
			QuestionID   uint   `json:"question_id"`
			QuestionName string `json:"question_name"`
			AnswerName   string `json:"answer_name"`
		]
	]
}
*/
