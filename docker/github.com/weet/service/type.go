package service

// sampleデータ
type Sample struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// ユーザ情報-----！
type User struct {
	UserBasics        UserBasics        `json:"user_basics"`
	UserSpecials      UserSpecials      `json:"user_specials"`
	UserIdealSpecials UserIdealSpecials `json:"user_ideal_specials"`
}

//基本的な情報
type UserBasics struct {
	MatchingFormatName string `json:"matching_format_name"`
	UserName           string `json:"user_name"`
	Image1             string `json:"image1"`
	Image2             string `json:"image2"`
	Image3             string `json:"image3"`
	Sex                string `json:"sex"`
	Age                uint   `json:"age"`
	Residence          string `json:"residence"`
	Hitokoto           string `json:"hitokoto"`
	Comment            string `json:"comment"`
}

//マッチング形式ごとのタイトルと質疑応答
type UserSpecial struct {
	MatchingFormatName      string                  `json:"matching_format_name"`
	UserQuestionsAndAnswers UserQuestionsAndAnswers `json:"user_questions_and_answers"`
}

type UserSpecials []UserSpecial

//１セットの質疑応答
//select user_basics.user_name, questions.name, answers.name from user_question_and_answers join user_basics on (user_question_and_answers.user_id = user_basics.id) left join questions on (user_question_and_answers.question_id = questions.id) left join answers on (user_question_and_answers.answer_id = answers.id);
type UserQuestionAndAnswer struct {
	QuestionID   uint   `json:"question_id"`
	QuestionName string `json:"question_name"`
	AnswerName   string `json:"answer_name"`
}

type UserQuestionsAndAnswers []UserQuestionAndAnswer

//マッチング形式ごとのタイトルと質疑応答
type UserIdealSpecial struct {
	MatchingFormatName           string                       `json:"matching_format_name"`
	UserIdealQuestionsAndAnswers UserIdealQuestionsAndAnswers `json:"user_ideal_questions_and_answers"`
}

type UserIdealSpecials []UserIdealSpecial

type UserIdealQuestionAndAnswer struct {
	QuestionID   uint   `json:"ideal_question_id"`
	QuestionName string `json:"ideal_question_name"`
	AnswerName   string `json:"ideal_answer_name"`
}

type UserIdealQuestionsAndAnswers []UserIdealQuestionAndAnswer

type AnswersByQuestion struct {
	QuestionID       uint              `json:"question_id"`
	CandidateAnswers []CandidateAnswer `json:"candidate_answer"`
}

//一つの質問に対する答えの候補-----！
type Answers []AnswersByQuestion

type CandidateAnswer struct {
	AnswerID   uint   `json:"answer_id"`
	AnswerName string `json:"answer_name"`
}

type MatchingUser struct {
	UserName  string `json:"user_name"`
	Image1    string `json:"image1"`
	Image2    string `json:"image2"`
	Image3    string `json:"image3"`
	Age       string `json:"age"`
	Residence string `json:"residence"`
	Hitokoto  string `json:"hitokoto"`
}

// {
// 	"answers": [
// 		{
// 			"question_id": "1",
// 			"candidate_answers": [
// 				{
// 					"answer_id": "1",
// 					"answer_name": "たこ焼き",
// 				},
// 				{
// 					"answer_id": "2",
// 					"answer_name": "お好み焼き",
// 				},
// 			]
// 		},
// 		{
// 			"question_id": "2",
// 			"candidate_answers": [
// 				{
// 					"answer_id": "1",
// 					"answer_name": "緑茶",
// 				},
// 				{
// 					"answer_id": "2",
// 					"answer_name": "麦茶",
// 				},
// 			]
// 		}
// 	]
// }
