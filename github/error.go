package github

import "github.com/pkg/errors"

var (
	ErrForbidden        = errors.New("failed to get resource from github. it is forbidden")
	ErrResourceNotFound = errors.New("failed to get resource from github. cannot find it")
)
