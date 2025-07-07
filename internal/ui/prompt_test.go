package ui

import (
    "testing"
)

func TestIsSlug(t *testing.T) {
    if !isSlug("foo-bar_123") {
        t.Error("should accept valid slug")
    }
    if isSlug("foo bar!") {
        t.Error("should reject invalid slug")
    }
}
