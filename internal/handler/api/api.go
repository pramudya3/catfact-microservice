package api

import (
	"context"
	"go-microservice/internal/service"
	"net/http"
)

type ApiServer struct {
	svc service.CatFactService
}

func NewApiServer(svc service.CatFactService) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (a *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", a.HandleGetCatFact)

	return http.ListenAndServe(listenAddr, nil)
}

func (a *ApiServer) HandleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := a.svc.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}
