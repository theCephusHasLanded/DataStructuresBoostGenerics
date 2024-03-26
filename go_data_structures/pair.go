package go_data_structures

import (
	"golang.org/x/exp/constraints"
)


type Pair[T constraints.Ordered, V any] struct {
  Key   T
  Priority T
  Value V
}
