package postgres

import (
	"context"
	"login-api/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type PostgresPaymentStore struct {
	DB *pgxpool.Pool
}

func NewPostgresPaymentStore(db *pgxpool.Pool) *PostgresPaymentStore {
	return &PostgresPaymentStore{DB: db}
}

// GetPayments mengambil semua data pembayaran dari database.
func (s *PostgresPaymentStore) GetPayments() ([]model.Payment, error) {
	query := `SELECT id, customer_name, amount, status, payment_date 
              FROM payments 
              ORDER BY payment_date DESC`

	rows, err := s.DB.Query(context.Background(), query)
	if err != nil {
		log.Error().Err(err).Msg("Gagal menjalankan query untuk mengambil pembayaran")
		return nil, err
	}
	defer rows.Close()

	var payments []model.Payment
	for rows.Next() {
		var p model.Payment
		if err := rows.Scan(&p.ID, &p.CustomerName, &p.Amount, &p.Status, &p.PaymentDate); err != nil {
			log.Error().Err(err).Msg("Gagal memindai baris pembayaran")
			return nil, err
		}
		payments = append(payments, p)
	}

	return payments, nil
}

// GetDashboardSummary menghitung data ringkasan dari tabel payments.
func (s *PostgresPaymentStore) GetDashboardSummary() (model.DashboardSummary, error) {
	var summary model.DashboardSummary
	query := `
        SELECT 
            COALESCE(SUM(CASE WHEN status = 'Lunas' THEN amount ELSE 0 END), 0) as total_revenue,
            COALESCE(SUM(CASE WHEN status = 'Lunas' THEN 1 ELSE 0 END), 0) as completed_payments,
            COALESCE(SUM(CASE WHEN status = 'Tertunda' THEN 1 ELSE 0 END), 0) as pending_payments
        FROM payments;
    `
	err := s.DB.QueryRow(context.Background(), query).Scan(
		&summary.TotalRevenue,
		&summary.CompletedPayments,
		&summary.PendingPayments,
	)

	if err != nil {
		log.Error().Err(err).Msg("Gagal menjalankan query untuk ringkasan dashboard")
		return model.DashboardSummary{}, err
	}

	return summary, nil
}

// GetChartData mengambil data pembayaran yang diagregasi per hari untuk 7 hari terakhir.
func (s *PostgresPaymentStore) GetChartData() ([]model.ChartData, error) {
	query := `
        SELECT 
            TO_CHAR(date_series, 'YYYY-MM-DD') as label,
            COALESCE(SUM(amount), 0) as value
        FROM 
            generate_series(
                CURRENT_DATE - INTERVAL '6 days', 
                CURRENT_DATE, 
                '1 day'
            ) AS date_series
        LEFT JOIN 
            payments ON DATE(payment_date) = date_series AND status = 'Lunas'
        GROUP BY 
            date_series
        ORDER BY 
            date_series;
    `
	rows, err := s.DB.Query(context.Background(), query)
	if err != nil {
		log.Error().Err(err).Msg("Gagal menjalankan query untuk data grafik")
		return nil, err
	}
	defer rows.Close()

	var chartData []model.ChartData
	for rows.Next() {
		var cd model.ChartData
		if err := rows.Scan(&cd.Label, &cd.Value); err != nil {
			log.Error().Err(err).Msg("Gagal memindai baris data grafik")
			return nil, err
		}
		chartData = append(chartData, cd)
	}
	return chartData, nil
}