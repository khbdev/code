package model

type ResultItem struct {
    TeacherTelegramID int64  `json:"teacher_telegram_id"`
    QuestionID        string `json:"question_id"`
    Correct           string `json:"correct"`
}