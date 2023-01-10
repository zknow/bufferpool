package main

import (
	"bytes"
	"log"

	"github.com/zknow/bufferpool/bufferpool"
)

func main() {
	poolSize := 3

	bfPool := bufferpool.NewBufferPool(poolSize)

	// 空的pool pool len=0
	cnt := bfPool.GetFreeBufferCount()
	log.Println(cnt) // count: 0

	// 產生一個buffer並回收後 pool len=1
	bf1 := bytes.NewBuffer(make([]byte, 0))
	bfPool.Put(bf1)
	cnt = bfPool.GetFreeBufferCount()
	log.Println(cnt) // count: 1

	// 產生兩個buffer並回收後 pool len=3
	bf2 := bytes.NewBuffer(make([]byte, 0))
	bf3 := bytes.NewBuffer(make([]byte, 0))
	bfPool.Put(bf2)
	bfPool.Put(bf3)
	cnt = bfPool.GetFreeBufferCount()
	log.Println(cnt) // count: 3

	// 達到pool size上限新buffer放入會失敗
	bf4 := bytes.NewBuffer(make([]byte, 0))
	bfPool.Put(bf4)
	cnt = bfPool.GetFreeBufferCount()
	log.Println(cnt) // count: 3
}
