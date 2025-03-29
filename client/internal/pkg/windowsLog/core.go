package windowsLog

import (
	"errors"
	"fmt"
	"github.com/0xrawsec/golang-evtx/evtx"
	"os"
)

type eventNameType string
type eventFunctionType func(evtxMap chan *evtx.GoEvtxMap) error

var (
	FunctionMap = map[eventNameType]eventFunctionType{
		LoginEvenType: loginEvent,
	}
)

func Run(eventName eventNameType) error {
	return base(eventName, FunctionMap[eventName])
}

func base(eventName eventNameType, f eventFunctionType) error {
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
		defer file.Close()
		eventReader, err := evtx.New(file)
		if err != nil {
			return err
		}
		defer eventReader.Close()
		if err := f(eventReader.FastEvents()); err != nil {
			return err
		}
	}
	return nil
}

func getBaseInfo(event *evtx.GoEvtxMap) map[string]interface{} {
	return map[string]interface{}{
		EventIdKey: event.EventID(),
	}
}

func loginEvent(evtxMap chan *evtx.GoEvtxMap) error {
	for event := range evtxMap {
		eventInfo := getBaseInfo(event)
		fmt.Printf("eventInfo: %v\n", eventInfo)
	}
	return nil
}
