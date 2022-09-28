package service

import (
	"alotoftypes"
	"alotoftypes/pkg/repository"
	"context"
	"net/mail"
)

type EmployeeService struct {
	repo repository.Employee
}

func (e *EmployeeService) GetEmployee(ctx context.Context, id string) (alotoftypes.EmployeeJSON, error) {
	return e.repo.GetEmployee(ctx, id)
}

func (e *EmployeeService) PostEmployee(ctx context.Context, employee alotoftypes.EmployeeJSON) error {
	_, err := mail.ParseAddress(employee.Email)
	if err != nil {
		return err
	}
	return e.repo.PostEmployee(ctx, employee)
}

func (e *EmployeeService) PutEmployee(ctx context.Context, employee alotoftypes.EmployeeUpdateJSON) error {
	if employee.Email != "" {
		_, err := mail.ParseAddress(employee.Email)
		if err != nil {
			return err
		}
	}
	return e.repo.PutEmployee(ctx, employee)
}

func (e *EmployeeService) DeleteEmployee(ctx context.Context, id string) error {
	return e.repo.DeleteEmployee(ctx, id)
}

func (e *EmployeeService) GetEmployeeWithParams(ctx context.Context, page uint, params interface{}) ([]*alotoftypes.ListOfEmployeesJSON, error) {
	return e.repo.GetEmployeeWithParams(ctx, page, params)
}

func NewEmployeeService(repo repository.Employee) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}
