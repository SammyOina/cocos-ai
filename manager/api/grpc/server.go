package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/ultravioletrs/manager/manager"
)

type grpcServer struct {
	createDomain kitgrpc.Handler
	run          kitgrpc.Handler
	manager.UnimplementedManagerServiceServer
}

// NewServer returns new AuthServiceServer instance.
func NewServer(svc manager.Service) manager.ManagerServiceServer {
	return &grpcServer{
		createDomain: kitgrpc.NewServer(
			createDomainEndpoint(svc),
			decodeCreateDomainRequest,
			encodeCreateDomainResponse,
		),
		run: kitgrpc.NewServer(
			runEndpoint(svc),
			decodeRunRequest,
			encodeRunResponse,
		),
	}
}

func decodeCreateDomainRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*manager.CreateDomainRequest)
	return createDomainReq{Pool: req.GetPool(), Volume: req.GetVolume(), Domain: req.GetDomain()}, nil
}

func encodeCreateDomainResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(createDomainRes)
	return &manager.CreateDomainResponse{Name: res.Name}, nil
}

func decodeRunRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*manager.RunRequest)
	return runReq{
		Computation: req.GetComputation(),
	}, nil
}

func encodeRunResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(runRes)
	return &manager.RunResponse{
		ID: res.ID,
	}, nil
}

func (s *grpcServer) CreateDomain(ctx context.Context, req *manager.CreateDomainRequest) (*manager.CreateDomainResponse, error) {
	_, res, err := s.createDomain.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	cdr := res.(*manager.CreateDomainResponse)
	return cdr, nil
}

func (s *grpcServer) Run(ctx context.Context, req *manager.RunRequest) (*manager.RunResponse, error) {
	_, res, err := s.run.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rr := res.(*manager.RunResponse)
	return rr, nil
}
