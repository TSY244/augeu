package convert

import (
	"augeu/agent/pkg/windowsLog"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/util/convert"
)

func LoginEvent2RLoginEventResq(event windowsLog.EventUnit) *models.LoginEvent {

	eventId := event[windowsLog.EventIdKey].(int64)
	eventTime := event[windowsLog.EventTimeKey].(string)
	machineUuid := event[windowsLog.MachineUUIDKey].(string)
	loginType := event[windowsLog.LoginTypeKey].(string)
	username := event[windowsLog.UsernameKey].(string)
	ipAddress := event[windowsLog.SourceIpKey].(string)
	subjectUserName := event[windowsLog.SubjectUsernameKey].(string)
	subjectDomain := event[windowsLog.SubjectDomainKey].(string)
	processName := event[windowsLog.ProcessNameKey].(string)

	// str -> *strfmt.DateTime
	tempTime, err := convert.StrTime2DateTime(eventTime)
	if err != nil {
		logger.Errorf("tempTime is error: %v", err)
		tempTime, _ = convert.StrTime2DateTime(`0-0-0 12:00:00`)
	}

	return &models.LoginEvent{
		EventID:       &eventId,
		EventTime:     tempTime,
		MachineUUID:   &machineUuid,
		LoginType:     &loginType,
		Username:      &username,
		SourceIP:      &ipAddress,
		SubjectUser:   &subjectUserName,
		SubjectDomain: &subjectDomain,
		ProcessName:   &processName,
	}
}
