package common

type Page struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	ParentId *int   `json:"parent_id"`
}
