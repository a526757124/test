package im

type Message struct {
	//接口版本号
	version string
	//设置uuid
	deviceId string
	//请求接接口命令字
	// 1绑定  2心跳   3上线   4下线
	cmd int
	//发送人
	sender string
	//接收人
	receiver string
	//请求1，应答2，通知3，响应4  format
	msgType int
	//1 rsa加密 2aes加密
	flag            int
	platform        string
	platformVersion string
	token           string
	appKey          string
	timeStamp       int64
	sign            string
	content         []byte
}
