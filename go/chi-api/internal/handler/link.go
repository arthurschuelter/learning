package handler

import (
	"chi-api/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LinksHandler struct {
	Service service.LinkServicer
}

func NewLinksHandler(s service.LinkServicer) *LinksHandler {
	return &LinksHandler{
		Service: s,
	}
}

func (l *LinksHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	links, err := l.Service.ListAllLinks()

	if err != nil {
		http.Error(w, "Could not retrieve links", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(links)
}

func (l *LinksHandler) FindLink(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	link, err := l.Service.GetURL(code)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, link.URL, http.StatusFound)
}

func (l *LinksHandler) SaveLink(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	link, err := l.Service.Save(req.URL)

	if err != nil {
		http.Error(w, "Could not create link", http.StatusInternalServerError)
		return
	}

	// longURL := chi.URLParam(r, "long_url")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(link)

}
