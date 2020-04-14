# GoLRU
 A constant time LRU cache written in Go

```go
package main

func main() {
 LRU := GoLRU.New(1000)   // empty (keys are of type int)
 LRU.Put(2, "b")          // 2->b
 LRU.Put(1, "x")          // 2->b, 1->x (insertion-order)
 LRU.Put(1, "a")          // 2->b, 1->a (insertion-order)
 LRU.Oldest()             // 2
 LRU.Newest()             // 1
 LRU.SetMaxSize(2000)     // MaxSize <- 2000
 LRU.GetMaxSize()         // 2000
 _, _ = LRU.Get(2)        // b, true
 _, _ = LRU.Get(3)        // nil, false
 _ = LRU.Values()         // []interface {}{"b", "a"} (insertion-order)
 _ = LRU.Keys()           // []interface {}{2, 1} (insertion-order)
 LRU.Remove(1)            // 2->b
 LRU.Clear()              // empty
 LRU.Empty()              // true
 LRU.Size()               // 0
}
```
