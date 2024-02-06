package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"log/slog"

	"google.golang.org/grpc"

	pb "github.com/theimes/proto_struct_message/message"

	structpb "google.golang.org/protobuf/types/known/structpb"

	"log"
	"net"
)

type server struct {
	pb.UnimplementedServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	logLevel := slog.LevelInfo
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
	}))
	slog.SetDefault(logger)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterServiceServer(grpcServer, &server{})

	slog.Info("Server started", "port", *port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) Create(ctx context.Context, in *pb.Message) (*pb.Message, error) {

	slog.Debug("Received", "message", in)

	dst := structpb.NewStructValue(in.Field)

	structValue := dst.GetStructValue()

	for k, v := range structValue.Fields {
		slog.Debug("loop", "key", k, "kind", v.Kind)
		slog.Debug("loop", "key", k, "interface", v.AsInterface())
	}

	addresses := structValue.Fields["addresses"].GetStructValue().Fields
	for k, v := range addresses {
		slog.Debug("loop", "key", k, "interface", v)
	}

	bill := pb.Invoice{
		BillId:   structValue.Fields["billId"].GetStringValue(),
		Amount:   structValue.Fields["amount"].GetNumberValue(),
		Currency: structValue.Fields["currency"].GetStringValue(),
		Addresses: map[string]pb.Address{
			"billing": {
				Street: addresses["billing"].GetStructValue().Fields["street"].GetStringValue(),
				City:   addresses["billing"].GetStructValue().Fields["city"].GetStringValue(),
				State:  addresses["billing"].GetStructValue().Fields["state"].GetStringValue(),
				Zip:    addresses["billing"].GetStructValue().Fields["zip"].GetStringValue(),
			},
			"shipping": {
				Street: addresses["shipping"].GetStructValue().Fields["street"].GetStringValue(),
				City:   addresses["shipping"].GetStructValue().Fields["city"].GetStringValue(),
				State:  addresses["shipping"].GetStructValue().Fields["state"].GetStringValue(),
				Zip:    addresses["shipping"].GetStructValue().Fields["zip"].GetStringValue(),
			},
		},
	}

	slog.Info("received", "bill", bill, "type", fmt.Sprintf("%T", bill))

	var gen pb.Generic
	gen = make(map[string]interface{})

	for k, v := range structValue.Fields {
		gen[k] = v.AsInterface()
	}

	slog.Info("received", "gen", gen, "type", fmt.Sprintf("%T", gen))

	return in, nil
}
