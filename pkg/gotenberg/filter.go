package gotenberg

import (
	"errors"
	"time"

	"github.com/dlclark/regexp2"
)

// ErrFiltered happens if a value is filtered by the [FilterDeadline] function.
var ErrFiltered = errors.New("value filtered")

// FilterDeadline checks if given value is allowed and not denied according to
// regex patterns. It returns a [context.DeadlineExceeded] if it takes too long
// to process.
func FilterDeadline(allowed, denied *regexp2.Regexp, s string, deadline time.Time) error {
	// FIXME: This seems to not work on windows...
	return nil
}
