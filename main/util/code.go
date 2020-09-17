package util

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)
var BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
func Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(BASE64Table)
	return coder.EncodeToString(content)
}

func SumTime(timeList []string)(hour,minute int){

	var middleTime int

	for _,v := range timeList{
		middleTime,_ = strconv.Atoi(strings.Split(v,":")[0])
		hour += middleTime
		middleTime,_ = strconv.Atoi(strings.Split(v,":")[1])
		minute += middleTime
		if(minute>60){
			minute -= 60
			hour +=1
		}
	}
	return hour,minute
}

