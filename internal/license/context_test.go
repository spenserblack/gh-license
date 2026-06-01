package license

import (
	"testing"
	"time"
)

func TestRender(t *testing.T) {
	context := Context{
		FullName: "Gopher",
		Date:     time.Date(2012, 3, 28, 12, 30, 45, 0, time.Local),
	}
	text := "Copyright (c) [year] [fullname]"
	want := "Copyright (c) 2012 Gopher"

	if got := context.Render(text); got != want {
		t.Fatalf(`Render() = %q, want %q`, got, want)
	}
}
