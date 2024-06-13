package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"

	"github.com/twilio/twilio-go/client/jwt"
	"github.com/twilio/twilio-go/twiml"
)

var twilioNumber = os.Getenv("TWILIO_NUMBER")

func voice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := fmt.Errorf("error: %q", err)
		fmt.Println(error.Error())
	}

	to := r.PostForm.Get("to")
	from := r.PostForm.Get("from")

	if to == twilioNumber {
		dial := &twiml.VoiceDial{
			Number:   to,
			CallerId: from,
		}

		twimlResult, err := twiml.Voice([]twiml.Element{dial})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			error := fmt.Errorf("error: %q", err)
			fmt.Println(error.Error())
		}

		xmlResponse, err := xml.MarshalIndent(twimlResult, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/xml")
			w.Write(xmlResponse)
		}
	}
}

func token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var twilioAccountSid string = os.Getenv("TWILIO_ACCOUNT_SID")
	var twilioApiKey string = os.Getenv("TWILIO_API_KEY")
	var twilioApiSecret string = os.Getenv("TWILIO_API_SECRET")

	var outgoingApplicationSid string = os.Getenv("TWIML_APP_SID")

	params := jwt.AccessTokenParams{
		AccountSid:    twilioAccountSid,
		SigningKeySid: twilioApiKey,
		Secret:        twilioApiSecret,
		Identity:      "system",
		Region:        "us1",
	}

	jwtToken := jwt.CreateAccessToken(params)
	voiceGrant := &jwt.VoiceGrant{
		Outgoing: jwt.Outgoing{
			ApplicationSid: outgoingApplicationSid,
		},
		Incoming: jwt.Incoming{
			Allow: true,
		},
	}

	jwtToken.AddGrant(voiceGrant)
	token, err := jwtToken.ToJwt()

	if err != nil {
		error := fmt.Errorf("error: %q", err)
		fmt.Println(error.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(token)
	}
}
