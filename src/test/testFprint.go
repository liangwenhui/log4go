package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main4() {
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	var iLoopTimes = 100
	var test_times = 10000
	start := time.Now()
	pLogFile, _ := os.OpenFile("./log4go.log", os.O_CREATE|os.O_RDWR, 0766)
	defer pLogFile.Close()
	writer := bufio.NewWriter(pLogFile)
	for j := 0; j < iLoopTimes; j++ {
		for i := j; i < test_times; i++ {
			fmt.Fprintf(writer, "INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsadawsadawsadsadwsadsadwsadawsadsadwsadawsadsadwsad", 3121.21212)
		}
		for i := test_times - 1; i > j; i-- {
			fmt.Fprintf(writer, "INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsadawawsadsadwsadsadsadwsadawsadsadwsadawsadsadwsad", 12314.21231)
		}
	}
	since := time.Since(start)
	fmt.Println(since)
	fmt.Printf("Press any key to exit...\n")

	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func main1() {
	//b, err := syscall.Mmap(-1, 0, syscall.Getpagesize(), syscall.PROT_NONE, syscall.MAP_ANON|syscall.MAP_PRIVATE)

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

func main3() {
	f, _ := os.OpenFile("./log4go.log", os.O_CREATE|os.O_RDWR, 0766)
	msg := "a"

	f.WriteAt([]byte(msg), 1)
	//data, err := syscall.Mmap(int(m.f.Fd()), int64(m.at), int(m.size), syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
}
