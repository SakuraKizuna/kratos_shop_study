package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	jwt2 "github.com/golang-jwt/jwt/v4"
	v1 "shop/api/shop/v1"
	"shop/internal/conf"
	"shop/internal/pkg/captcha"
	"shop/internal/pkg/middleware/auth"
	"time"
)

var (
	ErrPasswordInvalid     = errors.New("password invalid")
	ErrUsernameInvalid     = errors.New("username invalid")
	ErrCaptchaInvalid      = errors.New("verification code error")
	ErrMobileInvalid       = errors.New("mobile invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrLoginFailed         = errors.New("login failed")
	ErrGenerateTokenFailed = errors.New("generate token failed")
	ErrAuthFailed          = errors.New("authentication failed")
)

type User struct {
	ID        int64
	Mobile    string
	Password  string
	NickName  string
	Birthday  int64
	Gender    string
	Role      int
	CreatedAt time.Time
}

type UserRepo interface {
	CreateUser(c context.Context, u *User) (*User, error)
	//UserByMobile(ctx context.Context, mobile string) (*User, error)
	//UserById(ctx context.Context, Id int64) (*User, error)
	//CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
	//ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error)
	//UpdateUser(context.Context, *User) (bool, error)
}

type UserUsecase struct {
	uRepo      UserRepo
	log        *log.Helper
	signingKey string
}

func NewUserUsecase(repo UserRepo, logger log.Logger, conf *conf.Auth) *UserUsecase {
	helper := log.NewHelper(log.With(logger, "module", "usecase/shop"))
	return &UserUsecase{uRepo: repo, log: helper, signingKey: conf.JwtKey}
}

// GetCaptcha 验证码
func (uc *UserUsecase) GetCaptcha(ctx context.Context) (*v1.CaptchaReply, error) {
	captchaInfo, err := captcha.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.CaptchaReply{
		CaptchaId: captchaInfo.CaptchaId,
		PicPath:   captchaInfo.PicPath,
	}, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	newUser, err := NewUser(req.Mobile, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	createUser, err := uc.uRepo.CreateUser(ctx, &newUser)
	claims := auth.CustomClaims{
		ID:          createUser.ID,
		NickName:    createUser.NickName,
		AuthorityId: createUser.Role,
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix(),               // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
			Issuer:    "Gyl",
		},
	}
	token, err := auth.CreateToken(claims, uc.signingKey)
	if err != nil {
		return nil, err
	}

	return &v1.RegisterReply{
		Id:        createUser.ID,
		Mobile:    createUser.Mobile,
		Username:  createUser.NickName,
		Token:     token,
		ExpiredAt: time.Now().Unix() + 60*60*24*30,
	}, nil
}

func NewUser(mobile, username, password string) (User, error) {
	// check mobile
	if len(mobile) <= 0 {
		return User{}, ErrMobileInvalid
	}
	// check username
	if len(username) <= 0 {
		return User{}, ErrUsernameInvalid
	}
	// check password
	if len(password) <= 0 {
		return User{}, ErrPasswordInvalid
	}
	return User{
		Mobile:   mobile,
		NickName: username,
		Password: password,
	}, nil
}
