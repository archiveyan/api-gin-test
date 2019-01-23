package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

// // 自分のユーザー情報
// type User struct {
// 	//Model
// 	gorm.Model
// 	Uid   string
// 	Name  string
// 	Email string
// 	//        val githubStatus : GitHubStatus,// ログインで得られるGitHubの情報によってデータクラスを作る
// }

func CreateSeedUserBasics() {

	userBasics_infos := []map[string]string{
		map[string]string{
			"UserName": "かまやん",
			"Image1":   "https://ezumin.com/wp-content/uploads/2017/12/13395110_1166833013368528_105320497_n.jpg",
			"Image2":   "",
			"Image3":   "",
			"Age":      "22",
			"Sex":      "男性",
			"Hitokoto": "ガールフレンドを作りたいです！",
			"Comment":  "お酒を飲める人とバーや居酒屋に行きたいです。\n身長180cm以上でアクティブでポジティブなボーイッシュな人募集します！（周りからは理想高すぎっ！？と言われますが書いてみました。）",
		},
		map[string]string{
			"UserName": "橋本環奈",
			"Image1":   "https://pbs.twimg.com/media/DBl9HtdVoAAzt1A.jpg",
			"Image2":   "",
			"Image3":   "",
			"Age":      "20",
			"Sex":      "女性",
			"Hitokoto": "我、橋本環奈ぞ？",
			"Comment":  "1000年に一度の逸材です！",
		},
	}

	for _, info := range userBasics_infos {
		age, _ := strconv.Atoi(info["Age"])
		createUserBasics(model.UserBasics{
			UserName: info["UserName"],
			Image1:   info["Image1"],
			Image2:   info["Image2"],
			Image3:   info["Image3"],
			Age:      uint(age),
			Sex:      info["Sex"],
			Hitokoto: info["Hitokoto"],
			Comment:  info["Comment"],
		})
	}
}

func createUserBasics(userBasics model.UserBasics) {
	userBasics, err := model.CreateUserBasics(userBasics)
	if err != nil {
		panic(err)
	}
	fmt.Printf("userBasics created\n")
}
