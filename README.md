# LinuxProgramDesign
课设，写个linux上的通信服务,emmmmm，拿个~~c++~~ go搓搓


文件结构:

protocol 实现服务端和客户端统一的通信

服务端采用 生产者 - 消费者模型
生产者监听tcp连接，将获取的信息反序列化成 struct 然后统一处理


传输使用json
