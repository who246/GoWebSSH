package ssh_login

import (  
	"os"
	"os/exec"
	"github.com/kr/pty"
	"strings" 
	"errors"
	"syscall"
)
type SSH_Config struct {
	      Ip string
	Username string 
	Password string
	Command *exec.Cmd
	Pty *os.File
} 

func (this *SSH_Config) Close(){
	 if this.Command!=nil {
		this.Command.Process.Signal(syscall.SIGHUP)
	}
	if this.Pty!=nil {
		this.Pty.Close()
	}
}
func (this *SSH_Config) Login() (  error){ 
	var err error
	this.Command = exec.Command("ssh","-o","StrictHostKeyChecking=no",this.Username+"@"+this.Ip)
	this.Pty ,err = pty.Start(this.Command)
	if err !=nil{
		return  err;
	} 
	i:=0;
	for{
	if i >= 10 { 
		
		return  errors.New("login error") 
	}
	buf := make([]byte, 1024)
	size, err2 := this.Pty.Read(buf)
		if err2 != nil { 
			return  err;
		}   
		
	 
	if !strings.Contains(string([]byte(buf[:size])),"password"){
		i++;
		continue;
	}
	this.Pty.Write([]byte(this.Password+"\n"))
	 if err != nil {
        panic(err)
    }
	if err != nil { 
		return  errors.New("login error") 
	}
	return   nil;  
	}
}