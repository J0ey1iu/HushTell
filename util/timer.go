package util

import (
	"os"
	"time"
	"fmt"
)

func createFile() {
	os.Create("./file.test")
	createTime := time.Now()
	go checkTimer(createTime)
}

func checkTimer(createTime time.Time) {
	for 1==1 {
		fmt.Println(time.Now().Sub(createTime))
		time.Sleep(3*time.Second)
	}
}

