package msg



//func m() {
//
//}

/**


eventStore  interface{
   check event
   push event
   delete event
}

eventBus


check event status  {onDone}


push event  // time handler

handler event







*/
//
////FromChan 把一个chan转换成事件流
//func FromChan(source <-chan interface{}) event.Observable {
//	return func(sink *event.Observer) error {
//		for {
//			select {
//			case <-sink.Done():
//				return nil
//			case data, ok := <-source:
//				if ok {
//					sink.Next(data)
//				} else {
//					return nil
//				}
//			}
//		}
//	}
//}
//
////Range 产生一段范围内的整数序列
//func Range(start int, count uint) event.Observable {
//	end := start + int(count)
//	return func(sink *event.Observer) error {
//		for i := start; i < end && !sink.IsDisposed(); i++ {
//			sink.Next(i)
//		}
//		return nil
//	}
//}
