package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"strconv"
)

func main() {
	var regionId = flag.String("regionId", "", "区域标识")
	var accessKeyId = flag.String("id", "", "accessKeyId")
	var accessKeySecret = flag.String("secret", "", "accessKeySecret")
	var verifyCode = flag.String("code", "", "验证码")
	var phoneNumbers = flag.Int("phonenumbers", 0, "手机号")
	flag.Parse()

	if *phoneNumbers <= 0 {
		panic(fmt.Errorf("invalid phonenumbers"))
	}

	client, err := dysmsapi.NewClientWithAccessKey(*regionId, *accessKeyId, *accessKeySecret)
	if err != nil {
		panic(err)
	}

	params, _ := json.Marshal(map[string]interface{}{
		"code": verifyCode,
	})

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.TemplateCode = "SMS_209335004"
	request.SignName = "阿里大于测试专用"
	request.TemplateParam = string(params)
	request.PhoneNumbers = strconv.Itoa(*phoneNumbers)

	resp, err := client.SendSms(request)
	if err != nil {
		log.Printf("send sms failed resp=%v err=%v", resp, err)
		panic(err)
	}

	if !resp.IsSuccess() {
		log.Printf("send sms failed resp=%v err=%v", resp, err)
		panic(fmt.Errorf("failed: unknown reason"))
	}
}
