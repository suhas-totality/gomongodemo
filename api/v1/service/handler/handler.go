package handler

import (
	"github.com/suhas-totality/gomongodemo/api/v1/service"
	"github.com/suhas-totality/gomongodemo/api/v1/service/store"
)

type Handler struct {
	serviceRegistry *service.ServiceRegistry

	store store.Store
}

func NewHandler(svc *service.ServiceRegistry) Handler {
	return Handler{
		serviceRegistry: svc,
		store:           *store.NewStore(),
	}
}
