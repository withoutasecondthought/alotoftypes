package repository

import (
	"alotoftypes"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Employee interface {
	GetEmployee(ctx context.Context, id string) (alotoftypes.EmployeeJSON, error)
	PostEmployee(ctx context.Context, employee alotoftypes.EmployeeJSON) error
	PutEmployee(ctx context.Context, employee alotoftypes.EmployeeUpdateJSON) error
	DeleteEmployee(ctx context.Context, id string) error
	GetEmployeeWithParams(ctx context.Context, page uint, params interface{}) ([]*alotoftypes.ListOfEmployeesJSON, error)
}

type Repository struct {
	Employee
}

func NewRepository(db *mongo.Database, collection string) *Repository {
	return &Repository{
		Employee: NewEmployeeRepository(db, collection),
	}
}
