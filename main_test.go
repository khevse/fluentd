package main

import (
	"bufio"
	"bytes"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {

	data := `
{"level":"info","ts":1603360107.7447217,"caller":"p1/main.go:203","msg":"finished unary call with code OK","api":"internal","grpc.start_time":"2020-10-22T09:48:27Z","grpc.request.deadline":"2020-10-22T09:48:37Z","system":"grpc","span.kind":"server","grpc.service":"svc.Name1","grpc.method":"Method1","grpc.code":"OK","grpc.time_ms":4.697000026702881}
{"level":"info","ts":1603360112.184296,"caller":"p1/main.go:203","msg":"finished unary call with code OK","api":"internal","grpc.start_time":"2020-10-22T09:48:32Z","grpc.request.deadline":"2020-10-22T09:48:42Z","system":"grpc","span.kind":"server","grpc.service":"svc.Name1","grpc.method":"Method1","grpc.code":"OK","grpc.time_ms":8.45199966430664}
{"level":"info","ts":1603371032.2159479,"caller":"p3/main.go:227","msg":"incoming gRPC request","method":"/svc.Name2/Method2"}
{"level":"debug","ts":1603371032.2186007,"caller":"p1/main.go:144","msg":"invoke method v2","prop1":10,"duration":0.002403934,"count":1}
{"level":"info","ts":1603371032.218727,"caller":"p1/main.go:203","msg":"finished unary call with code OK","api":"internal","grpc.start_time":"2020-10-22T12:50:32Z","grpc.request.deadline":"2020-10-22T12:50:32Z","system":"grpc","span.kind":"server","grpc.service":"svc.Name2","grpc.method":"Method2","grpc.code":"OK","grpc.time_ms":2.6070001125335693}
`

	conn, err := net.Dial("tcp", "localhost:5000")
	require.NoError(t, err)

	defer func() {
		require.NoError(t, conn.Close())
	}()

	r := bytes.NewReader([]byte(data))
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		require.NoError(t, scanner.Err())

		if strings.TrimSpace(line) == "" {
			continue
		}

		_, err := conn.Write([]byte(line + "\n"))
		require.NoError(t, err)
	}
}
