package handlers

import (
	r "cook_book/backend/internal/controllers/repositories"
	s "cook_book/backend/internal/controllers/services"
	"cook_book/backend/internal/db"
)

type Handler struct {
	cookService s.CookBookServices
}

type HandlerConfig struct {
	CookService s.CookBookServices
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		cookService: c.CookService,
	}
}

func InitAllHandlers() *Handler {
	db := db.Get()

	cookRepo := r.NewCookBookRepositories((r.CookBookRepoConfig{DB: db}))

	cookService := s.NewCookBookServices((&s.CookBookServiceConfig{Repo: cookRepo}))

	h := New(&HandlerConfig{
		CookService: cookService,
	})

	return h
}
