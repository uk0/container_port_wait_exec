package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05";

var wg = sync.WaitGroup{}

var IS_START = false

func main() {
	wg.Add(1)
	addrs := flag.String("addrs", "127.0.0.1:8080", "addr arrays")
	cmds := flag.String("cmds", "ls -al ", "command some")
	flag.Parse()
	t1 := time.NewTicker(time.Duration(1000) * time.Millisecond)
	go func() {
		for t := range t1.C {
			if !IS_START {
				fmt.Println("Test Port Wait.... ", t.Format(TIME_FORMAT))
				for _, addr := range strings.Split(*addrs, ",") {
					fmt.Println(addr)
					if 1 == CheckPortStatus(addr, "tcp") {
						// 改成启动
						IS_START = true
						WindowsPacSet(fmt.Sprintf(`%s`,*cmds))
					}
				}
			}else {
				fmt.Println(" Server is running.....")
			}
		}
		wg.Done()
	}()
	wg.Wait()
	t1.Stop()

}
func WindowsPacSet(CmdLine string) {
	cmd := exec.Command("/bin/sh")
	cmd.Stdout = os.Stdout
	input, _ := cmd.StdinPipe()
	_ = cmd.Start()
	_, _ = fmt.Fprintln(input, CmdLine)
	_ = cmd.Wait()
}

func CheckPortStatus(address string, types string) (int) {

	switch types {
	case "tcp":
		return tcp(address);
		break;
	case "udp":
		return udp(address);
		break;
	default:
		return 0
	}
	return 0
}

func tcp(url string) int {
	_, err := net.Dial("tcp", url)
	if err != nil {
		return 0
	} else {
		return 1
	}
}

func udp(url string) int {
	_, err := net.Dial("udp", url)
	if err != nil {
		return 0
	} else {
		return 1
	}
}
