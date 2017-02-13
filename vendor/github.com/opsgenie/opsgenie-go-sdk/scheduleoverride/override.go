/*
Copyright 2016. All rights reserved.
Use of this source code is governed by a Apache Software
license that can be found in the LICENSE file.
*/

//Package override provides requests and response structures to achieve Schedule Override API actions.
package override

// AddScheduleOverrideRequest provides necessary parameter structure for adding Schedule Override
type AddScheduleOverrideRequest struct {
	APIKey      string   `json:"apiKey,omitempty"`
	Alias       string   `json:"alias,omitempty"`
	Schedule    string   `json:"schedule,omitempty"`
	User        string   `json:"user,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
	RotationIds []string `json:"rotationIds,omitempty"`
	Timezone    string   `json:"timezone,omitempty"`
}

// UpdateScheduleOverrideRequest provides necessary parameter structure for updating Schedule Override
type UpdateScheduleOverrideRequest struct {
	APIKey      string   `json:"apiKey,omitempty"`
	Alias       string   `json:"alias,omitempty"`
	Schedule    string   `json:"schedule,omitempty"`
	User        string   `json:"user,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
	RotationIds []string `json:"rotationIds,omitempty"`
	Timezone    string   `json:"timezone,omitempty"`
}

// DeleteScheduleOverrideRequest provides necessary parameter structure for deleting Schedule Override
type DeleteScheduleOverrideRequest struct {
	APIKey   string `url:"apiKey,omitempty"`
	Alias    string `url:"alias,omitempty"`
	Schedule string `url:"schedule,omitempty"`
}

// GetScheduleOverrideRequest provides necessary parameter structure for getting Schedule Override
type GetScheduleOverrideRequest struct {
	APIKey   string `url:"apiKey,omitempty"`
	Alias    string `url:"alias,omitempty"`
	Schedule string `url:"schedule,omitempty"`
}

// ListScheduleOverridesRequest provides necessary parameter structure for listing Schedule Override
type ListScheduleOverridesRequest struct {
	APIKey   string `url:"apiKey,omitempty"`
	Schedule string `url:"schedule,omitempty"`
}
