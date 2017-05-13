package syscall

import "sync"

// Global pool for requests and responses
var p = newPools()

func newPools() *pools {
	var p pools

	// Requests
	p.preadReqs.New = func() interface{} {
		return newPreadRequest()
	}

	p.rwResps.New = func() interface{} {
		return newRWResp()
	}

	return &p
}

type pools struct {
	preadReqs sync.Pool

	rwResps sync.Pool
}

func (p *pools) acquirePreadReq() (req *preadRequest) {
	var ok bool
	if req, ok = p.preadReqs.Get().(*preadRequest); !ok {
		panic("invalid pool type")
	}

	return
}

func (p *pools) acquireRWResp() (resp *RWResp) {
	var ok bool
	if resp, ok = p.rwResps.Get().(*RWResp); !ok {
		panic("invalid pool type")
	}

	return
}
func (p *pools) releasePreadReq(req *preadRequest) {
	req.fd = 0
	req.b = nil
	req.offset = 0
	p.preadReqs.Put(req)
}

func (p *pools) releaseRWResp(resp *RWResp) {
	resp.N = 0
	resp.Err = nil
	p.rwResps.Put(resp)
}
