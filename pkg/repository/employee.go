package repository

import (
	"alotoftypes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmployeeRepository struct {
	db *mongo.Collection
}

func (e *EmployeeRepository) GetEmployee(ctx context.Context, id string) (alotoftypes.EmployeeJSON, error) {
	obid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return alotoftypes.EmployeeJSON{}, err
	}

	var res alotoftypes.EmployeeJSON
	err = e.db.FindOne(ctx, bson.M{"_id": obid}).Decode(&res)
	if err != nil {
		return alotoftypes.EmployeeJSON{}, err
	}

	return res, nil
}

func (e *EmployeeRepository) PostEmployee(ctx context.Context, employee alotoftypes.EmployeeJSON) error {
	_, err := e.db.InsertOne(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepository) PutEmployee(ctx context.Context, employee alotoftypes.EmployeeUpdateJSON) error {
	_, err := e.db.UpdateOne(ctx, bson.M{"_id": employee.Id}, bson.D{
		{"$set", employee},
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepository) DeleteEmployee(ctx context.Context, id string) error {
	obid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = e.db.DeleteOne(ctx, bson.M{"_id": obid})
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepository) GetEmployeeWithParams(ctx context.Context, page uint, params interface{}) ([]*alotoftypes.ListOfEmployeesJSON, error) {
	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSkip(int64(page) * 20)

	var employees []*alotoftypes.ListOfEmployeesJSON

	cur, err := e.db.Find(ctx, params, opts)
	if err != nil {
		return employees, err
	}

	for cur.Next(ctx) {
		var emp *alotoftypes.ListOfEmployeesJSON
		err := cur.Decode(&emp)
		if err != nil {
			return employees, err
		}

		employees = append(employees, emp)
	}

	if len(employees) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return employees, nil
}

func NewEmployeeRepository(db *mongo.Database, collection string) *EmployeeRepository {
	return &EmployeeRepository{
		db: db.Collection(collection),
	}
}
