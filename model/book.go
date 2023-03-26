package model

// BookRequest struct (Model)
type BookRequest struct {
	//Title   string   `json:"title" binding: "required"`
	//Authors []string `json:"authors" binding: "required" `
	Title   string   `json:"title" binding:"required"`
	Authors []string `json:"author" binding:"required"`
}

// Book struct (Model)
type Book struct {
	BookId  string   `json:"book_id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
