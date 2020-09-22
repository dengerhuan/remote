package instance

import (
	"errors"
	"log"
	"strings"
	"sync"
)

/*
“status”:[
  {“carid”: “12345678901234567”, “adminState”:true, “rdState”:true, “speed”:17.3, “delay”：57, “distance”:23.3, “time”:41, “longitude”:43.3133, “latitude”: 123.323436, “direction”:53.3, “timeStamp”: “2000-07-05 14:35:14.99” },
  {“carid”: “12345678901234568”, “adminState”:true, “rdState”:false, “speed”:17.3, “delay”：0, “distance”:23.3, “time”:41, “longitude”:43.3123, “latitude”: 123.323343, “direction”:53.3, “timeStamp”: “2000-07-05 14:35:14.97”  }

],

	Id         string
	Statue     bool
	CaseInfo   string
	Step       int //  0 init |1 add car |2 add cockpit| 3 word
	Car        *RdLogic
	Cockpit    *RdLogic
	adminState bool //
*/

type RdInstance interface {
	// get cockpit and lock and validate
	SetCockpit(cockpitId string) error

	// 设置 台驾
	SetVehicle(vid string) error

	GetCockpit() *RdLogic

	// 设置 台驾
	GetVehicle() *RdLogic

	//  是否在结束远程驾驶
	GetState() bool
	// 设置 台驾
	GetConsole() *RdLogic

	GetCockpitId() string

	Start()
	Stop()
	Cancer(cancerId int)

	//Cancelled() (int, string, error)
	IsCancelled() bool
}

func NewRdInstance(orderId string) RdInstance {
	return &rdInstance{Id: orderId}
}

type rdInstance struct {
	mu        sync.RWMutex
	Id        string
	State     bool // work no work
	VehicleId string
	CockpitId string
	Vehicle   *RdLogic
	Cockpit   *RdLogic

	StartTime  int64
	EndTime    int64
	Cancelled  bool
	CancerCode int
	CancerInfo string

	Console *RdLogic
}

func (r *rdInstance) GetState() bool {
	return r.State
}

func (r *rdInstance) GetCockpitId() string {
	return r.CockpitId
}

func (r *rdInstance) GetConsole() *RdLogic {
	return r.Console
}
func (r *rdInstance) GetCockpit() *RdLogic {
	return r.Cockpit
}

// 设置 台驾
func (r *rdInstance) GetVehicle() *RdLogic {
	return r.Vehicle
}

//
func (r *rdInstance) IsCancelled() bool {
	return r.Cancelled
}

// set 保持 id // start 绑定 tunnel
// create empty order with id and carid
func (r *rdInstance) SetCockpit(cockpitId string) error {

	log.Println("set cockpit", cockpitId)
	// cockpit 为空 随机指定， 不为空按照id指定

	if !strings.EqualFold(cockpitId, "") {
		log.Println("get cockpit by cockpit id not support")
	}

	logic, err := CockpitGroup.GetOne()

	if err != nil {
		return err
	}

	if !logic.ChannelValid() {
		return errors.New("rpc nor ok")
	}
	//r.Cockpit = logic
	r.CockpitId = logic.id
	return nil
}

// 设置 台驾
func (r *rdInstance) SetVehicle(vid string) error {

	logic, err := VehicleGroup.GetById(vid)
	if err != nil {
		return err
	}

	if !logic.ChannelValid() {
		return errors.New("rpc nor ok")
	}

	//r.Vehicle = logic
	r.VehicleId = logic.id
	return nil
}

/**

 */
func (r *rdInstance) Start() {
	if r.State {
		return
	}
	r.State = true

	// set console
	console, err := getConsole()

	// set vehicle
	vehicle, err := VehicleGroup.GetById(r.VehicleId)

	if err != nil {
		log.Println(err)
	}
	r.Vehicle = vehicle

	// set cockpiy
	cockpit, err := CockpitGroup.GetById(r.CockpitId)
	if err != nil {
		log.Println(err)
	}
	r.Cockpit = cockpit

	if err != nil {

		r.Vehicle.vg = append(r.Vehicle.vg, r.Cockpit)
	} else {
		r.Console = console
		r.Vehicle.vg = append(r.Vehicle.vg, r.Cockpit, console)
	}

	r.Cockpit.vg = append(r.Cockpit.vg, r.Vehicle)

	// set state in use
	r.Vehicle.SetState(2)
	r.Cockpit.SetState(2)
}
func (r *rdInstance) Stop() {

	if !r.State {
		return
	}

	r.State = false

	r.Vehicle.Unlock()
	r.Cockpit.Unlock()

	r.Vehicle.vg = make([]*RdLogic, 0)
	r.Cockpit.vg = make([]*RdLogic, 0)

	r.Cancelled = true
	r.CancerCode = 222
	r.CancerInfo = "用户强制停止"
}

func (r *rdInstance) Cancer(cancerId int) {

	r.Cancelled = true
	r.CancerCode = cancerId

}

func getConsole() (*RdLogic, error) {

	// status 0 unlock  1 lock  2 working

	rd, ok := ConsoleGroup.list["001"]
	if ok {
		return rd, nil
	}
	return nil, errors.New("no console")
}
