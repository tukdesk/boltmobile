package boltmobile

type KeyValue struct {
	key   []byte
	value []byte
}

func (this *KeyValue) Key() []byte {
	return this.key
}

func (this *KeyValue) Value() []byte {
	return this.value
}
