package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Location    string             `bson:"location" json:"location"`
	Price       float64            `bson:"price" json:"price"`
}

type Reservation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName string             `bson:"full_name" json:"full_name"`
	Email    string             `bson:"email" json:"email"`
	CheckIn  string             `bson:"check_in" json:"check_in"`
	CheckOut string             `bson:"check_out" json:"check_out"`
	HotelID  string             `bson:"hotel_id" json:"hotel_id"`
}
