package main

import (
	"math/rand"
	"strconv"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type DBSuite struct{}

var _ = Suite(&DBSuite{})

func testDBGet(c *C, db DB) {
	db := NewSliceDB()
	_, err := db.Get([]byte("key"))
	c.Assert(err, ErrorMatches, "key not found:.*")

	// Put something else in the store
	db.Put([]byte("other_key"), []byte("value"))
	_, err = db.Get([]byte("key"))
	c.Assert(err, ErrorMatches, "key not found:.*")

	// Now put the key, and make sure we get it back
	db.Put([]byte("key"), []byte("value"))
	val, err := db.Get([]byte("key"))
	c.Assert(err, IsNil)
	c.Assert(string(val), Equals, "value")
}


func (s *DBSuite) TestDBGet(c *C) {
	testDBGet(c, NewSliceDB())
}

func (s *DBSuite) TestDBDelete(c *C) {
	db := NewSliceDB()
	err := db.Delete([]byte("key"))
	c.Assert(err, ErrorMatches, "key not found:.*")

	// Put something else in the store
	db.Put([]byte("other_key"), []byte("value"))
	err = db.Delete([]byte("key"))
	c.Assert(err, ErrorMatches, "key not found:.*")

	err = db.Delete([]byte("other_key"))
	c.Assert(err, IsNil)
	_, err = db.Get([]byte("other_key"))
	c.Assert(err, ErrorMatches, "key not found:.*")
}

func (s *DBSuite) TestDBHas(c *C) {
	db := NewSliceDB()
	hasKey, err := db.Has([]byte("key"))
	c.Assert(hasKey, Equals, false)
	c.Assert(err, IsNil)

	db.Put([]byte("key"), []byte("value"))
	hasKey, err = db.Has([]byte("key"))
	c.Assert(hasKey, Equals, true)
	c.Assert(err, IsNil)
}

func (s *DBSuite) TestRangeScan(c *C) {
	db := NewSliceDB()
	iter, err := db.RangeScan([]byte("0"), []byte("1"))
	c.Assert(err, IsNil)
	c.Assert(iter.Error(), IsNil)
	c.Assert(iter.Key(), IsNil)
	c.Assert(iter.Value(), IsNil)
	c.Assert(iter.Next(), Equals, false)

	db.Put([]byte("2"), []byte("value2"))
	db.Put([]byte("5"), []byte("value5"))
	db.Put([]byte("4"), []byte("value4"))
	db.Put([]byte("1"), []byte("value1"))
	db.Put([]byte("3"), []byte("value3"))

	iter, err = db.RangeScan([]byte("1"), []byte("6"))
	c.Assert(err, IsNil)
	c.Assert(iter.Error(), IsNil)
	c.Assert(string(iter.Key()), Equals, "1")
	c.Assert(string(iter.Value()), Equals, "value1")
	c.Assert(iter.Next(), Equals, true)

	c.Assert(string(iter.Key()), Equals, "2")
	c.Assert(string(iter.Value()), Equals, "value2")
	c.Assert(iter.Next(), Equals, true)

	c.Assert(string(iter.Key()), Equals, "3")
	c.Assert(string(iter.Value()), Equals, "value3")
	c.Assert(iter.Next(), Equals, true)

	c.Assert(string(iter.Key()), Equals, "4")
	c.Assert(string(iter.Value()), Equals, "value4")
	c.Assert(iter.Next(), Equals, true)

	c.Assert(string(iter.Key()), Equals, "5")
	c.Assert(string(iter.Value()), Equals, "value5")
	c.Assert(iter.Next(), Equals, false)

	c.Assert(string(iter.Key()), Equals, "")
	c.Assert(string(iter.Value()), Equals, "")
	c.Assert(iter.Next(), Equals, false)
}

func benchmarkPut(b *testing.B, db DB, keys [][]byte) [][]byte {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		keyStr := "key" + strconv.Itoa(rand.Intn(b.N))
		key := []byte(keyStr)
		value := []byte("value" + strconv.Itoa(rand.Intn(b.N)))
		b.StartTimer()
		db.Put(key, value)
		b.StopTimer()
		keys = append(keys, key)
	}
	return keys
}

func benchmarkGet(b *testing.B, db DB, keys [][]byte) {
	for i := 0; i < b.N; i++ {
		db.Get(keys[i%len(keys)])
	}
}

func benchmarkDelete(b *testing.B, db DB, keys [][]byte) {
	for i := 0; i < b.N; i++ {
		k := keys[i%len(keys)]
		b.StartTimer()
		db.Delete(k)
		b.StopTimer()
	}
}

func benchmarkRangeScan(b *testing.B, db DB, keys [][]byte) {
	for i := 0; i < b.N; i++ {
		start := []byte("key" + strconv.Itoa(rand.Intn(b.N)))
		end := []byte("key" + strconv.Itoa(rand.Intn(b.N)+1))
		b.StartTimer()
		iter, err := db.RangeScan(start, end)
		if err == nil {
			for iter.Next() {
				_ = iter.Key()
				_ = iter.Value()
			}
		}
	}
}

func BenchmarkDB(b *testing.B) {
	db := NewSliceDB()
	keys := make([][]byte, 0, b.N)
	b.Run("Put", func(b *testing.B) {
		keys = benchmarkPut(b, db, keys)
	})
	b.Run("Get", func(b *testing.B) {
		benchmarkGet(b, db, keys)
	})
	b.Run("RangeScan", func(b *testing.B) {
		benchmarkRangeScan(b, db, keys)
	})
	b.Run("Delete", func(b *testing.B) {
		benchmarkDelete(b, db, keys)
	})
}
