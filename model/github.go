package model

type CreateFileRequest struct {
	Message     string `json:"message"`
	Content     string `json:"content"`
	Branch      string `json:"branch"`
	AuthorName  string `json:"committer"`
	AuthorEmail string `json:"committer_email"`
}
