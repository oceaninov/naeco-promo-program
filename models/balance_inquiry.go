package models

type BalanceInquiryFailedResponse struct {
	Pan          string `json:"pan"`
	ProCode      string `json:"proCode"`
	Gmt          string `json:"gmt"`
	TraceNo      string `json:"traceNo"`
	ReqTime      string `json:"reqTime"`
	ReqDate      string `json:"reqDate"`
	MerchantType string `json:"merchantType"`
	AcqId        string `json:"acqId"`
	FwdId        string `json:"fwdId"`
	ReffNo       string `json:"reffNo"`
	TerminalId   string `json:"terminalId"`
	TerminalType string `json:"terminalType"`
	TerminalDesc string `json:"terminalDesc"`
	Currency     string `json:"currency"`
	ResponseCode string `json:"responseCode"`
	ErrorCode    string `json:"errorCode"`
}

type BalanceInquiryFailed struct {
	BalanceInquiryResp BalanceInquiryFailedResponse `json:"BalanceInquiryResp"`
}

type BalanceInquirySuccessResponse struct {
	Pan              string `json:"pan"`
	ProCode          string `json:"proCode"`
	Gmt              string `json:"gmt"`
	TraceNo          string `json:"traceNo"`
	ReqTime          string `json:"reqTime"`
	ReqDate          string `json:"reqDate"`
	MerchantType     string `json:"merchantType"`
	AcqId            string `json:"acqId"`
	FwdId            string `json:"fwdId"`
	ReffNo           string `json:"reffNo"`
	TerminalId       string `json:"terminalId"`
	TerminalType     string `json:"terminalType"`
	TerminalDesc     string `json:"terminalDesc"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	AvailableBalance string `json:"availableBalance"`
	ResponseCode     string `json:"responseCode"`
	ErrorCode        string `json:"errorCode"`
}

type BalanceInquirySuccess struct {
	BalanceInquiryResp BalanceInquirySuccessResponse `json:"BalanceInquiryResp"`
}

type BalanceInquiryReq struct {
	ProCode       string `json:"proCode"`
	TraceNo       string `json:"traceNo"`
	ReqTime       string `json:"reqTime"`
	ReqDate       string `json:"reqDate"`
	YearTrx       string `json:"yearTrx"`
	MerchantType  string `json:"merchantType"`
	AcqId         string `json:"acqId"`
	FwdId         string `json:"fwdId"`
	ReffNo        string `json:"reffNo"`
	TerminalId    string `json:"terminalId"`
	TerminalType  string `json:"terminalType"`
	TerminalDesc  string `json:"terminalDesc"`
	Currency      string `json:"currency"`
	AppCode       string `json:"appCode"`
	FromAccNumber string `json:"fromAccNumber"`
}

type BalanceInquiryRequest struct {
	BalanceInquiryReq BalanceInquiryReq `json:"BalanceInquiryReq"`
}
