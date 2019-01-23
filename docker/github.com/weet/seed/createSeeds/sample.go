package createSeeds

import (
	"fmt"

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

func CreateSeedSamples() {

	sample_infos := []map[string]string{
		map[string]string{
			"Name":    "野口英世",
			"Comment": "1000円",
		},
		map[string]string{
			"Name":    "樋口一葉",
			"Comment": "5000円",
		},
		map[string]string{
			"Name":    "福沢諭吉",
			"Comment": "我諭吉ぞ？",
		},
	}

	for _, info := range sample_infos {
		createSample(model.Sample{
			Name:    info["Name"],
			Comment: info["Comment"],
		})
	}
}

func createSample(sample model.Sample) {
	sample, err := model.CreateSample(sample)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sample created\n")
}
