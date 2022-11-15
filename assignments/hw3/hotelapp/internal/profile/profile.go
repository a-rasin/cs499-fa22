package profile

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	pb "github.com/ucy-coast/hotel-app/internal/profile/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// Profile implements the profile service
type Profile struct {
	port      int
	addr      string
	dbsession *DatabaseSession
	tracer    opentracing.Tracer
}

// NewProfile returns a new Profile service
func NewProfile(a string, p int, db *DatabaseSession, tr opentracing.Tracer) *Profile {
	return &Profile{
		addr:      a,
		port:      p,
		dbsession: db,
		tracer:    tr,
	}
}

// Run starts the server
func (s *Profile) Run() error {
	// DONE: Implement me
	if s.port == 0 {
		return fmt.Errorf("server port must be set")
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: 120 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.tracer),
		),
	}

	srv := grpc.NewServer(opts...)
	// pb.RegisterSearchServer(srv, s)

	// Register reflection service on gRPC server.
	reflection.Register(srv)

	// listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Start Profile server. Addr: %s:%d\n", s.addr, s.port)
	return srv.Serve(lis)
}

// GetProfiles returns hotel profiles for requested IDs
func (s *Profile) GetProfiles(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	// DONE: Implement me
	out := new(pb.Result)
	err := grpc.Invoke(ctx, "/profile.Memdb/GetProfiles", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
