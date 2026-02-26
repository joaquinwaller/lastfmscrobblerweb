package auth

import "net/http"

type Handler struct{
	Service *Service
}

func (h *Handler) StartLastFM (w http.ResponseWriter, r *http.Response){


}