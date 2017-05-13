package syscall

import (
	"github.com/itsmontoya/async"
)

// Pread will perform a pread
func Pread(fd int, b []byte, offset int64) (n int, err error) {
	resp := <-PreadAsync(fd, b, offset)
	n = resp.N
	err = resp.Err
	p.releaseRWResp(resp)
	return
}

// PreadAsync will perform syscall.Pread asynchronously
func PreadAsync(fd int, b []byte, offset int64) <-chan *RWResp {
	// Acquire request
	req := p.acquirePreadReq()

	// Set values
	req.fd = fd
	req.b = b
	req.offset = offset

	// Send request to queue
	async.Queue(req)
	return req.resp
}
