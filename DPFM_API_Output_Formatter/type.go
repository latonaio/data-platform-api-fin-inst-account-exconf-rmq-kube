package dpfm_api_output_formatter

type MetaData struct {
	ConnectionKey        string                `json:"connection_key"`
	Result               bool                  `json:"result"`
	RedisKey             string                `json:"redis_key"`
	Filepath             string                `json:"filepath"`
	APIStatusCode        int                   `json:"api_status_code"`
	RuntimeSessionID     string                `json:"runtime_session_id"`
	BusinessPartnerID    *int                  `json:"business_partner"`
	ServiceLabel         string                `json:"service_label"`
	FinInstAccountHeader *FinInstAccountHeader `json:"FinInstAccountHeader,omitempty"`
	FinInstAccountItem   *FinInstAccountItem   `json:"FinInstAccountItem,omitempty"`
	APISchema            string                `json:"api_schema"`
	Accepter             []string              `json:"accepter"`
	Deleted              bool                  `json:"deleted"`
}

type FinInstAccountHeader struct {
	FinInstCountry             string `json:"FinInstCountry"`
	FinInstCode                string `json:"FinInstCode"`
	FinInstBusinessPartner     int    `json:"FinInstBusinessPartner"`
	InternalFinInstCustomerID  int    `json:"InternalFinInstCustomerID"`
	AccountBusinessPartner     int    `json:"AccountBusinessPartner"`
	FinInstCustomerIDByFinInst string `json:"FinInstCustomerIDByFinInst"`
	ValidityEndDate            string `json:"ValidityEndDate"`
	ValidityStartDate          string `json:"ValidityStartDate"`
	IsMarkedForDeletion        bool   `json:"IsMarkedForDeletion"`
	ExistenceConf              bool   `json:"ExistenceConf"`
}

type FinInstAccountItem struct {
	FinInstCountry            string `json:"FinInstCountry"`
	FinInstCode               string `json:"FinInstCode"`
	FinInstBranchCode         string `json:"FinInstBranchCode"`
	FinInstFullCode           string `json:"FinInstFullCode"`
	InternalFinInstCustomerID int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  int    `json:"InternalFinInstAccountID"`
	ValidityEndDate           string `json:"ValidityEndDate"`
	ValidityStartDate         string `json:"ValidityStartDate"`
	FinInstControlKey         string `json:"FinInstControlKey"`
	FinInstAccountName        string `json:"FinInstAccountName"`
	FinInstAccount            string `json:"FinInstAccount"`
	IsMarkedForDeletion       bool   `json:"IsMarkedForDeletion"`
	ExistenceConf             bool   `json:"ExistenceConf"`
}
