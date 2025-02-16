// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package set

import (
	"slices"
	"testing"
)

func TestSet(t *testing.T) {
	s := Set[int]{}
	s.Add(1)
	s.Add(2)
	if !s.Contains(1) {
		t.Error("missing 1")
	}
	if !s.Contains(2) {
		t.Error("missing 2")
	}
	if s.Contains(3) {
		t.Error("shouldn't have 3")
	}
	if s.Len() != 2 {
		t.Errorf("wrong len %d; want 2", s.Len())
	}

	more := []int{3, 4}
	s.AddSlice(more)
	if !s.Contains(3) {
		t.Error("missing 3")
	}
	if !s.Contains(4) {
		t.Error("missing 4")
	}
	if s.Contains(5) {
		t.Error("shouldn't have 5")
	}
	if s.Len() != 4 {
		t.Errorf("wrong len %d; want 4", s.Len())
	}

	es := s.Slice()
	if len(es) != 4 {
		t.Errorf("slice has wrong len %d; want 4", len(es))
	}
	for _, e := range []int{1, 2, 3, 4} {
		if !slices.Contains(es, e) {
			t.Errorf("slice missing %d (%#v)", e, es)
		}
	}
}

func TestSetOf(t *testing.T) {
	s := SetOf[int]([]int{1, 2, 3, 4, 4, 1})
	if s.Len() != 4 {
		t.Errorf("wrong len %d; want 2", s.Len())
	}
	for _, n := range []int{1, 2, 3, 4} {
		if !s.Contains(n) {
			t.Errorf("should contain %d", n)
		}
	}
}
