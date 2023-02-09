package serializer

//http://api.k780.com/?app=ip.get&ip=202.96.128.166&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json

const (
	APP_KEY = "67594"
	SIGN    = "58dc7aa641433d7475b7ff3140032bde"

	Success = 1
)

//错误的返回体
type IpFailData struct {
	Success string `json:"success"`
	Msgid   string `json:"msgid"`
	Msg     string `json:"msg"`
}

//成功的返回体
type IpSuccessData struct {
	Success string `json:"success"`
	Result  struct {
		Status           string `json:"status"`
		Ip               string `json:"ip"`
		IpStr            string `json:"ip_str"`
		IpEnd            string `json:"ip_end"`
		InetIp           string `json:"inet_ip"`
		InetStr          string `json:"inet_str"`
		InetEnd          string `json:"inet_end"`
		Areano           string `json:"areano"`
		Postno           string `json:"postno"`
		Operators        string `json:"operators"`
		Att              string `json:"att"`
		Detailed         string `json:"detailed"`
		AreaStyleSimcall string `json:"area_style_simcall"`
		AreaStyleAreanm  string `json:"area_style_areanm"`
	} `json:"result"`
}
