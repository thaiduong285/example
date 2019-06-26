package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/user/examples_app/src/proto/live_entity"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	proto "github.com/user/examples_app/src/proto/rectangle"
	"google.golang.org/grpc"
)

var (
	endpoint = flag.String("entity_endpoint", "localhost:50051", "endpoint of EntityService")
)

func RunEndPoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println("Starting server on port :8080.")
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterRectangleServiceHandlerFromEndpoint(ctx, mux, *endpoint, dialOpts)
	if err != nil {
		return err
	}

	liveError := live_entity.RegisterLiveServiceHandlerFromEndpoint(ctx, mux, *endpoint, dialOpts)

	if liveError != nil {
		return liveError
	}

	http.ListenAndServe(address, mux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := RunEndPoint(":8080"); err != nil {
		glog.Fatal(err)
	}
}
