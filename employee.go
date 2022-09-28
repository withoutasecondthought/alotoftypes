package alotoftypes

import "go.mongodb.org/mongo-driver/bson/primitive"

type MessageJSON struct {
	Message string `json:"message"`
}

type ListOfEmployeesJSON struct {
	Id    *primitive.ObjectID `json:"_id" bson:"_id" binding:"required"`
	Name  string              `json:"name" bson:"name"`
	Phone string              `json:"phone" bson:"phone"`
	Work  `json:"work" bson:"work"`
}

type EmployeeJSON struct {
	Id                  *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                string              `json:"name" bson:"name" binding:"required"`
	Surname             string              `json:"surname" bson:"surname" binding:"required"`
	Age                 int                 `json:"age" bson:"age" binding:"required"`
	Phone               string              `json:"phone" bson:"phone" binding:"required,e164"`
	PhoneModel          string              `json:"phoneModel" bson:"phoneModel" binding:"required"`
	Email               string              `json:"email" json:"email" binding:"required,email"`
	EmailInbox          string              `json:"emailInbox" bson:"emailInbox" binding:"required"`
	Family              `json:"family" bson:"family" binding:"required"`
	Dates               `json:"dates" bson:"dates" binding:"required"`
	Work                `json:"work" bson:"work" binding:"required"`
	Address             string           `json:"address" bson:"address" binding:"required"`
	RegistrationAddress string           `json:"registrationAddress" bson:"registrationAddress" binding:"required"`
	FavoriteLetter      primitive.Symbol `json:"favoriteLetter" bson:"favoriteLetter" binding:"required"`
}

type Family struct {
	Mother       string   `json:"mother,omitempty" bson:"mother,omitempty"`
	Father       string   `json:"father,omitempty" bson:"father,omitempty"`
	Children     []Child  `json:"children,omitempty" bson:"children,omitempty"`
	OtherMembers []string `json:"otherMembers,omitempty" bson:"otherMembers,omitempty"`
}

type Child struct {
	Name            string `json:"name" bson:"name"`
	IsBoy           bool   `json:"isBoy" bson:"isBoy"`
	TrainingGrounds string `json:"trainingGrounds" bson:"trainingGrounds"`
}

type Dates struct {
	Start    *primitive.DateTime `json:"start" bson:"start"`
	End      *primitive.DateTime `json:"end,omitempty" bson:"end,omitempty"`
	Birth    *primitive.DateTime `json:"birth" bson:"birth"`
	Death    *primitive.DateTime `json:"death,omitempty" bson:"death,omitempty"`
	Birthday *primitive.DateTime `json:"birthday" bson:"birthday"`
}

type Work struct {
	Hours     uint64 `json:"hours" bson:"hours"`
	DealsDone uint64 `json:"dealsDone" bson:"dealsDone"`
	WeekHours uint64 `json:"weekHours" bson:"weekHours"`
	WeekDeals uint64 `json:"weekDeals" bson:"weekDeals"`
}

type EmployeeUpdateJSON struct {
	Id                  *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                string              `json:"name,omitempty,omitempty" bson:"name,omitempty" `
	Surname             string              `json:"surname,omitempty" bson:"surname,omitempty"`
	Age                 int                 `json:"age,omitempty" bson:"age,omitempty" `
	Phone               string              `json:"phone,omitempty" bson:"phone,omitempty" binding:"required,e164"`
	PhoneModel          string              `json:"phoneModel,omitempty" bson:"phoneModel,omitempty" `
	Email               string              `json:"email,omitempty" bson:"email,omitempty" binding:"email"`
	EmailInbox          string              `json:"emailInbox,omitempty" bson:"emailInbox,omitempty"`
	Family              *Family             `json:"family,omitempty" bson:"family,omitempty"`
	Dates               *Dates              `json:"dates,omitempty" bson:"dates,omitempty"`
	Work                *Work               `json:"work,omitempty" bson:"work,omitempty"`
	Address             string              `json:"address,omitempty" bson:"address,omitempty" `
	RegistrationAddress string              `json:"registrationAddress,omitempty" bson:"registrationAddress,omitempty" `
	FavoriteLetter      *primitive.Symbol   `json:"favoriteLetter,omitempty" bson:"favoriteLetter,omitempty" `
}
