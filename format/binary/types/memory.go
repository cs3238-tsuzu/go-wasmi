package bintypes

import (
	"io"

	"github.com/cs3238-tsuzu/go-wasmi/types"
	"github.com/pkg/errors"
)

// ReadMemoryType parses memory type from reader
func ReadMemoryType(r io.Reader) (*types.MemoryType, error) {
	limits, err := ReadLimits(r)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &types.MemoryType{
		Limits: *limits,
	}, nil
}
