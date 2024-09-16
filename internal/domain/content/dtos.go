package content

type CreateContentRequestDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ContentDto struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
