package searchstore

type User struct {
	EmpNum       string  `json:"EmpNum"`
	FullName     string  `json:"FullName"`
	FullNameKana string  `json:"FullNameKana"`
	Email        string  `json:"email"`
	Score        float64 `json:"score"`
}
