package tencent

import (
	"github.com/ahui2012/go-ddns-client/config"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

type TencentCloud struct {
	Config  config.DomainConfig
	Client  *dnspod.Client
	Records map[string]*dnspod.RecordListItem
}

func (t *TencentCloud) Init() {
	credential := common.NewCredential(t.Config.SecretId, t.Config.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := dnspod.NewClient(credential, "", cpf)
	t.Client = client

	t.getRecordList()
}

func (t *TencentCloud) getRecordList() {
	t.Records = make(map[string]*dnspod.RecordListItem)
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(t.Config.Domain)
	response, err := t.Client.DescribeRecordList(request)
	if err != nil {
		panic(err)
	}

	records := response.Response.RecordList
	for _, item := range records {
		if *item.Type == "A" {
			t.Records[*item.Name] = item
		}
	}
}

func (t *TencentCloud) UpdateRecord(ip string) error {
	for _, name := range t.Config.SubDomains {
		request := dnspod.NewModifyRecordRequest()
		request.Domain = common.StringPtr(t.Config.Domain)
		request.SubDomain = common.StringPtr(name)
		request.RecordType = common.StringPtr("A")
		request.RecordLine = common.StringPtr("默认")
		request.Value = common.StringPtr(ip)
		request.RecordId = common.Uint64Ptr(*t.Records[name].RecordId)
		_, err := t.Client.ModifyRecord(request)
		return err
	}

	return nil
}
