package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"wenqianIm/conf"
	"wenqianIm/pkg/e"
)

func (manger *ClientManager) Start() {
	for{
		log.Println("监听管道")
		select {
		case conn:=<-Manager.Register:
			log.Printf("建立新连接:%v",conn.ID)
			Manager.Clients[conn.ID]=conn
			replyMsg:=&ReplyMsg{
				Code:e.WebsocketSuccess,
				Content: "已连接至服务器",
			}
			msg,_:=json.Marshal(replyMsg)
			_=conn.Socket.WriteMessage(websocket.TextMessage,msg)

		case conn:=<-Manager.Unregister:
			log.Printf("连接失败:%v",conn.ID)
			if _,ok:=Manager.Clients[conn.ID];ok{
				replyMsg:=&ReplyMsg{
					Code: e.WebsocketEnd,
					Content: "连接已断开",
				}
				msg,_:=json.Marshal(replyMsg)
				_=conn.Socket.WriteMessage(websocket.TextMessage,msg)
				close(conn.Send)
				delete(Manager.Clients,conn.ID)
			}
			//广播消息
			case broadcast:=<-Manager.Broadcast:
				message:=broadcast.Message
				sendID:=broadcast.Client.SendID
				flag:=false//默认对方不在线
				for id ,conn:=range Manager.Clients{
					if id!=sendID{
						continue
					}
					select {
					case conn.Send<-message:
						flag=true
					default:
						close(conn.Send)
						delete(Manager.Clients,conn.ID)
					}
				}
				id:=broadcast.Client.ID
				if flag{
					log.Println("对方在线应答")
					replyMsg:=&ReplyMsg{
						Code: e.WebsocketOnlineReply,
						Content: "对方在线应答",
					}
					msg,err:=json.Marshal(replyMsg)
					_=broadcast.Client.Socket.WriteMessage(websocket.TextMessage,msg)
					err=InsertMsg(conf.MongoDBName,id,string(message),1,int64(3*month))
					if err!=nil{
						fmt.Println("InsertOneMsg Err",err)
					}
				}else{
					log.Println("对方不在线")
					replyMsg:=ReplyMsg{
						Code: e.WebsocketOnlineReply,
						Content: "对方不在线应答",
					}
					msg,err:=json.Marshal(replyMsg)
					_=broadcast.Client.Socket.WriteMessage(websocket.TextMessage,msg)
					err=InsertMsg(conf.MongoDBName,id,string(message),0,int64(3*month))
					if err!=nil{
						fmt.Println("InsertOnMsg Err",err)
					}
				}
		}
	}
}


























