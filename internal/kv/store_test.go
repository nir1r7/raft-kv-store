
package kv

import "testing"

func TestPutGetDelete(t *testing.T) {
      s := NewStore()

      // missing key before any write
      if _, found := s.Get("a"); found {
              t.Fatalf("expected a to be missing")
      }

      // Put then Get
      s.Put("a", "1")
      v, found := s.Get("a")
      if !found || v != "1" {
              t.Fatalf("got (%q, %v), want (\"1\", true)", v, found)
      }

      // overwrite
      s.Put("a", "2")
      if v, _ := s.Get("a"); v != "2" {
              t.Fatalf("got %q, want \"2\"", v)
      }

      // Delete then Get
      s.Delete("a")
      if _, found := s.Get("a"); found {
              t.Fatalf("expected a to be gone after delete")
      }

      // Delete of a missing key
      s.Delete("nope")
}