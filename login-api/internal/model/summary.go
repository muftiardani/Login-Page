package model

// DashboardSummary merepresentasikan data ringkasan untuk dashboard.
type DashboardSummary struct {
	TotalRevenue      float64 `json:"total_revenue"`
	CompletedPayments int64   `json:"completed_payments"`
	PendingPayments   int64   `json:"pending_payments"`
}