package dummy

import (
	"io"
	"sync"
)

const (
	closeBacklog    = 128
	closeConcurrent = 32
)

// START OMIT
func (p *Pool) manage() {
	closers := make(chan io.Closer, closeBacklog)
	var closeWG sync.WaitGroup
	closeWG.Add(closeConcurrent)
	for i := 0; i < closeConcurrent; i++ {
		go func() {
			defer closeWG.Done()
			for c := range closers {
				p.CloseErrorHandler(c.Close())
			}
		}()
	}
	defer func() {
		close(closers)
		closeWG.Wait()
	}()

	for {
		select {
		case c := <-p.closeConn:
			closers <- c
		}
	}
}

// END OMIT
