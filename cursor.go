package boltmobile

import (
	"github.com/boltdb/bolt"
)

type BoltCursor struct {
	boltBucket *BoltBucket
	cursor     *bolt.Cursor
}

func newBoltCursor(boltBucket *BoltBucket, cursor *bolt.Cursor) *BoltCursor {
	return &BoltCursor{
		boltBucket: boltBucket,
		cursor:     cursor,
	}
}

func (this *BoltCursor) Bucket() *BoltBucket {
	return this.boltBucket
}

func (this *BoltCursor) First() *KeyValue {
	k, v := this.cursor.First()
	if k == nil {
		return nil
	}

	return &KeyValue{
		key:   k,
		value: v,
	}
}

func (this *BoltCursor) Last() *KeyValue {
	k, v := this.cursor.Last()
	if k == nil {
		return nil
	}

	return &KeyValue{
		key:   k,
		value: v,
	}
}

func (this *BoltCursor) Next() *KeyValue {
	k, v := this.cursor.Next()
	if k == nil {
		return nil
	}

	return &KeyValue{
		key:   k,
		value: v,
	}
}

func (this *BoltCursor) Prev() *KeyValue {
	k, v := this.cursor.Prev()
	if k == nil {
		return nil
	}

	return &KeyValue{
		key:   k,
		value: v,
	}
}

func (this *BoltCursor) Seek(seek []byte) *KeyValue {
	k, v := this.cursor.Seek(seek)
	if k == nil {
		return nil
	}

	return &KeyValue{
		key:   k,
		value: v,
	}
}

func (this *BoltCursor) Delete() error {
	return this.cursor.Delete()
}

// TODO:

// NO EXPORTING:
