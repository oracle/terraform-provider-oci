// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Group Represents the current state of a consumer group, including partition reservations and committed offsets.
type Group struct {

	// The streamId for which the group exists.
	StreamId *string `mandatory:"true" json:"streamId"`

	// The name of the consumer group.
	GroupName *string `mandatory:"true" json:"groupName"`

	// An array of the partition reservations of a group.
	Reservations []PartitionReservation `mandatory:"false" json:"reservations"`
}

func (m Group) String() string {
	return common.PointerString(m)
}
