package dto

import "time"

// DocumentResponseDTO represents the struct that is returned by rest endpoints
type DocumentResponseDTO struct {

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

// DocumentRequestDTO represents the struct that is accepted as input for the rest endpoint
type DocumentRequestDTO struct {

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
