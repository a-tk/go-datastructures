package lru

import "testing"

func TestCache_Get(t *testing.T) {
	c := New[int, int](5)

	c.Put(1, 1)

	v, found := c.Get(1)

	if !found {
		t.Errorf("reported incorrect found")
	}

	if v != 1 {
		t.Errorf("reported incorrect value, v=%d", v)
	}
}

func TestCache_PutFull(t *testing.T) {
	c := New[int, int](4)

	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	c.Put(5, 5)

	// 1 should not be in the cache
	_, found := c.Get(1)

	if found {
		t.Errorf("found 1 for size 4!")
	}

	v, _ := c.Get(2)
	if v != 2 {
		t.Errorf("found something wierd for 2. v=%d", v)
	}
}

func TestCache_LRU(t *testing.T) {
	c := New[int, int](4)

	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)

	v, replaced := c.Put(5, 5)

	if !replaced {
		t.Errorf("should have been full and evicted a value")
	}
	if v != 1 {
		t.Errorf("value should have been 1, got=%d", v)
	}

}

func TestCache_MRU(t *testing.T) {
	c := New[int, int](3)

	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)

	c.Get(1) // 1, 3, 2

	v, replaced := c.Put(4, 4)

	if !replaced {
		t.Errorf("should have been full and evicted a value")
	}
	if v != 2 {
		t.Errorf("value should have been 2, got=%d", v)
	}

	v, found := c.Get(2)

	if found {
		t.Errorf("should have not found 2")
	}

	v, found = c.Get(1)

	if v != 1 {
		t.Errorf("value should have been 1, got=%d", v)
	}
}
