// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package wallet

import (
	"net/http"

	"github.com/web3-wallet-org/go-zero-demo/internal/logic/wallet"
	"github.com/web3-wallet-org/go-zero-demo/internal/svc"
	"github.com/web3-wallet-org/go-zero-demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateWalletHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateWalletReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := wallet.NewCreateWalletLogic(r.Context(), svcCtx)
		resp, err := l.CreateWallet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
