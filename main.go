package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Hello Tracee!")
	for i := 0; i < 20; i++ { // Need to wait until Tracee is ready, TODO: how do we profile for shorter than this runs?
		fmt.Printf("Tick %d\n", i)
		time.Sleep(time.Second)
	}
	doSomethingMalicious()
}

func doSomethingMalicious() {
	cmd := exec.Command("./poc.py")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("exploit failed, ", err)
	}
	fmt.Println("exploit successful!")
}
