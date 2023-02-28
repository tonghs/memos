package store

import (
	"github.com/usememos/memos/api"
)

type SystemSettingRaw struct {
	Name        api.SystemSettingName
	Value       string
	Description string
}

func (raw *SystemSettingRaw) ToSystemSetting() *api.SystemSetting {
	return &api.SystemSetting{
		Name:        raw.Name,
		Value:       raw.Value,
		Description: raw.Description,
	}
}
