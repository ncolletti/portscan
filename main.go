package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type ScanConfig struct {
	port    string
	next    bool
	network string
}

var err error
var port string
var next bool
var network string
var help bool
var verboose bool
var osStdOut *os.File

func main() {
	registerFlags()

	flag.Parse()

	verifyInput()

	scan := NewScan(port, next, network)

	scan.port, err = checkPortIsAvailable(*scan)
	if err != nil {
		os.Exit(2)
	}

	// return port
	fmt.Fprintln(osStdOut, scan.port)
	os.Exit(0)
}

func NewScan(p string, na bool, n string) *ScanConfig {
	return &ScanConfig{port: p, next: na, network: n}
}

func usage() {
	fmt.Fprintf(os.Stdout, `Usage:	portscan -p 8888 -n -nt tcp6
				portscan -p - port - Port to check if available
				portscan -n - next - Return next available closest port
				portscan -nt - network - Specify a network - 'tcp'(default) 'tcp4' 'tcp6'
				portscan -v - verboose - Enable logging
`)
}

func registerFlags() {
	flag.StringVar(&port, "p", "", "Port to check if available")
	flag.BoolVar(&next, "n", false, "Return next available closest port")
	flag.StringVar(&network, "nt", "tcp4", "Specify a network - 'tcp'(default) 'tcp4' 'tcp6'")
	flag.BoolVar(&help, "h", false, "Instructions for use")
	flag.BoolVar(&verboose, "v", false, "Enable logging")
}

func verifyInput() {
	osStdOut = os.Stdout

	if help {
		usage()
		os.Exit(1)
	}

	if !verboose {
		os.Stdout, _ = os.Open("os.DevNull")
		os.Stderr, _ = os.Open("os.DevNull")
	}

	if port == "" && verboose {
		fmt.Fprintf(os.Stdout, "No port was given.\n")
		reader := bufio.NewReader(os.Stdin)

		fmt.Fprintln(os.Stdout, "Enter port number you wish to scan: ")

		port, err = readFromUser(reader)

		if err != nil {
			fmt.Fprintf(os.Stdout, "Error reading port: %v", err)
			return
		}
	}
}

func checkPortIsAvailable(s ScanConfig) (string, error) {
	ln, err := net.Listen(s.network, ":"+s.port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s\n", s.port, err)
		if s.next {
			i, err := strconv.Atoi(s.port)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to convert to int")
				os.Exit(2)
			}
			i++
			s.port = strconv.Itoa(i)
			fmt.Fprintf(os.Stdout, "Checking for next available")
			return checkPortIsAvailable(s)
		}
		return "", err
	}

	err = ln.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't stop listening on port. %q: %s\n", s.port, err)
		return "", err
	}
	return s.port, nil
}

func readFromUser(reader *bufio.Reader) (string, error) {
	result, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	result = strings.TrimRight(result, "\n")

	return result, nil
}
