package application

import (
	"github.com/serdarkalayci/gitdoc/domain"
)

// DocumentRepository is the interface that we expect to be fulfilled to be used as a backend for Document Service
type DocumentRepository interface {
	List() ([]domain.Document, error)
	Add(document domain.Document) (domain.Document, error)
	Get(string) (domain.Document, error)
	Update(string, domain.Document) error
	Delete(string) error
}

// DocumentService represents the struct which contains a DocumentRepository and exports methods to access the data
type DocumentService struct {
	documentRepo DocumentRepository
}

// NewDocumentService creates a new DocumentService instance and sets its repository
func NewDocumentService(dr DocumentRepository) DocumentService {
	if dr == nil {
		panic("missing documentRepository")
	}
	return DocumentService{
		documentRepo: dr,
	}
}

// List loads all the data from the included repository and returns them
// Returns an error if the repository returns one
func (ps DocumentService) List() ([]domain.Document, error) {
	documents, err := ps.documentRepo.List()
	return documents, err
}

// Add adds a new document to the included repository, and returns it
// Returns an error if the repository returns one
func (ps DocumentService) Add(p domain.Document) (domain.Document, error) {
	document, err := ps.documentRepo.Add(p)
	return document, err
}

// Get selects the document from the included repository with the given unique identifier, and returns it
// Returns an error if the repository returns one
func (ps DocumentService) Get(id string) (domain.Document, error) {
	document, err := ps.documentRepo.Get(id)
	return document, err
}

// Update updates the document on the included repository with the given unique identifier, and returns it
// Returns an error if the repository returns one
func (ps DocumentService) Update(id string, p domain.Document) error {
	err := ps.documentRepo.Update(id, p)
	return err
}

// Delete deletes the document from the included repository with the given unique identifier
// Returns an error if the repository returns one
func (ps DocumentService) Delete(id string) error {
	err := ps.documentRepo.Delete(id)
	return err
}
