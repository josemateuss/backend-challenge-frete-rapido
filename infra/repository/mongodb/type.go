package mongodb

import "time"

type quoteDB struct {
	Carriers  []carrierDB `bson:"carrier"`
	CreatedAt time.Time   `bson:"created_at"`
}

type carrierDB struct {
	Name     string  `bson:"name"`
	Service  string  `bson:"service"`
	Deadline uint    `bson:"deadline"`
	Price    float64 `bson:"price"`
}
