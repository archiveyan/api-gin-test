package main

import (
	_ "github.com/jinzhu/gorm"
	"github.com/weet/model"
)

func main() {

	db := model.GetDBConn()

	// migrate

	db.DropTableIfExists(&model.Sample{})
	db.DropTableIfExists(&model.UserBasics{})
	db.DropTableIfExists(&model.UserQuestionAndAnswer{})
	db.DropTableIfExists(&model.Question{})
	db.DropTableIfExists(&model.Answer{})
	db.DropTableIfExists(&model.MatchingFormat{})
	db.DropTableIfExists(&model.UserIdealQuestionAndAnswer{})

	db.CreateTable(&model.Sample{})
	db.CreateTable(&model.UserBasics{})
	db.CreateTable(&model.UserQuestionAndAnswer{})
	db.CreateTable(&model.Question{})
	db.CreateTable(&model.Answer{})
	db.CreateTable(&model.MatchingFormat{})
	db.CreateTable(&model.UserIdealQuestionAndAnswer{})
}
