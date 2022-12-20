package dpfm_api_input_reader

import (
	"data-platform-api-fin-inst-account-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *HeaderSDC) ConvertToFinInstAccountHeader() *requests.FinInstAccountHeader {
	data := sdc.FinInstAccountHeader
	return &requests.FinInstAccountHeader{
		FinInstCountry:             data.FinInstCountry,
		FinInstCode:                data.FinInstCode,
		FinInstBusinessPartner:     data.FinInstBusinessPartner,
		InternalFinInstCustomerID:  data.InternalFinInstCustomerID,
		AccountBusinessPartner:     data.AccountBusinessPartner,
		FinInstCustomerIDByFinInst: data.FinInstCustomerIDByFinInst,
		ValidityEndDate:            data.ValidityEndDate,
		ValidityStartDate:          data.ValidityStartDate,
		IsMarkedForDeletion:        data.IsMarkedForDeletion,
	}
}

func (sdc *ItemSDC) ConvertToFinInstAccountItem() *requests.FinInstAccountItem {
	data := sdc.FinInstAccountItem
	return &requests.FinInstAccountItem{
		FinInstCountry:             data.FinInstCountry,
		FinInstCode:                data.FinInstCode,
		FinInstBranchCode:          data.FinInstBranchCode,
		FinInstFullCode:            data.FinInstFullCode,
		InternalFinInstCustomerID:  data.InternalFinInstCustomerID,
		InternalFinInstAccountID:   data.InternalFinInstAccountID,
		ValidityEndDate:            data.ValidityEndDate,
		ValidityStartDate:          data.ValidityStartDate,
		FinInstControlKey:          data.FinInstControlKey,
		FinInstAccountName:         data.FinInstAccountName,
		FinInstAccount:             data.FinInstAccount,
		IsMarkedForDeletion:        data.IsMarkedForDeletion,
	}
}
