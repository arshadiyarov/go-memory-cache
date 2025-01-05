package main

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(k string, v interface{}) {
	c.data[k] = v
}

func (c *Cache) Get(k string) interface{} {
	return c.data[k]
}

func (c *Cache) Delete(k string) {
	delete(c.data, k)
}

func (c *Cache) GetAll() map[string]interface{} {
	return c.data
}
