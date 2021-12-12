export type Topic = {
    Tags: string[],
    isAnswered: boolean,
    ViewCount: number,
    AnswerCount: number,
    Score: number,
    LastActivityDate: number,
    CreationDate: number,
    QuestionID: number,
    Link: string,
    Title: string
}


// Tags             []string `json:"tags"`
// 	IsAnswered       bool     `json:"is_answered"`
// 	ViewCount        int      `json:"view_count"`
// 	AnswerCount      int      `json:"answer_count"`
// 	Score            int      `json:"score"`
// 	LastActivityDate int      `json:"last_activity_date"`
// 	CreationDate     int      `json:"creation_date"`
// 	QuestionID       int      `json:"question_id"`
// 	ContentLicense   string   `json:"content_license"`
// 	Link             string   `json:"link"`
// 	Title            string   `json:"title"`