package pool

import (
	"base/log"
	"errors"
	"os"
	"os/signal"
	"time"
)

var logger = log.GetLogger()

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

var ErrTimeout = errors.New("timeout")

var ErrInterrupt = errors.New("interrupt")

func NewRunner(d time.Duration) *Runner {

	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d)}
}

func (r *Runner) Add(task ...func(int)) {

	r.tasks = append(r.tasks, task...)
}

///////
//
func (r *Runner) run() error {

	// range 遍历执行，， 如果终止了  返回error
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}

	return nil
}

func (r *Runner) goInterrupt() bool {

	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Start() error {
	// 注册中断事件
	signal.Notify(r.interrupt, os.Interrupt)

	// 启动task
	go func() {
		r.complete <- r.run()
	}()

	// 接受异常
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
