package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var queue []string

func enqueue(queue []string, element string) []string {
	queue = append(queue, element) // Simply append to enqueue.
	//fmt.Println("Enqueued:", element)
	return queue
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		newdata, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(newdata)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("Received message -> ", string(newdata))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
		queue = enqueue(queue, string(newdata))
		//fmt.Println("Queue:", queue)

		if strings.TrimSpace(string(newdata)) == "ALL" {
			//queue = enqueue(queue, string(newdata))
			fmt.Println("Queue:", queue)
		}

	}
}
