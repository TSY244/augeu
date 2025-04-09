package windowsLog

import "testing"

func TestRun(t *testing.T) {
	if err := Run(SysmonEventType); err != nil {
		t.Error(err)
	}
}
