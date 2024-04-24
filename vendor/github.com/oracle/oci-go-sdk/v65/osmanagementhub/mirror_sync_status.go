// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MirrorSyncStatus Status summary of the mirror sync.
type MirrorSyncStatus struct {

	// Total number of software sources that have not yet been synced.
	Unsynced *int `mandatory:"true" json:"unsynced"`

	// Total number of software sources that are queued for sync.
	Queued *int `mandatory:"true" json:"queued"`

	// Total number of software sources currently syncing.
	Syncing *int `mandatory:"true" json:"syncing"`

	// Total number of software sources that successfully synced.
	Synced *int `mandatory:"true" json:"synced"`

	// Total number of software sources that failed to sync.
	Failed *int `mandatory:"true" json:"failed"`
}

func (m MirrorSyncStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MirrorSyncStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
