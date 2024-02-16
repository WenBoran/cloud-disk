package logic

import (
	"clou-disk/models"
	"context"
	"errors"

	"clou-disk/core/internal/svc"
	"clou-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	resp = new(types.UserDetailReply)
	uc := new(models.UserBasic)
	has, err := models.Engine.Where("identity = ?", req.Identity).Get(uc)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("没有该用户")
	}
	resp.Name = uc.Name
	resp.Email = uc.Email
	return
}
