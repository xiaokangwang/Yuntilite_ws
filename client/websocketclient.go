package main

import "golang.org/x/net/websocket"
import "net"
import "io"
import "os"
import golanglog "log"

func OnHandleConnection(localsocket net.Conn, intentaddr string) {
	wsc, _ := ConnectToServer(sever_addr, "discard")
	var theYuntiLeapAction YuntiLeapAction
	theYuntiLeapAction.StreamId = "$CTLMSG"
	theYuntiLeapAction.ActionType = "ConnectImmerse"
	theYuntiLeapAction.Data = intentaddr
	websocket.JSON.Send(wsc, theYuntiLeapAction)
	go io.Copy(wsc, localsocket)
	io.Copy(localsocket, wsc)
}

func ConnectToServer(saddr string, token string) (*websocket.Conn, error) {
	wsc, err := websocket.Dial(saddr, "", "https://kkdev.org/?spec=yuntilite_ws_v1&token="+token)
	if err != nil {
		golanglog.Println(err)
	}
	return wsc, err
}

var sever_addr string

func main() {
	//var local_addr string
	local_addr := os.Getenv("local_addr")
	sever_addr = os.Getenv("sever_addr")
	Run_socks5(local_addr)
}
