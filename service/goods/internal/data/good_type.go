package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
	"gorm.io/gorm"
	"time"
)

// GoodsType 商品类型表
type GoodsType struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;comment:商品类型名称" json:"name"`
	TypeCode  string         `gorm:"type:varchar(50);not null;comment:商品类型编码" json:"type_code"`
	NameAlias string         `gorm:"type:varchar(50);not null;comment:商品类型别名" json:"name_alias"`
	IsVirtual bool           `gorm:"comment:是否是虚拟商品显示;default:false" json:"is_virtual"`
	Desc      string         `gorm:"type:varchar(50);not null;comment:商品类型描述" json:"desc"`
	Sort      int32          `gorm:"comment:类型排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type goodsTypeRepo struct {
	data *Data
	log  *log.Helper
}

// NewGoodsTypeRepo .
func NewGoodsTypeRepo(data *Data, logger log.Logger) biz.GoodsTypeRepo {
	return &goodsTypeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateGoodsType 创建基本的商品类型
func (g *goodsTypeRepo) CreateGoodsType(ctx context.Context, req *domain.GoodsType) (int64, error) {
	goodsType := &GoodsType{
		Name:      req.Name,
		TypeCode:  req.TypeCode,
		NameAlias: req.NameAlias,
		IsVirtual: req.IsVirtual,
		Desc:      req.Desc,
		Sort:      req.Sort,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	result := g.data.db.Save(goodsType)
	return goodsType.ID, result.Error
}
