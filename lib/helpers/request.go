package helpers

import "connectrpc.com/connect"

// NewRequestWithMetadata creates a new Connect request with the provided message and metadata.
func NewRequestWithMetadata[T any](message *T, metadata map[string]string) *connect.Request[T] {
	req := connect.NewRequest(message)
	for key, value := range metadata {
		req.Header().Set(key, value)
	}
	return req
}
