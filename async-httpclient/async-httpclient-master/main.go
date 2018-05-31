package main

import (
	"fmt"
	"time"

	"github.com/weimumu/async-httpclient/myAgency"
)

func main() {
	STime1 := time.Now()
	myAgency.SyncTask()
	fmt.Println("SyncTask Total Time:", time.Since(STime1))

	STime2 := time.Now()
	myAgency.AsyncTask()
	fmt.Println("AsyncTask Total Time:", time.Since(STime2))
}
