package pgctl

import (
	"fmt"

	"github.com/octohelm/courier/pkg/statuserror"
)

type ErrPostgresNotReady struct {
	statuserror.FailedDependency

	Reason error
}

func (e *ErrPostgresNotReady) Error() string {
	return fmt.Sprintf("postgres is not ready: %s", e.Reason)
}
