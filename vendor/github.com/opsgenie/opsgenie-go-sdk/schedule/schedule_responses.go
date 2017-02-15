package schedule

// Create schedule response structure
type CreateScheduleResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// Update schedule response structure
type UpdateScheduleResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// Delete schedule response structure
type DeleteScheduleResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// Participant
type Participant struct {
	Participant string `json:"participant,omitempty"`
	Type        string `json:"type,omitempty"`
}

// RotationInfo defines the structure for each rotation definition
type RotationInfo struct {
	Id             string        `json:"id,omitempty"`
	StartDate      string        `json:"startDate,omitempty"`
	EndDate        string        `json:"endDate,omitempty"`
	RotationType   string        `json:"rotationType,omitempty"`
	Participants   []Participant `json:"participants,omitempty"`
	Name           string        `json:"name,omitempty"`
	RotationLength int           `json:"rotationLength,omitempty"`
	Restrictions   []Restriction `json:"restrictions,omitempty"`
}

// Get schedule structure
type GetScheduleResponse struct {
	Id    string         `json:"id,omitempty"`
	Name  string         `json:"name,omitempty"`
	Team  string         `json:"team,omitempty"`
	Rules []RotationInfo `json:"rules,omitempty"`
}

// List schedule response structure
type ListSchedulesResponse struct {
	Schedules []GetScheduleResponse `json:"schedules,omitempty"`
}

// Get Timeline schedule structure
type GetTimelineScheduleResponse struct {
	Schedule TimelineScheduleResponse `json:"schedule,omitempty"`
	Took     int                      `json:"took,omitempty"`
	Timeline TimelineTimelineResponse `json:"timeline,omitempty"`
}

type TimelineScheduleResponse struct {
	Id       string `json:"id,omitempty"`
	Team     string `json:"team,omitempty"`
	Name     string `json:"name,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
}

type TimelineTimelineResponse struct {
	StartTime     uint64           			`json:"startTime,omitempty"`
	EndTime       uint64           			`json:"endTime,omitempty"`
	FinalSchedule TimelineFinalScheduleResponse 	`json:"finalSchedule,omitempty"`
}

type TimelineFinalScheduleResponse struct {
	Rotations []TimelineRotation `json:"rotations,omitempty"`
}

// Rotation defines the structure for each rotation definition
type TimelineRotation struct {
	Name    string    `json:"name,omitempty"`
	Id      string    `json:"id,omitempty"`
	Order   float64   `json:"order,omitempty"`
	Periods []Periods `json:"periods,omitempty"`
}

type Periods struct {
	StartTime           uint64		  `json:"startTime,omitempty"`
	EndTime             uint64 		  `json:"endTime,omitempty"`
	Type                string                `json:"type,omitempty"`
	FromUsers 	    []FromUsers 	  `json:"fromUsers,omitempty"`
	Recipients          []Recipients          `json:"recipients,omitempty"`
}

type FromUsers struct {
	DisplayName string `json:"displayName,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
}

type Recipients struct {
	DisplayName string `json:"displayName,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
}

// WhoIsOnCallParticipant
type WhoIsOnCallParticipant struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Forwarded bool `json:"forwarded,omitempty"`
	Participants []*WhoIsOnCallParticipant `json:"participants,omitempty"`
	NotifyType string `json:"notifyType,omitempty"`
}

// WhoIsOnCall response structure
type WhoIsOnCallResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Participants []WhoIsOnCallParticipant `json:"participants,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	IsEnabled bool `json:"isEnabled,omitempty"`
}
