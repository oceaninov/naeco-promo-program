package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/nbdgocean6/nobita-promo-program/models"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/repositories/interface"
	"io"
	"io/ioutil"
	"crypto/rand"
	"net/http"
	"strconv"
	"time"
)

type api struct{}

func (a *api) errorDefinition(balanceDetailFailed models.BalanceInquiryFailed) error {
	var errorCode = balanceDetailFailed.BalanceInquiryResp.ErrorCode
	var responseCode = balanceDetailFailed.BalanceInquiryResp.ResponseCode
	if errorCode == "DDA0103" && responseCode == "76" {
		return errors.New("account number not recognized or not registered in system")
	} else if errorCode == "CDS1404" && responseCode == "94" {
		return errors.New("request cannot be processed because request is not relevant or conflict")
	} else {
		return errors.New(fmt.Sprintf("failed to get balance of promo engine, %s %s", errorCode, responseCode))
	}
}

func (a *api) generateNumbers(digit int) string {
	if digit > 39 {
		digit = 39
	}
	randomCrypto, _ := rand.Prime(rand.Reader, 128)
	return randomCrypto.String()[:digit]
}

func (a *api) CheckBalance(account string) (res int64, err error) {
	t := time.Now()
	var balanceDetail models.BalanceInquirySuccess
	var balanceDetailFailed models.BalanceInquiryFailed
	request := models.BalanceInquiryRequest{
		BalanceInquiryReq: models.BalanceInquiryReq{
			ProCode:       "310080",
			TraceNo:       a.generateNumbers(6),
			ReffNo:        a.generateNumbers(12),
			ReqTime:       t.Format("150405"),
			ReqDate:       t.Format("20060102"),
			YearTrx:       t.Format("2006"),
			MerchantType:  "9800",
			AcqId:         "100300",
			FwdId:         "503",
			TerminalId:    "PROMOCHL",
			TerminalType:  "WEB",
			TerminalDesc:  "DASHBOARD PROMO",
			Currency:      "360",
			AppCode:       "S",
			FromAccNumber: account,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest(http.MethodPost, "http://35.219.67.179:31000/omni/esb", bytes.NewBuffer(body))
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bodyBytes, &balanceDetail)
	if err != nil {
		return res, err
	}
	if len(balanceDetail.BalanceInquiryResp.AvailableBalance) > 0 {
		unformatedBalance := balanceDetail.BalanceInquiryResp.AvailableBalance[9:]
		balance, err := strconv.Atoi(unformatedBalance[:len(unformatedBalance)-2])
		if err != nil {
			return res, err
		}
		return int64(balance), nil
	} else {
		err = json.Unmarshal(bodyBytes, &balanceDetailFailed)
		if err != nil {
			return res, errors.New(fmt.Sprintf("system error %+v", err))
		}
		return res, a.errorDefinition(balanceDetailFailed)
	}
}

func NewAPI() _interface.Api {
	return &api{}
}
