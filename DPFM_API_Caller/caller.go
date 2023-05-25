package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-fin-inst-account-exconf-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-fin-inst-account-exconf-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type ExistenceConf struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewExistenceConf(ctx context.Context, db *database.Mysql, l *logger.Logger) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (e *ExistenceConf) Conf(msg rabbitmq.RabbitmqMessage) interface{} {
	var ret interface{}
	ret = map[string]interface{}{
		"ExistenceConf": false,
	}
	input := make(map[string]interface{})
	err := json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return ret
	}

	_, ok := input["FinInstAccountHeader"]
	if ok {
		input := &dpfm_api_input_reader.HeaderSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confFinInstAccountHeader(input)
		goto endProcess
	}
	_, ok = input["FinInstAccountItem"]
	if ok {
		input := &dpfm_api_input_reader.ItemSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.ConfFinInstAccountItem(input)
		goto endProcess
	}

endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}

func (e *ExistenceConf) confFinInstAccountHeader(input *dpfm_api_input_reader.HeaderSDC) *dpfm_api_output_formatter.FinInstAccountHeader {
	exconf := dpfm_api_output_formatter.FinInstAccountHeader{
		ExistenceConf: false,
	}
	if input.FinInstAccountHeader.FinInstCode == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.FinInstCountry == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.FinInstBusinessPartner == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.InternalFinInstCustomerID == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.AccountBusinessPartner == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.FinInstCustomerIDByFinInst == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.ValidityEndDate == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.ValidityStartDate == nil {
		return &exconf
	}
	if input.FinInstAccountHeader.IsMarkedForDeletion == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.FinInstAccountHeader{
		FinInstCountry:             *input.FinInstAccountHeader.FinInstCountry,
		FinInstCode:                *input.FinInstAccountHeader.FinInstCode,
		FinInstBusinessPartner:     *input.FinInstAccountHeader.FinInstBusinessPartner,
		InternalFinInstCustomerID:  *input.FinInstAccountHeader.InternalFinInstCustomerID,
		AccountBusinessPartner:     *input.FinInstAccountHeader.AccountBusinessPartner,
		FinInstCustomerIDByFinInst: *input.FinInstAccountHeader.FinInstCustomerIDByFinInst,
		ValidityEndDate:            *input.FinInstAccountHeader.ValidityEndDate,
		ValidityStartDate:          *input.FinInstAccountHeader.ValidityStartDate,
		IsMarkedForDeletion:        *input.FinInstAccountHeader.IsMarkedForDeletion,
		ExistenceConf:              false,
	}

	rows, err := e.db.Query(
		`SELECT FinInstName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_account_header_data 
		WHERE (finInstCountry, finInstCode,	finInstBusinessPartner, internalFinInstCustomerID, accountBusinessPartner, finInstCustomerIDByFinInst, validityEndDate, validityStartDate, isMarkedForDeletion) = (?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		exconf.FinInstCountry, exconf.FinInstCode, exconf.FinInstBusinessPartner, exconf.InternalFinInstCustomerID, exconf.AccountBusinessPartner, exconf.FinInstCustomerIDByFinInst, exconf.ValidityEndDate, exconf.ValidityStartDate, exconf.IsMarkedForDeletion,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfFinInstAccountItem(input *dpfm_api_input_reader.ItemSDC) *dpfm_api_output_formatter.FinInstAccountItem {
	exconf := dpfm_api_output_formatter.FinInstAccountItem{
		ExistenceConf: false,
	}
	if input.FinInstAccountItem.FinInstCode == nil {
		return &exconf
	}
	if input.FinInstAccountItem.FinInstCountry == nil {
		return &exconf
	}
	if input.FinInstAccountItem.FinInstBranchCode == nil {
		return &exconf
	}
	if input.FinInstAccountItem.FinInstFullCode == nil {
		return &exconf
	}
	if input.FinInstAccountItem.InternalFinInstCustomerID == nil {
		return &exconf
	}
	if input.FinInstAccountItem.InternalFinInstAccountID == nil {
		return &exconf
	}
	if input.FinInstAccountItem.ValidityEndDate == nil {
		return &exconf
	}
	if input.FinInstAccountItem.ValidityStartDate == nil {
		return &exconf
	}
	if input.FinInstAccountItem.FinInstControlKey == nil {
		return &exconf
	}
	if input.FinInstAccountItem.FinInstAccountName == nil {
		return &exconf
	}
	if input.FinInstAccountItem.FinInstAccount == nil {
		return &exconf
	}
	if input.FinInstAccountItem.IsMarkedForDeletion == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.FinInstAccountItem{
		FinInstCountry:            *input.FinInstAccountItem.FinInstCountry,
		FinInstCode:               *input.FinInstAccountItem.FinInstCode,
		FinInstBranchCode:         *input.FinInstAccountItem.FinInstBranchCode,
		FinInstFullCode:           *input.FinInstAccountItem.FinInstFullCode,
		InternalFinInstCustomerID: *input.FinInstAccountItem.InternalFinInstCustomerID,
		InternalFinInstAccountID:  *input.FinInstAccountItem.InternalFinInstAccountID,
		ValidityEndDate:           *input.FinInstAccountItem.ValidityEndDate,
		ValidityStartDate:         *input.FinInstAccountItem.ValidityStartDate,
		FinInstControlKey:         *input.FinInstAccountItem.FinInstControlKey,
		FinInstAccountName:        *input.FinInstAccountItem.FinInstAccountName,
		FinInstAccount:            *input.FinInstAccountItem.FinInstAccount,
		IsMarkedForDeletion:       *input.FinInstAccountItem.IsMarkedForDeletion,
		ExistenceConf:             false,
	}
	rows, err := e.db.Query(
		`SELECT FinInstCode 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_account_item_data 
		WHERE (finInstCountry, finInstCode, finInstBranchCode, finInstFullCode, internalFinInstCustomerID, internalFinInstAccountID, validityEndDate, validityStartDate, finInstControlKey, finInstAccountName, finInstAccount, isMarkedForDeletion) = (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		exconf.FinInstCountry, exconf.FinInstCode, exconf.FinInstBranchCode, exconf.FinInstFullCode, exconf.InternalFinInstCustomerID, exconf.InternalFinInstAccountID, exconf.ValidityEndDate, exconf.ValidityStartDate, exconf.FinInstControlKey, exconf.FinInstAccountName, exconf.FinInstAccount, exconf.IsMarkedForDeletion,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
