package accounts

import "github.com/dominik-najberg/form3-interview/pkg/client"

func (c client.Form3Client) ListAccounts(request ListAccountsRequest) (ListAccountsResponse, error) {
	panic("implement me")
}

type ListAccountsRequest struct {
	PageNumber int
	PageSize   int64
}

type ListAccountsResponse struct {
	Data []ListAccountsItem `json:"data"`
}

type ListAccountsItem struct {
	Type           string `json:"type"`
	ID             string `json:"id"`
	OrganisationID string `json:"organisation_id"`
	Version        int    `json:"version"`
	Attributes     struct {
		Country               string `json:"country"`
		BaseCurrency          string `json:"base_currency"`
		AccountNumber         string `json:"account_number"`
		BankID                string `json:"bank_id"`
		BankIDCode            string `json:"bank_id_code"`
		Bic                   string `json:"bic"`
		Iban                  string `json:"iban"`
		AccountClassification string `json:"account_classification"`
		JointAccount          bool   `json:"joint_account"`
		AccountMatchingOptOut bool   `json:"account_matching_opt_out"`
	} `json:"attributes"`
}
