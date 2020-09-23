package instance

import (
	"errors"
	"log"
	"strings"
	"sync"
)

type LogicManager interface {
	Register(logic *RdLogic)
}

var VehicleGroup = logicGroup{list: make(map[string]*RdLogic)}
var CockpitGroup = logicGroup{list: make(map[string]*RdLogic)}
var ConsoleGroup = logicGroup{list: make(map[string]*RdLogic)}
var MonitorGroup = logicGroup{list: make(map[string]*RdLogic)}

func registerByLogic(ty byte) logicGroup {

	switch ty {
	case 0:
		log.Print("device type is :car ")
		return VehicleGroup
	case 1:
		log.Print("device type is : cockpit ")
		return CockpitGroup
	case 2:
		log.Print(" device type is :console ")
		return ConsoleGroup
	default:
		return MonitorGroup
	}

}

type logicGroup struct {
	mutex sync.RWMutex
	list  map[string]*RdLogic
}

func (l *logicGroup) Register(logic *RdLogic) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	l.list[logic.id] = logic
	log.Println("device detail info ", logic)
}

func (l *logicGroup) ReleaseOne(carId string) error {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	logic, ok := l.list[carId]

	if ok {
		logic.vg = make([]*RdLogic, 0)
		logic.Unlock()
	}
	return nil
}

func (l *logicGroup) GetById(carId string) (*RdLogic, error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	logic, ok := l.list[carId]

	if !ok {
		return nil, errors.New("car state no ok")
	}

	logic.Lock()
	return logic, nil
}

// check available cockpit
func (l *logicGroup) GetOne() (*RdLogic, error) {

	// status 0 unlock  1 lock  2 working

	l.mutex.RLock()
	defer l.mutex.RUnlock()

	log.Println(l.list)
	cockpitId := ""
	for _, c := range l.list {
		if c.state == 0 {
			cockpitId = c.id
			break
		}
	}

	log.Println(cockpitId)

	if strings.EqualFold(cockpitId, "") {
		log.Println("no available cockpit device")
		return nil, errors.New("no available cockpit")
	} else {
		logic, _ := l.list[cockpitId]
		logic.Lock()
		return logic, nil
	}
}
