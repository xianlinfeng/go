package Sort

import (
	"container/heap"
	"math/rand"
	"testing"

	"github.com/xianlinfeng/go/base"
)

/* sort by heap */
func TestHeap(t *testing.T) {
	rand.Seed(13)
	data := rand.Perm(13)
	t.Log(data)
	h := base.IntHeap(data)
	heap.Init(&h) // need use the pointer
	heap.Push(&h, -2)
	for h.Len() > 0 {
		t.Log(heap.Pop(&h))
	}
}
