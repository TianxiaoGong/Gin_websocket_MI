package utils

import (
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}
func RespOK(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, SUCCESS, data, msg)
}
func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, FAILED, nil, msg)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}
func RespOKList(w http.ResponseWriter, data interface{}, total interface{}) {
	RespList(w, SUCCESS, data, total)
}
func RespFailList(w http.ResponseWriter, data interface{}, total interface{}) {
	RespList(w, FAILED, data, total)
}
