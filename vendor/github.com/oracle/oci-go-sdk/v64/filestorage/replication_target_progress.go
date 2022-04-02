// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// ReplicationTargetProgress The progress percentage of target job and source-target combined progress percentage.
type ReplicationTargetProgress struct {

	// The Progress percentage of the target job.
	TargetProgressPercentage *int `mandatory:"true" json:"targetProgressPercentage"`

	// The average Progress percentage between the source job and target job.
	CombinedProgressPercentage *int `mandatory:"true" json:"combinedProgressPercentage"`

	// The num of the last snapshot which was completely applied to the target file system.
	LastSnapshotNum *string `mandatory:"true" json:"lastSnapshotNum"`

	// The num of the new snapshot which is currently being applied to the target file system.
	NewSnapshotNum *string `mandatory:"true" json:"newSnapshotNum"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of replication.
	ReplicationId *string `mandatory:"true" json:"replicationId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of replicationTarget.
	ReplicationTargetId *string `mandatory:"true" json:"replicationTargetId"`
}

func (m ReplicationTargetProgress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationTargetProgress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
