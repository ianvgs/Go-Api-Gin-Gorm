package mongomodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"id,omitempty"`
	Username string             `json:"username,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
