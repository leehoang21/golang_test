package report

import (
	"base/internal/repository"
	"base/internal/utils/validator"
	"context"
)

type Service interface {
	GenerateMorningReport(ctx context.Context) (string, error)
}

type reportService struct {
	taskRepo  repository.Task
	validator validator.Validator
}

func NewReportService(taskRepo repository.Task,
	validator validator.Validator) Service {
	return &reportService{
		taskRepo:  taskRepo,
		validator: validator,
	}
}
