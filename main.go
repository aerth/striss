package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
	"strconv"
)

var num_conn uint = 1014

var svr_address string
var tick_interval uint = 500
const buffer_size = 4 * 1024
const print_progress = false

func create_getter(uri string,
	bufsize int,
	quit chan uint16,
	death chan uint16) func(uint) {
	return func(id uint) {
		defer func() { death <- 1 }()
		if print_progress {
			fmt.Println()
		}
		fmt.Println("starting go routine", id)
		var cnt uint = 0
		buff := make([]byte, bufsize)
		conn, err := net.Dial("tcp", uri)
		for err != nil {
			fmt.Fprintln(os.Stderr, err)
			time.Sleep(5*time.Second)
			conn, err = net.Dial("tcp", uri)
		}
		defer conn.Close()

// This is where you set the mountpoint (until i can figure out the CLI args)

		fmt.Fprint(conn, "GET /stream HTTP/1.0\r\n\r\n")
		for {
			select {
			case <-quit:
				if print_progress {
					fmt.Println()
				}
				fmt.Println("terminating go routine", id)
				return
			default:
				if n, err := conn.Read(buff); n == 0 {
					if print_progress {
						fmt.Println()
					}
					fmt.Println("Reading connection", id, "failed:", err)
					return
				}
				if cnt++; print_progress && cnt%num_conn == 0 {
					fmt.Print(".")
				}
			}
		}
	}
}

func printUsage() {
	fmt.Print("\n\nicecast-stress\n\nUsage: ", os.Args[0], " <host.server.com:port> <num_conn> [interval]\n")
	fmt.Print("Defaults to mountpoint: /stream -- If you need, go in and rebuild!\n\n")
//	fmt.Print("Example: ", os.Args[0], " example.com:8000 /stream 250\n")
}

func init() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(-1)
	}

	svr_address = os.Args[1]
	if n, e := strconv.ParseUint(os.Args[2], 10, 0); e != nil {
		printUsage()
		os.Exit(-1)
	} else {
		num_conn = uint(n)
	}

	if len(os.Args) > 3 {
		if n, e := strconv.ParseUint(os.Args[3], 10, 0); e != nil {
			printUsage()
			os.Exit(-1)
		} else {
			tick_interval = uint(n)
		}
	}
}

func main() {
	quit := make(chan uint16)
	death := make(chan uint16, 8)

	get := create_getter(svr_address, buffer_size, quit, death)

	s := make(chan os.Signal, 2)
	signal.Notify(s)
	var sig os.Signal

	ticker := time.NewTicker(time.Duration(tick_interval) * time.Millisecond)
	var born, dead uint
M:
	for born < num_conn {
		select {
		case <-ticker.C:
			born++
			go get(born)
		case sig = <-s:
			break M
		}
	}
	ticker.Stop()

	for sig == nil && dead < born {
		select {
		case sig = <-s:
		case <-death:
			dead++
		}
	}
	signal.Stop(s)

	if print_progress {
		fmt.Println()
	}
	fmt.Println("Terminating", born-dead, "go routine(s)..")

	for dead < born {
		select {
		case quit <- uint16(1):
		case <-death:
			dead++
			fmt.Println("Alive:", born-dead, "go routine(s).")
		}
	}

	if print_progress {
		fmt.Println()
	}
	fmt.Println("Got signal:", sig)
}
