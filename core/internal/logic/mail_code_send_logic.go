package logic

import (
	"clou-disk/core/define"
	"clou-disk/core/helper"
	"clou-disk/core/internal/svc"
	"clou-disk/core/internal/types"
	"clou-disk/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendLogic {
	return &MailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendLogic) MailCodeSend(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {

	count, err := models.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("邮箱已被注册")
	}
	//未注册
	resp = new(types.MailCodeSendReply)
	code := helper.RandCode()
	models.RDB.Set(l.ctx, req.Email, code, define.CodeExpire)
	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
