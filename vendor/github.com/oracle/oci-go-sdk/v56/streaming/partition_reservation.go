// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
