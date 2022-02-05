package google

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	constant "github.com/NpoolPlatform/third-gateway/pkg/const"

	"github.com/go-resty/resty/v2"

	"golang.org/x/xerrors"
)

type verifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func VerifyGoogleRecaptchaV3(ctx context.Context, in *npool.VerifyGoogleRecaptchaV3Request) (*npool.VerifyGoogleRecaptchaV3Response, error) {
	hostname := config.GetStringValueWithNameSpace("", config.KeyHostname)
	recaptchaURL := config.GetStringValueWithNameSpace(hostname, constant.GoogleRecaptchaV3URL)
	recaptchaSecret := config.GetStringValueWithNameSpace(hostname, constant.GoogleRecaptchaV3Secret)

	if recaptchaURL == "" || recaptchaSecret == "" || in.GetRecaptchaToken() == "" {
		return nil, xerrors.Errorf("invalid recaptcha parameter")
	}

	url := fmt.Sprintf("%v?secret=%v&response=%v", recaptchaURL, recaptchaSecret, in.GetRecaptchaToken())
	resp, err := resty.New().R().SetResult(&verifyResponse{}).Post(url)
	if err != nil {
		return nil, xerrors.Errorf("fail verify google recaptcha v3: %v", err)
	}

	vResp, ok := resp.Result().(*verifyResponse)
	if !ok {
		return nil, xerrors.Errorf("fail get response")
	}

	if !vResp.Success {
		return nil, xerrors.Errorf("fail verify google recaptcha v3: %v", vResp.ErrorCodes)
	}

	return &npool.VerifyGoogleRecaptchaV3Response{}, nil
}
