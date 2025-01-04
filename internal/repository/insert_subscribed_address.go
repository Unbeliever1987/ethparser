package repository

import "context"

// InsertSubscribedAddress implements Repository.
func (i *impl) InsertSubscribedAddress(ctx context.Context, address string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.subscribedAddresses = append(i.subscribedAddresses, address)

	return nil
}
