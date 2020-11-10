package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main1() {
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	var iLoopTimes = 100
	var test_times = 10000
	start := time.Now()
	pLogFile, _ := os.OpenFile("D:\\log4go.log", os.O_CREATE|os.O_RDWR, 0766)
	defer pLogFile.Close()
	writer := bufio.NewWriter(pLogFile)
	for j := 0; j < iLoopTimes; j++ {
		for i := j; i < test_times; i++ {
			fmt.Fprintf(writer, "INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsad", 3.2)
		}
		for i := test_times - 1; i > j; i-- {
			fmt.Fprintf(writer, "INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsad", 4.2)
		}
	}
	since := time.Since(start)
	fmt.Print(since)

}

func main() {
	var iLoopTimes = 100
	var test_times = 10000
	start := time.Now()
	pLogFile, _ := os.OpenFile("D:\\log4go.log", os.O_CREATE|os.O_RDWR, 0766)
	defer pLogFile.Close()
	for j := 0; j < iLoopTimes; j++ {
		for i := j; i < test_times; i++ {
			fmt.Sprintf("INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsad", 4.2)
		}
		for i := test_times - 1; i > j; i-- {
			fmt.Sprintf("INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsad", 4.2)
		}
	}
	since := time.Since(start)
	fmt.Print(since)
}

//func RandString(len int) string {
//	bytes := make([]byte, len)
//	for i := 0; i < len; i++ {
//		b := r.Intn(26) + 65
//		bytes[i] = byte(b)
//	}
//	return string(bytes)
//}
