package util

type Resp struct {
	Errno  string      `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

const (
	RECODE_OK              = "0"
	RECODE_UNKNOWERR       = "1"
	RECODE_BCOSERR         = "4001"
	RECODE_NETWORK         = "4002"
	RECODE_PASSERR         = "4003"
	RECODE_ACCERR          = "4004"
	RECODE_CONTRACTERR     = "4005"
	RECODE_REGERR          = "4006"
	RECODE_CREATEWALLETERR = "4007"
	RECODE_LOGINERR        = "4008"
)

var recodeText = map[string]string{
	RECODE_OK:              "成功",
	RECODE_UNKNOWERR:       "未知错误",
	RECODE_BCOSERR:         "bcos平台错误",
	RECODE_NETWORK:         "网络错误",
	RECODE_PASSERR:         "密码错误",
	RECODE_ACCERR:          "账户错误",
	RECODE_CONTRACTERR:     "合约调用错误",
	RECODE_REGERR:          "注册错误",
	RECODE_CREATEWALLETERR: "创建钱包错误",
	RECODE_LOGINERR:        "登陆错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
