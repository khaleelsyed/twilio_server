package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/twilio/twilio-go/client/jwt"
)

func voice(w http.ResponseWriter, r *http.Request) {}

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
	}

	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(token)
	return
}
