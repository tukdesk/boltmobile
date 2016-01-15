package boltmobile

import (
	"github.com/boltdb/bolt"
)

func NewBoltDB(path string) (boltDB *BoltDB, err error) {
	db, err := bolt.Open(path, 0644, nil) // 非 0644 在 ios simulator 上报错
	if err != nil {
		return nil, err
	}

	return &BoltDB{
		db: db,
	}, nil
}

type BoltTxHandler interface {
	Handle(boltTx *BoltTx) error
}

type BoltDB struct {
	db *bolt.DB
}

func (this *BoltDB) Path() string {
	return this.db.Path()
}

func (this *BoltDB) Close() error {
	return this.db.Close()
}

func (this *BoltDB) Begin(writable bool) (boltTx *BoltTx, err error) {
	tx, err := this.db.Begin(writable)

	if err != nil {
		return nil, err
	}

	return newBoltTx(this, tx), nil
}

func (this *BoltDB) Update(handler BoltTxHandler) error {
	return this.db.Update(this.genarateTxFn(handler))
}

func (this *BoltDB) View(handler BoltTxHandler) error {
	return this.db.View(this.genarateTxFn(handler))
}

func (this *BoltDB) Batch(handler BoltTxHandler) error {
	return this.db.Batch(this.genarateTxFn(handler))
}

func (this *BoltDB) genarateTxFn(handler BoltTxHandler) func(*bolt.Tx) error {
	fn := func(tx *bolt.Tx) error {
		if handler == nil {
			return nil
		}

		boltTx := newBoltTx(this, tx)

		return handler.Handle(boltTx)
	}

	return fn
}

func (this *BoltDB) IsReadOnly() bool {
	return this.db.IsReadOnly()
}

// TODO
// bolt.DB.Stats

// NO EXPORTING:
// bolt.DB.Sync
// bolt.DB.Info
