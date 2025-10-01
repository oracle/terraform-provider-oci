// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QueueFailureDestinationDetails The destination queue or channel in the Queue service to which to send the response of the failed detached function invocation.
// Example: `{"kind": "QUEUE", "queueId": "queue_OCID", "channelId": "channel_Id"}`
type QueueFailureDestinationDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the queue.
	QueueId *string `mandatory:"true" json:"queueId"`

	// The ID of the channel in the queue.
	ChannelId *string `mandatory:"false" json:"channelId"`
}

func (m QueueFailureDestinationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueueFailureDestinationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m QueueFailureDestinationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeQueueFailureDestinationDetails QueueFailureDestinationDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeQueueFailureDestinationDetails
	}{
		"QUEUE",
		(MarshalTypeQueueFailureDestinationDetails)(m),
	}

	return json.Marshal(&s)
}
