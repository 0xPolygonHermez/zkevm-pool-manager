package server

import (
	"context"

	"github.com/0xPolygonHermez/zkevm-pool-manager/db"
)

type poolDBInterface interface {
	AddL2Transaction(ctx context.Context, tx *db.L2Transaction) error
	UpdateL2TransactionStatus(ctx context.Context, txHash string, newStatus string, errorMsg string) error
}

type senderInterface interface {
	SendL2Transaction(l2Tx *db.L2Transaction) error
}
