// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package wallet

import (
	"context"

	"github.com/web3-wallet-org/go-zero-demo/internal/svc"
	"github.com/web3-wallet-org/go-zero-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWalletLogic {
	return &GetWalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWalletLogic) GetWallet(req *types.GetWalletReq) (resp *types.WalletResp, err error) {
	wallet, err := l.svcCtx.WalletStore.Get(req.Address)
	if err != nil {
		return nil, err
	}

	return toWalletResp(wallet), nil
}
