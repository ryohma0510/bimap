package bimap

import (
	"fmt"

	"github.com/samber/lo"
)

// BiMap is a bidirectional map, which allows you to look up a value by its key
type BiMap[K, V comparable] struct {
	forward map[K]V
	inverse map[V]K
}

// NewFromMap creates a new BiMap from a forward map.
// The forward map must not contain duplicate values.
func NewFromMap[K, V comparable](forward map[K]V) (*BiMap[K, V], error) {
	inverse := lo.Invert(forward)
	if len(forward) != len(inverse) {
		return nil, fmt.Errorf("duplicate values in forward map")
	}

	return &BiMap[K, V]{
		forward: forward,
		inverse: inverse,
	}, nil
}

// Get returns the value associated with the key, and a boolean indicating
func (b *BiMap[K, V]) Get(k K) (V, bool) {
	v, ok := b.forward[k]
	return v, ok
}

// InverseGet returns the key associated with the value, and a boolean indicating
func (b *BiMap[K, V]) InverseGet(v V) (K, bool) {
	k, ok := b.inverse[v]
	return k, ok
}
