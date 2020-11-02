package common

import (
	"github.com/sony/sonyflake"
	"strconv"
	"time"
)

var sf *sonyflake.Sonyflake

func init() {
	st := sonyflake.Settings{
		StartTime: time.Now(),
	}

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

// Generate return a distributed unique ID by sonyflake
func Generate() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(id, 10), nil
}
