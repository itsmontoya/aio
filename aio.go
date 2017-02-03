package aio

import (
	"os"
	//	"github.com/missionMeteora/toolkit/errors"
)

// New returns a new AIO
func New() *AIO {
	a := AIO{
		rq: make(chan interface{}, 32),
	}

	t := newThread(a.rq)
	go t.listen()
	return &a
}

// AIO does stuff
type AIO struct {
	rq chan interface{}
}

// Open will open a new file for reading
func (a *AIO) Open(key string) <-chan *OpenResp {
	return a.OpenFile(key, os.O_RDONLY, 0)
}

// OpenFile will open a new file with flag and perm
func (a *AIO) OpenFile(key string, flag int, perm os.FileMode) <-chan *OpenResp {
	var or openRequest
	or.key = key
	or.flag = flag
	or.perm = perm
	or.resp = make(chan *OpenResp, 1)
	a.rq <- &or
	return or.resp
}

// Delete will delete a file
func (a *AIO) Delete(key string) <-chan error {
	var dr deleteRequest
	dr.key = key
	dr.errCh = make(chan error, 1)
	a.rq <- &dr
	return dr.errCh
}

// OpenResp is a response for open requests
type OpenResp struct {
	f   *File
	err error
}

// RWResp is a response for read/write requests
type RWResp struct {
	n   int
	err error
}

type openRequest struct {
	key  string
	flag int
	perm os.FileMode

	resp chan *OpenResp
}

type readRequest struct {
	f *os.File
	b []byte

	resp chan *RWResp
}

type writeRequest struct {
	f *os.File
	b []byte

	resp chan *RWResp
}

type closeRequest struct {
	f *os.File

	errCh chan error
}

type deleteRequest struct {
	key string

	errCh chan error
}

type sema chan struct{}

func newFile(r *openRequest, rq chan<- interface{}) (fp *File, err error) {
	var f File
	if f.f, err = os.OpenFile(r.key, r.flag, r.perm); err != nil {
		return
	}

	f.rq = rq
	fp = &f
	return
}
