package domain

import (
	"time"
)

// Document represents a document to be indexed.
type Document struct {
	// ID is the unique identifier of the document.
	ID string `json:"id"`
	// Name is the name of the document.
	Name string `json:"name"`
	// Content is the content of the document.
	Content string `json:"content"`
	// CreatedAt is the creation date of the document.
	CreatedAt time.Time `json:"createdAt"`
	// LastUpdatedAt is the last update date of the document.
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	// LastUpdatedBy is the last user who updated the document.
	LastUpdatedBy string `json:"lastUpdatedBy"`
}
