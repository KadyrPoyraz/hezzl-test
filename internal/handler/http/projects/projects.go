package projectshandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KadyrPoyraz/httplayout/internal/handler/http/handler"
	"github.com/KadyrPoyraz/httplayout/internal/service/projects"
	"github.com/gorilla/mux"
)

type projectsHandler struct {
	router *mux.Router

    projectsService projects.Service
}

func New(router *mux.Router, projectsService projects.Service) handler.Handler {
    handler := &projectsHandler{
		router: router,
        projectsService: projectsService,
	}
    handler.fill()

    return handler
}

func (h *projectsHandler) fill() {
	base := "/projects"
	r := h.router.PathPrefix(base).Subrouter()

	r.HandleFunc("/", h.handleHello)
}

func (h *projectsHandler) handleHello(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    projects, err := h.projectsService.GetProjects(ctx)
    if err != nil {
		fmt.Println(err)
        return
    }

	b, err := json.Marshal(projects)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(b))
}
