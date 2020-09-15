package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
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

func Read(path string)*os.File{
	//var file *os.File
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

func ReadLog(c *gin.Context){
	var file *os.File
	file ,err:=os.Open("nohup.out")
	if err!=nil{
		log.Println(err)
	}
	http.ServeContent(c.Writer, c.Request, file.Name(), time.Now(), file)


}

//解析一个学号+姓名的表
func AnalyzeExcel(path string) [1000][2]string{
	xlFile, err:= xlsx.OpenFile(path)
	if(err!=nil){
		log.Println(err)
	}


	var result [1000][2] string


	for _,sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if (rowIndex > 0) {
				if len(row.Cells) <2{
					break
				}
				for cellIndex, cell := range row.Cells {
					text := cell.String()
					if(len(text)==0){
						break
					}
					result[rowIndex-1][cellIndex] = text
				}
			}
		}
		break
	}

	for i:=0;i<3;i++{
		fmt.Println(result[i][0]+result[i][1])
	}

	return result

}