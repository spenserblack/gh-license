package license

import (
	"strconv"
	"strings"
	"time"

	"github.com/spenserblack/gh-license/internal/git"
)

// Context is the context used for filling in placeholders in license text.
type Context struct {
	// FullName is the full name of the user.
	FullName string
	// Time is the current time.
	Date time.Time
}

// DefaultContext builds the default context.
func DefaultContext(git git.Git) (ctx Context, err error) {
	ctx.FullName, err = git.GetConfig("user.name")
	if err != nil {
		return
	}
	ctx.Date = time.Now()
	return
}

// Render replaces text's placeholders with the actual values according to the context.
func (ctx Context) Render(s string) string {
	replacer := strings.NewReplacer(
		"[year]", strconv.Itoa(ctx.Date.Year()),
		"[fullname]", ctx.FullName,
	)
	return replacer.Replace(s)
}
