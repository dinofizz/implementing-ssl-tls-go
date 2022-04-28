package main

import (
	"errors"
	"fmt"
	"github.com/dinofizz/impl-tsl-go/pkg/common"
	"io"
	"net"
	"os"
	"strings"
)

var HTTP_PORT = "80"

type DestinationParams struct {
	host string
	port string
	path string
}

type ProxyParams struct {
	host     string
	port     string
	username string
	password string
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage %s: [-p http://[username:password@]proxy-host:proxy-port] <URL>\n", args[0])
		os.Exit(1)
	}

	i := 1

	var proxyParams *ProxyParams

	if args[i] == "-p" {
		arg := args[i+1]
		var err error
		proxyParams, err = parseProxyParams(arg)
		if err != nil {
			fmt.Printf("Error - malformed proxy parameter: %s\n", arg)
			os.Exit(2)
		}
		i = i + 2
	}

	fmt.Println(proxyParams)

	destParams, err := parseUrl(args[i])
	if err != nil {
		fmt.Printf("Error - malformed URL '%s'.\n", args[1])
		os.Exit(1)
	}

	var tcpHost, tcpPort string
	if proxyParams != nil {
		tcpHost = proxyParams.host
		tcpPort = proxyParams.port
	} else {
		tcpHost = destParams.host
		tcpPort = destParams.port
	}

	hostName := fmt.Sprintf("%s:%s", tcpHost, tcpPort)

	fmt.Printf("Connecting to host '%s'\n", hostName)

	tcpAddr, err := net.ResolveTCPAddr("tcp", hostName)
	if err != nil {
		fmt.Printf("Error in name resolution: %s\n", err)
		os.Exit(3)
	}

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("Unable to connect to host: %s\n", err)
		os.Exit(4)
	}
	fmt.Println(tcpConn)

	fmt.Printf("Retrieving document: %s\n", destParams.path)

	err = httpGet(tcpConn, destParams.path, destParams.host, proxyParams)
	if err != nil {
		fmt.Printf("Error writing GET command: %s\n", err)
		os.Exit(5)
	}

	displayResult(tcpConn)

	err = tcpConn.Close()
	if err != nil {
		fmt.Printf("Error closing client connection: %s\n", err)
		os.Exit(5)
	}
}

func parseProxyParams(proxySpec string) (*ProxyParams, error) {
	proxySpec = strings.TrimPrefix(proxySpec, "http://")

	var proxyParams ProxyParams

	var colonSep int
	loginSep := strings.Index(proxySpec, "@")

	if loginSep != -1 {
		colonSep = strings.Index(proxySpec, ":")
		if colonSep == -1 || colonSep > loginSep {
			fmt.Printf("Expected password in %s\n", proxySpec)
			return nil, errors.New("Malformed proxy argument.")
		}
		proxyParams.username = proxySpec[:colonSep]
		proxyParams.password = proxySpec[colonSep+1 : loginSep]
		proxySpec = proxySpec[loginSep+1:]
	}

	proxySpec = strings.TrimSuffix(proxySpec, "/")
	colonSep = strings.Index(proxySpec, ":")
	if colonSep != -1 {
		proxyParams.host = proxySpec[:colonSep]
		proxyParams.port = proxySpec[colonSep+1:]
	} else {
		proxyParams.port = HTTP_PORT
		proxyParams.host = proxySpec
	}

	return &proxyParams, nil
}

func displayResult(tcpConn *net.TCPConn) {
	for {
		recvBuf := make([]byte, 255)
		n, err := tcpConn.Read(recvBuf)
		if err != nil && err == io.EOF{
			break
		}
		if err != nil && err != io.EOF{
			fmt.Printf("Error reading from TCP connection: %s", err)
			os.Exit(5)
		}
		if n == 0 {
			break
		}

		fmt.Print(string(recvBuf))
	}
	fmt.Println()
}

func httpGet(tcpConn *net.TCPConn, path, host string, proxyParams *ProxyParams) error {
	var getCommand string
	if proxyParams != nil {
		getCommand = fmt.Sprintf("GET http://%s/%s HTTP/1.1\r\n", host, path)
	} else {
		getCommand = fmt.Sprintf("GET /%s HTTP/1.1\r\n", path)
	}

	_, err := tcpConn.Write([]byte(getCommand))
	if err != nil {
		return err
	}

	getCommand = fmt.Sprintf("Host: %s\r\n", host)
	_, err = tcpConn.Write([]byte(getCommand))
	if err != nil {
		return err
	}

	if proxyParams != nil && proxyParams.username != "" {
		plaintextCredentials := fmt.Sprintf("%s:%s", proxyParams.username, proxyParams.host)
		inBuf := []byte(plaintextCredentials)
		outBuf := common.Base64Encode(&inBuf)
		getCommand = fmt.Sprintf("Proxy-Authorization: BASIC %s\r\n", string(*outBuf))
		_, err = tcpConn.Write([]byte(getCommand))
		if err != nil {
			return err
		}
	}

	getCommand = "Connection: close\r\n\r\n"
	_, err = tcpConn.Write([]byte(getCommand))
	if err != nil {
		return err
	}

	return nil
}

func parseUrl(uri string) (*DestinationParams, error) {
	i := strings.Index(uri, "//")
	destParams := new(DestinationParams)
	if i == -1 {
		return destParams, errors.New("Malformed URL")
	}

	uri = uri[i+2:]

	j := strings.Index(uri, "/")
	if j == -1 {
		destParams.host = uri
		destParams.path = ""
	} else {
		destParams.host = uri[:j]
		destParams.path = uri[j+1:]
	}

	k := strings.Index(destParams.host, ":")
	if k != -1 {
		destParams.port = destParams.host[k+1:]
		destParams.host = destParams.host[:k]
	}

	return destParams, nil
}
