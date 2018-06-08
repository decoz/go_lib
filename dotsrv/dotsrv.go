package dotsrv

import (
	"strconv"
	"log"
	"golang.org/x/net/websocket"
	//"github.com/gorilla/websocket"
	"github.com/decoz/go_lib/dot"
	"net/http"
)

type dotsrv struct {
	dsrv service
}

type service interface{
	Put(*dot.Dot) *dot.Dot
}

func (ds *dotsrv) listen(ws *websocket.Conn){
	defer func() {
	 log.Println("connection closed")
 	}()
 
 for ws.IsServerConn() {
	 	//size, err := ws.Read(msg)
		var msg string
		websocket.Message.Receive(ws, &msg)
		d_msg := dot.Make(msg)
	 	r := ds.dsrv.Put(d_msg)

	 	if r != nil {
	 		ws.Write([]byte(r.String()))
		}

 }
}


func Lunch(port int , srv service) error{
	wssrv := new( dotsrv )
	wssrv.dsrv = srv
	log.Println("start srv")
	http.Handle("/ws", websocket.Handler(wssrv.listen))
	err := http.ListenAndServe(":" + strconv.Itoa(port) , nil)

	if err != nil {
		panic("Listen and serve:" + err.Error())
	}

	return err


}
