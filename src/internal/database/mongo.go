package database

import (
	"context"
	"log"
	"time"

	"backend/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client                *mongo.Client
	HotelCollection       *mongo.Collection
	ReservationCollection *mongo.Collection
}

// ConnectDB initializes connection to MongoDB
func ConnectDB(uri string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database("devops_booking")

	mongoDB := &MongoDB{
		Client:                client,
		HotelCollection:       db.Collection("hotels"),
		ReservationCollection: db.Collection("reservations"),
	}

	// Seed data if database is empty
	mongoDB.seedHotels()

	return mongoDB, nil
}

// seedHotels inserts 15 real-world hotels inspired by Booking.com
func (m *MongoDB) seedHotels() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, _ := m.HotelCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		hotels := []interface{}{
			model.Hotel{Name: "The Plaza Hotel", Description: "Iconic luxury hotel near Central Park with world-class dining.", Location: "New York, USA", Price: 850.0},
			model.Hotel{Name: "Marina Bay Sands", Description: "Features the world's largest rooftop infinity pool and a massive casino.", Location: "Singapore", Price: 620.0},
			model.Hotel{Name: "Burj Al Arab Jumeirah", Description: "Ultra-luxury 7-star experience on a private artificial island.", Location: "Dubai, UAE", Price: 1500.0},
			model.Hotel{Name: "Hotel Ritz Paris", Description: "Elegant rooms reflecting French art de vivre, located in the heart of Paris.", Location: "Paris, France", Price: 1100.0},
			model.Hotel{Name: "Claridge's", Description: "Art Deco masterpiece in Mayfair, offering legendary service and afternoon tea.", Location: "London, UK", Price: 780.0},
			model.Hotel{Name: "Atlantis The Palm", Description: "Ocean-themed resort featuring an underground aquarium and waterpark.", Location: "Dubai, UAE", Price: 450.0},
			model.Hotel{Name: "Bellagio Las Vegas", Description: "Famous for its synchronized fountains, high-stakes casino, and botanical gardens.", Location: "Las Vegas, USA", Price: 290.0},
			model.Hotel{Name: "Amangiri", Description: "Secluded modernist resort blended into Utah's dramatic desert canyon landscape.", Location: "Utah, USA", Price: 1900.0},
			model.Hotel{Name: "Cuca Canela Resort", Description: "Tropical paradise with private villas overlooking rice terraces.", Location: "Bali, Indonesia", Price: 180.0},
			model.Hotel{Name: "Soneva Jani", Description: "Overwater villas with retractable roofs and private water slides into the lagoon.", Location: "Maldives", Price: 2200.0},
			model.Hotel{Name: "Hotel Danieli", Description: "Historic Venetian palace steps away from St. Mark's Square.", Location: "Venice, Italy", Price: 540.0},
			model.Hotel{Name: "Keio Plaza Hotel", Description: "High-rise hotel in Shinjuku with stunning views of the Tokyo skyline.", Location: "Tokyo, Japan", Price: 240.0},
			model.Hotel{Name: "Taj Mahal Palace", Description: "Flagship heritage hotel offering breathtaking views of the Gateway of India.", Location: "Mumbai, India", Price: 310.0},
			model.Hotel{Name: "Belmond Hotel das Cataratas", Description: "The only hotel located inside Brazil's Iguaçu National Park, next to the falls.", Location: "Iguazu, Brazil", Price: 490.0},
			model.Hotel{Name: "Grand Hotel Tremezzo", Description: "Art Nouveau palace offering spectacular views over Lake Como and the Alps.", Location: "Lake Como, Italy", Price: 950.0},
		}
		_, err := m.HotelCollection.InsertMany(ctx, hotels)
		if err != nil {
			log.Printf("Failed to seed database: %v", err)
		} else {
			log.Println("Database successfully seeded with 15 Booking.com hotels!")
		}
	}
}
