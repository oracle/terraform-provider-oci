// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// A description of the Queue API
//

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Stats The stats for a queue or a dead letter queue.
type Stats struct {

	// The approximate number of visible messages (available for delivery) currently in the queue.
	VisibleMessages *int64 `mandatory:"true" json:"visibleMessages"`

	// The approximate number of messages delivered to a consumer but not yet deleted and so unavailable for re-delivery.
	InFlightMessages *int64 `mandatory:"true" json:"inFlightMessages"`

	// The approximate size of the queue in bytes. Sum of the size of visible of in-flight messages.
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`
}

func (m Stats) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Stats) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
