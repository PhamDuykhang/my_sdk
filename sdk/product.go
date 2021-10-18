package sdk

import "gopkg.in/mgo.v2/bson"

type (
	Product struct {
		ID          bson.ObjectId `json:"id,omitempty"`
		Name        string        `json:"name,omitempty"`
		Price       float64       `json:"price,omitempty"`
		Description string        `json:"description,omitempty"`
	}
)
