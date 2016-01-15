package boltmobile

import (
	"github.com/boltdb/bolt"
)

type BoltTxEachHandler interface {
	Handle(name []byte, boltBucket *BoltBucket) error
}

type BoltTxCommitHandler interface {
	Handle()
}

type BoltTx struct {
	boltDB *BoltDB
	tx     *bolt.Tx
}

func newBoltTx(boltDB *BoltDB, tx *bolt.Tx) *BoltTx {
	return &BoltTx{
		boltDB: boltDB,
		tx:     tx,
	}
}

func (this *BoltTx) ID() int {
	return this.tx.ID()
}

func (this *BoltTx) DB() (db *BoltDB) {
	db = this.boltDB
	return
}

func (this *BoltTx) Size() int64 {
	return this.tx.Size()
}

func (this *BoltTx) Writable() bool {
	return this.Writable()
}

func (this *BoltTx) Cursor() *BoltCursor {
	cursor := this.tx.Cursor()
	bucket := cursor.Bucket()

	return newBoltCursor(newBoltBucket(this, bucket), cursor)
}

func (this *BoltTx) Bucket(name []byte) *BoltBucket {
	bucket := this.tx.Bucket(name)

	return newBoltBucket(this, bucket)
}

func (this *BoltTx) CreateBucket(name []byte) (boltBucket *BoltBucket, err error) {
	bucket, err := this.tx.CreateBucket(name)
	if err != nil {
		return nil, err
	}

	return newBoltBucket(this, bucket), nil
}

func (this *BoltTx) CreateBucketIfNotExists(name []byte) (boltBucket *BoltBucket, err error) {
	bucket, err := this.tx.CreateBucketIfNotExists(name)
	if err != nil {
		return nil, err
	}

	return newBoltBucket(this, bucket), nil
}

func (this *BoltTx) DeleteBucket(name []byte) error {
	return this.tx.DeleteBucket(name)
}

func (this *BoltTx) Commit() error {
	return this.tx.Commit()
}

func (this *BoltTx) Rollback() error {
	return this.tx.Rollback()
}

func (this *BoltTx) ForEach(handler BoltTxEachHandler) error {
	fn := func(name []byte, b *bolt.Bucket) error {
		if handler == nil {
			return nil
		}

		boltBucket := newBoltBucket(this, b)
		return handler.Handle(name, boltBucket)
	}

	return this.tx.ForEach(fn)
}

func (this *BoltTx) OnCommit(handler BoltTxCommitHandler) {
	if handler != nil {
		this.tx.OnCommit(handler.Handle)
	}
}

// TODO:
// bolt.Tx.Stats
// bolt.Tx.WriteTo
// bolt.Tx.CopyFile

// NO EXPORTING:
// bolt.Tx.Check
// bolt.Tx.Copy
// bolt.Tx.Page
