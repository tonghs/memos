package store

import "github.com/usememos/memos/api"

// ActivityRaw is the sqlite model for an Activity.
// Fields have exactly the same meanings as Activity.
type ActivityRaw struct {
	ID int

	// Standard fields
	CreatorID int
	CreatedTs int64

	// Domain specific fields
	Type    api.ActivityType
	Level   api.ActivityLevel
	Payload string
}

// ToActivity creates an instance of Activity based on the ActivityRaw.
func (raw *ActivityRaw) ToActivity() *api.Activity {
	return &api.Activity{
		ID: raw.ID,

		CreatorID: raw.CreatorID,
		CreatedTs: raw.CreatedTs,

		Type:    raw.Type,
		Level:   raw.Level,
		Payload: raw.Payload,
	}
}
