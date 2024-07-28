package searchstore

type User struct {
	EmpNum       string  `json:"EmpNum"`
	FullName     string  `json:"FullName"`
	FullNameKana string  `json:"FullNameKana"`
	Email        string  `json:"Email"`
	Score        float64 `json:"Score"`
}
