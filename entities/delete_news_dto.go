package entities

type DeleteNewsDto struct {
	Title string `json:"title"`
	Contents string `json:"contents"`
	AuthorID int `json:"author_id"`
}
