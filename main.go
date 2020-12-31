package main


import (
	"fmt"
	"press_test/tool"
	"sync"
	"time"
)


func main() {
	startTime := time.Now().UnixNano()
    wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		//go tool.GetCMDBInfo(&wg)
		go tool.SetCMDBInfo(&wg)
	}
	wg.Wait()
    endTime := time.Now().UnixNano()
    fmt.Printf("花费时间秒数：%e\n", float64((endTime - startTime) / 1e9))
}