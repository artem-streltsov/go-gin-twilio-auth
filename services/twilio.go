package services

import (
	twilio "github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
	"os"
)

type TwilioService struct {
	client     *twilio.RestClient
	serviceSID string
}

func NewTwilioService() *TwilioService {
	return &TwilioService{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: os.Getenv("TWILIO_ACCOUNT_SID"),
			Password: os.Getenv("TWILIO_AUTH_TOKEN"),
		}),
		serviceSID: os.Getenv("TWILIO_SERVICE_SID"),
	}
}

func (ts *TwilioService) StartVerification(phone string) error {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	_, err := ts.client.VerifyV2.CreateVerification(ts.serviceSID, params)
	return err
}

func (ts *TwilioService) CheckVerification(phone, code string) (bool, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(code)

	resp, err := ts.client.VerifyV2.CreateVerificationCheck(ts.serviceSID, params)
	if err != nil {
		return false, err
	}
	return *resp.Status == "approved", nil
}
