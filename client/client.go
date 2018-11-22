package main

import (
	"context"
	pb "grpcDemo/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func client() {
	//创建一条连接链接
	conn, err := grpc.Dial("localhost:8848", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//创建一天grpc的客户端连接
	client := pb.NewCalculate_Client(conn)
	//创建超时上下文，超时处理时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	returnMsg, err := client.NumberMul(ctx, &pb.CalculateInt{A: &pb.Number_{Num: 20}, B: &pb.Number_{Num: 3}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("addValue: %d", (*returnMsg).Num)
}

func main() {
	client()
}
