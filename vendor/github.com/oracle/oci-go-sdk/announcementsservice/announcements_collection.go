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

// AnnouncementsCollection Results of annoucements search. Contains both announcements, and user specific status of the announcments
type AnnouncementsCollection struct {

	// collection of announcements
	Items []AnnouncementSummary `mandatory:"false" json:"items"`

	// user specific status of found announcements
	UserStatuses []AnnouncementUserStatusDetails `mandatory:"false" json:"userStatuses"`
}

func (m AnnouncementsCollection) String() string {
	return common.PointerString(m)
}
