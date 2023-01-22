package main

import (
	"errors"
	"fmt"
	"sort"
)

type DB interface {
	// Get gets the value for the given key. It returns an error if the
	// DB does not contain the key.
	Get(key []byte) (value []byte, err error)

	// Has returns true if the DB contains the given key.
	Has(key []byte) (ret bool, err error)

	// Put sets the value for the given key. It overwrites any previous value
	// for that key; a DB is not a multi-map.
	Put(key, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key []byte) error

	// RangeScan returns an Iterator (see below) for scanning through all
	// key-value pairs in the given range, ordered by key ascending.
	RangeScan(start, limit []byte) (Iterator, error)
}

type Iterator interface {
	// Next moves the iterator to the next key/value pair.
	// It returns false if the iterator is exhausted.
	Next() bool

	// Error returns any accumulated error. Exhausting all the key/value pairs
	// is not considered to be an error.
	Error() error

	// Key returns the key of the current key/value pair, or nil if done.
	Key() []byte

	// Value returns the value of the current key/value pair, or nil if done.
	Value() []byte
}

type sliceDB struct {
	sortedKeys []string
	innerMap   map[string][]byte
}

/*
Properties:
- always errs on an empty store
- errs if key does not exist
- forall k, v; Put(k, v) -> Get(k) = v
*/
func (db *sliceDB) Get(key []byte) ([]byte, error) {
	val, ok := db.innerMap[string(key)]
	if !ok {
		return nil, errors.New(fmt.Sprintf("key not found: 0x%x", key))
	}
	return val, nil
}

/*
Properties:
- Always errs on an empty store
- errs if key does not exist
- forall k, v; Put(k, v) /\ Delete(k) -> ~Get(k)
- forall k1, k2, v1, v2; Put(k1, v1) /\ Put(k2, v2) /\ Delete(k1) -> Get(k2) = v2
*/
func (db *sliceDB) Delete(key []byte) error {
	sKey := string(key)
	_, ok := db.innerMap[sKey]
	if !ok {
		return errors.New(fmt.Sprintf("key not found: 0x%x", key))
	}
	idx := sort.SearchStrings(db.sortedKeys, sKey)
	db.sortedKeys = append(db.sortedKeys[:idx], db.sortedKeys[idx+1:]...)
	delete(db.innerMap, sKey)
	return nil
}

/*
Properties:
- If store is empty, returns false
- forall k; Has(k) <-> exists v; Get(k) = v
*/
func (db sliceDB) Has(key []byte) (bool, error) {
	_, ok := db.innerMap[string(key)]
	return ok, nil
}

func (db *sliceDB) Put(key, value []byte) error {
	sKey := string(key)
	db.innerMap[sKey] = value
	idx := sort.SearchStrings(db.sortedKeys, sKey)
	db.sortedKeys = append(db.sortedKeys, "")
	copy(db.sortedKeys[idx+1:], db.sortedKeys[idx:])
	db.sortedKeys[idx] = sKey
	return nil
}

/*
Properties:
- If store is empty, returns empty iterator
- If keys are out of order, returns empty iterator
- Returns an Iterator over keys that are in the range [start, limit) when called.
*/
func (db sliceDB) RangeScan(start, limit []byte) (Iterator, error) {
	startIdx := sort.SearchStrings(db.sortedKeys, string(start))
	limitIdx := sort.SearchStrings(db.sortedKeys, string(limit))
	if startIdx > limitIdx {
		return nil, errors.New(fmt.Sprintf("Range inversion. start idx: %d, end idx: %d", startIdx, limitIdx))
	}
	sortedKeys := db.sortedKeys[startIdx:limitIdx]
	innerMap := make(map[string][]byte)
	for _, k := range sortedKeys {
		innerMap[k] = db.innerMap[k]
	}
	return &iterator{
		cursor:     0,
		sortedKeys: sortedKeys,
		innerMap:   innerMap,
	}, nil
}

func NewSliceDB() DB {
	return &sliceDB{
		innerMap: make(map[string][]byte),
	}
}

type iterator struct {
	cursor     int
	sortedKeys []string
	innerMap   map[string][]byte
	err        error
}

func (i iterator) Error() error {
	return i.err
}

func (i iterator) hasCurrent() bool {
	return i.cursor < len(i.sortedKeys)
}

func (i iterator) Key() []byte {
	if !i.hasCurrent() {
		return nil
	}
	return []byte(i.sortedKeys[i.cursor])
}

func (i iterator) Value() []byte {
	if !i.hasCurrent() {
		return nil
	}
	return i.innerMap[i.sortedKeys[i.cursor]]
}

func (i *iterator) Next() bool {
	i.cursor++
	if !i.hasCurrent() {
		return false
	}
	return true
}
