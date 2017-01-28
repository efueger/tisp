package rbt

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSetInsertRemoveRandomly(t *testing.T) {
	s := NewSet(less)

	for i := 0; i < MAX_ITERS; i++ {
		x := generateKey()
		insert := rand.Int()%2 == 0

		if insert {
			s = s.Insert(x)
		} else {
			s, _ = s.Remove(x)
		}

		ok := s.Include(x)

		if insert && !ok || !insert && ok {
			t.Fail()
		}
	}
}

func TestSetFirstRest(t *testing.T) {
	xs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := NewSet(less)

	for _, x := range xs {
		s = s.Insert(x)
	}

	x, f := s.FirstRest()

	for _, expected := range xs {
		t.Log(x)
		assert.Equal(t, expected, x)
		x, f = f()
	}

	assert.Equal(t, nil, x)
	assert.Equal(t, FirstRestFunc(nil), f)
}
