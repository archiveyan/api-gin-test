package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weet/controller"
)

func apiRouter(api *gin.RouterGroup) {

	// sampleAPI
	api.GET("/sample/:sample_id", controller.GetSampleById)

	//---マイページ画面---
	//ID別にユーザページ情報を取得
	api.GET("/user/:user_id", controller.GetUserById)

	//全ての回答を取得
	api.GET("/answers", controller.GetAnswers)

	//マッチング相手
	//性別、居住地（複数）、マッチ項目数
	//api.GET("/matching/player/:user_id/matching-format/:matching_format_id", controller.GetMatchingUser)

	/*
		//基本情報の取得
		api.GET("/mypage/basic/:user_id", controller.GetBasicMypageFilteredById)

		//フレンドマッチング情報の取得
		api.GET("/mypage/friend/:user_id", controller.GetFriendMypageFilteredById)

		//恋愛マッチング情報の取得
		api.GET("/mypage/love/:user_id", controller.GetLoveMypageFilteredById)

		//婚活マッチング情報の取得
		api.GET("/mypage/marriage/:user_id", controller.GetMarriageMypageFilteredById)

		//ルームメイトマッチング情報の取得
		api.GET("/mypage/roommate/:user_id", controller.GetRoommateMypageFilteredById)

		//---マッチング画面---
		//選択したマッチング形式でのマッチング相手を取得
		api.GET("/matching/:matching_id/:user_id", controller.GetMatchingFormatFilteredById)
	*/
}
