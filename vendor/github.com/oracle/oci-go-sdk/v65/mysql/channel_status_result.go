// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChannelStatusResult Status information for the channel attached to the DB System.
type ChannelStatusResult struct {

	// The OCID of the Channel for which the status is gathered.
	ChannelId *string `mandatory:"true" json:"channelId"`

	// Specifies if the channel is healthy or not. If healthy, replication target DB System is connected to the
	// source and no replication errors are seen.
	IsHealthy *bool `mandatory:"false" json:"isHealthy"`

	// Specifies if all transactions received by this channel are executed and their GTIDs are part of gtid_executed
	// set.
	IsReceivedGtidSetApplied *bool `mandatory:"false" json:"isReceivedGtidSetApplied"`

	// Channel errors identified, if there are any.
	Errors []string `mandatory:"false" json:"errors"`

	// The channel lag, with respect to the immediate source of the channel.
	// If the channel is configured with replication delay, the channel lag includes the replication delay.
	LagDuration *string `mandatory:"false" json:"lagDuration"`
}

func (m ChannelStatusResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChannelStatusResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
