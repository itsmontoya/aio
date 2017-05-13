package syscall

func newRWResp() *RWResp {
	return &RWResp{}
}

// RWResp is a response for a read or write request
type RWResp struct {
	N   int
	Err error
}
