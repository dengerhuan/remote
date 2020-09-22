package pool

import (
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewWork(maxgoroutine int) *Pool {

	p := Pool{work: make(chan Worker)}

	p.wg.Add(maxgoroutine)

	for i := 0; i < maxgoroutine; i++ {

		go func() {
			// 一直阻塞等待 ，close 后 结束阻塞
			for w := range p.work {
				w.Task()
			}

			logger.Info("shutdown go r")
			p.wg.Done()

		}()

	}
	return &p
}

func (p *Pool) Run(worker Worker) {
	p.work <- worker
}
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
