package override

// AddScheduleOverrideResponse provides the response structure for adding Schedule Override
type AddScheduleOverrideResponse struct {
	Alias  string `json:"alias"`
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// UpdateScheduleOverrideResponse provides the response structure for updating Schedule Override
type UpdateScheduleOverrideResponse struct {
	Alias  string `json:"alias"`
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// DeleteScheduleOverrideResponse provides the response structure for deleting Schedule Override
type DeleteScheduleOverrideResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// GetScheduleOverrideResponse provides the response structure for getting Schedule Override
type GetScheduleOverrideResponse struct {
	Alias       string   `json:"alias,omitempty"`
	User        string   `json:"user,omitempty"`
	RotationIds []string `json:"rotationIds,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
	Timezone    string   `json:"timezone,omitempty"`
}

// ListScheduleOverridesResponse provides the response structure for listing Schedule Override
type ListScheduleOverridesResponse struct {
	Overrides []GetScheduleOverrideResponse `json:"overrides,omitempty"`
}
