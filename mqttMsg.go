package escapeRoomDataTypes

import (
	"encoding/json"
	"fmt"
)

type MqttMsg struct {
	PUBLISHTOPIC   string
	MSGTYPE        string
	PUBER          string
	SUBER          []string
	DATA           []string
	SUPPLEMENTDATA []string
	CHIPID         string
	UNIQUEID       string
}

var DefaultMqttMsg = MqttMsg{
	PUBLISHTOPIC:   "EMPTY",
	MSGTYPE:        "EMPTY",
	PUBER:          "EMPTY",
	SUBER:          []string{"EMPTY"},
	DATA:           []string{"EMPTY"},
	SUPPLEMENTDATA: []string{"EMPTY"},
	CHIPID:         "EMPTY",
	UNIQUEID:       "EMPTY",
}

func PayloadToMqttMsg(inputPayload []byte) (MqttMsg, int) {
	var MqttMsgReturn MqttMsg
	err := json.Unmarshal(inputPayload, &MqttMsgReturn)
	if err != nil {
		fmt.Println("Payload to json error")
		return DefaultMqttMsg, -1
	}
	return MqttMsgReturn, 1

}
