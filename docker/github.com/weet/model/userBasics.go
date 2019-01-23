package model

import "github.com/weet/service"

// userBasicsTableにInsertする
func CreateUserBasics(userBasics UserBasics) (UserBasics, error) {
	err := db.Create(&userBasics).Error
	if err != nil {
		return UserBasics{}, err
	}
	return userBasics, nil
}

func GetUserBasicsById(userId uint) (service.UserBasics, error) {
	userBasics := service.UserBasics{}
	err := db.Where("id = ?", userId).First(&userBasics).Error
	userBasics.MatchingFormatName = "基本情報"
	return userBasics, err
}

/*
//基本的な情報
type UserBasics struct {
	FormatName string `json:"format_name"`
	UserName   string `json:"user_name"`
	Image1     string `json:"image1"`
	Image2     string `json:"image2"`
	Image3     string `json:"image3"`
	Age        string `json:"age"`
	Hitokoto   string `json:"hitokoto"`
	Comment    string `json:"comment"`
}
*/
