package windowsLog

import "testing"

func TestRun(t *testing.T) {
	if err := Run(RdpEventType); err != nil {
		t.Error(err)
	}
}
