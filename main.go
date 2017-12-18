package main

import (
	"os/exec"
	"fmt"
	"regexp"
	"time"
)

func main() {
	c := time.Tick(5 * time.Minute)
	for range c {
		dockerMonitor()
	}
}

func dockerMonitor() {
	out, err := exec.Command("docker", "ps", "-a").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("out is %s\n", out)
	redis_up, _ := regexp.MatchString("redis.*?Up.*?leon-redis", string(out))
	// fmt.Println("redis is up:", redis_up)
	if !redis_up {
		if err := exec.Command("docker", "start", "leon-redis").Start(); err != nil {
			fmt.Println(err)
		}
	}
	mongo_up, _ := regexp.MatchString("mongo.*?Up.*?leon-mongo", string(out))
	// fmt.Println("mongo is up:", mongo_up)
	if !mongo_up {
		exec.Command("docker", "start", "leon-mongo").Start()
	}
}
