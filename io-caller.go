package dummy

import "io"

type sentinelCloser int

func (s sentinelCloser) Close() error {
	panic("should never get called")
}

// START OMIT
var dialConnSentinel = sentinelCloser(1) // HL

func (p *Pool) Acquire() (io.Closer, error) {
	r := make(chan io.Closer)
	p.acquire <- r
	c <- r
	if c == dialConnSentinel { // HL
		return p.dial()
	}
	return c, nil
}

func (p *Pool) manager() {
	for {
		select {
		case r := <-p.acquire:
			if connAvailable {
				r <- conn
				return
			}
			r <- dialConnSentinel // HL
		}
	}
}

// END OMIT
