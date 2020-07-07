package client

import (
	pb "entry_task/src/proto"
	"entry_task/src/tools"
	"fmt"
	log "github.com/cihub/seelog"
	"google.golang.org/grpc"
	"os"
	"time"
)

var client *pb.UserServiceClient
var GrpcPool *tools.Pool

func InitClient(file string) {
	// Set up src connection to the server.
	fmt.Printf(os.Getwd())
	tools.InitHttpConfig(file)
	address := tools.GetHttpConfig().Rpc.Address
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error("did not connect: %v", err)
	}
	//defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	client = &c

	var factory tools.Factory
	factory = func() (*grpc.ClientConn, error) {

		conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
		if err != nil {

		}
		return conn, err
	}

	p, e := tools.New(factory, 10, 100, time.Second*100, time.Second*100)
	if e != nil {

	}
	GrpcPool = p
}

func GetGRClient() pb.UserServiceClient {

	return *client;
}
