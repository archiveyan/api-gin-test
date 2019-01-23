package main

import "github.com/weet/seed/createSeeds"

func main() {
	createSeeds.CreateSeedSamples()
	createSeeds.CreateSeedUserBasics()
	createSeeds.CreateSeedQuestions()
	createSeeds.CreateSeedAnswers()
	createSeeds.CreateSeedMatchingFormats()
	createSeeds.CreateSeedUserQuestionAndAnswers()
	createSeeds.CreateSeedUserIdealQuestionAndAnswers()
}
