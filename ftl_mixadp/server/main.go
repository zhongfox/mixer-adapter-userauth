package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"fmt"
	"net"
	"strings"
	"userauth/ftl_mixadp"
)

type server struct{}

func (s *server) Verify(ctx context.Context, req *ftl_mixadp.VerifyReq) (*ftl_mixadp.VerifyRsp, error) {
	// Read metadata from client.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "Verify: failed to get metadata")
	}
	if c, ok := md["cookie"]; ok {
		fmt.Sprintf("cookie from metadata: %s\n", c)
		if len(c) == 1 && strings.HasPrefix(c[0], "user=fox") {
			return &ftl_mixadp.VerifyRsp{ErrCode: "ok"}, nil
		}
		return &ftl_mixadp.VerifyRsp{ErrCode: "sig_err"}, nil
	}
	return &ftl_mixadp.VerifyRsp{ErrCode: "no-sig"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()
	ftl_mixadp.RegisterSigServer(s, &server{})
	s.Serve(lis)
}