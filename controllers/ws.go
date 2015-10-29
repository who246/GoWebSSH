package controllers
import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"io/ioutil"  
	"encoding/base64"   
	"github.com/who246/GoWebSSH/ssh_login"
	"github.com/who246/GoWebSSH/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
type WSController struct {
	beego.Controller
	ws *websocket.Conn 
	send chan []byte
	ssh_login.SSH_Config
}
 const (
	writeWait = 10 * time.Second
	readWait = 5*60 * time.Second
	pingPeriod = (60 * time.Second * 9) / 10
	)
func (this *WSController) Get() {
	id,err := this.GetInt("id",-1) 
	cmdIdStr := this.GetString("cmdId","") 
	m := models.Server{Id: id}
	err = orm.NewOrm().Read(&m,"id")
	if err!=nil {
		beego.Error(err) 
		return;
	}
	
	this.Ip = m.Ip
	this.Username = m.LoginName
	this.Password = m.LoginPassword 
	this.ws, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	this.send = make(chan []byte, 256)

	defer func() { 
	     
		if this.ws!=nil {
		this.ws.Close()
		} 
	    this.Close() 
	}()
	if err = this.Login(); err != nil {
		http.Error(this.Ctx.ResponseWriter, err.Error(), 400)
		return
	}
	if cmdIdStr != "" {
	var cmdId int;
	cmdId,err =strconv.Atoi(cmdIdStr)
	c := models.Command{Id: cmdId}
	err = orm.NewOrm().Read(&c,"id")
	if err!=nil {
		beego.Error(err) 
		return;
	}
	this.cmdWrite([]byte(c.Cmd+"\n"));
	}
	go this.wsWrite()
	go this.cmdRead()
	this.wsRead();
	
}

func (this *WSController)cmdRead(){
	defer func() {
	close(this.send)
	}()
    buf := make([]byte, 1024)
 
	for {
		 
		size, err := this.Pty.Read(buf)
		if err != nil {
			beego.Error(err)
			return
		}
		safeMessage := base64.StdEncoding.EncodeToString([]byte(buf[:size])) 
		
		this.send <- []byte(string(safeMessage))
		 
	}
} 

func (this *WSController)cmdWrite(b []byte) error{
	_, err := this.Pty.Write(b)
	 return err
} 


func (this *WSController) wsRead() {
	
	for{
		 this.ws.SetReadDeadline(time.Now().Add(readWait))
		op, r, err := this.ws.NextReader()
		if err != nil {
			beego.Error(err)
			return
		}
		switch op {
		case websocket.PongMessage:
			this.ws.SetReadDeadline(time.Now().Add(readWait))
		case websocket.TextMessage:
			message, err := ioutil.ReadAll(r)
			err = this.cmdWrite(message);
			if err != nil {
				beego.Error(err)
				return
			}
			 
		}
	}
}

func (this *WSController) wsWrite() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {		
		ticker.Stop()		
	}()

	for {
		select {
		case message, ok := <-this.send:
			if !ok {
				this.w(websocket.CloseMessage, []byte{})
				return
			} 
			if err := this.w(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C: 
			if err := this.w(websocket.PingMessage, []byte{}); err != nil {
				return
			} 
		}
	}
}
func (this *WSController) w(opCode int, data []byte) error {
	this.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return this.ws.WriteMessage(opCode, data)
}