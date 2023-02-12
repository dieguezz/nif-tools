package main

import (
	"context"
	"log"
	"net"
	"net/http"

	nifparser "github.com/dieguezz/nif-tools/parser"

	pb "github.com/dieguezz/nif-tools/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNifApiServer
}

func (s *server) GetControlDigit(ctx context.Context, in *pb.ControlDigitRequest) (*pb.ControlDigitResponse, error) {
	nif, err := nifparser.GetParsedNIF(in.GetNif())
	return &pb.ControlDigitResponse{ControlDigit: nif.Control}, err
}

func (s *server) GetType(ctx context.Context, in *pb.TypeRequest) (*pb.TypeResponse, error) {
	nif, err := nifparser.GetParsedNIF(in.GetNif())
	return &pb.TypeResponse{Type: nif.Kind}, err
}

func (s *server) GetParsedNIF(ctx context.Context, in *pb.ParsedNIFRequest) (*pb.ParsedNIFResponse, error) {
	nif, err := nifparser.GetParsedNIF(in.GetNif())
	return &pb.ParsedNIFResponse{Number: int32(nif.Number), Kind: nif.Kind, Control: nif.Control}, err
}

func main() {
	// Grpc server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	go func() {
		mux := runtime.NewServeMux()
		pb.RegisterNifApiHandlerServer(context.Background(), mux, &server{})
		// Rest server
		log.Fatalln(http.ListenAndServe("localhost:8080", mux))
	}()

	pb.RegisterNifApiServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
