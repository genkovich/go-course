package video_game

type Question struct {
	text         string
	variants     map[int]string
	answerNumber int
}

type QuestionPack struct {
	title     string
	questions []Question
}

func (q Question) String() string {
	return q.text
}

func GenerateQuestionPack() QuestionPack {
	return QuestionPack{
		title: "Math",
		questions: []Question{
			{
				text: "2+2",
				variants: map[int]string{
					1: "1",
					2: "2",
					3: "3",
					4: "4",
				},
				answerNumber: 4,
			},
			{
				text: "2+3",
				variants: map[int]string{
					1: "5",
					2: "2",
					3: "7",
					4: "0",
				},
				answerNumber: 1,
			},
			{
				text: "2+4",
				variants: map[int]string{
					1: "1",
					2: "6",
					3: "3",
					4: "5",
				},
				answerNumber: 2,
			},
		},
	}
}
