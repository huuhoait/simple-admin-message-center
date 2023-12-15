package smslog

import (
	"context"

	"github.com/huuhoait/simple-admin-message-center/internal/svc"
	"github.com/huuhoait/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/huuhoait/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogByIdLogic {
	return &GetSmsLogByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsLogByIdLogic) GetSmsLogById(in *mcms.UUIDReq) (*mcms.SmsLogInfo, error) {
	result, err := l.svcCtx.DB.SmsLog.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.SmsLogInfo{
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:   pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:   pointy.GetPointer(result.UpdatedAt.Unix()),
		PhoneNumber: &result.PhoneNumber,
		Content:     &result.Content,
		SendStatus:  pointy.GetPointer(uint32(result.SendStatus)),
		Provider:    &result.Provider,
	}, nil
}
