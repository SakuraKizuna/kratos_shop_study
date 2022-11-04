package service

import (
	"context"

	pb "goods/api/goods/v1"
)

type GoodsService struct {
	pb.UnimplementedGoodsServer
}

func NewGoodsService() *GoodsService {
	return &GoodsService{}
}

func (s *GoodsService) CreateGoods(ctx context.Context, req *pb.CreateGoodsRequest) (*pb.CreateGoodsReply, error) {
	return &pb.CreateGoodsReply{}, nil
}
func (s *GoodsService) UpdateGoods(ctx context.Context, req *pb.UpdateGoodsRequest) (*pb.UpdateGoodsReply, error) {
	return &pb.UpdateGoodsReply{}, nil
}
func (s *GoodsService) DeleteGoods(ctx context.Context, req *pb.DeleteGoodsRequest) (*pb.DeleteGoodsReply, error) {
	return &pb.DeleteGoodsReply{}, nil
}
func (s *GoodsService) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsReply, error) {
	return &pb.GetGoodsReply{}, nil
}
func (s *GoodsService) ListGoods(ctx context.Context, req *pb.ListGoodsRequest) (*pb.ListGoodsReply, error) {
	return &pb.ListGoodsReply{}, nil
}
