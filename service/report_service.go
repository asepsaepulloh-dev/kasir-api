package service

import (
	"kasir-api/model"
	"kasir-api/repository"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*model.DailyReport, error) {
	return s.repo.GetDailyReport()
}
