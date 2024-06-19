package clients

import (
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	ucprovider "github.com/CloudSilk/usercenter/provider"
)

func Init(serviceMode string) {
	if serviceMode == "ALL" {
		userProvider := new(ucprovider.UserProvider)
		UserClient.Add = userProvider.Add
		UserClient.Delete = userProvider.Delete
		UserClient.Export = userProvider.Export
		UserClient.GetDetail = userProvider.GetDetail
		UserClient.Query = userProvider.Query
		UserClient.Update = userProvider.Update
		UserClient.LoginByStaffNo = userProvider.LoginByStaffNo
		UserClient.LogoutByUserName = userProvider.LogoutByUserName
	} else {
		if os.Getenv("MES_DISABLE_AUTH") != "true" {
			config.SetConsumerService(UserClient)
		}
		config.SetConsumerService(ProductionStationClient)
		config.SetConsumerService(ProductionStationSignupClient)
		config.SetConsumerService(ProductionLineClient)
		config.SetConsumerService(ProductAttributeClient)
		config.SetConsumerService(ProductionCrosswayClient)
		config.SetConsumerService(MaterialTrayClient)
		config.SetConsumerService(MaterialTrayBindingRecordClient)
		config.SetConsumerService(MaterialChannelLayerClient)
		config.SetConsumerService(ProductPackageRecordClient)
		config.SetConsumerService(ProductInfoClient)
		config.SetConsumerService(ProductOrderClient)
		config.SetConsumerService(ProductRhythmRecordClient)
		config.SetConsumerService(ProductProcessRouteClient)
		config.SetConsumerService(ProductionProcessClient)
		config.SetConsumerService(ProductionProcessSopClient)
		config.SetConsumerService(ProductModelClient)
		config.SetConsumerService(PersonnelQualificationClient)
		config.SetConsumerService(SystemEventClient)
		config.SetConsumerService(SystemEventTriggerClient)
		config.SetConsumerService(SystemEventTriggerParameterClient)
		config.SetConsumerService(ProductReworkRecordClient)
		config.SetConsumerService(ProductOrderProcessClient)
		config.SetConsumerService(ProductOrderProcessStepClient)
		config.SetConsumerService(ProductionStationOutputClient)
		config.SetConsumerService(ProductOrderAttributeClient)
		config.SetConsumerService(ProcessStepParameterClient)
		config.SetConsumerService(ProductionProcessStepClient)
		config.SetConsumerService(ProductTestRecordClient)
		config.SetConsumerService(ProductWorkRecordClient)
	}
}
