package vatansms

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/tiaguinho/gosoap"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (data OneToN) Send1N() (SendResult, error) {
	values := PrepareXml(data)
	fmt.Println(values)
	req, err := http.PostForm(Url1N, values)
	if err != nil {
		panic(err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(req.Body)
	response, _ := SmsResponse(string(bodyBytes))
	return response, nil

}

func (data NToN) SendNN() (SendResult, error) {
	values := PrepareXml(data)
	req, err := http.PostForm(UrlNN, values)
	if err != nil {
		panic(err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(req.Body)
	response, _ := SmsResponse(string(bodyBytes))
	return response, nil
}

func SmsResponse(response string) (SendResult, error) {
	parse := strings.Split(response, ":")
	if len(parse) < 2 {
		return SendResult{}, errors.New("SMS sonucu parçalanamadı. Servis yanıtı: " + response)
	}
	code, _ := strconv.Atoi(parse[0])

	if code == 1 {
		reportID, _ := strconv.Atoi(parse[1])
		count, _ := strconv.Atoi(parse[3])
		return SendResult{
			Status:      true,
			ReportID:    reportID,
			Description: parse[2],
			Count:       count,
		}, nil
	} else {
		return SendResult{
			Status:      false,
			Description: parse[1],
		}, nil
	}
}

func (data Report) GetReport() ([]ReportDetailResult, error) {
	var reportDetails ReportDetail
	var httpClient = &http.Client{
		Timeout: 1500 * time.Millisecond,
	}

	soap, err := gosoap.SoapClient(WebServiceUrl, httpClient)
	if err != nil {
		return reportDetails.Result, errors.New("Soap başlatılamadı. " + err.Error())
	}

	date, _ := time.Parse("2006-01-02", data.Date)

	fmt.Println(date)

	params := gosoap.Params{
		"kullanicino":    strconv.Itoa(int(data.UserID)),
		"kullaniciadi":   data.Username,
		"sifre":          data.Password,
		"baslangictarih": data.Date,
		"bitistarih":     date.AddDate(0, 0, 2).String()[0:10],
		"smsid":          strconv.Itoa(data.ID),
	}

	res, err := soap.Call("ikitariharasisorgulaXMLverID", params)
	if err != nil {
		return reportDetails.Result, errors.New("Soap çağrısı yapılamadı. " + err.Error())
	}

	reportDetailReturn := struct {
		XMLName xml.Name
		Return  []string `xml:"return"`
	}{}

	err = res.Unmarshal(&reportDetailReturn)

	//err = xml.Unmarshal([]byte(reportDetailReturn), &reportDetailReturn)
	if err != nil {
		return reportDetails.Result, errors.New("Kayıt bulunamadı. ")
	}

	if strings.Contains(reportDetailReturn.Return[0], "HATA:Kullanici bulunamadi") {
		return reportDetails.Result, errors.New("Kullanıcı bulunamadı, kullanıcı bilgileri, rapor ID ve tarih bilgilerini doğru girmeye özen gösteriniz. ")
	}

	err = xml.Unmarshal([]byte(reportDetailReturn.Return[0]), &reportDetails)
	if err != nil {
		return reportDetails.Result, errors.New("Kayıt bulunamadı. ")
	}

	return reportDetails.Result, nil
}

func (data UserInfo) GetUser() (UserInfoResult, error) {
	var user UserInfoResult
	var httpClient = &http.Client{
		Timeout: 1500 * time.Millisecond,
	}

	soap, err := gosoap.SoapClient(WebServiceUrl, httpClient)
	if err != nil {
		return user, errors.New("Soap başlatılamadı. " + err.Error())
	}

	params := gosoap.Params{
		"kullanicino":  strconv.Itoa(int(data.UserID)),
		"kullaniciadi": data.Username,
		"sifre":        data.Password,
	}

	res, err := soap.Call("UyeBilgisiSorgula", params)
	if err != nil {
		return user, errors.New("Soap çağrısı yapılamadı. " + err.Error())
	}

	userInfo := struct {
		XMLName xml.Name
		Return  []string `xml:"return"`
	}{}

	err = res.Unmarshal(&userInfo)
	if err != nil {
		return user, errors.New("Kayıt bulunamadı. ")
	}

	if strings.Contains(userInfo.Return[0], "Kullanici bulunamadi") {
		return user, errors.New("Kullanıcı bulunamadı, kullanıcı bilgilerini kontrol ediniz. ")
	}

	splitBr := strings.Split(userInfo.Return[0], "<br>")
	for _, row := range splitBr {
		splitField := strings.Split(row, "=")
		splitField[0] = strings.TrimSpace(splitField[0])
		splitField[1] = strings.TrimSpace(splitField[1])
		if strings.Contains(splitField[0], "Toplam SMS") {
			credit, _ := strconv.Atoi(splitField[1])
			user.Credit = uint(credit)
		} else if strings.Contains(splitField[0], "Firma") {
			user.Company = splitField[1]
		} else if strings.Contains(splitField[0], "Yetkili") {
			user.Author = splitField[1]
		} else {
			continue
		}
	}

	return user, nil
}
