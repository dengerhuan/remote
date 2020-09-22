package event

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

/**


middleWare(event)-> Handler ->

   -> receive event
   -> send event


type funcObject func(middleWare)

funcObject.func(Handler)-->  decorate middleWare-handler. bind handler ->receive data

frontFunc -> funcObject(middleWare) ->middleWare-buildData

*/
func (ob Observable) Listen(onNext NextHandler, onError func(error), onComplete func()) *Observer {

	observer := &Observer{nil, nil, onNext}

	result := ob(observer)

	if result != nil {
		onError(result)
	} else {
		onComplete()
	}
	return observer
}

func (ob Observable) Map(f func(interface{}) interface{}) Observable {

	return func(observer *Observer) error {
		return ob(observer.Create(func(e *Event) {
			observer.Next(f(e.Data))
		}))
	}
}

func NewEvent(eventId string, data interface{}) Observable {

	return func(observer *Observer) error {

		observer.Context, observer.ctxFunc = context.WithTimeout(context.Background(), EventStore.timeout)

		if EventStore.replayAtackCheck(eventId) {
			return errors.New("on process")
		}
		//  记录时间
		EventStore.add(eventId)

		go observer.Next(data)

		select {
		case <-observer.Done():
			// 删除事件
			EventStore.remove(eventId)

			if !strings.Contains(observer.Err().Error(), "canceled") {
				fmt.Println(observer.Err())
				return errors.New("process time out")
			}

		}

		// 下发事件
		return nil
	}
}

func s() {
	NewEvent("s", "s").Listen(NextFunc(func(e *Event) {}), func(err error) {

	}, func() {

	})
}

//Subscribe 对外同步订阅Observable，可以在收到的事件中访问Observer的Dispose函数来终止事件流
func (ob Observable) Subscribe(onNext NextHandler) error {
	ctx, cancel := context.WithCancel(context.Background())
	return ob(&Observer{ctx, cancel, onNext})
}

//SubscribeAsync 对外异步订阅模式，返回一个可用于终止事件流的控制器,可以在未收到数据时也能终止事件流
func (ob Observable) SubscribeAsync(onNext NextHandler, onError func(error), onComplete func()) *Observer {
	ctx, cancel := context.WithCancel(context.Background())
	source := &Observer{ctx, cancel, onNext}
	go func() {
		err := ob(source)
		if !source.IsDisposed() {
			if err != nil {
				onError(err)
			} else {
				onComplete()
			}
		}
	}()
	return source
}
