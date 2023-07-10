# LinuxProgramDesign
## Build :

### Client:
```
cd ./client/cmd
go build main.go
```

### Server:
```
cd ./server/cmd
go build main.go
```

## 设计思路:

### Server:

主线程函数`Run`，循环监听客户端请求，根据设定的容量大小，如果剩余空位，将客户端链接处理函数提交到工作池中，若无剩余空位，就断开客户端链接，并返回断开链接提示。
```
digraph G {
    node [shape=box]

    subgraph cluster_main {
        label="主线程函数 Run"
        style=dashed

        listen [label="循环监听客户端请求"]
        check_capacity [label="检查剩余空位"]
        submit_to_pool [label="提交客户端链接处理函数到工作池"]
        disconnect [label="断开客户端链接"]
        return_message [label="返回断开链接提示"]

        listen -> check_capacity
        check_capacity -> submit_to_pool [label="有剩余空位"]
        check_capacity -> disconnect [label="无剩余空位"]
        disconnect -> return_message
        
       
    }

    subgraph cluster_pool {
        label="工作池"
        style=dashed

        worker1 [label="HandleConnect 1"]
        worker2 [label="HandleConnect 2"]
        worker3 [label="HandleConnect 3"]
        worker4 [label="HandleConnect ..."]

        submit_to_pool -> worker1
        submit_to_pool -> worker2
        submit_to_pool -> worker3
        submit_to_pool -> worker4
    }
    
    subgraph cluster_pool2 {
        label=""
        style=dashed

        
        messageChan[label = "全局OnlineMap"]
        
        messageChan1[label = "chan1"]
        messageChan2[label = "chan2"]
        messageChan3[label = "chan3"]
        messageChan4[label = "chan4"]
        
        listen -> messageChan [label= "启动消息分发"]
        worker1 -> messageChan1[dir = both]
        worker2 -> messageChan2[dir = both]
        worker3 -> messageChan3[dir = both]
        worker4 -> messageChan4[dir = both]
        
        messageChan1 -> messageChan [dir = both]
        messageChan2 -> messageChan [dir = both]
        messageChan3 -> messageChan [dir = both]
        messageChan4 -> messageChan [dir = both]
        
    }
}


```
处理客户端请求函数`HandleConnect`， 先进行登录请求处理，登陆成功，启动`WriteMsgToClient`和`UserMsgHandler`两个协程异步运行来进行信息分发，最后是超时检测以及断连检测
```
digraph G {
    node [shape=box]
    rankdir=LR;
    subgraph cluster_handle {
        label="处理客户端请求函数 HandleConnect"
        style=dashed

        login [label="登录请求处理"]
        login_success [label="登录成功"]
        write_msg [label="启动 WriteMsgToClient 协程"]
        user_msg [label="启动 UserMsgHandler 协程"]
        timeout [label="超时检测"]
        disconnect [label="断连检测"]

        login -> login_success [label="登录成功"]
        login_success -> write_msg
        login_success -> user_msg
        write_msg -> timeout
        user_msg -> timeout
        timeout -> disconnect
    }

    subgraph cluster_coroutines {
        label="协程"
        style=dashed

        write_msg [label="WriteMsgToClient"]
        user_msg [label="UserMsgHandler"]
    }
}
```
### Client:
主线程函数`Run`，启动两个协程函数`MessageLinster`，`SendMessageHandler` 异步进行，`MessageLinster`负责处理服务端通信，`SendMessageHandler`负责处理用户输入，最后for循环判断是否退出并阻塞主线程结束。
```
digraph G {
    node [shape=box]

    subgraph cluster_main {
        label="主线程函数 Run"
        style=dashed

        message_listener [label="启动 MessageListener 协程"]
        send_message_handler [label="启动 SendMessageHandler 协程"]
        exit_check [label="退出判断"]
        main_thread [label="阻塞主线程结束"]

        message_listener -> exit_check
        send_message_handler -> exit_check
        exit_check -> main_thread [label="退出"]
        exit_check -> message_listener [label="继续"]
    }

    subgraph cluster_coroutines {
        label="协程"
        style=dashed

        message_listener [label="MessageListener"]
        send_message_handler [label="SendMessageHandler"]
    }
}
```


## 碎碎念
课设，写个linux上的通信服务,emmmmm，拿个~~c++~~ go搓搓

写了个一坨答辩，再也不干为了这点醋 学如何包饺子的事情了(

服务端一边学go的多线程、并发一边写得，采用`master-worker`模式设计，没写好，代码的可维护性太低了，就简单的实现了接收客户端链接，并把客户端消息广播给其他客户端的功能，并且写了个超时检测，自动关闭长时间无消息的客户端链接。

客户端写的还好（自认为）

没啥安全性可言（，私聊信息的处理逻辑是写在客户端的，服务端代码写的太屎了，根本没法改了
