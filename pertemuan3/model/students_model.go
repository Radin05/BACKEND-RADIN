package model

type Students struct {
	StudentsID int64  `json:"students_id"`
	Name       string `json:"name"`
	Age        int64 `json:"age"`
	Address    string `json:"address"`
	Hoby       string `json:"hoby"`
}
