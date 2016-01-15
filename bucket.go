package boltmobile

import (
	"github.com/boltdb/bolt"
)

type BoltBucketEachHandler interface {
	Handle(key, value []byte) error
}

type BoltBucket struct {
	boltTx *BoltTx
	bucket *bolt.Bucket
}

func newBoltBucket(boltTx *BoltTx, bucket *bolt.Bucket) *BoltBucket {
	return &BoltBucket{
		boltTx: boltTx,
		bucket: bucket,
	}
}

func (this *BoltBucket) Tx() *BoltTx {
	return this.boltTx
}

func (this *BoltBucket) Writable() bool {
	return this.boltTx.Writable()
}

func (this *BoltBucket) Cursor() *BoltCursor {
	cursor := this.bucket.Cursor()
	return newBoltCursor(this, cursor)
}

func (this *BoltBucket) Bucket(name []byte) *BoltBucket {
	bucket := this.bucket.Bucket(name)
	return newBoltBucket(this.boltTx, bucket)
}

func (this *BoltBucket) CreateBucket(name []byte) (boltBucket *BoltBucket, err error) {
	bucket, err := this.bucket.CreateBucket(name)
	if err != nil {
		return nil, err
	}

	return newBoltBucket(this.boltTx, bucket), nil
}

func (this *BoltBucket) CreateBucketIfNotExists(name []byte) (boltBucket *BoltBucket, err error) {
	bucket, err := this.bucket.CreateBucketIfNotExists(name)
	if err != nil {
		return nil, err
	}

	return newBoltBucket(this.boltTx, bucket), nil
}

func (this *BoltBucket) DeleteBucket(name []byte) error {
	return this.bucket.DeleteBucket(name)
}

func (this *BoltBucket) Get(key []byte) []byte {
	return this.bucket.Get(key)
}

func (this *BoltBucket) Put(key, value []byte) error {
	return this.bucket.Put(key, value)
}

func (this *BoltBucket) Delete(key []byte) error {
	return this.bucket.Delete(key)
}

// TODO: 等待 gobind 支持 uint64
func (this *BoltBucket) NextSequence() (seq int64, err error) {
	next, err := this.bucket.NextSequence()
	return int64(next), err
}

func (this *BoltBucket) ForEach(handler BoltBucketEachHandler) error {
	fn := func(key []byte, value []byte) error {
		if handler == nil {
			return nil
		}

		return handler.Handle(key, value)
	}

	return this.bucket.ForEach(fn)
}

// TODO:
// bolt.Bucket.Root
// bolt.Bucket.Stats

// NO EXPORTING:
