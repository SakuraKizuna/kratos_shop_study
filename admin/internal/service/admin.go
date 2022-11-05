package service

import (
	"context"

	pb "admin/api/admin/v1"
)

type AdminService struct {
	pb.UnimplementedAdminServer
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (s *AdminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.CreateAdminReply, error) {
	return &pb.CreateAdminReply{}, nil
}
func (s *AdminService) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminRequest) (*pb.UpdateAdminReply, error) {
	return &pb.UpdateAdminReply{}, nil
}
func (s *AdminService) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminRequest) (*pb.DeleteAdminReply, error) {
	return &pb.DeleteAdminReply{}, nil
}
func (s *AdminService) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.GetAdminReply, error) {
	return &pb.GetAdminReply{}, nil
}
func (s *AdminService) ListAdmin(ctx context.Context, req *pb.ListAdminRequest) (*pb.ListAdminReply, error) {
	return &pb.ListAdminReply{}, nil
}
