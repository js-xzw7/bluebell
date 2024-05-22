package sonyflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake *sonyflake.Sonyflake
)

func Init(machineid uint16) (err error) {
	sonyMachineID := machineid

	t, _ := time.Parse("2006-01-02", "2024-01-01")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: func() (uint16, error) { return sonyMachineID, nil },
	}

	sonyFlake, err = sonyflake.New(settings)
	return
}

// 返回生成的id值
func GetId() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}
