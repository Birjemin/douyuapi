package douyuapi

const tokenUri = "/api/thirdPart/token"

// Token
type Token struct {
	BaseClient
}

// TokenResponse
type TokenResponse struct {
	ErrorResponse
	Data struct {
		Token  string `json:"token"`
		Expire int    `json:"expire"`
	} `json:"data"`
}

// Handle
func (t *Token) Handle(timestamp string) (*TokenResponse, error) {
	return t.do(DouYuDomain+tokenUri, timestamp)
}

// do
func (t *Token) do(url, timestamp string) (*TokenResponse, error) {
	var params = map[string]string{
		"aid":  t.AID,
		"time": timestamp,
	}
	params["auth"] = GetSign(t.Secret, tokenUri, params)

	if err := t.Client.HttpGet(url, params); err != nil {
		return nil, err
	} else {
		var ret, errResp = new(TokenResponse), new(ErrorResponse)
		if err = t.Client.GetResponseJson(ret, errResp); err != nil {
			return nil, err
		}
		if errResp.Code != 0 {
			ret.Code = errResp.Code
			ret.Msg = errResp.Msg
		}
		return ret, nil
	}
}
