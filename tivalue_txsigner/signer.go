package tivalue_txsigner

import (
	"github.com/blocktree/virtualeconomy-adapter/virtualeconomy_txsigner"
)

var Default = &TransactionSigner{}

type TransactionSigner struct {

}

// SignTransactionHash 交易哈希签名算法
// required
func (singer *TransactionSigner) SignTransactionHash(msg []byte, privateKey []byte, eccType uint32) ([]byte, error) {
	return virtualeconomy_txsigner.Default.SignTransactionHash(msg, privateKey, eccType)
}
