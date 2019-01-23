package model

import (
	"github.com/jinzhu/gorm"
)

//type //model struct {
//    ID      uint `gorm:"primary_key"`
//    CreatedAt time.Time
//    UpdatedAt time.Time
//}

// Sample情報
type Sample struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Comment string `gorm:"not null"`
}

//基本的なユーザ情報
type UserBasics struct {
	gorm.Model
	UserName string `gorm:"not null"`
	Image1   string `gorm:"not null"`
	Image2   string
	Image3   string
	Age      uint   `gorm:"not null"`
	Sex      string `gorm:"not null"`
	Hitokoto string
	Comment  string
}

//ユーザごとの質疑応答
type UserQuestionAndAnswer struct {
	gorm.Model
	MatchingFormatID uint `gorm:"not null"`
	UserID           uint `gorm:"not null"`
	QuestionID       uint `gorm:"not null"`
	AnswerID         uint
	Question         Question
	Answer           Answer
}

//ユーザごとの質疑応答(理想像)
type UserIdealQuestionAndAnswer struct {
	gorm.Model
	MatchingFormatID uint `gorm:"not null"`
	UserID           uint `gorm:"not null"`
	QuestionID       uint `gorm:"not null"`
	AnswerID         uint
	Question         Question
	Answer           Answer
}

//マイページの質問
type Question struct {
	gorm.Model
	Name             string `gorm:"not null"`
	MatchingFormatID uint   `gorm:"not null"`
	Answer           []Answer
}

//マイページの質問の答えの選択肢
type Answer struct {
	gorm.Model
	Name       string `gorm:"not null"`
	QuestionID uint   `gorm:"not null"`
}

//マッチング形式
type MatchingFormat struct {
	gorm.Model
	Name string `gorm:"not null"`
}
