package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPortalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPortalLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserPortalLogic {
	return GetUserPortalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPortalLogic) GetUserPortal(req types.GetUserPortalReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
