package deepmap_test

import (
	"testing"

	"github.com/LiamMartens/konf-vlt/deepmap"
)

func TestDeepInsertShallow(t *testing.T) {
	dest := make(map[string]any)
	deepmap.DeepInsert(dest, []string{"foo"}, "bar")
	if dest["foo"] != "bar" {
		t.Fatalf("failed to insert key into map %q", "foo")
	}
}

func TestDeepInsertDeep(t *testing.T) {
	dest := make(map[string]any)
	deepmap.DeepInsert(dest, []string{"foo", "bar", "baz"}, "foobarbaz")

	foo_map := dest["foo"].(map[string]any)
	bar_map := foo_map["bar"].(map[string]any)
	if bar_map["baz"] != "foobarbaz" {
		t.Fatalf("failed to insert key into map %q", "foo.bar.baz")
	}
}
