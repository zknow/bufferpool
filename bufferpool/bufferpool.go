package bufferpool

import (
	"bytes"
)

type BufferPool struct {
	pool chan *bytes.Buffer
}

func NewBufferPool(size int) (bp *BufferPool) {
	return &BufferPool{
		pool: make(chan *bytes.Buffer, size),
	}
}

// 拿取buffer
func (bp *BufferPool) Get() (b *bytes.Buffer) {
	select {
	case b = <-bp.pool:
		//如果pool有空閑的buffer則回傳可使用的buffer
	default:
		//如果當下沒有空閑的 buffer 則重新建構一個出來
		b = bytes.NewBuffer(make([]byte, 0))
	}
	return
}

// 回收buffer
func (bp *BufferPool) Put(b *bytes.Buffer) {
	b.Reset() // 重置buffer空間

	select {
	case bp.pool <- b:
		// 把buffer放回pool中
	default:
		// 如果pool空間已滿，丟棄掉此buffer
	}
}

// 取得pool目前空閒的buffer數量
func (bp *BufferPool) GetFreeBufferCount() int {
	return len(bp.pool)
}
