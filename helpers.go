package vatansms

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"
)

func PrepareXml(data interface{}) url.Values {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	xmlString := string(xmlData)
	values := url.Values{}
	values.Add("data", xmlString)
	return values
}

func PhoneVerify(phone string) string {
	re := regexp.MustCompile("[0-9]") //
	noChar := re.FindAllString(phone, -1)
	if len(noChar) < 10 {
		return ""
	}
	getLast10 := noChar[len(noChar)-10:]
	number := ""
	for char := range getLast10 {
		number += getLast10[char]
	}
	return number
}

func NumbersArrayToString(phones []string) string {
	var replacedNumbers []string
	for i := range phones {
		if phones[i] != "" {
			replacedNumbers = append(replacedNumbers, PhoneVerify(phones[i]))
		}
	}
	return strings.Join(replacedNumbers, ",")
}

func CharReplace(message string) string {
	find := []string{"@", "!", "(", "/", ">", "^", "'", ")", ":", "?", "{", "}", "~", "=", "<", ";", "*", "+", "-", "&", "_", "€", "$", "£", "#", "%", "ö", "ü", "ç", "ş", "ğ", "ı", "Ö", "Ü", "Ç", "Ş", "İ", "Ğ", "\n", "’", "é"}
	replace := []string{"|01|", "|26|", "|33|", "|39|", "|44|", "|51|", "|27|", "|34|", "|40|", "|45|", "|46|", "|47|", "|49|", "|43|", "|42|", "|41|", "|35|", "|36|", "|38|", "|31|", "|14|", "|05|", "|03|", "|02|", "|28|", "|30|", "|62|", "|63|", "|64|", "|65|", "|66|", "|67|", "|68|", "|69|", "|70|", "|71|", "|72|", "|73|", "|61|", "|27|", ""}
	for i := range find {
		message = strings.Replace(message, find[i], replace[i], -1)
	}

	return message
}
