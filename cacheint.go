package mfc

import "sync"

// CacheInt32 - simple Cache (key is int)
type CacheInt32 struct {
	data map[int32]interface{}

	mu sync.Mutex

	// KeyGet - get key by item
	KeyGet func(item interface{}) (key int32, err error)
	// ValueGet - get value by item if it necessary. If nil then used item
	ValueGet func(item interface{}) (val interface{})

	// OnAppend do if not nil
	OnAppend func(key int32, item interface{})
	// OnDelete do if not nil
	OnDelete func(key int32, item interface{})
}

// CacheInt32Create create CacheInt32
func CacheInt32Create() (c *CacheInt32) {
	c = &CacheInt32{
		data: make(map[int32]interface{}),
	}
	return c
}

// Append - append item into cache
// if exists then delete befor
func (c *CacheInt32) Append(item interface{}) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	key, err := c.KeyGet(item)
	if err != nil {
		return err
	}

	return c.AppendKVUnSave(key, item)
}

// AppendKV - append item into cache
// if exists then delete befor
func (c *CacheInt32) AppendKV(key int32, item interface{}) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.AppendKVUnSave(key, item)
}

// AppendKVUnSave - append item into cache
// if exists then delete befor
func (c *CacheInt32) AppendKVUnSave(key int32, item interface{}) (err error) {
	val, ok := c.data[key]
	if ok {
		c.DeleteKVUnSave(key, val)
	}

	if c.ValueGet == nil {
		c.data[key] = item
	} else {
		c.data[key] = c.ValueGet(item)
	}

	c.OnAppend(key, item)

	return nil
}

// Delete - delete element from cache
func (c *CacheInt32) Delete(item interface{}) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	key, err := c.KeyGet(item)
	if err != nil {
		return err
	}

	return c.DeleteKVUnSave(key, item)
}

// DeleteKV - delete element from cache
func (c *CacheInt32) DeleteKV(key int32, item interface{}) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.DeleteKVUnSave(key, item)
}

// DeleteKVUnSave - delete element from cache
func (c *CacheInt32) DeleteKVUnSave(key int32, item interface{}) (err error) {
	_, ok := c.data[key]
	if ok {
		delete(c.data, key)
		c.OnDelete(key, item)
	}

	return nil
}
