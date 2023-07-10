package protocol

import "encoding/json"

//	type User struct {
//		Username string `json:"username"`
//		UserIp   string `json:"user_ip"`
//	}

type Login struct {
	Username string `json:"username"`
}

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Status struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// TODO
type SercetMessage struct {
	UsernameFrom string `json:"usernameFrom"`
	UsernameTo   string `json:"usernameTo"`
	Content      string `json:"content"`
}

// 序列化函数，塞入一个任意类型的数据，返回一个byte类型的数据
func SerializeData(data interface{}) ([]byte, error) {
	serializedData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return serializedData, nil
}

// 反序列化数据，第一个参数为byte类型的数据，第二个参数为目标数据的地址
// 例如：err = protocol.UnserializeData(data, &u2)
func UnserializeData(data []byte, target interface{}) error {
	err := json.Unmarshal(data, &target)
	if err != nil {
		return err
	}
	return nil
}

func MakeMsg(username, content string) ([]byte, error) {
	//var msg = Message{
	//	Username: username,
	//	Content:  content,
	//}
	var msg = SercetMessage{
		UsernameFrom: "Server",
		UsernameTo:   "All",
		Content:      content,
	}
	data, err := SerializeData(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//func UnserializeData(data []byte) (interface{}, error) {
//	err := json.Unmarshal(data, target)
//	if err != nil {
//		return target, err
//	}
//	return target, nil
//}

//func seralizeUser(u user) ([]byte, error) {
//	data, err := json.Marshal(u)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func unSeralizeUser(data []byte) (user, error) {
//	var u user
//	err := json.Unmarshal(data, &u)
//	if err != nil {
//		return u, err
//	}
//	return u, nil
//}
//
//func seralizeMessage(m message) ([]byte, error) {
//	data, err := json.Marshal(m)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func unSeralizeMessage(data []byte) (message, error) {
//	var m message
//	err := json.Unmarshal(data, &m)
//	if err != nil {
//		return m, err
//	}
//	return m, nil
//}
