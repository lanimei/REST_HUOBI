package config

// API KEY
const (
	// todo: replace with your own AccessKey and Secret Key
	ACCESS_KEY string = "677650c7-c7b63f1a-a1fcba4d-11117"
	SECRET_KEY string = "26764c72-eb1a9e06-8fbc2908-a344a"

	// default to be disabled, please DON'T enable it unless it's officially announced.
	ENABLE_PRIVATE_SIGNATURE bool = false

	// generated the key by: openssl ecparam -name prime256v1 -genkey -noout -out privatekey.pem
	// only required when Private Signature is enabled
	// todo: replace with your own PrivateKey from privatekey.pem
	PRIVATE_KEY_PRIME_256 string = `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`

)

// API请求地址, 不要带最后的/
const (
	//todo: replace with real URLs and HostName
	MARKET_URL string = "https://api.huobi.pro/market"
	TRADE_URL  string = ""
	HOST_NAME  string = ""
)