package receive

//kf_close_session

/*
{"ToUserName":"gh_b068735876a5","FromUserName":"oMBo6wh75aCAcnIzpmQ3WVgvA4R8","CreateTime":"1525769301","MsgType":"event","Event":"kf_close_session","KfAccount":"Lincoln1@lincoln","CloseType":"TIMEOUT"}
*/

type EventKfCloseSession struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	KfAccount    string `json:"KfAccount"`
	CloseType    string `json:"CloseType"`
}
