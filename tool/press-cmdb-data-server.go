package tool

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)


func SetCMDBInfo(wg *sync.WaitGroup) {
	cmdb_url := "http://dev-cmdb.ops.ipfsyuanli.com/api/v1/data/server"

	client := &http.Client{}
	cmdbToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjA5OTgyOTY2LCJlbWFpbCI6IiJ9.GARIfIGGiQs6SosrYHG26WaAW-5_K7Jg5Our3cn_f1k"
	var cmdbSearchParamsFormat = `{
		"pool__key":"b",
		"group_ids":"bb",
		"group_id":"2",
		"p4_group_id":"b",
		"seq_group_id":"b",
		"bk_host_innerip":"b",
		"ssh_port":"1",
		"bk_state":"b",
		"config_item__key":"b",
		"cpu_type":"b",
		"cluster_type":"b"
	}`

	var cmdbSearchParamsTmp = fmt.Sprintf(cmdbSearchParamsFormat)
	var cmdbSearchParams = []byte(cmdbSearchParamsTmp)


	for i := 0; i< 1000; i++ {
		req1, err := http.NewRequest("POST", cmdb_url, bytes.NewBuffer(cmdbSearchParams))
		req1.Header.Set("Content-Type", "application/json")
		my_jwt := fmt.Sprintf("JWT %s", cmdbToken)
		req1.Header.Set("Authorization", my_jwt)
		req1.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6")

		//_, err = client.Do(req1)
		//if err != nil {
		//	panic(err)
		//}

		cmdbSearchResult1, err := client.Do(req1)
		if err != nil {
			panic(err)
		}

		cmdbSearchResultBody, err := ioutil.ReadAll(cmdbSearchResult1.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer cmdbSearchResult1.Body.Close()
		cmdbSearchResult2 := string(cmdbSearchResultBody)
		println("%s",cmdbSearchResult2)
	}

	wg.Done()
}
