// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicationEstimate Details for response from replication estimation.
type ReplicationEstimate struct {

	// The rate of change on source filesystem which was used to provide the estimate in MegaBytes per second.
	ChangeRateInMBps *int `mandatory:"true" json:"changeRateInMBps"`

	// Specifies whether replication can be enabled on the file system.
	IsReplicationSupported *bool `mandatory:"true" json:"isReplicationSupported"`

	// The minimum supported replication interval for specified file system in minutes.
	MinimumSupportedIntervalInMinutes *int `mandatory:"true" json:"minimumSupportedIntervalInMinutes"`

	// The approximate time required for the base sync between source and target to finish.
	EstimatedBaseCopyTimeInMinutes *int `mandatory:"true" json:"estimatedBaseCopyTimeInMinutes"`

	// Array of allowed target region names which can be paired with source file system.
	AllowedTargetRegions []string `mandatory:"true" json:"allowedTargetRegions"`
}

func (m ReplicationEstimate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationEstimate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
