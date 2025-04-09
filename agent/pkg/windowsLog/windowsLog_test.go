package windowsLog

import "testing"

func TestRun(t *testing.T) {
	if err := Run(UserEventType); err != nil {
		t.Error(err)
	}
}
