package handlers

import (
    "bytes"
    "encoding/json"
    "encoding/xml"
    "gin-soap-service/models"
    "gin-soap-service/utils"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "time"
)

func ConvertHandler(c *gin.Context) {
    // Generate X-Request-ID
    requestID := utils.GenerateRequestID()

    // Read JSON request body
    var requestPayload models.RequestPayload
    if err := c.ShouldBindJSON(&requestPayload); err != nil {
        utils.LogError(requestID, "Failed to parse JSON request", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
        return
    }

    // Build XML payload
    xmlPayload, err := buildXMLPayload(requestPayload)
    if err != nil {
        utils.LogError(requestID, "Failed to build XML payload", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // Send XML payload to SOAP endpoint
    // soapResponse, err := sendSOAPRequest(xmlPayload)
    // if err != nil {
    //     utils.LogError(requestID, "Failed to send SOAP request", err)
    //     c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
    //     return
    // }
	soapResponse := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
      <ord:QueryOrderRspMsg xmlns:ord="http://www.huawei.com/bss/soaif/interface/OrderService/" xmlns:com="http://www.huawei.com/bss/soaif/interface/common/">
         <com:RspHeader>
            <com:Version>1</com:Version>
            <com:ReturnCode>0000</com:ReturnCode>
            <com:ReturnMsg>Success</com:ReturnMsg>
            <com:RspTime>20240718200041</com:RspTime>
         </com:RspHeader>
         <ord:TotalRowNum>5</ord:TotalRowNum>
         <ord:OrderInfo>
            <com:OrderId>16036917449</com:OrderId>
            <com:OrderType>CO004</com:OrderType>
            <com:OrderStatus>5</com:OrderStatus>
            <com:AcceptTime>20240711062811</com:AcceptTime>
            <com:ExecutionTime>20240711062845</com:ExecutionTime>
            <com:ChangeTime>20240711062845</com:ChangeTime>
         </ord:OrderInfo>
         <ord:OrderInfo>
            <com:OrderId>16036917450</com:OrderId>
            <com:OrderType>CO011</com:OrderType>
            <com:OrderStatus>5</com:OrderStatus>
            <com:AcceptTime>20240711062811</com:AcceptTime>
            <com:ExecutionTime>20240711062845</com:ExecutionTime>
            <com:ChangeTime>20240711062845</com:ChangeTime>
         </ord:OrderInfo>
         <ord:OrderInfo>
            <com:OrderId>16036917451</com:OrderId>
            <com:OrderType>CO026</com:OrderType>
            <com:OrderStatus>5</com:OrderStatus>
            <com:AcceptTime>20240711062811</com:AcceptTime>
            <com:ExecutionTime>20240711062845</com:ExecutionTime>
            <com:ChangeTime>20240711062845</com:ChangeTime>
         </ord:OrderInfo>
         <ord:OrderInfo>
            <com:OrderId>16036917452</com:OrderId>
            <com:OrderType>CO030</com:OrderType>
            <com:OrderStatus>5</com:OrderStatus>
            <com:AcceptTime>20240711062811</com:AcceptTime>
            <com:ExecutionTime>20240711062845</com:ExecutionTime>
            <com:ChangeTime>20240711062845</com:ChangeTime>
         </ord:OrderInfo>
         <ord:OrderInfo>
            <com:OrderId>16036917453</com:OrderId>
            <com:OrderType>CV105</com:OrderType>
            <com:OrderStatus>5</com:OrderStatus>
            <com:AcceptTime>20240711062811</com:AcceptTime>
            <com:ExecutionTime>20240711062845</com:ExecutionTime>
            <com:ChangeTime>20240711062845</com:ChangeTime>
         </ord:OrderInfo>
      </ord:QueryOrderRspMsg>
   </soapenv:Body>
</soapenv:Envelope>`

    // Convert SOAP response to JSON
    jsonResponse, err := convertSOAPResponse(soapResponse)
    if err != nil {
        utils.LogError(requestID, "Failed to convert SOAP response", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // Log response
    utils.LogResponse(requestID, jsonResponse)

    // Return JSON response
    c.JSON(http.StatusOK, jsonResponse)
}

func buildXMLPayload(payload models.RequestPayload) (string, error) {
    request := models.SOAPEnvelope{
        XMLNS: struct {
            SoapEnv string `xml:"xmlns:soapenv,attr"`
            Ord     string `xml:"xmlns:ord,attr"`
            Com     string `xml:"xmlns:com,attr"`
        }{
            SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
            Ord:     "http://www.huawei.com/bss/soaif/interface/OrderService/",
            Com:     "http://www.huawei.com/bss/soaif/interface/common/",
        },
        Body: models.Body{
            QueryOrderReqMsg: models.QueryOrderReqMsg{
                ReqHeader: models.ReqHeader{
                    Version:        "1",
                    BusinessCode:   "QueryOrder",
                    TransactionId:  payload.TransactionId,
                    Channel:        payload.Channel,
                    PartnerId:      payload.PartnerId,
                    BrandId:        payload.BrandId,
                    ReqTime:        payload.ReqTime,
                    TimeFormat: models.TimeFormat{
                        TimeType:   payload.TimeType,
                        TimeZoneID: payload.TimeZoneID,
                    },
                    AccessUser:     payload.AccessUser,
                    AccessPassword: payload.AccessPassword,
                    OperatorId:     payload.OperatorId,
                    AdditionalProperty: payload.AdditionalProperties,
                },
                OrderId:          payload.OrderId,
                CustId:           payload.CustId,
                TimePeriod:       models.TimePeriod{
                    StartTime: payload.StartTime,
                    EndTime:   payload.EndTime,
                },
                PagingInfo: models.PagingInfo{
                    TotalRowNum: payload.TotalRowNum,
                    BeginRowNum: payload.BeginRowNum,
                    FetchRowNum: payload.FetchRowNum,
                },
                AdditionalProperty: payload.AdditionalProperties,
            },
        },
    }

    xmlData, err := xml.MarshalIndent(request, "", "  ")
    if err != nil {
        return "", err
    }

    return string(xmlData), nil
}

func sendSOAPRequest(xmlPayload string) (string, error) {
    soapURL := utils.GetEnv("SOAP_ENDPOINT_URL", "http://default.url")
    resp, err := http.Post(soapURL, "text/xml; charset=utf-8", bytes.NewBufferString(xmlPayload))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}

func convertSOAPResponse(soapResponse string) (models.JSONResponse, error) {
    var response models.JSONResponse
    err := xml.Unmarshal([]byte(soapResponse), &response)
    if err != nil {
        return response, err
    }
    return response, nil
}
