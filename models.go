package vatansms

import (
	"encoding/xml"
)

type SmsDefaults struct {
	UserID   uint   `xml:"kno"`
	Username string `xml:"kulad"`
	Password string `xml:"sifre"`
	Sender   string `xml:"gonderen"`
	Type     string `xml:"tur"`
	Time     string `xml:"zaman,omitempty"`
}

type NumberAndMessage struct {
	Number  string `xml:"tel"`
	Message string `xml:"mesaj"`
}

type NToN struct {
	XMLName xml.Name `xml:"sms"`
	SmsDefaults
	NumberAndMessages []NumberAndMessage `xml:"telmesajlar>telmesaj"`
}

type OneToN struct {
	XMLName xml.Name `xml:"sms"`
	SmsDefaults
	Message string `xml:"mesaj"`
	Numbers string `xml:"numaralar"`
}

type Report struct {
	UserID   uint   `xml:"kullanicino"`
	Username string `xml:"kullaniciadi"`
	Password string `xml:"sifre"`
	Date     string `xml:"baslangictarih"`
	EndDate  string `xml:"bitistarih"`
	ID       int    `xml:"smsid"`
}

type UserInfo struct {
	UserID   uint   `xml:"kullanicino"`
	Username string `xml:"kullaniciadi"`
	Password string `xml:"sifre"`
}

type UserInfoResult struct {
	Company string
	Author  string
	Credit  uint
}

type SendResult struct {
	Status      bool
	Description string
	ReportID    int
	Count       int
}

type ReportResult struct {
	Number string
	State  string
}

type ReportDetail struct {
	XMLName xml.Name             `xml:"tum_sonuclar"`
	Result  []ReportDetailResult `xml:"sms_sonucu"`
}

type ReportDetailResult struct {
	Number            string `xml:"tel"`
	Operator          string `xml:"operator"`
	PackageNumber     string `xml:"paketno"`
	Sender            string `xml:"orginator"`
	Message           string `xml:"mesaj"`
	Status            string `xml:"sonuc"`
	StatusDescription string `xml:"sonucaciklama"`
	ReceivedAt        string `xml:"iletimtarihi"`
	SentAt            string `xml:"gonderimtarihi"`
	Type              string `xml:"tur"`
	Len               string `xml:"boy"`
	Price             string `xml:"fiyat"`
}
