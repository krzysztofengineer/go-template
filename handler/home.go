package handler

import (
	"database/sql"
	"go-template/components"
	"go-template/pages"
	"go-template/store"
	"net/http"
	"strconv"
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

func (h *Home) Increase(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	value := r.PostFormValue("value")
	i, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	i++

	components.Counter(strconv.Itoa(i)).Render(r.Context(), w)
}

func (h *Home) Decrease(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	value := r.PostFormValue("value")
	i, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	i--

	components.Counter(strconv.Itoa(i)).Render(r.Context(), w)
}
