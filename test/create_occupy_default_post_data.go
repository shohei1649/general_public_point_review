package test

import (
	"encoding/json"
	"fmt"

	"github.com/work/general_public_point_review/util"
)

var securityCode = "00000000"

func CreateDefaultPostJSON() PostJSON {

	var defaultContactInfo = ContactInfo{
		Name:           "毛沢東",
		Phone:          "13900000000",
		Email:          "xinmeida@meituan.com",
		PickupAddress:  "新美大",
		HotelAddressEn: "香格里拉北门",
		MailAddress:    "xinmeida@gmail.com",
		MailZipCode:    "210000",
		StartDate:      "2017-10-18 12:00:00",
		EndDate:        "2017-10-18 15:00:00",
	}

	var defaultVisaInfo = VisaInfo{
		VisaCountry: "签证国家",
		VisaType:    "签证国家",
	}

	var defaultWifiInfo = WifiInfo{
		GetAddress:    "127.0.0.1",
		ReturnAddress: "192.10.1.5",
		WifiCityName:  "Wi-Fi城市",
	}

	var defaultOrderItemList = []OrderItem{

		{
			OrderID:  1000200030004000,
			SkuID:    10065,
			OtaSkuID: "RS0001",
			SaleType: "1",
			Quantity: "1",
			SkuPrice: "10.0",
		},
		{
			OrderID:  112233445566778899,
			SkuID:    10089,
			OtaSkuID: "RT0002",
			SaleType: "2",
			Quantity: "3",
			SkuPrice: "50.0",
		},
	}
	var defaultPassengerInfoList = []PassengerInfo{

		{
			SaleType:         1,
			Name:             "新美大",
			FirstName:        "美大",
			SecondName:       "新",
			FirstEnName:      "meida",
			SecondEnName:     "xin",
			IDCardNumber:     "800000000",
			PassportNumber:   "P1234567",
			HkPassportNumber: "HK1234567",
			TwPassportNumber: "T1234567",
		},
		{
			SaleType:         2,
			Name:             "エビソル太郎",
			FirstName:        "太郎",
			SecondName:       "エビソル",
			FirstEnName:      "tarou",
			SecondEnName:     "ebisol",
			IDCardNumber:     "11111111",
			PassportNumber:   "P7654321",
			HkPassportNumber: "HK7654321",
			TwPassportNumber: "T7654321",
		},
	}

	var defaultData = Data{
		OrderID:            1000200030004000,
		OrderPrice:         20.0,
		ProductID:          12569,
		OtaPid:             "HB009",
		OtaPackageID:       "HBC001",
		Contact:            defaultContactInfo,
		OrderItems:         defaultOrderItemList,
		PassengerInfos:     defaultPassengerInfoList,
		WifiInfo:           defaultWifiInfo,
		FlightInfo:         "アシアナ航空",
		HotelInfo:          "上海世茂皇家艾美酒店",
		VisaInfo:           defaultVisaInfo,
		ReserveConsumeTime: "2017/10/1 12:34:56",
		Remark:             "你可以提前为孩子准备一个座位吗？",
		CarType:            "丰田",
	}

	defaultDataStr, err := json.Marshal(defaultData)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
	}

	var dataBase64 = util.EncodeBase64(string(defaultDataStr))
	var otaID = 10086
	var defaultPostJSON = PostJSON{
		OtaID:   otaID,
		Data:    defaultData,
		Sign:    util.ToMD5Hash(securityCode + string(otaID) + dataBase64),
		AgentID: otaID, //OtaIDと同じ値を使用
	}

	return defaultPostJSON
}

func CreateDefaultResponseJSON() string {

	defaultResponse := Response{
		Code:           string(CodeOK),
		IsSuccess:      "true",
		Msg:            "OK",
		OtaOrderID:     "11111",
		OrderID:        22222,
		OtaOrderStatus: "102",
	}

	defaultResponseStr, _ := json.Marshal(defaultResponse)
	defaultResponseIndent := util.ToJSONIndent(defaultResponseStr)

	return defaultResponseIndent
}

func CreateBase64PostJSON(postJSON PostJSON) string {

	dataStr, _ := json.Marshal(postJSON.Data)
	dataBase64 := util.EncodeBase64(string(dataStr))
	postJSONBase64 := PostJSONBase64{
		OtaID:   postJSON.OtaID,
		Data:    dataBase64,
		Sign:    util.ToMD5Hash(securityCode + string(postJSON.OtaID) + dataBase64),
		AgentID: postJSON.AgentID,
	}

	postJSONBase64Str, _ := json.Marshal(postJSONBase64)
	postJSONBase64Indent := util.ToJSONIndent(postJSONBase64Str)

	return postJSONBase64Indent

}
