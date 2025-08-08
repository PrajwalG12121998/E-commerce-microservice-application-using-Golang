package order

// Client represents a client for the order service
type Client struct {
	url string
}

// NewClient creates a new order service client
func NewClient(url string) (*Client, error) {
	return &Client{url: url}, nil
}

// Close closes the client connection
func (c *Client) Close() error {
	return nil
}
