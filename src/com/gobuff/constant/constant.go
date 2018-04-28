package constant

const (
	HEART_BEAT_PERIOD = 2 //心跳间隔，以秒计算
	HEART_BEAT_RATIO  = 2 //服务端心跳间隔系数,服务端心跳超时时间要比客户端长

	//指令集合
	CodeHeartBeat int32 = 11 //心跳指令
	CodeUserId    int32 = 12 //发送uid指令
)
