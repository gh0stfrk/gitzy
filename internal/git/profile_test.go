package git

import (
    "os"
    "testing"
)

func TestIndexSaveLoad(t *testing.T) {
    dir := t.TempDir()
    os.Setenv("GITZY_HOME", dir)
    idx := &Index{Profiles: map[string]*Profile{
        "foo": {Name: "foo", UserName: "bar"},
    }}
    if err := SaveIndex(idx); err != nil {
        t.Fatal(err)
    }
    idx2, err := LoadIndex()
    if err != nil {
        t.Fatal(err)
    }
    if idx2.Profiles["foo"].UserName != "bar" {
        t.Error("profile not saved/loaded correctly")
    }
}
