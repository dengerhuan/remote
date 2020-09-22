package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var remoteAdd = "http://192.168.100.249:8096"

// application/json
func Post(url string, data interface{}, contentType string) (res gin.H) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		//panic(err)
		log.Println(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(result)
	json.Unmarshal(result, &res)
	return res
}

/**
"{
“orderId”:“2731489313”
   }"
*/
func ErdApplyReq(data map[string]interface{}) {
	res := Post(remoteAdd+"/erd/apply", data, "application/json")

	fmt.Println(res)
}

func ErdApplyAck(data map[string]interface{}) {

	req := gin.H{"orderId": data["orderId"].(string),
		"distance":  37.3,
		"time":      53,
		"latitude":  125.3249352,
		"longitude": 43.8593245,
		"direction": 37.3,
		"cause":     "用户强制结束",
		"causeCode": 222}

	res := Post(remoteAdd+"/erd/ack", req, "application/json")

	fmt.Println(res)
}

/*
" {
“orderId”:“2731489313”，
“distance”:37.3,
“time”:53,
“latitude”: 125.3249352,
“longitude”: 43.8593245,
“direction”: 37.3，
“cause”: “用户强制结束”,
“causeCode”:222
   }"
*/
func ersAck() {
	res := Post("", gin.H{"orderId": "2731489313"}, "application/json")

	fmt.Println(res)
}
