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

type EventUnit map[string]interface{}
type ExternalFunctionForMapChan func(evtxMap chan *evtx.GoEvtxMap) error

type ExternalFunctionForMap func(evtxMap *evtx.GoEvtxMap)

var (
	FunctionMap = map[EventNameType]ExternalFunctionForMapChan{
		LoginEvenType: loginEvent,
	}
)

// -------------------------------------- public --------------------------------------

// Run 用于启动 eventName 对应的处理方式
//
// 参数：
//   - eventName: 事件名称
//   - mapChanFunctions: 外部函数列表，用于处理事件，默认使用 FunctionMap 中的函数，也可以通过 RegisterFunctionMap 注册新的函数
func Run(eventName EventNameType, mapChanFunctions ...ExternalFunctionForMapChan) error {
	eventFunc, ok := FunctionMap[eventName]
	if !ok {
		return errors.New("event not found")
	}
	if mapChanFunctions != nil {
		mapChanFunctions = append(mapChanFunctions, eventFunc)
	} else {
		mapChanFunctions = []ExternalFunctionForMapChan{
			eventFunc,
		}
	}
	return runBase(eventName, mapChanFunctions)
}

func RegisterFunctionMap(eventName EventNameType, function ExternalFunctionForMapChan) {
	FunctionMap[eventName] = function
}

func GetEventsForLoginEvent(evtxMap chan *evtx.GoEvtxMap) []EventUnit {
	events := make([]EventUnit, 0)
	for event := range evtxMap {
		if _, ok := LoginEvent[event.EventID()]; !ok {
			continue
		}
		eventInfo := getBaseInfo(event)
		addLoginEventInfoForEvent(event, &eventInfo)
		events = append(events, eventInfo)
	}
	return events
}

// base functions

func Wrapper(path string) *evtx.GoEvtxPath {
	p := evtx.Path(path)
	return &p
}

func GetString(evtxMap *evtx.GoEvtxMap, path *evtx.GoEvtxPath) string {
	value, err := evtxMap.GetString(path)
	if err != nil {
		logger.Error("GetString -> get string error: ", err)
		return ""
	}
	return value
}

// -------------------------------------- private --------------------------------------

func runBase(eventName EventNameType, mapChanFunctions []ExternalFunctionForMapChan) error {
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
		eventMapChan := eventReader.FastEvents()
		for _, f := range mapChanFunctions {
			if err := f(eventMapChan); err != nil {
				return err
			}
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
	events := GetEventsForLoginEvent(evtxMap)
	fmt.Println(events)
	return nil
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

func addLoginEventInfoForEvent(evtxMap *evtx.GoEvtxMap, eventUnit *EventUnit) {
	(*eventUnit)[LoginTypeKey] = getLoginType(GetString(evtxMap, Wrapper(logonTypePath)))
	(*eventUnit)[UsernameKey] = GetString(evtxMap, Wrapper(usernamePath))
	(*eventUnit)[SourceIpKey] = GetString(evtxMap, Wrapper(ipAddressPath))
	(*eventUnit)[SubjectUsernameKey] = GetString(evtxMap, Wrapper(subjectUserNamePath))
	(*eventUnit)[SubjectDomainKey] = GetString(evtxMap, Wrapper(subjectDomainPath))
	(*eventUnit)[ProcessNameKey] = GetString(evtxMap, Wrapper(processNamePath))
}
