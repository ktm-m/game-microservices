package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Credential struct {
		Id          primitive.ObjectID `json:"id" bson:"id,omitempty"`
		PlayerId    string             `json:"player_id" bson:"player_id"`
		RoleCode    int                `json:"role_code" bson:"role_code"`
		AccessToken string             `json:"access_token" bson:"access_token"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	}

	Role struct {
		Id    primitive.ObjectID `json:"id" bson:"id,omitempty"`
		Title string             `json:"title" bson:"title"`
		Code  int                `json:"code" bson:"code"`
	}
)
