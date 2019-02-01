// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// A description of the AnnouncementsService API
//

package announcementsservice

import (
	"github.com/oracle/oci-go-sdk/common"
)

// NotificationFollowupDetails Information represents a notification follow-up
type NotificationFollowupDetails struct {

	// The follow-up message, a markdown format input
	Message *string `mandatory:"false" json:"message"`

	// When the update is made
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m NotificationFollowupDetails) String() string {
	return common.PointerString(m)
}
