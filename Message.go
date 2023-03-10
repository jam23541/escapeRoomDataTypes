package escapeRoomDataTypes

import "sync"

type Message struct {
	message MqttMsg
	once    sync.Once
}

func (m *Message) Duplicate() bool {
	return false
}

func (m *Message) Qos() byte {
	return 0
}

func (m *Message) Retained() bool {
	return false
}
func (m *Message) Topic() string {
	return m.message.PUBLISHTOPIC
}
func (m *Message) MessageID() uint16 {
	return 0
}
func (m *Message) Payload() []byte {
	result, err := m.message.ToPayload()
	if err != nil {
		return []byte{}
	}
	return result

}
func (m *Message) Ack() {
	m.once.Do(m.Ack)
}
func (m *Message) GetMqttMsg() MqttMsg {
	return m.message
}
