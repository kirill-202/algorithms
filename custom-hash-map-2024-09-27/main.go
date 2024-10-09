package main


/*
custom hash map implementation
*/
import (
	"fmt"
	"crypto/sha512"
	"hash"
	"encoding/binary"
)

const HASHCAPACITY int = 100

type HashFunc func() hash.Hash


type AllowedTypes interface {
	int | float64 | string
}

type HashMap[T AllowedTypes] struct {
	keys []T
	values []any
	HashFunc HashFunc
}

func NewHashMap[T AllowedTypes](hashAlgoFunc HashFunc) *HashMap[T] {
	return &HashMap[T]{
		keys: make([]T, HASHCAPACITY),
		values: make([]any, HASHCAPACITY),
		HashFunc: hashAlgoFunc,
	}
}


func (m *HashMap[T]) hashKey(key T) (hashIndex int) {
	hash := m.HashFunc()
	hash.Write([]byte(fmt.Sprintf("%v", key)))
	rHash := hash.Sum(nil)
	//fmt.Printf("Hash of key %v: %x\n", key, rHash)
	hashIndex = int(binary.BigEndian.Uint64(rHash[:8])) % HASHCAPACITY
	//fmt.Printf("This is index output %d\n", hashIndex)
	if hashIndex < 0 {
        hashIndex = -hashIndex
    }
	return 
}

func (m *HashMap[T]) Put(key T, value any) {
	hashIndex := m.hashKey(key)

	m.keys[hashIndex] = key
	m.values[hashIndex] = value
}

func (m *HashMap[T]) Delete(key T) error {
	hashIndex := m.hashKey(key)
	
	if m.keys[hashIndex] != key {
		return fmt.Errorf("key %v does not exist", key)
	}
	m.keys = append(m.keys[:hashIndex], m.keys[hashIndex+1:]...)
	m.values =  append(m.values[:hashIndex], m.values[hashIndex+1:]...)

	return nil
}

func (m *HashMap[T]) Get(key T) (value any, err error) {
	hashIndex := m.hashKey(key)
	if m.keys[hashIndex] != key {
		return nil, fmt.Errorf("key '%v' does not exist", key)
	}
	return m.values[hashIndex], nil
}

func main() {
	hashMap := NewHashMap[string](sha512.New)
	hashMap.Put("First", 5)
	hashMap.Put("Second", 5)
	hashMap.Put("Third", 10)
	hashMap.Delete("Second")
	value1, err1 := hashMap.Get("First")

	fmt.Printf("testing hash... %v, %v\n", value1, err1)
	value2, err2 := hashMap.Get("Second")
	fmt.Printf("testing hash... %v, %v\n", value2, err2)
	value3, err3 := hashMap.Get("Third")
	fmt.Printf("testing hash... %v, %v\n", value3, err3)

}