package syscall

import "syscall"

func newPreadRequest() *preadRequest {
	return &preadRequest{
		resp: make(chan *RWResp),
	}
}

type preadRequest struct {
	fd     int
	b      []byte
	offset int64

	resp chan *RWResp
}

func (req *preadRequest) Action() {
	resp := p.acquireRWResp()
	resp.N, resp.Err = syscall.Pread(req.fd, req.b, req.offset)
	req.resp <- resp
	p.releasePreadReq(req)
}
