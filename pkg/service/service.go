package service

import (
	"alotoftypes"
	"alotoftypes/pkg/repository"
	"context"
)

type Employee interface {
	GetEmployee(ctx context.Context, id string) (alotoftypes.EmployeeJSON, error)
	PostEmployee(ctx context.Context, employee alotoftypes.EmployeeJSON) error
	PutEmployee(ctx context.Context, employee alotoftypes.EmployeeUpdateJSON) error
	DeleteEmployee(ctx context.Context, id string) error
	GetEmployeeWithParams(ctx context.Context, page uint, params interface{}) ([]*alotoftypes.ListOfEmployeesJSON, error)
}

type Service struct {
	Employee
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(r.Employee),
	}
}
