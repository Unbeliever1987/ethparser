package repository

import "context"

// RetrieveAllSubcribedAddresses implements Repository.
func (i *impl) RetrieveAllSubcribedAddresses(ctx context.Context) ([]string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.subscribedAddresses, nil
}
