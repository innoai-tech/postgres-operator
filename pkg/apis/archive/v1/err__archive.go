package v1

import (
	"fmt"

	"github.com/octohelm/courier/pkg/statuserror"
)

type ErrArchiveAlreadyRunning struct {
	statuserror.Conflict
}

func (e *ErrArchiveAlreadyRunning) Error() string {
	return fmt.Sprintf("archive is already running")
}

type ErrArchiveNotFound struct {
	statuserror.NotFound
}

func (e *ErrArchiveNotFound) Error() string {
	return fmt.Sprintf("archive not found")
}
