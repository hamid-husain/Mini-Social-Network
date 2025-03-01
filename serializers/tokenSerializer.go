package serializers

type TokenSerializer struct {
	Key        string `json:"key"`
	ExpiryTime int64  `json:"expiry_time"`
}

func SerializeToken(token string, expiryTime int64) TokenSerializer {
	return TokenSerializer{
		Key:        token,
		ExpiryTime: expiryTime,
	}
}
