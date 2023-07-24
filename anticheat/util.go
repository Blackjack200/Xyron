package anticheat

import "sort"

type BufferedData[T any] struct {
	prev, cur T
}

func NewBufferedData[T any](cur T) *BufferedData[T] {
	return &BufferedData[T]{prev: cur, cur: cur}
}

func (b *BufferedData[T]) Previous() T {
	return b.prev
}

func (b *BufferedData[T]) Current() T {
	return b.cur
}

func (b *BufferedData[T]) Set(v T) {
	b.prev, b.cur = b.cur, v
}

type TickedData[T any] struct {
	dv T
	v  T
}

func (t *TickedData[T]) Default(dv T) {
	t.dv = dv
}

func (t *TickedData[T]) Set(v T) {
	t.v = v
}
func (t *TickedData[T]) Get() T {
	return t.v
}

func (t *TickedData[T]) Reset() {
	t.v = t.dv
}

func NewTickedData[T any](dv T) *TickedData[T] {
	return &TickedData[T]{dv: dv, v: dv}
}

func NewTickedBool(d bool) *TickedData[bool] {
	return &TickedData[bool]{
		dv: d,
		v:  d,
	}
}

type TimestampedData[T any] struct {
	t int64
	v T
}

func (t *TimestampedData[T]) Timestamp() int64 {
	return t.t
}

func (t *TimestampedData[T]) Get() T {
	return t.v
}

func NewTimestampedData[T any](timestamp int64, v T) *TimestampedData[T] {
	return &TimestampedData[T]{t: timestamp, v: v}
}

type ComparableSlice[T ~int |
	~int8 | ~int32 | ~int64 |
	~uint8 | ~uint32 | ~uint64 |
	~float32 | ~float64,
] []T

func (x ComparableSlice[T]) Len() int           { return len(x) }
func (x ComparableSlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x ComparableSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x ComparableSlice[T]) Sort() { sort.Sort(x) }
