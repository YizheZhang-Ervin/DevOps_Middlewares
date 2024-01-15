package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "xxxx", "xxxx")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.Domain = "xx.yy.com" //域名
	request.PhoneNumbers = "123123"
	request.SignName = "xxxx"
	request.TemplateCode = "SMS_123123"
	request.TemplateParam = `{"code":232323}`

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
