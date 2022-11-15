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
	// TODO: Implement me
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
