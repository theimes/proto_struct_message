package main

import (
	"context"
	"log/slog"
	"os"

	pb "github.com/theimes/proto_struct_message/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/protobuf/types/known/structpb"

	"log"
)

func main() {

	logLevel := slog.LevelDebug
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}))
	slog.SetDefault(logger)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Println("It is fine, this is not a complete example.")
	}

	defer conn.Close()

	msgClient := pb.NewServiceClient(conn)
	ctx := context.Background()

	val, err := structpb.NewStruct(map[string]interface{}{
		"billId":   "123",
		"amount":   100,
		"currency": "USD",
		"addresses": map[string]interface{}{
			"billing": map[string]interface{}{
				"street": "123 Main St",
				"city":   "Springfield",
				"state":  "IL",
				"zip":    "62701",
			},
			"shipping": map[string]interface{}{
				"street": "123 Main St",
				"city":   "Springfield",
				"state":  "IL",
				"zip":    "62701",
			},
		},
		"customer": map[string]interface{}{
			"firstName": "John",
			"lastName":  "Doe",
		},
	})
	if err != nil {
		slog.Error("Error creating structpb", "error", err)
	}

	a := pb.Message{
		Field: val,
	}

	slog.Debug("a", "a", a)

	resp, err := msgClient.Create(ctx, &a)
	if err != nil {
		slog.Error("Error sending message", "error", err)
	} else {
		dst := structpb.NewStructValue(resp.Field)
		//if err != nil {
		//	slog.Error("could not deserialize anything")
		//}
		slog.Info("Response", "billId", dst)
	}
}
