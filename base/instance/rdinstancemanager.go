package instance

import (
	"sync"
)

var InstanceManager = &instanceManager{store: map[string]RdInstance{}}

type instanceManager struct {
	mutex sync.RWMutex
	store map[string]RdInstance
}

func (m *instanceManager) Create(orderId string) RdInstance {

	instance := NewRdInstance(orderId)

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	m.store[orderId] = instance
	return instance
}

func (m *instanceManager) Get(orderId string) (RdInstance, bool) {

	instance, ok := m.store[orderId]

	return instance, ok

}




// check
// set cockpit
// set vehicle
// start
// stop
