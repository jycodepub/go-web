package book

type Book struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Isbn    string   `json:"isbn"`
}
