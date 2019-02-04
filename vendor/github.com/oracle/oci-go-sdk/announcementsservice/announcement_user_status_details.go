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

// AnnouncementUserStatusDetails An announcement status
type AnnouncementUserStatusDetails struct {

	// The OCID of the announcement this status belongs to
	UserStatusAnnouncementId *string `mandatory:"true" json:"userStatusAnnouncementId"`

	// The OCID of the user this status belongs to
	UserId *string `mandatory:"true" json:"userId"`

	// The date and time the announcement was acknowledged, in the format defined by RFC3339
	// Example: `2016-07-22T17:43:01.389+0000`
	TimeAcknowledged *common.SDKTime `mandatory:"false" json:"timeAcknowledged"`
}

func (m AnnouncementUserStatusDetails) String() string {
	return common.PointerString(m)
}
