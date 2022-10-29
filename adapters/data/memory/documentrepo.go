package memory

import (
	"github.com/serdarkalayci/gitdoc/domain"
)

type DocumentRepository struct {
}

func newDocumentRepository() DocumentRepository {
	return DocumentRepository{}
}

// List loads all the document records from tha database and returns it
// Returns an error if database fails to provide service
func (pr DocumentRepository) List() ([]domain.Document, error) {

	documents := make([]domain.Document, 0)
	return documents, nil
}

// Add adds a new document to the underlying database.
// It returns the document inserted on success or error
func (pr DocumentRepository) Add(p domain.Document) (domain.Document, error) {

	return domain.Document{}, nil
}

// Get selects a single document from the database with the given unique identifier
// Returns an error if database fails to provide service
func (pr DocumentRepository) Get(id string) (domain.Document, error) {
	return domain.Document{}, nil
}

// Update updates fields of a single document from the database with the given unique identifier
// Returns an error if database fails to provide service
func (pr DocumentRepository) Update(id string, p domain.Document) error {

	return nil
}

// Delete selects a single document from the database with the given unique identifier
// Returns an error if database fails to provide service
func (pr DocumentRepository) Delete(id string) error {

	return nil
}
