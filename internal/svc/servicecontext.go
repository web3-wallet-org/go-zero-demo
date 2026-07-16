// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/web3-wallet-org/go-zero-demo/internal/config"
	"github.com/web3-wallet-org/go-zero-demo/internal/store"
)

type ServiceContext struct {
	Config      config.Config
	WalletStore *store.WalletStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		WalletStore: store.NewWalletStore(),
	}
}
