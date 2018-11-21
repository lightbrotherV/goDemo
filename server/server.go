package main

import (
	"context"
	pb "grpcDemo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) NumberAdd(ctx context.Context, in *pb.CalculateInt) (*pb.Number_, error) {
	res := ((*((*in).GetA())).Num) + ((*((*in).GetB())).Num)

	return &pb.Number_{Num: res}, nil
}

func (*server) NumberMul(ctx context.Context, in *pb.CalculateInt) (*pb.Number_, error) {
	res := ((*((*in).GetA())).Num) - ((*((*in).GetB())).Num)

	return &pb.Number_{Num: res}, nil
}

func (*server) StringAdd(ctx context.Context, in *pb.CalculateString) (*pb.String_, error) {
	res := ((*((*in).GetA())).Str) + ((*((*in).GetB())).Str)

	return &pb.String_{Str: res}, nil
}

func main() {
	//创建tcp监听进程，用于接受客户端连接
	lis, err := net.Listen("tcp", "5198")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//创建grpc服务器
	grpcServer := grpc.NewServer()
	pb.RegisterCalculate_Server(grpcServer, &server{})
	// 反射grpc服务，在测试是有用，此处可注释
	reflection.Register(grpcServer)

}
