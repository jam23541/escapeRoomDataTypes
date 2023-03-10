package escapeRoomDataTypes

import (
	"encoding/json"
	"fmt"
	"sync"
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

func (m *MqttMsg) ToPayload() ([]byte, error) {
	result, err := json.Marshal(m)
	return result, err
}

func (m *MqttMsg) ToMessage() Message {
	return Message{
		message: *m,
		once:    sync.Once{},
	}
}
