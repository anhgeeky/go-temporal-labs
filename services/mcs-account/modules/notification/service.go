package notification

type Service struct {
	Repo Repository
}

func (r Service) SendOTP() (*SendOtpRes, error) {
	return nil, nil
}

func (r Service) SendSms() (*SendSmsRes, error) {
	return nil, nil
}

func (r Service) SendAppNotification() (*SendOtpRes, error) {
	return nil, nil
}
