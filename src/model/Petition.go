package model

type Petition struct {
    CreatedAt    DateTime `json:"createdAt"`
	Description  string   `json:"description"`
    Categories   []string `json:"categories"`
    Image        string   `json:"image"`
	ID			 string   `json:"id"`
}