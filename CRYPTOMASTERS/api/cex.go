package api

import (
	datatypes "cryptomasters/data"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error){
	if len(currency)!=3{
		return nil, fmt.Errorf("3 characters requiered!")
	}
	res, err:= http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	//check for the error
	if err != nil{
		return nil,err //if you want to send nil back you need to use the pointer
	}
	var response CEXResponse
	//check if the response is the right one
	if res.StatusCode == http.StatusOK{
		//response is in chunks of packets from tcp or udp
		//we want to wait for the whole thing to download before using it
		bodyBytes, err := io.ReadAll(res.Body)// Read everything synchronously wait for the complete download before moving down, could use routine
		
		//return error if could not download it
		if err!=nil{
			return nil, err
		}

		//Json := string(bodyBytes)//convert the whole https response to a json
		//fmt.Println(json)
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil{
			return nil, err
		}
	
	} else{
		return nil, fmt.Errorf("status code received: %v", res.StatusCode) //We should create this error to tell where there was a fault
	}
	rate := datatypes.Rate{Currency: currency, Price:response.Bid} 

	return &rate, nil 

}
