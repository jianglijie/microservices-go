package main

import (
	"config"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"internal/finance"
	"localdb"
	"log"
	"net"
	pb "services/proto"
	"strconv"
)

const (
	Port          = ":50051"
)

type financeServer struct{}
type testServer struct{}

func (s *financeServer) GetCoinFinance(ctx context.Context, in *pb.CoinFinanceRequest) (*pb.CoinFinanceReply, error) {
	fmt.Println(config.Config())
	data := finance.GetCoinFinance(in.Coin, in.StartTime, in.EndTime)
	var retList []*pb.FinanceItem
	for _, item := range data {
		smallIn, _ := strconv.ParseFloat(item["small_in"], 64)
		middleIn, _ := strconv.ParseFloat(item["middle_in"], 64)
		bigIn, _ := strconv.ParseFloat(item["big_in"], 64)
		superIn, _ := strconv.ParseFloat(item["super_in"], 64)
		smallOut, _ := strconv.ParseFloat(item["small_out"], 64)
		middleOut, _ := strconv.ParseFloat(item["middle_out"], 64)
		bigOut, _ := strconv.ParseFloat(item["big_out"], 64)
		superOut, _ := strconv.ParseFloat(item["super_out"], 64)
		financeDate, _ := strconv.Atoi(item["finance_date"])
		updated, _ := strconv.Atoi(item["updated"])
		retList = append(retList, &pb.FinanceItem{CoinKey: item["coin_key"], SmallIn: smallIn, MiddleIn: middleIn,
			BigIn: bigIn, SuperIn: superIn, SmallOut: smallOut, MiddleOut: middleOut, BigOut: bigOut, SuperOut: superOut, FinanceDate: uint32(financeDate), UpdateTime: uint32(updated)})
	}
	return &pb.CoinFinanceReply{FinanceList: retList}, nil
}

func (s *testServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) { // 服务端实现 proto 中定义的方法
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil //拼接客户端发送过来的消息，并返回给客户端
}


func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer localdb.Mysql.GetInstance().Close()

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &testServer{})
	pb.RegisterFinanceServer(s, &financeServer{})
	s.Serve(lis)
}
