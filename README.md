# data-platform-api-fin-inst-account-exconf-rmq-kube
data-platform-api-fin-inst-account-exconf-rmq-kube は、データ連携基盤において、API で 金融機関口座の存在性チェックを行うためのマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 存在確認先テーブル名
以下のsqlファイルに対して、金融機関口座の存在確認が行われます。

* data-platform-fin-inst-account-header-data.sql（データ連携基盤 金融機関口座 - ヘッダデータ）
* data-platform-fin-inst-account-item-data.sql（データ連携基盤 金融機関口座 - 明細データ）

## caller.go による存在性確認
Input で取得されたファイルに基づいて、caller.go で、 API がコールされます。
caller.go の 以下の箇所が、指定された API をコールするソースコードです。

```
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
```

## Input
data-platform-api-fin-inst-account-exconf-rmq-kube では、以下のInputファイルをRabbitMQからJSON形式で受け取ります。  

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"api_status_code": 200,
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"service_label": "FIN_INST_ACCOUNT",
	"FinInstAccountHeader": {
		"FinInstCountry": "JP",
        "FinInstCode":
        "FinInstBusinessPartner":
        "InternalFinInstCustomerID":
		"AccountBusinessPartner":
        "FinInstCustomerIDByFinInst":
        "ValidityEndDate":
        "ValidityStartDate":
        "IsMarkedForDeletion":
	},
	"api_schema": "DPFMFinInstAccountCreates",
	"accepter": ["All"],
	"deleted": false
}

```
