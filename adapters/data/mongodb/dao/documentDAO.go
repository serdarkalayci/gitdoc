package dao

import "time"

// DocumentDAO represents the struct of document type to be stored in mongoDB
type DocumentDAO struct {
	ID            string    `bson:"uuid"`
	Name          string    `bson:"Name"`
	Content       string    `bson:"Content"`
	CreatedAt     time.Time `bson:"CreatedAt"`
	LastUpdatedAt time.Time `bson:"LastUpdatedAt"`
	LastUpdatedBy string    `bson:"Survived"`
}
