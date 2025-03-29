package windowsLog

import (
	"augeu/client/pkg/machine"
	"augeu/public/pkg/logger"
	_const "augeu/public/util/const"
	"errors"
	"fmt"
	"github.com/0xrawsec/golang-evtx/evtx"
	"os"
)

type EventNameType string
type EventFunctionType func(evtxMap chan *evtx.GoEvtxMap) error
type EventUnit map[string]interface{}

var (
	FunctionMap = map[EventNameType]EventFunctionType{
		LoginEvenType: loginEvent,
	}
)

// -------------------------------------- public --------------------------------------

func Run(eventName EventNameType) error {
	return runBase(eventName, FunctionMap[eventName])
}

// -------------------------------------- private --------------------------------------

func runBase(eventName EventNameType, f EventFunctionType) error {
	if f == nil {
		return errors.New("event not found")
	}
	pathList, ok := EventToFilePath[eventName]
	if !ok {
		return errors.New("event file path not found")
	}
	for _, path := range pathList {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		eventReader, err := evtx.New(file)
		if err != nil {
			return err
		}
		if err := f(eventReader.Events()); err != nil {
			return err
		}

		// 资源释放
		file.Close()
		eventReader.Close()
	}
	return nil
}

func getBaseInfo(event *evtx.GoEvtxMap) EventUnit {
	// time : year-month-day hour:minute:second
	machineUUid, err := machine.GetWindowsGuid()
	if err != nil {
		fmt.Println("getBaseInfo -> get windows guid error: ", err)
	}
	return EventUnit{
		EventIdKey:     event.EventID(),
		EventTimeKey:   event.TimeCreated().Format(_const.TimeFormat),
		MachineUUIDKey: machineUUid,
	}
}

// login event functions

// loginEvent 用于处理登录事件
func loginEvent(evtxMap chan *evtx.GoEvtxMap) error {
	//fmt.Println("total event count: ", len(evtxMap))
	events := getEvents(evtxMap)
	fmt.Println(events)
	return nil
}

func getEvents(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := LoginEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		addLoginEventInfo(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

func addLoginEventInfo(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	(*eventUnit)[LoginTypeKey] = getLoginType(getString(evtxMap, wapper(logonTypePath)))
	(*eventUnit)[UsernameKey] = getString(evtxMap, wapper(usernamePath))
	(*eventUnit)[SourceIpKey] = getString(evtxMap, wapper(ipAddressPath))
	(*eventUnit)[SubjectUsernameKey] = getString(evtxMap, wapper(subjectUserNamePath))
	(*eventUnit)[SubjectDomainKey] = getString(evtxMap, wapper(subjectDomainPath))
	(*eventUnit)[ProcessNameKey] = getString(evtxMap, wapper(processNamePath))
}

func getLoginType(typeId string) string {
	switch typeId {
	case "2":
		return "Interactive（交互式登录）"
	case "3":
		return "Network（网络登录）"
	case "4":
		return "Batch（批处理登录）"
	case "5":
		return "Service（服务登录）"
	case "7":
		return "Unlock（解锁登录）"
	case "8":
		return "NetworkCleartext（网络明文登录）"
	case "9":
		return "NewCredentials（新凭证登录）"
	case "10":
		return "RemoteInteractive（远程交互登录）"
	case "11":
		return "CachedInteractive（缓存交互登录）"
	default:
		logger.Error("getLoginType -> login type not found: ", typeId)
		return "Unknown"
	}
}

// base functions

func wapper(path string) *evtx.GoEvtxPath {
	p := evtx.Path(path)
	return &p
}

func getString(evtxMap *evtx.GoEvtxMap, path *evtx.GoEvtxPath) string {
	value, err := evtxMap.GetString(path)
	if err != nil {
		logger.Error("getString -> get string error: ", err)
		return ""
	}
	return value
}
