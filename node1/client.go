package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Cluster_Resource_Table/node1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStatsServiceClient(conn)

	for {
		req := &pb.StatsRequest{}
		res, err := client.GetStats(context.Background(), req)
		if err != nil {
			log.Fatalf("could not get stats: %v", err)
		}

		fmt.Println("节点: ", res.NodeId)
		fmt.Printf("CPU 使用率: %.2f%%\n", res.GetCpuPercent())
		fmt.Printf("内存 使用率: %.2f%%\n", res.GetMemPercent())

		time.Sleep(1 * time.Second)
	}
}
