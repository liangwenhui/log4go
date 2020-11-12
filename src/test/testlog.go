package main

import (
	"fmt"
	"lwhlog"
	"os"
	"time"
)

//func main() {
//	config := lwhlog.NewLogConfig("./log4g0.out",10)
//	log := lwhlog.NewLog(*config)
//		start := time.Now()
//	for  i:=0;i<2;i++{
//		//log.Infof("INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsadawsadawsadsadwsadsadwsadawsadsadwsadawsadsadwsad", 3121.21212)
//		go p(log)
//	}
//
//		since := time.Since(start)
//		fmt.Println("耗时：",since)
//		fmt.Printf("Press any key to exit...\n")
//
//		b := make([]byte, 1)
//		os.Stdin.Read(b)
//}
//
//func p(log *lwhlog.Logger){
//	for  i:=0;i<1;i++{
//		log.Infof("INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsadawsadawsadsadwsadsadwsadawsadsadwsadawsadsadwsad", 3121.21212)
//	}
//
//}

func main() {
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	config := lwhlog.NewLogConfig("./log4g0.out", 10)
	log := lwhlog.NewLog(*config)
	var iLoopTimes = 100
	var test_times = 10000
	start := time.Now()
	for j := 0; j < iLoopTimes; j++ {
		for i := j; i < test_times; i++ {
			log.Infof("INTERGET:%d|STRING:%s|FLOAT:%f\n", AiRandomInt[i], AsRandomStr[i], AfRandomFloat[i])
		}
		for i := test_times - 1; i > j; i-- {
			log.Infof("INTERGET:%d|STRING:%s|FLOAT:%f\n", AiRandomInt[i], AsRandomStr[i], AfRandomFloat[i])
		}
		//for i := j; i < test_times; i++ {
		//	log.Infof( "INTERGET:%d|STRING:%s|FLOAT:%f\n", i, "awsadsadwsadawsadawsadsadwsadsadwsadawsadsadwsadawsadsadwsad", 3121.21212)
		//}
	}
	since := time.Since(start)
	fmt.Println("耗时：", since)
	fmt.Printf("Press any key to exit...\n")

	b := make([]byte, 1)
	os.Stdin.Read(b)

}
