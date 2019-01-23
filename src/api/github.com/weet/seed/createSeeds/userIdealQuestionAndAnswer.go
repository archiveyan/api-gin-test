package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedUserIdealQuestionAndAnswers() {

	idealQuestionAndAnswer_infos := []map[string]string{
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
			"MatchingFormatID": "2",
			"UserID":           "1",
			"QuestionID":       "3",
			"AnswerID":         "10",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "1",
			"QuestionID":       "3",
			"AnswerID":         "11",
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
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "1",
			"AnswerID":         "1",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "2",
			"AnswerID":         "6",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "2",
			"QuestionID":       "3",
			"AnswerID":         "10",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "2",
			"QuestionID":       "3",
			"AnswerID":         "11",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "4",
			"AnswerID":         "13",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "2",
			"QuestionID":       "5",
			"AnswerID":         "17",
		},
	}

	for _, info := range idealQuestionAndAnswer_infos {
		matchingFormatID, _ := strconv.Atoi(info["MatchingFormatID"])
		userID, _ := strconv.Atoi(info["UserID"])
		questionID, _ := strconv.Atoi(info["QuestionID"])
		answerID, _ := strconv.Atoi(info["AnswerID"])

		createUserIdealQuestionAndAnswer(model.UserIdealQuestionAndAnswer{
			MatchingFormatID: uint(matchingFormatID),
			UserID:           uint(userID),
			QuestionID:       uint(questionID),
			AnswerID:         uint(answerID),
		})
	}
}

func createUserIdealQuestionAndAnswer(userIdealQuestionAndAnswer model.UserIdealQuestionAndAnswer) {
	userIdealQuestionAndAnswer, err := model.CreateUserIdealQuestionAndAnswer(userIdealQuestionAndAnswer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("userIdealQuestionAndAnswer created\n")
}

// answerIDsString := strings.Split(info["AnswerIDs"], ",")
// var answerIDsUint []uint
// var answerIDInt int
// for _, answerID := range answerIDsString {
// 	answerIDInt, _ = strconv.Atoi(answerID)
// 	answerIDsUint = append(answerIDsUint, uint(answerIDInt))
// }
