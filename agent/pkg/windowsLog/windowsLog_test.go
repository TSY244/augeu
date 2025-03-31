package windowsLog

import "testing"

func TestRun(t *testing.T) {
	if err := Run(LoginEvenType); err != nil {
		t.Error(err)
	}
}
