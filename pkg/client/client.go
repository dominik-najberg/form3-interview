package client

import (
	"github.com/dominik-najberg/form3-interview/pkg/client/accounts"
	"net/http"
)

// Form3Client Api client
type Form3Client interface {
	ListAccounts(request accounts.ListAccountsRequest) (accounts.ListAccountsResponse, error)
}

type form3Client struct {
	Client *http.Client
}

// New creates a new Form3Client client
func New() Form3Client {
	return &form3Client{&http.Client{}}
}

// NewWithClient creates a new Form3Client client with custom http.Form3Client
func NewWithClient(client *http.Client) Form3Client {
	return &form3Client{client}
}
