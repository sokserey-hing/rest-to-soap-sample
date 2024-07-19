package models

import "encoding/xml"

// Define the structures to match the XML payload
type SOAPEnvelope struct {
    XMLName xml.Name      `xml:"soapenv:Envelope"`
    XMLNS   struct {
        SoapEnv string `xml:"xmlns:soapenv,attr"`
        Ord     string `xml:"xmlns:ord,attr"`
        Com     string `xml:"xmlns:com,attr"`
    } `xml:",attr"`
    Header struct{} `xml:"soapenv:Header"`
    Body   struct {
        QueryOrderReqMsg QueryOrderReqMsg `xml:"ord:QueryOrderReqMsg"`
    } `xml:"soapenv:Body"`
}

type QueryOrderReqMsg struct {
    ReqHeader         ReqHeader          `xml:"com:ReqHeader"`
    OrderId           string             `xml:"ord:OrderId"`
    CustId            string             `xml:"ord:CustId"`
    TimePeriod        TimePeriod         `xml:"ord:TimePeriod"`
    PagingInfo        PagingInfo         `xml:"ord:PagingInfo"`
    AdditionalProperty []AdditionalProperty `xml:"ord:AdditionalProperty"`
}

type ReqHeader struct {
    Version           string              `xml:"com:Version"`
    BusinessCode      string              `xml:"com:BusinessCode"`
    TransactionId     string              `xml:"com:TransactionId"`
    Channel           string              `xml:"com:Channel"`
    PartnerId         string              `xml:"com:PartnerId"`
    BrandId           string              `xml:"com:BrandId"`
    ReqTime           string              `xml:"com:ReqTime"`
    TimeFormat        TimeFormat          `xml:"com:TimeFormat"`
    AccessUser        string              `xml:"com:AccessUser"`
    AccessPassword    string              `xml:"com:AccessPassword"`
    OperatorId        string              `xml:"com:OperatorId"`
    AdditionalProperty []AdditionalProperty `xml:"com:AdditionalProperty"`
}

type TimeFormat struct {
    TimeType   string `xml:"com:TimeType"`
    TimeZoneID string `xml:"com:TimeZoneID"`
}

type TimePeriod struct {
    StartTime string `xml:"com:StartTime"`
    EndTime   string `xml:"com:EndTime"`
}

type PagingInfo struct {
    TotalRowNum   string `xml:"com:TotalRowNum"`
    BeginRowNum   string `xml:"com:BeginRowNum"`
    FetchRowNum   string `xml:"com:FetchRowNum"`
}

type AdditionalProperty struct {
    Code  string `xml:"com:Code"`
    Value string `xml:"com:Value"`
}

type RequestPayload struct {
    TransactionId     string              `json:"transactionId"`
    Channel           string              `json:"channel"`
    PartnerId         string              `json:"partnerId"`
    BrandId           string              `json:"brandId"`
    ReqTime           string              `json:"reqTime"`
    TimeType          string              `json:"timeType"`
    TimeZoneID        string              `json:"timeZoneID"`
    AccessUser        string              `json:"accessUser"`
    AccessPassword    string              `json:"accessPassword"`
    OperatorId        string              `json:"operatorId"`
    OrderId           string              `json:"orderId"`
    CustId            string              `json:"custId"`
    StartTime         string              `json:"startTime"`
    EndTime           string              `json:"endTime"`
    TotalRowNum       string              `json:"totalRowNum"`
    BeginRowNum       string              `json:"beginRowNum"`
    FetchRowNum       string              `json:"fetchRowNum"`
    AdditionalProperties []AdditionalProperty `json:"additionalProperties"`
}

type JSONResponse struct {
    // Define the structure of the JSON response if needed
}
