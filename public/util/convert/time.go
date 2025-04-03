package convert

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"time"
)

func ConvertStrfmtToTime(dt *strfmt.DateTime) (time.Time, error) {
	if dt == nil {
		return time.Time{}, fmt.Errorf("input is nil pointer")
	}
	// strfmt.DateTime 本质是 time.Time 的别名，可直接取值
	return time.Time(*dt), nil
}
