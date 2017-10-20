package test

const (
	SaleTypeAdult     = 1
	SaleTypeChild     = 2
	SaleTypeBaby      = 3
	SaleTypeOldPeople = 4

	OrderStatusSuccess                  = 102
	OrderStatusFailure                  = 103
	OrderStatusPositionSuccess          = 202
	OrderStatusClearedFail              = 203
	OrderStatusConfirming               = 301
	OrderStatusConfirmSuccess           = 302
	OrderStatusValidationFails          = 303
	OrderStatusDuringCancellation       = 401
	OrderStatusPartialRevocationSuccess = 404
	OrderStatusParticalLiftingOfFailure = 405

	CodeOK                  = 200
	CodeUnknowException     = 305
	CodeBadRequest          = 400
	CodeUnAuthorization     = 401
	CodeForbidden           = 403
	CodeNotFound            = 404
	CodeInternalServerError = 500
	CodeSignValidationFails = 501
)

type APITestCase struct {
	StatusCode               int
	Description              string
	RequestParamJSONFileName string
	ResponseJSONFileName     string
	PostJSONFileName         string
	DBDataJSONFileName       string
}

type PassengerInfo struct {
	SaleType         int    `json:"saleType"`
	Name             string `json:"name"`
	FirstName        string `json:"firstName"`
	SecondName       string `json:"secondName"`
	FirstEnName      string `json:"firstEnName"`
	SecondEnName     string `json:"secondEnName"`
	IDCardNumber     string `json:"idCardNumber"`
	PassportNumber   string `json:"passportNumber"`
	HkPassportNumber string `json:"hkPassportNumber"`
	TwPassportNumber string `json:"twPassportNumber"`
}

type OrderItem struct {
	OrderID  int64  `json:"orderId"`
	SkuID    int    `json:"skuId"`
	OtaSkuID string `json:"otaSkuId"`
	SaleType string `json:"saleType"`
	Quantity string `json:"quantity"`
	SkuPrice string `json:"skuPrice"`
}

type WifiInfo struct {
	GetAddress    string `json:"getAddress"`
	ReturnAddress string `json:"returnAddress"`
	WifiCityName  string `json:"wifiCityName"`
}

type VisaInfo struct {
	VisaCountry string `json:"visaCountry"`
	VisaType    string `json:"visaType"`
}

type PostJSON struct {
	OtaID   int    `json:"otaId"`
	Data    Data   `json:"data"`
	Sign    string `json:"sign"`
	AgentID int    `json:"agentId"`
}

type PostJSONBase64 struct {
	OtaID   int    `json:"otaId"`
	Data    string `json:"data"`
	Sign    string `json:"sign"`
	AgentID int    `json:"agentId"`
}

type Response struct {
	Code           string `json:"code"`
	IsSuccess      string `json:"isSuccess"`
	Msg            string `json:"msg"`
	OtaOrderID     string `json:"otaOrderId"`
	OrderID        int64  `json:"orderId"`
	OtaOrderStatus string `json:"otaOrderStatus"`
}

type ContactInfo struct {
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	PickupAddress  string `json:"pickupAddress"`
	HotelAddressEn string `json:"hotelAddressEn"`
	MailAddress    string `json:"mailAddress"`
	MailZipCode    string `json:"mailZipCode"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
}

type Data struct {
	OrderID            int64           `json:"orderId"`
	OrderPrice         float64         `json:"orderPrice"`
	ProductID          int             `json:"productId"`
	OtaPid             string          `json:"otaPid"`
	OtaPackageID       string          `json:"otaPackageId"`
	Contact            ContactInfo     `json:"contact"`
	OrderItems         []OrderItem     `json:"orderItems"`
	PassengerInfos     []PassengerInfo `json:"passengerInfos"`
	WifiInfo           WifiInfo        `json:"wifiInfo"`
	FlightInfo         string          `json:"flightInfo"`
	HotelInfo          string          `json:"hotelInfo"`
	VisaInfo           VisaInfo        `json:"visaInfo"`
	ReserveConsumeTime string          `json:"reserveConsumeTime"`
	Remark             string          `json:"remark"`
	CarType            string          `json:"carType"`
}
