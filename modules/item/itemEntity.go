package item

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Item struct {
		Id          primitive.ObjectID `json:"id" bson:"id,omitempty"`
		Title       string             `json:"title" bson:"title"`
		Price       float64            `json:"price" bson:"price"`
		Damage      int                `json:"damage" bson:"damage"`
		ImageUrl    string             `json:"image_url" bson:"image_url"`
		UsageStatus bool               `json:"usage_status" bson:"usage_status"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
