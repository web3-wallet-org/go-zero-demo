// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package wallet

import (
	"context"

	"github.com/web3-wallet-org/go-zero-demo/internal/svc"
	"github.com/web3-wallet-org/go-zero-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWalletLogic {
	return &CreateWalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWalletLogic) CreateWallet(req *types.CreateWalletReq) (resp *types.WalletResp, err error) {
	wallet, err := l.svcCtx.WalletStore.Create(req.Address, req.Balance)
	if err != nil {
		return nil, err
	}

	return toWalletResp(wallet), nil
}
