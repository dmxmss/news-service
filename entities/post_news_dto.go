package entities

type PostNewsDto struct {
	Title string `json:"title"`
	Contents string `json:"contents"`
	Tags []Tag `json:"tags"`
	Source string `json:"source"`
}
