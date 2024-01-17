package dynamo

import (
	"context"
)

func defaultContext() (context.Context, context.CancelFunc) {
	return context.Background(), (func() {})
}

func (db *DB) retry(ctx context.Context, f func() error) error {
	return f()
}
