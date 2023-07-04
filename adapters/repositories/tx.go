package repositories

import (
	"context"

	"github.com/daichi1002/go-graphql/adapters"
	"github.com/daichi1002/go-graphql/constants"
)

type TxRepositoryDependencies struct {
	sqlHandler adapters.SqlHandler
}

func NewTxRepository(sqlHandler adapters.SqlHandler) TxRepository {
	return &TxRepositoryDependencies{
		sqlHandler,
	}
}

func (dep *TxRepositoryDependencies) DoInTx(f func(ctx context.Context) error) error {
	tx, err := dep.sqlHandler.Begin()

	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), constants.TxCtxKey, tx)

	if err := f(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil
}
