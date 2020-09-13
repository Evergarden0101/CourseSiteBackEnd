package util

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func Write(file *multipart.FileHeader,path string) bool{

	filestream,err :=file.Open()
	defer filestream.Close()
	if err!=nil{
		log.Println(err)
		return false
	}
	os.Mkdir(path,os.ModePerm)
	out,err :=os.Create(path+"/"+file.Filename)
	defer out.Close()
	if err!=nil{
		log.Println(err)
		return false
	}
	_,err=io.Copy(out,filestream)
	if err!=nil{
		log.Println(err)
		return false
	}
	return true
}
//func Write1(filestream multipart.File,path string) bool{
//
//
//	defer filestream.Close()
//	//if err!=nil{
//	//	log.Println(err)
//	//	return false
//	//}
//	os.Mkdir(path,os.ModePerm)
//	out,err :=os.Create(path+"/"+"test")
//	defer out.Close()
//	if err!=nil{
//		log.Println(err)
//		return false
//	}
//	_,err=io.Copy(out,filestream)
//	if err!=nil{
//		log.Println(err)
//		return false
//	}
//	return true
//}
func Read(path string)*os.File{
	var file *os.File
	log.Println(path)
	file ,err:=os.Open(path)
	if err!=nil{
		log.Println(err)
		return nil
	}
	return file
}
func Image(c *gin.Context){
	var file *os.File
	file ,err:=os.Open("upload1.png")
	if err!=nil{
		log.Println(err)
	}
	http.ServeContent(c.Writer, c.Request, file.Name(), time.Now(), file)



}