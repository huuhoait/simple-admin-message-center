// Code generated by goctl. DO NOT EDIT.
// Source: mcms.proto

package server

import (
	"context"

	"github.com/suyuan32/simple-admin-message-center/internal/logic/base"
	"github.com/suyuan32/simple-admin-message-center/internal/logic/email"
	"github.com/suyuan32/simple-admin-message-center/internal/logic/emaillog"
	"github.com/suyuan32/simple-admin-message-center/internal/logic/emailprovider"
	"github.com/suyuan32/simple-admin-message-center/internal/logic/sms"
	"github.com/suyuan32/simple-admin-message-center/internal/logic/smslog"
	"github.com/suyuan32/simple-admin-message-center/internal/logic/smsprovider"
	"github.com/suyuan32/simple-admin-message-center/internal/svc"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"
)

type McmsServer struct {
	svcCtx *svc.ServiceContext
	mcms.UnimplementedMcmsServer
}

func NewMcmsServer(svcCtx *svc.ServiceContext) *McmsServer {
	return &McmsServer{
		svcCtx: svcCtx,
	}
}

func (s *McmsServer) InitDatabase(ctx context.Context, in *mcms.Empty) (*mcms.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

func (s *McmsServer) SendEmail(ctx context.Context, in *mcms.EmailInfo) (*mcms.BaseUUIDResp, error) {
	l := email.NewSendEmailLogic(ctx, s.svcCtx)
	return l.SendEmail(in)
}

// EmailLog management
func (s *McmsServer) CreateEmailLog(ctx context.Context, in *mcms.EmailLogInfo) (*mcms.BaseUUIDResp, error) {
	l := emaillog.NewCreateEmailLogLogic(ctx, s.svcCtx)
	return l.CreateEmailLog(in)
}

func (s *McmsServer) UpdateEmailLog(ctx context.Context, in *mcms.EmailLogInfo) (*mcms.BaseResp, error) {
	l := emaillog.NewUpdateEmailLogLogic(ctx, s.svcCtx)
	return l.UpdateEmailLog(in)
}

func (s *McmsServer) GetEmailLogList(ctx context.Context, in *mcms.EmailLogListReq) (*mcms.EmailLogListResp, error) {
	l := emaillog.NewGetEmailLogListLogic(ctx, s.svcCtx)
	return l.GetEmailLogList(in)
}

func (s *McmsServer) GetEmailLogById(ctx context.Context, in *mcms.UUIDReq) (*mcms.EmailLogInfo, error) {
	l := emaillog.NewGetEmailLogByIdLogic(ctx, s.svcCtx)
	return l.GetEmailLogById(in)
}

func (s *McmsServer) DeleteEmailLog(ctx context.Context, in *mcms.UUIDsReq) (*mcms.BaseResp, error) {
	l := emaillog.NewDeleteEmailLogLogic(ctx, s.svcCtx)
	return l.DeleteEmailLog(in)
}

// EmailProvider management
func (s *McmsServer) CreateEmailProvider(ctx context.Context, in *mcms.EmailProviderInfo) (*mcms.BaseIDResp, error) {
	l := emailprovider.NewCreateEmailProviderLogic(ctx, s.svcCtx)
	return l.CreateEmailProvider(in)
}

func (s *McmsServer) UpdateEmailProvider(ctx context.Context, in *mcms.EmailProviderInfo) (*mcms.BaseResp, error) {
	l := emailprovider.NewUpdateEmailProviderLogic(ctx, s.svcCtx)
	return l.UpdateEmailProvider(in)
}

func (s *McmsServer) GetEmailProviderList(ctx context.Context, in *mcms.EmailProviderListReq) (*mcms.EmailProviderListResp, error) {
	l := emailprovider.NewGetEmailProviderListLogic(ctx, s.svcCtx)
	return l.GetEmailProviderList(in)
}

func (s *McmsServer) GetEmailProviderById(ctx context.Context, in *mcms.IDReq) (*mcms.EmailProviderInfo, error) {
	l := emailprovider.NewGetEmailProviderByIdLogic(ctx, s.svcCtx)
	return l.GetEmailProviderById(in)
}

func (s *McmsServer) DeleteEmailProvider(ctx context.Context, in *mcms.IDsReq) (*mcms.BaseResp, error) {
	l := emailprovider.NewDeleteEmailProviderLogic(ctx, s.svcCtx)
	return l.DeleteEmailProvider(in)
}

func (s *McmsServer) SendSms(ctx context.Context, in *mcms.SmsInfo) (*mcms.BaseUUIDResp, error) {
	l := sms.NewSendSmsLogic(ctx, s.svcCtx)
	return l.SendSms(in)
}

// SmsLog management
func (s *McmsServer) CreateSmsLog(ctx context.Context, in *mcms.SmsLogInfo) (*mcms.BaseUUIDResp, error) {
	l := smslog.NewCreateSmsLogLogic(ctx, s.svcCtx)
	return l.CreateSmsLog(in)
}

func (s *McmsServer) UpdateSmsLog(ctx context.Context, in *mcms.SmsLogInfo) (*mcms.BaseResp, error) {
	l := smslog.NewUpdateSmsLogLogic(ctx, s.svcCtx)
	return l.UpdateSmsLog(in)
}

func (s *McmsServer) GetSmsLogList(ctx context.Context, in *mcms.SmsLogListReq) (*mcms.SmsLogListResp, error) {
	l := smslog.NewGetSmsLogListLogic(ctx, s.svcCtx)
	return l.GetSmsLogList(in)
}

func (s *McmsServer) GetSmsLogById(ctx context.Context, in *mcms.UUIDReq) (*mcms.SmsLogInfo, error) {
	l := smslog.NewGetSmsLogByIdLogic(ctx, s.svcCtx)
	return l.GetSmsLogById(in)
}

func (s *McmsServer) DeleteSmsLog(ctx context.Context, in *mcms.UUIDsReq) (*mcms.BaseResp, error) {
	l := smslog.NewDeleteSmsLogLogic(ctx, s.svcCtx)
	return l.DeleteSmsLog(in)
}

// SmsProvider management
func (s *McmsServer) CreateSmsProvider(ctx context.Context, in *mcms.SmsProviderInfo) (*mcms.BaseIDResp, error) {
	l := smsprovider.NewCreateSmsProviderLogic(ctx, s.svcCtx)
	return l.CreateSmsProvider(in)
}

func (s *McmsServer) UpdateSmsProvider(ctx context.Context, in *mcms.SmsProviderInfo) (*mcms.BaseResp, error) {
	l := smsprovider.NewUpdateSmsProviderLogic(ctx, s.svcCtx)
	return l.UpdateSmsProvider(in)
}

func (s *McmsServer) GetSmsProviderList(ctx context.Context, in *mcms.SmsProviderListReq) (*mcms.SmsProviderListResp, error) {
	l := smsprovider.NewGetSmsProviderListLogic(ctx, s.svcCtx)
	return l.GetSmsProviderList(in)
}

func (s *McmsServer) GetSmsProviderById(ctx context.Context, in *mcms.IDReq) (*mcms.SmsProviderInfo, error) {
	l := smsprovider.NewGetSmsProviderByIdLogic(ctx, s.svcCtx)
	return l.GetSmsProviderById(in)
}

func (s *McmsServer) DeleteSmsProvider(ctx context.Context, in *mcms.IDsReq) (*mcms.BaseResp, error) {
	l := smsprovider.NewDeleteSmsProviderLogic(ctx, s.svcCtx)
	return l.DeleteSmsProvider(in)
}
