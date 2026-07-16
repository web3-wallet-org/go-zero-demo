// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package wallet

import (
	"context"

	"github.com/web3-wallet-org/go-zero-demo/internal/svc"
	"github.com/web3-wallet-org/go-zero-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferLogic {
	return &TransferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferLogic) Transfer(req *types.TransferReq) (resp *types.TransferResp, err error) {
	from, to, err := l.svcCtx.WalletStore.Transfer(req.From, req.To, req.Amount)
	if err != nil {
		return nil, err
	}

	return &types.TransferResp{
		From: *toWalletResp(from),
		To:   *toWalletResp(to),
	}, nil
}
