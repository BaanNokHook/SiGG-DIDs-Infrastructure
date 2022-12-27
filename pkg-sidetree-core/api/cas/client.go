/*
SiGG-DIDs-Infrastructure
*/

package cas

// Client defines interface for accessing the underlying content addressable storage.
type Client interface {
	// Write writes the given content to CASClient.
	Write(content []byte) (string, error)

	// Read reads the content of the given address in CASClient.
	// returns the content of the given address.
	Read(address string) ([]byte, error)
}
