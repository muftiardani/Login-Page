package handler

import (
	"encoding/json"
	"login-api/internal/storage/postgres"
	"net/http"

	"github.com/rs/zerolog/log"
)

type PaymentHandler struct {
	Store *postgres.PostgresPaymentStore
}

func NewPaymentHandler(store *postgres.PostgresPaymentStore) *PaymentHandler {
	return &PaymentHandler{Store: store}
}

// GetPaymentsHandler menangani permintaan untuk mengambil data pembayaran.
func (h *PaymentHandler) GetPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	payments, err := h.Store.GetPayments()
	if err != nil {
		http.Error(w, `{"message":"Gagal mengambil data pembayaran."}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(payments); err != nil {
		log.Error().Err(err).Msg("Gagal melakukan encode response pembayaran")
	}
}