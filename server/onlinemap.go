package server

import "sync"

// 定义全局 map 存储在线用户 key:username, value: Client 涉及到并发，带锁，做到线程安全
type onlineMap struct {
	Clients map[string]Client
	Mutex   sync.Mutex
}

func AddMap(onlineMap *onlineMap, username string, clnt Client) {
	onlineMap.Mutex.Lock()
	defer onlineMap.Mutex.Unlock()
	onlineMap.Clients[username] = clnt
}

func DelMap(onlineMap *onlineMap, username string) {
	onlineMap.Mutex.Lock()
	defer onlineMap.Mutex.Unlock()
	delete(onlineMap.Clients, username)
}

func GetMap(onlineMap *onlineMap, username string) Client {
	onlineMap.Mutex.Lock()
	defer onlineMap.Mutex.Unlock()
	return onlineMap.Clients[username]
}

func GetAllMap(onlineMap *onlineMap) map[string]Client {
	onlineMap.Mutex.Lock()
	defer onlineMap.Mutex.Unlock()
	return onlineMap.Clients
}
