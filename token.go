package douyuapi

const TokenUri = "/api/thirdPart/token"

// Token
type Token struct {
	BaseClient
}

// TokenResponse
type TokenResponse struct {
	ErrorResponse
	Data struct {
		Token  string
		Expire int
	}
}

// Handle
func (t *Token) Handle(timestamp string) (*TokenResponse, error) {
	return t.do(DouYuDomain+TokenUri, timestamp)
}

// do
func (t *Token) do(url, timestamp string) (*TokenResponse, error) {
	var params = map[string]string{
		"aid":  t.AID,
		"time": timestamp,
	}
	params["auth"] = GetSign(t.Secret, TokenUri, params)

	if err := t.Client.HttpGet(url, params); err != nil {
		return nil, err
	} else {
		var ret = new(TokenResponse)
		if err = t.Client.GetResponseJson(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
