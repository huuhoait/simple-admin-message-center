package smslog

import (
	"context"

	"github.com/huuhoait/simple-admin-message-center/internal/svc"
	"github.com/huuhoait/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/huuhoait/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSmsLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSmsLogLogic {
	return &UpdateSmsLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSmsLogLogic) UpdateSmsLog(in *mcms.SmsLogInfo) (*mcms.BaseResp, error) {
	query := l.svcCtx.DB.SmsLog.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilPhoneNumber(in.PhoneNumber).
		SetNotNilContent(in.Content).
		SetNotNilProvider(in.Provider)

	if in.SendStatus != nil {
		query.SetNotNilSendStatus(pointy.GetPointer(uint8(*in.SendStatus)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
