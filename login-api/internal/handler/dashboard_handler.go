package handler

import (
	"encoding/json"
	"login-api/internal/storage/postgres"
	"net/http"

	"github.com/rs/zerolog/log"
)

type DashboardHandler struct {
	Store *postgres.PostgresPaymentStore
}

func NewDashboardHandler(store *postgres.PostgresPaymentStore) *DashboardHandler {
	return &DashboardHandler{Store: store}
}

// GetSummaryHandler menangani permintaan untuk data ringkasan dashboard.
func (h *DashboardHandler) GetSummaryHandler(w http.ResponseWriter, r *http.Request) {
	summary, err := h.Store.GetDashboardSummary()
	if err != nil {
		http.Error(w, `{"message":"Gagal mengambil data ringkasan."}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		log.Error().Err(err).Msg("Gagal melakukan encode response ringkasan")
	}
}

// GetChartDataHandler menangani permintaan untuk data grafik.
func (h *DashboardHandler) GetChartDataHandler(w http.ResponseWriter, r *http.Request) {
	chartData, err := h.Store.GetChartData()
	if err != nil {
		http.Error(w, `{"message":"Gagal mengambil data grafik."}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(chartData); err != nil {
		log.Error().Err(err).Msg("Gagal melakukan encode response data grafik")
	}
}