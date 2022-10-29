package memory

// HealthRepository represent a structure that will communicate to memory data store to accomplish health related transactions
type HealthRepository struct {
}

func newHealthRepository() HealthRepository {
	return HealthRepository{}
}

// Ready checks the memory data store connection, hence always returns true
func (hr HealthRepository) Ready() bool {
	return true

}
