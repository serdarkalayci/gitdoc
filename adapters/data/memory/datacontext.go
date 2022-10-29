package memory

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	DocumentRepository DocumentRepository
	HealthRepository   HealthRepository
}

// NewDataContext returns a new memory backed DataContext
func NewDataContext() (DataContext, error) {

	dataContext := DataContext{}
	dataContext.DocumentRepository = newDocumentRepository()
	dataContext.HealthRepository = newHealthRepository()
	return dataContext, nil
}
