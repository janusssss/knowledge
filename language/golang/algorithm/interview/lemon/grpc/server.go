package main

import (
	"context"
	"flag"
	"fmt"
	"go-interview/arithmetic/lemon/pd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			AuthInterceptor,
			LogInterceptor,
		),
	)
	pd.RegisterLoginServiceServer(s, &server{})

	// 启动服务器
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var (
	port = flag.Int("port", 50051, "服务端口")
)

// server 是实现的LoginServiceServer接口
type server struct {
	pd.UnimplementedLoginServiceServer
}

// Login 实现LoginServiceServer接口的Login方法
func (s *server) Login(ctx context.Context, req *pd.LoginRequest) (*pd.LoginResponse, error) {
	response := &pd.LoginResponse{}
	if req.Username == "Janus" {
		response.Result = &pd.LoginResponse_Account{Account: "Janus"}
	} else {
		response.Result = &pd.LoginResponse_ErrorMessage{ErrorMessage: "用户不存在"}
	}
	return response, nil
}

// AuthInterceptor 是一个鉴权拦截器
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "metadata is missing")
	}
	token := md.Get("Authorization")
	if token == nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	return handler(ctx, req)
}

// LogInterceptor 是一个日志记录器
func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("[ERROR] %s took %v and returned error: %v", info.FullMethod, time.Since(startTime), err)
		return resp, err
	}
	log.Printf("[INFO] %s took %v", info.FullMethod, time.Since(startTime))
	return resp, err
}
