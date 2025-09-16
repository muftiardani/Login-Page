package model

// ChartData merepresentasikan data untuk satu titik pada grafik.
type ChartData struct {
	Label string  `json:"label"` // Contoh: "15 Sep"
	Value float64 `json:"value"` // Contoh: 1500000
}