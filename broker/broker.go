package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a port!")
		return
	}
	PORT := ":" + args[1]
	t := RoutingTable{table: map[string]string{}, log: map[time.Time]string{}, members: map[string]bool{}}
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(c, t)
	}
}

func handleConnection(c net.Conn, t RoutingTable) {
	addr := c.RemoteAddr().String()
	fmt.Printf("Serving %s\n", addr)
	t.addMember(addr)
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		temp := strings.TrimSpace(string(netData))
		// TODO: differentiate incoming requests
		if temp == "STOP" {
			c.Write([]byte("Shutting down peer..."))
			break
		} else {
			result := strconv.Itoa(404) + "\n"
			c.Write([]byte(string(result)))
		}
	}
	t.removeMember(addr)
	c.Close()
}
