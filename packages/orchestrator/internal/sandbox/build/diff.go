package build

import (
	"io"

	"github.com/khulnasoft/infra/packages/shared/pkg/storage"
)

type DiffType string

type ErrNoDiff struct{}

func (ErrNoDiff) Error() string {
	return "the diff is empty"
}

const (
	Memfile DiffType = storage.MemfileName
	Rootfs  DiffType = storage.RootfsName
)

type Diff interface {
	io.Closer
	io.ReaderAt
	Slice(off, length int64) ([]byte, error)
	CachePath() (string, error)
}

type NoDiff struct{}

func (n *NoDiff) CachePath() (string, error) {
	return "", ErrNoDiff{}
}

func (n *NoDiff) Slice(off, length int64) ([]byte, error) {
	return nil, ErrNoDiff{}
}

func (n *NoDiff) Close() error {
	return nil
}

func (n *NoDiff) ReadAt(p []byte, off int64) (int, error) {
	return 0, ErrNoDiff{}
}
