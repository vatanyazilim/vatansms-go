 ### Kurulum
 ```go get github.com/vatanyazilim/vatansms-go```
 
 ### Toplu SMS
  ##### Her numaraya farklı mesajlar göndermek için;
  ``` 
          var numberAndMessages []vatansms.NumberAndMessage
     	numberAndMessages = append(numberAndMessages, vatansms.NumberAndMessage{
     		Number:  vatansms.PhoneVerify("5666991012"),
     		Message: "Hi 5666991012",
     	})
     
     	numberAndMessages = append(numberAndMessages, vatansms.NumberAndMessage{
     		Number:  vatansms.PhoneVerify("5356992106"),
     		Message: "Hi 5356992106",
     	})
     
     	smsNN := vatansms.NToN{}
     	smsNN.UserID = 123
     	smsNN.Username = "test"
     	smsNN.Password = "12345"
     	smsNN.Sender = "VATAN SMS"
     	smsNN.NumberAndMessages = numberAndMessages
     	smsNN.Type = "normal"
     	smsNN.SendNN()
```

 ### Toplu SMS
  ##### Bir veya daha fazla numaraya aynı mesajı göndermek için;
  ```
          numbers := []string{"5356992106","5666991012"}
     	numbersString := vatansms.NumbersArrayToString(numbers)
     
     	sms1N := vatansms.OneToN{}
     	sms1N.UserID = 123
     	sms1N.Username = "test"
     	sms1N.Password = "12345"
     	sms1N.Sender = "VATAN SMS"
     	sms1N.Numbers = numbersString
     	sms1N.Message = "Hi"
     	sms1N.Type = "normal"
     	sms1N.Send1N()
```

 ### Rapor Sorgulama
  ##### Gönderilen bir mesajdan sonra dönen ID bilgisi ile rapor sorgulamak için;
  ```
    	report := vatansms.Report{}
    	report.UserID = 123
    	report.Username = "test"
    	report.Password = "12345"
    	report.Date = "2020-04-10" //SMS id's create date
    	report.ID = 250675671
    	res, err := report.GetReport()
    	if err != nil {
    		panic(err.Error())
    	}
    	fmt.Println(res)
```


 ### Üye Bilgisi ve Kalan Kredi
  ##### Üyeye ait bilgileri ve kalan krediyi sorgulamak için;
  ```
          userInfo := UserInfo{}
    	userInfo.UserID = 123
    	userInfo.Username = "test"
    	userInfo.Password = "12345"
    	userInfoResult, err := userInfo.GetUser()
    	if err != nil {
    		panic(err.Error())
    	}
    	fmt.Println(userInfoResult)
