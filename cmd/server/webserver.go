package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage %s: <host>:<port>\n", args[0])
		os.Exit(1)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", args[1])
	if err != nil {
		fmt.Printf("Error resolving TCP addrress: %s\n", err)
		os.Exit(0)
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("Error creating TCP listener: %s\n", err)
		os.Exit(0)
	}

	defer l.Close()

	for {
		tcpConn, err := l.AcceptTCP()
		if err != nil {
			fmt.Printf("Error accepting TCP connection: %s\n", err)
			os.Exit(0)
		}

		go handleTCPConnection(tcpConn)
	}
}

func handleTCPConnection(tcpConn *net.TCPConn) {
	defer tcpConn.Close()
	requestLine := readLine(tcpConn)
	if !strings.HasPrefix(requestLine, "GET") {
		buildErrorResponse(tcpConn, http.StatusNotImplemented)
	}

	for readLine(tcpConn) != "" {
	}

	buildSuccessResponse(tcpConn)
}

func buildSuccessResponse(tcpConn *net.TCPConn) {
	var responseBuf bytes.Buffer

	responseBuf.WriteString("HTTP/1.1 200 Success\r\n")
	responseBuf.WriteString("Connection: Close\r\n")
	responseBuf.WriteString("Content-Type: text/html\r\n\r\n")
	responseBuf.WriteString("<html><head><title>Test Page</title></head><body>Nothing here</body></html>\r\n")

	_, err := tcpConn.Write(responseBuf.Bytes())
	if err != nil {
		fmt.Println("Unable to respond to TCP request.")
	}
}

func buildErrorResponse(tcpConn *net.TCPConn, status int) {
	var responseBuf bytes.Buffer

	responseBuf.WriteString(fmt.Sprintf("HTTP/1.1 %d Error Occurred\r\n\r\n", status))

	_, err := tcpConn.Write(responseBuf.Bytes())
	if err != nil {
		fmt.Println("Unable to respond to TCP request.")
	}

}

func readLine(tcpConn *net.TCPConn) string {
	var line string
	var lineBuf bytes.Buffer
	recvBuf := make([]byte, 1)
	var pos int
	for {
		n, err := tcpConn.Read(recvBuf)
		if (err != nil && err == io.EOF) || n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading from TCP connection: %s", err)
			break
		}

		if string(recvBuf[0]) == "\n" && string(lineBuf.Bytes()[pos-1]) == "\r" {
			line = string(lineBuf.Bytes()[:pos-1])
			break
		}

		pos++
		lineBuf.WriteByte(recvBuf[0])
	}

	return line
}
