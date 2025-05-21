package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexander777hub/parserclient-go/ps/parser"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <file.pdf>")
	}
	filePath := os.Args[1]

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*50)),
	)
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC-серверу: %v", err)
	}
	defer conn.Close()

	client := parser.NewParserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.ParseTable(ctx, &parser.ParseRequest{FilePath: filePath})
	if err != nil {
		log.Fatalf("Ошибка при вызове ParseTable: %v", err)
	}

	for i, row := range resp.Rows {
		fmt.Printf("Row %d: %v\n", i+1, row.Cells)
	}
}
