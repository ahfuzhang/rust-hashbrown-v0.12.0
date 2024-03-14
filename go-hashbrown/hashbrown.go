package main

import (
	// "constraints"
	"fmt"
	"reflect"
	"unsafe"
	"github.com/cespare/xxhash/v2"
)

const GROUP_SIZE = 16

const (
	EMPTY   byte = 0xFF
	DELETED byte = 0x80
)

type hashBrownHeader struct {
	totalSize  uint64
	bucketMask uint64
	growthLeft uint64
	itemsCount uint64
}

type hashBrown[KEY comparable, VALUE any] struct {
	hashBrownHeader
}

type bucket struct { // 桶的格式
	data []byte
}

func newHashBrown[KEY comparable, VALUE any](capacity uint64) *hashBrown[KEY, VALUE] {
	// capacity = capacity*8 / 7
	// capacity = GetPowOfTwo(capacity)
	capacity = CapacityToBuckets(capacity)
	// 找到capacity的最高位
	// out := &hashBrown[KEY, VALUE]{
	// 	hashBrownHeader{
	// 		bucketMask: capacity-1,
	// 		growthLeft: capacity*8/7, // 87.5%的装载率
	// 		itemsCount:0,
	// 	},
	// }
	headerSize := uint64(unsafe.Sizeof(hashBrownHeader{}))
	headerPadSize := ((headerSize + 15) / 16) * 16
	var k KEY
	var v VALUE
	bucketSize := uint64(unsafe.Sizeof(k)+unsafe.Sizeof(v)) * capacity
	// bucketSize *= capacity
	bucketPadSize := ((bucketSize + 15) / 16) * 16
	totalSize := headerPadSize + bucketPadSize + capacity + GROUP_SIZE
	buf := make([]byte, totalSize)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	out := (*hashBrown[KEY, VALUE])(unsafe.Pointer(sh.Data))
	out.totalSize = totalSize
	out.bucketMask = capacity - 1
	out.growthLeft = BucketMaskToCapacity(capacity - 1)
	out.itemsCount = 0
	ctrls := out.ctrls()
	for i := range ctrls {
		ctrls[i] = EMPTY
	}
	return out
}

// BucketMaskToCapacity
func BucketMaskToCapacity(bucketMask uint64) uint64 { // 87.5%的装载率
	if bucketMask < 16 {
		return 16
	}
	return (bucketMask + 1) / 8 * 7
}

// CapacityToBuckets 获得正确的桶的容量
func CapacityToBuckets(capacity uint64) uint64 {
	if capacity < 16 {
		capacity = 16
	}
	adjustedCap := capacity * 8 / 7
	return NextPowerOfTwo(adjustedCap)
}

// NextPowerOfTwo 与2的幂对齐
func NextPowerOfTwo(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func GetPowOfTwo(capacity uint64) uint64 {
	// 返回向上对齐的2的幂
	// for i := 63; i >= 4; i-- {
	// 	if (capacity >> i & 0x01) == 1 {
	// 		return 1 << i
	// 	}
	// }
	// return 16
	v := capacity
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	if v < 16 {
		return 16
	}
	return v
}

// 根据下标获取内容
func (h *hashBrown[KEY, VALUE]) getBucket(index int) []byte {
	ptr := unsafe.Pointer(h)
	offset := uint64(unsafe.Sizeof(hashBrownHeader{}))
	offset = ((offset + 15) / 16) * 16
	var k KEY
	var v VALUE
	bucketSize := uint64(unsafe.Sizeof(k) + unsafe.Sizeof(v))
	header := reflect.SliceHeader{
		Data: uintptr(ptr) + uintptr(offset),
		Len:  int(bucketSize),
		Cap:  int(bucketSize),
	}
	return *(*[]byte)(unsafe.Pointer(&header))
}

func (h *hashBrown[KEY, VALUE]) ctrls() []byte {
	ptr := unsafe.Pointer(h)
	offset := uint64(unsafe.Sizeof(hashBrownHeader{}))
	offset = ((offset + 15) / 16) * 16
	var k KEY
	var v VALUE
	bucketSize := uint64(unsafe.Sizeof(k)+unsafe.Sizeof(v)) * (h.bucketMask + 1)
	offset += bucketSize
	offset = ((offset + 15) / 16) * 16
	header := reflect.SliceHeader{
		Data: uintptr(ptr) + uintptr(offset),
		Len:  int(h.bucketMask + 1),
		Cap:  int(h.bucketMask + 1),
	}
	return *(*[]byte)(unsafe.Pointer(&header))
}

func (h *hashBrown[KEY, VALUE]) String() string {
	return fmt.Sprintf("total size=%d\nbucketCount=%d\ngrowthLeft=%d\nitems=%d\n",
		h.totalSize,
		h.bucketMask+1,
		h.growthLeft,
		h.itemsCount,
	)
}

// 插入数据  string类型应该怎么计算hashcode呢？  golang写通用hash还真的是困难啊
// func (h *hashBrown[KEY, VALUE]) insert(k KEY, v VALUE) {
// 	header := reflect.SliceHeader{
// 		Data: uintptr(unsafe.Pointer(&k)),
// 		Len:  int(h.bucketMask + 1),
// 		Cap:  int(h.bucketMask + 1),
// 	}
// 	return *(*[]byte)(unsafe.Pointer(&header))
// 	h := xxhash.Sum64(k)
// }

// type item struct {
// 	Key   uint64
// 	Value uint64
// }

// const MAX_ITEMS = 1024
// const GROUP_SIZE = 16

// type hashBrown struct {
// 	bucketMask uint64
// 	growthLeft uint64
// 	itemsCount uint64
// 	buckets    [MAX_ITEMS]item
// 	ctrls      [MAX_ITEMS + GROUP_SIZE]byte
// }

// func newHashBrown(capacity uint64) *hashBrown {
// 	if capacity != MAX_ITEMS {
// 		capacity = MAX_ITEMS
// 	}
// 	out := &hashBrown{
// 		bucketMask: capacity - 1,
// 		growthLeft: capacity * 8 / 7, //装载率 87.5%
// 		itemsCount: 0,
// 	}
// 	for i := uint64(0); i < capacity; i++ {
// 		out.ctrls[i] = EMPTY
// 	}
// 	return out
// }

// func (h *hashBrown) insert(k uint64, v uint64) {

// }
func main() {
	fmt.Println("hello")
	h := newHashBrown[string, string](500)
	fmt.Println(h)
	arr := h.ctrls()
	fmt.Println(len(arr))
	fmt.Println(arr[0], arr[len(arr)-1])
	fmt.Println(h.getBucket(0))
	fmt.Println(h.getBucket(len(arr) - 1))
}
