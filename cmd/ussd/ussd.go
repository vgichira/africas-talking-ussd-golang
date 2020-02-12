package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func ussd(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
		return
	}

	queryString := string(data)

	params, err := url.ParseQuery(queryString)

	if err != nil {
		log.Fatalln(err)
		return
	}

	// sessionID := params.Get("sessionId")
	// sessionCode := params.Get("serviceCode")
	phoneNumber := params.Get("phoneNumber")
	accountNo := "TEST16278"
	balance := "20,000"

	text := params.Get("text")

	var response string

	// when a user dials the USSD
	if text == "" {
		response = "CON Hello there, what would you want to do today?\n"
		response += "1. My Account\n"
		response += "2. Check Phone Number\n"
	} else if text == "1" {
		// first level logic
		response = "CON Please choose one option to proceed\n"
		response += "1. Check Account Number\n"
		response += "2. Check Balance\n"
	} else if text == "1*1" {
		// second level logic
		response = "END Your account number is " + accountNo
	} else if text == "1*2" {
		// second level logic
		response = "END Your account balance is KES " + balance
	} else if text == "2" {
		response = "END Your phone number is " + phoneNumber
	}

	w.Header().Set("Content-type", "text/plain")

	fmt.Fprintln(w, response)
}
