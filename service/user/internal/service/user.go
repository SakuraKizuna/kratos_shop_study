package service

import (
	"context"
	"go.opentelemetry.io/otel"
	v1 "user/api/user/v1"
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// CreateUser create a user
func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserInfo) (*v1.UserInfoResponse, error) {
	user, err := u.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := UserResponse(user)
	return &userInfoRsp, nil
}

func (u *UserService) GetUserByMobile(ctx context.Context, req *v1.MobileRequest) (*v1.UserInfoResponse, error) {
	tr := otel.Tracer("user service")
	ctx, span := tr.Start(ctx, "GetUserByMobile service")
	defer span.End()
	user, err := u.uc.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	rsp := UserResponse(user)
	return &rsp, nil
}

// CheckPassword .
func (u *UserService) CheckPassword(ctx context.Context, req *v1.PasswordCheckInfo) (*v1.CheckResponse, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "check user  password")
	defer span.End()
	check, err := u.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &v1.CheckResponse{Success: check}, nil
}

func UserResponse(user *biz.User) v1.UserInfoResponse {
	userInfoRsp := v1.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}
