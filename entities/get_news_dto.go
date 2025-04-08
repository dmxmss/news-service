package entities

type GetNewsDto struct {
	Title string `json:"title"`
	Contents string `json:"contents"`
	AuthorID int `json:"author_id"`
	Approved bool `json:"approved"`
	ApprovedAt string `json:"approved_at"`
	Tags []Tag `json:"tags,omitempty"`
	Source string `json:"source"`
}
