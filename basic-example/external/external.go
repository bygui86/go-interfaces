package external

// This package is considered as an external library, so UNMODIFIABLE

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetData() (string, error) {
	return "data", nil
}
