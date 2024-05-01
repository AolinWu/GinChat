package log

import (
	"log"
	"os"
)

var Logger *log.Logger=nil
var f *os.File=nil
const (
	DEBUG_LEVEL = "DEBUG"
	INFO_LEVEL = "INFO"
	ERROR_LEVEL ="ERROR"
	FATAL_LEVEL = "FATAL"
)
func withOutputFile(path string){
	 var err error
     f,err=os.OpenFile(path,os.O_APPEND|os.O_WRONLY|os.O_CREATE,0666)
	 if err!=nil{
		panic(err)
	 }
     if Logger==nil{
		panic("uninitialized logger")
	 }
	 Logger.SetOutput(f)
}
func Clean() error{
    if f!=nil{
		err:=f.Close()
		return err
	}
	return nil
}
func Debug(msg interface{}){
	Logger.Printf("====%s===%v\n",DEBUG_LEVEL,msg)
}
func Info(msg interface{}){
	Logger.Printf("====%s===%v\n",INFO_LEVEL,msg)
}

func Fatal(msg interface{}){
    Logger.Printf("====%s===%v\n",FATAL_LEVEL,msg)
}
func Error(msg interface{}){
	Logger.Printf("====%s===%v\n",ERROR_LEVEL,msg)
}

func init(){
	Logger=log.Default()
	withOutputFile("log.txt")
}