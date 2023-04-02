package dao

import (
	"Gin_WebSocket_IM/models"
	"Gin_WebSocket_IM/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gopkg.in/fatih/set.v0"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type MsgDao struct {
}

func NewMsgDao() *MsgDao {
	return &MsgDao{}
}

type Node struct {
	Conn      *websocket.Conn
	Addr      string
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func (md *MsgDao) MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println("Subscribe err:", err)
			return
		}
		t := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", t, msg)
		err = ws.WriteMessage(1, []byte(m))
	}
}

func (md *MsgDao) Chat(writer http.ResponseWriter, request *http.Request) {
	//1、校验token等合法性
	query := request.URL.Query()
	//token := query.Get("token")
	isValida := true //checkToken()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//targetId := query.Get("targetId")
	//msgType := query.Get("type")
	//context := query.Get("context")
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isValida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//2、获取连接
	node := &Node{
		Conn:      conn,
		Addr:      conn.RemoteAddr().String(),
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe), //线程安全
	}
	//3、用户关系
	//4、userid跟node绑定 并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	//5、完成发送逻辑
	go sendProc(node)
	//6、完成接收逻辑
	go receiveProc(node)
	//sendMsg(userId, []byte("欢迎进入聊天室"))
	models.SetUserOnlineInfo("online_"+Id, []byte(node.Addr), time.Duration(viper.GetInt("timeout.RedisOnlineTime"))*time.Hour)
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("[ws]sendProc >>> msg:", string(data))
			if err := node.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func receiveProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws]receiveProc <<< msg:", string(data))
	}
}

var udpsendChan = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	fmt.Println("init goroutine")
	go udpSendProc()
	go udpReceiveProc()
}

// 完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//IP:   net.IPv4(172, 26, 208, 1),
		IP:   net.IPv4(10, 100, 158, 189),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		select {
		case data := <-udpsendChan:
			fmt.Println("udpSendProc data :", string(data))
			if _, err = con.Write(data); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// 完成udp数据接收协程
func udpReceiveProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpReceiveProc data :", string(buf[0:n]))
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := models.Message{}
	msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1:
		fmt.Println("dispatch data :", string(data))
		fmt.Println("dispatch data :", string(data))
		sendMsg(msg.TargetId, data)
		//case 2:
		//	sendGroupMsg()
		//case 3:
		//	sendAllMsg()
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	jsonMsg := models.Message{}
	json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	//fmt.Println("jsonMsg", jsonMsg)
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(jsonMsg.FromId))
	r, err := utils.Rdb.Get(ctx, "online_"+userIdStr).Result()
	if err != nil {
		fmt.Println(err) //没有找到
	}
	if r != "" {
		if ok {
			fmt.Println("[ws]sendMsg >>> userId:", userId, " msg:", string(msg))
			node.DataQueue <- msg
		}
	}
	var key string
	if userId > jsonMsg.FromId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}
	fmt.Println("sendMsg__key:", key)
	res, e := utils.Rdb.ZAdd(ctx, key, &redis.Z{Score: 1, Member: msg}).Result()
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(res)
}

//// MarshalBinary 需要重写此方法才能完整的msg转byte[]
//func (msg Message) MarshalBinary() ([]byte, error) {
//	return json.Marshal(msg)
//}

// GetRedisMsg 获取缓存里面的消息
func (md *MsgDao) GetRedisMsg(userIdA int64, userIdB int64) []interface{} {
	RdbMsg := make([]interface{}, 0)
	rwLocker.RLock()
	node, _ := clientMap[userIdA]
	rwLocker.RUnlock()
	//jsonMsg := Message{}
	//json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	//fmt.Println("AAA", userIdA)
	//fmt.Println("BBB", userIdB)
	userIdStr := strconv.Itoa(int(userIdA))
	targetIdStr := strconv.Itoa(int(userIdB))
	var key string
	if userIdA > userIdB {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}
	fmt.Println("RedisMsg__key:", key)
	rels, err := utils.Rdb.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println(err) //没有找到
	}
	for _, val := range rels {
		RdbMsg = append(RdbMsg, val)
		fmt.Println("sendMsg >>> userID: ", userIdA, " msg:", val)
		node.DataQueue <- []byte(val)
	}
	return RdbMsg
}
