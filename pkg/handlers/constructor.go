package handlers

type Handler struct {
	Repo Repo
}

func NewHandler(repo Repo) *Handler {
	return &Handler{Repo: repo}
}
