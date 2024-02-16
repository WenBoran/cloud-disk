package logic

import (
	"clou-disk/core/helper"
	"clou-disk/models"
	"context"
	"errors"
	"fmt"

	"clou-disk/core/internal/svc"
	"clou-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	//1.从数据中查询当前用户
	user := new(models.UserBasic)
	has, err := models.Engine.Where("name = ? AND password = ?", req.Name, helper.MD5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	fmt.Println(helper.MD5(req.Password))
	if !has {
		return nil, errors.New("用户名或密码错误")
	}
	//2.生成token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Password)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	return
	return
}
