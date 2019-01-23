package model

// questionTableにInsertする
func CreateQuestion(question Question) (Question, error) {
	err := db.Create(&question).Error
	if err != nil {
		return Question{}, err
	}
	return question, nil
}
