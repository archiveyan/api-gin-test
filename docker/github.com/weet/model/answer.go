package model

import (
	"github.com/weet/service"
)

// answerTableにInsertする
func CreateAnswer(answer Answer) (Answer, error) {
	err := db.Create(&answer).Error
	if err != nil {
		return Answer{}, err
	}
	return answer, nil
}

func GetAnswers() (service.Answers, error) {
	answers := service.Answers{}
	answer := service.AnswersByQuestion{}

	candidateAnswers := []service.CandidateAnswer{}
	candidateAnswer := service.CandidateAnswer{}

	modelAnswers := []Answer{}
	err := db.Find(&modelAnswers).Error

	var i int
	for n, modelAnswer := range modelAnswers { //DBの結果（配列）を一つずつ変数に格納

		//DBの結果の一つのQuestionIDをServiceで渡す構造体のQuestionIDに代入
		if n == 0 || answers[i-1].QuestionID != modelAnswer.QuestionID { //一つ前のレコードのQuestionIDと新しいレコードのQuestionIDが同じでないならば
			answer.QuestionID = modelAnswer.QuestionID
			candidateAnswers = nil
			candidateAnswer.AnswerID = modelAnswer.ID
			candidateAnswer.AnswerName = modelAnswer.Name
			candidateAnswers = append(candidateAnswers, candidateAnswer)
			answer.CandidateAnswers = candidateAnswers
			answers = append(answers, answer)
			i++
		} else {
			candidateAnswer.AnswerID = modelAnswer.ID
			candidateAnswer.AnswerName = modelAnswer.Name
			candidateAnswers = append(candidateAnswers, candidateAnswer)
			answers[i-1].CandidateAnswers = candidateAnswers
		}
	}

	return answers, err
}
