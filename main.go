package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcddd/myservice/myservice/pbfile"
	"log"
	"net/http"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:4567", "")
)

func main() {
	flag.Parse()
	mux := runtime.NewServeMux()
	mux.HandlePath("GET", "/ping", pingHandler)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pbfile.RegisterMyServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "{\"msg\":\"pong\"}")
}
