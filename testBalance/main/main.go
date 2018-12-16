package main

import (
	"fmt"
	"gotest/testBalance/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	var balancer balance.Balancer
	var conf = "random"
	if len(os.Args) > 1 {
		conf = os.Args[1]
	}

	if conf == "random" {
		balancer = &balance.RandomBalance{}
		fmt.Println("random")
	} else if conf == "roundrobin" {
		balancer = &balance.RoundRobinBalance{}
		fmt.Println("roundrobin")
	}

	for {
		inst, err := balancer.DoBalance(insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
