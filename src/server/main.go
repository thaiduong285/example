package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/user/examples_app/src/proto/live_entity"
	proto "github.com/user/examples_app/src/proto/rectangle"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type RectangleServiceServer struct{}
type LiveServiceServer struct{}

var db *mongo.Client
var liveEntity *mongo.Collection
var mongoCtx context.Context

func main() {
	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	// non-nil empty context
	mongoCtx = context.Background()

	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// Handle potential errors
	if err != nil {
		log.Fatal(err)
	}

	// Check whether the connection was succesful by pinging the MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	// Bind our collection to our global variable for use in other methods
	liveEntity = db.Database("examples_app").Collection("live_entity")

	//grcp server
	fmt.Println("localhost:50051")
	flag.Parse()
	lis, err := net.Listen("tcp", ":50051")
	fmt.Println(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	rectangleServer := &RectangleServiceServer{}
	proto.RegisterRectangleServiceServer(grpcServer, rectangleServer)
	liveServiceServer := &LiveServiceServer{}
	live_entity.RegisterLiveServiceServer(grpcServer, liveServiceServer)
	grpcServer.Serve(lis)
}
