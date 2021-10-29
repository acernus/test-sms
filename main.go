package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://58.33.66.182:8888/GeneralSMS/pointSendMsg"
	method := "POST"

	payload := strings.NewReader("{\n  \"loginName\": \"hjapp\",\n  \"orgId\": \"10001\",\n  \"batchNo\": 1,\n  \"p\": [\n    {\n      \"text\": \"code=5120\",\n      \"password\": \"HJSHadmin@2020\",\n      \"to\": \"18515154760\",\n      \"smsId\": \"esk-00108810617858187967\"\n    }\n  ]\n}")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}