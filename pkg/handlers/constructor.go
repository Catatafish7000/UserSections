package handlers

type Handler struct {
	repo repo
}

func NewHandler(repo repo) *Handler {
	return &Handler{repo: repo}
}
