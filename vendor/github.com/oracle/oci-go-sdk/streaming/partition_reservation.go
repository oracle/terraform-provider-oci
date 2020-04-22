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

// PartitionReservation Represents the state of a single partition reservation.
type PartitionReservation struct {

	// The partition for which the reservation applies.
	Partition *string `mandatory:"false" json:"partition"`

	// The latest offset which has been committed for this partition.
	CommittedOffset *int64 `mandatory:"false" json:"committedOffset"`

	// The consumer instance which currently has the partition reserved.
	ReservedInstance *string `mandatory:"false" json:"reservedInstance"`

	// A timestamp when the current reservation expires.
	TimeReservedUntil *common.SDKTime `mandatory:"false" json:"timeReservedUntil"`
}

func (m PartitionReservation) String() string {
	return common.PointerString(m)
}
