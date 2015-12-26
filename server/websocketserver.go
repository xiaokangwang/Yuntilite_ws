package main

import "golang.org/x/net/websocket"
import "io"

//import "net/url"
//import "errors"
import "os"
import "net"
import "net/http"
import "log"

func yuntilite_ws_v1_Server(ws *websocket.Conn) {
	var theYuntiLeapAction YuntiLeapAction
	websocket.JSON.Receive(ws, &theYuntiLeapAction)
	if !(theYuntiLeapAction.StreamId == "$CTLMSG" && theYuntiLeapAction.ActionType == "ConnectImmerse") {
		return
	}
	conn, err := net.Dial("tcp", theYuntiLeapAction.Data)
	if err != nil {
		return
	}
	go io.Copy(conn, ws)
	io.Copy(ws, conn)
	return
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	wss := new(websocket.Server)
	wss.Handler = yuntilite_ws_v1_Server
	http.Handle("/"+os.Getenv("yuntilite_ws_v1_token"), wss)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
