package wallet

import (
	"github.com/web3-wallet-org/go-zero-demo/internal/store"
	"github.com/web3-wallet-org/go-zero-demo/internal/types"
)

func toWalletResp(wallet store.Wallet) *types.WalletResp {
	return &types.WalletResp{
		Address: wallet.Address,
		Balance: wallet.Balance,
	}
}
