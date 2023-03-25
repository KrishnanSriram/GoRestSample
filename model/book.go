package model

type Book struct {
	BookId  string   `json:"book_id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

type BookRequest struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
