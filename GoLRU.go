package GoLRU

import (
	"github.com/emirpasic/gods/maps/linkedhashmap"
	_ "github.com/emirpasic/gods/utils"
)


type LRU struct {
	cache    *linkedhashmap.Map
	size     int
}

func New(s int) *LRU {
	return &LRU{
		cache:    linkedhashmap.New(),
		size:     s,
	}
}

// Get oldest key
func (lm *LRU) Oldest() (key interface{}, found bool) {
	if lm.cache.Size() == 0 {
		return nil, false
	}
	iterator := lm.cache.Iterator()
	_ = iterator.Last()
	return iterator.Key(), true
}

// Get newest key
func (lm *LRU) Newest() (key interface{}, found bool) {
	if lm.cache.Size() == 0 {
		return nil, false
	}
	iterator := lm.cache.Iterator()
	_ = iterator.First()
	return iterator.Key(), true
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (lm *LRU) Put(key interface{}, value interface{}) {
	if lm.cache.Size() >= lm.size {
		key, _ := lm.Oldest()
		lm.Remove(key)
	}
	lm.cache.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (lm *LRU) Get(key interface{}) (value interface{}, found bool) {
	return lm.cache.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (lm *LRU) Remove(key interface{}) {
	lm.cache.Remove(key)
}

// Empty returns true if map does not contain any elements
func (lm *LRU) Empty() bool {
	return lm.cache.Size() == 0
}

// Size returns number of elements in the map.
func (lm *LRU) Size() int {
	return lm.cache.Size()
}

// Set the new max size
func (lm *LRU) SetMaxSize(s int) {
	lm.size = s
}

// Get the max size
func (lm *LRU) GetMaxSize() int {
	return lm.size
}

// Keys returns all keys in-order
func (lm *LRU) Keys() []interface{} {
	return lm.cache.Keys()
}

// Values returns all values in-order based on the key.
func (lm *LRU) Values() []interface{} {
	return lm.cache.Values()
}

// Clear removes all elements from the map.
func (lm *LRU) Clear() {
	lm.cache.Clear()
}

// String returns a string representation of container
func (lm *LRU) String() string {
	return lm.cache.String()
}

// Return Iterator
func (lm *LRU) Iterator() linkedhashmap.Iterator {
	return lm.cache.Iterator()
}
