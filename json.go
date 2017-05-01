package trends

import (
	"fmt"
	"strconv"
	"time"
)

type JSONTime struct {
	time.Time
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	unix, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return fmt.Errorf("could not parse time %s: %v", string(b), err)
	}
	t.Time = time.Unix(unix, 0)

	return nil
}
