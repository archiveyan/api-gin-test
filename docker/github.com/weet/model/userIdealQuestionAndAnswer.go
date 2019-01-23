package model

import (
	"github.com/weet/service"
)

// userQuestionAndAnswerTableにInsertする
func CreateUserIdealQuestionAndAnswer(userIdealQuestionAndAnswer UserIdealQuestionAndAnswer) (UserIdealQuestionAndAnswer, error) {
	err := db.Create(&userIdealQuestionAndAnswer).Error
	if err != nil {
		return UserIdealQuestionAndAnswer{}, err
	}
	return userIdealQuestionAndAnswer, nil
}

//user_idとFormat_Idで条件を絞り、理想像の質問と答えが１セットの配列が欲しい(N+1問題の解決のためにクエリを手打ち)
func GetUserIdealQuestionAndAnswerByUserIDAndFormatID(userId uint, matchingFormatId uint) (service.UserIdealQuestionsAndAnswers, error) {
	userIdealQuestionsAndAnswers := service.UserIdealQuestionsAndAnswers{}
	userIdealQAndA := service.UserIdealQuestionsAndAnswers{}
	count := 0

	err := db.Raw("select questions.id AS question_id, questions.name AS question_name , answers.name AS answer_name from user_ideal_question_and_answers join user_basics on (user_ideal_question_and_answers.user_id = user_basics.id) left join questions on (user_ideal_question_and_answers.question_id = questions.id) left join answers on (user_ideal_question_and_answers.answer_id = answers.id) where user_ideal_question_and_answers.user_id = ? and user_ideal_question_and_answers.matching_format_id = ?", userId, matchingFormatId).Scan(&userIdealQuestionsAndAnswers).Error

	//ABB
	userIdealQAndA = append(userIdealQAndA, userIdealQuestionsAndAnswers[0])

	for i := 0; i < len(userIdealQuestionsAndAnswers)-1; i++ {
		// if 1 ==

		if userIdealQuestionsAndAnswers[i].QuestionID == userIdealQuestionsAndAnswers[i+1].QuestionID {
			userIdealQAndA[count].AnswerName = userIdealQAndA[count].AnswerName + "," + userIdealQuestionsAndAnswers[i+1].AnswerName
		} else {
			count++
			userIdealQAndA = append(userIdealQAndA, userIdealQuestionsAndAnswers[i+1])
		}

	}
	return userIdealQAndA, err
}

/*
	map[string]string{
		"MatchingFormatID": "1",
		"UserID":           "1",
		"QuestionID":       "1",
		"AnswerID":         "1",
	},
	map[string]string{
		"MatchingFormatID": "1",
		"UserID":           "1",
		"QuestionID":       "1",
		"AnswerID":         "2",
	},
	map[string]string{
		"MatchingFormatID": "1",
		"UserID":           "1",
		"QuestionID":       "2",
		"AnswerID":         "5",
	},
	map[string]string{
		"MatchingFormatID": "2",
		"UserID":           "1",
		"QuestionID":       "3",
		"AnswerID":         "8",
	},
	map[string]string{
		"MatchingFormatID": "2",
		"UserID":           "1",
		"QuestionID":       "3",
		"AnswerID":         "9",
	},
	map[string]string{
		"MatchingFormatID": "3",
		"UserID":           "1",
		"QuestionID":       "4",
		"AnswerID":         "13",
	},
	map[string]string{
		"MatchingFormatID": "4",
		"UserID":           "1",
		"QuestionID":       "5",
		"AnswerID":         "15",
	},
	map[string]string{
		"MatchingFormatID": "4",
		"UserID":           "1",
		"QuestionID":       "5",
		"AnswerID":         "16",
	},
*/

// 	for i, _ := range userIdealQuestionsAndAnswers {
// 		fmt.Println(userIdealQuestionsAndAnswers[i+1].QuestionID)

// 		if userIdealQuestionsAndAnswers[i].QuestionID == userIdealQuestionsAndAnswers[i].QuestionID {
// 			// 	userIdealQAndA[count].AnswerName = userIdealQAndA[count].AnswerName + "," + userIdealQuestionsAndAnswers[i+1].AnswerName
// 			// } else {
// 			// 	count++
// 			// 	userIdealQAndA = append(userIdealQAndA, userIdealQuestionsAndAnswers[i+1])
// 		}

// 	}
// 	return userIdealQAndA
// }
