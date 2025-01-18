package handler

import (
	"database/sql"
	"go-template/pages"
	"go-template/store"
	"net/http"
)

type Home struct {
	DB *sql.DB
}

func NewHome(db *sql.DB) *Home {
	return &Home{
		DB: db,
	}
}

func (h *Home) Index(w http.ResponseWriter, r *http.Request) {
	u, err := store.New(h.DB).GetUserByEmail(r.Context(), "test@example.com")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	pages.Home(u).Render(r.Context(), w)
}
