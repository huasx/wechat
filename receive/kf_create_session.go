package receive

//kf_create_session

/*
{"ToUserName":"gh_5b746904e970","FromUserName":"oaB7kjq7ADujTxQlf4y6eNuNhZqA","CreateTime":"1525769569","MsgType":"event","Event":"kf_create_session","KfAccount":"kf2008@youyuewuxian"}
*/

type EventKfCreateSession struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	KfAccount    string `json:"KfAccount"`
}
