package repository

import (
	"database/sql"
	"kasir-api/model"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetDailyReport() (*model.DailyReport, error) {
	report := &model.DailyReport{}

	query := `SELECT 
				COALESCE(SUM(total_amount), 0) as total_revenue,
				COUNT(*) as total_transaksi
			  FROM transactions 
			  WHERE DATE(created_at) = CURRENT_DATE`

	err := repo.db.QueryRow(query).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	topProductQuery := `SELECT 
							p.name,
							COALESCE(SUM(td.quantity), 0) as qty_terjual
						FROM transaction_details td
						JOIN transactions t ON td.transaction_id = t.id
						JOIN products p ON td.product_id = p.id
						WHERE DATE(t.created_at) = CURRENT_DATE
						GROUP BY p.id, p.name
						ORDER BY qty_terjual DESC
						LIMIT 1`

	var topProduct model.TopProduct
	err = repo.db.QueryRow(topProductQuery).Scan(&topProduct.Nama, &topProduct.QtyTerjual)
	if err == sql.ErrNoRows {
		return report, nil
	}
	if err != nil {
		return nil, err
	}

	report.ProdukTerlaris = &topProduct
	return report, nil
}
